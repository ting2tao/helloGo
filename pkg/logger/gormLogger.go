package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

type gormLogger struct {
	SlowThreshold                       time.Duration
	SourceField                         string
	SkipErrRecordNotFound               bool
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New() *gormLogger {
	var (
		infoStr      = Green + "%s\n" + Reset + Green + "[info] " + Reset
		warnStr      = BlueBold + "%s\n" + Reset + Magenta + "[warn] " + Reset
		errStr       = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
		traceStr     = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Cyan + "[trace_id:%v]" + Reset + " %s"
		traceWarnStr = Green + "%s " + Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Cyan + "[trace_id:%v]" + Magenta + " %s" + Reset
		traceErrStr  = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Cyan + "[trace_id:%v]" + Reset + " %s"
	)
	return &gormLogger{
		SlowThreshold:         time.Millisecond * 10,
		SkipErrRecordNotFound: false,
		infoStr:               infoStr,
		warnStr:               warnStr,
		errStr:                errStr,
		traceStr:              traceStr,
		traceWarnStr:          traceWarnStr,
		traceErrStr:           traceErrStr,
	}
}

func (l *gormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Infof(s, args)
	fmt.Println(fmt.Sprintf(l.infoStr+s, append([]interface{}{utils.FileWithLineNum()}, args...)...))
}

func (l *gormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Warnf(s, args)
	fmt.Println(fmt.Sprintf(l.warnStr+s, append([]interface{}{utils.FileWithLineNum()}, args...)...))
}

func (l *gormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Errorf(s, args)
	fmt.Println(fmt.Sprintf(l.errStr+s, append([]interface{}{utils.FileWithLineNum()}, args...)...))

}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	fields := log.Fields{}
	traceID := ctx.Value("trace_id")
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[log.ErrorKey] = err
		log.WithContext(ctx).WithFields(fields).Errorf("[%s] [trace_id:%s] [rows:%v %s] [%s]", utils.FileWithLineNum(), traceID, rows, elapsed, sql)
		fmt.Println(fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, traceID, sql))

		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		log.WithContext(ctx).WithFields(fields).Warnf("[%s] [trace_id:%s] [rows:%v %s] [%s]", utils.FileWithLineNum(), traceID, rows, elapsed, sql)
		fmt.Println(fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, traceID, sql))

		return
	}

	log.WithContext(ctx).WithFields(fields).Debugf("[%s] [trace_id:%s] [rows:%v %s] [%s]", utils.FileWithLineNum(), traceID, rows, elapsed, sql)
	fmt.Println(fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, traceID, sql))
}
