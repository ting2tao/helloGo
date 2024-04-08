package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Logger = logrus.Logger
type Hook = logrus.Hook
type Entry = logrus.Entry

// Define key
const (
	TraceIDKey = "trace_id"
	UserIDKey  = "user_id"
	TagKey     = "tag"
	VersionKey = "version"
	StackKey   = "stack"
)

var version string

// StandardLogger 获取标准日志
func StandardLogger() *Logger {
	return logrus.StandardLogger()
}

// SetLevel 设定日志级别
func SetLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}

// SetFormatter 设定日志输出格式
func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}

func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

func SetVersion(v string) {
	version = v
}

func AddHook(hook Hook) {
	logrus.AddHook(hook)
}

type (
	traceIDKey struct{}
	userIDKey  struct{}
	tagKey     struct{}
	stackKey   struct{}
)

func Check(err error, backLevel ...int) bool {
	level := 1
	if len(backLevel) != 0 {
		level = backLevel[0]
	}
	if err != nil {
		_, filename, line, _ := runtime.Caller(level)
		_, name := filepath.Split(filename)
		fmt.Printf("%s \033[31m[E] [%s:%d]\033[32m\033[0m %s\n", time.Now().Format(time.RFC3339), name, line, err.Error())
		logrus.Errorf("[%s:%d] %s", name, line, err.Error())
		return false
	}
	return true
}

func CheckWithContext(ctx context.Context, err error, backLevel ...int) bool {
	level := 1
	if len(backLevel) != 0 {
		level = backLevel[0]
	}
	if err != nil {
		_, filename, line, _ := runtime.Caller(level)
		_, name := filepath.Split(filename)
		traceID, ok := ctx.Value("trace_id").(string)
		if !ok {
			traceID = ""
		}
		fmt.Printf("%s \033[31m[C] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, err.Error())
		logrus.Errorf("[%s:%d] [trace_id:%s] %s", name, line, traceID, err.Error())
		return false
	}
	return true
}

func WarnCheck(err error, backLevel ...int) bool {
	level := 1
	if len(backLevel) != 0 {
		level = backLevel[0]
	}
	if err != nil {
		_, filename, line, _ := runtime.Caller(level)
		_, name := filepath.Split(filename)
		fmt.Printf("%s \033[31m[E] [%s:%d]\033[32m\033[0m %s\n", time.Now().Format(time.RFC3339), name, line, err.Error())
		logrus.Warnf("[%s:%d] %s", name, line, err.Error())
		return false
	}
	return true
}

func Debug(values ...interface{}) {
	var datas []string
	for _, value := range values {
		datas = append(datas, fmt.Sprintf("%#v", value))
	}
	_, filename, line, _ := runtime.Caller(1)
	_, name := filepath.Split(filename)

	fmt.Printf("%s \033[34m[D] [%s:%d]\033[32m\033[0m %s\n", time.Now().Format(time.RFC3339), name, line, strings.Join(datas, " "))
	logrus.Debugf("[%s:%d] %s", name, line, strings.Join(datas, " "))
}

func DebugWithCtx(ctx context.Context, values ...interface{}) {
	var datas []string
	for _, value := range values {
		datas = append(datas, fmt.Sprintf("%#v", value))
	}
	_, filename, line, _ := runtime.Caller(1)
	_, name := filepath.Split(filename)

	fmt.Printf("%s \033[34m[D] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, ctx.Value("trace_id").(string), strings.Join(datas, " "))
	logrus.Debugf("[%s:%d] [trace_id:%s] %s", name, line, ctx.Value("trace_id").(string), strings.Join(datas, " "))
}

//修改颜色
func Warn(values ...interface{}) {
	var datas []string
	for _, value := range values {
		datas = append(datas, fmt.Sprintf("%#v", value))
	}
	_, filename, line, _ := runtime.Caller(1)
	_, name := filepath.Split(filename)

	fmt.Printf("%s \033[34m[D] [%s:%d]\033[32m\033[0m %s\n", time.Now().Format(time.RFC3339), name, line, strings.Join(datas, " "))
	logrus.Debugf("[%s:%d] %s", name, line, strings.Join(datas, " "))
}

// NewTraceIDContext 创建跟踪ID上下文
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// FromTraceIDContext 从上下文中获取跟踪ID
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value("trace_id")
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewUserIDContext 创建用户ID上下文
func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// FromUserIDContext 从上下文中获取用户ID
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value("user_id")
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewTagContext 创建Tag上下文
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

// FromTagContext 从上下文中获取Tag
func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewStackContext 创建Stack上下文
func NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

// FromStackContext 从上下文中获取Stack
func FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

// WithContext Use context create entry
func WithContext(ctx context.Context) *Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := map[string]interface{}{
		VersionKey: version,
	}

	if v := FromTraceIDContext(ctx); v != "" {
		fields[TraceIDKey] = v
	}

	if v := FromUserIDContext(ctx); v != "" {
		fields[UserIDKey] = v
	}

	if v := FromTagContext(ctx); v != "" {
		fields[TagKey] = v
	}

	if v := FromStackContext(ctx); v != nil {
		fields[StackKey] = fmt.Sprintf("%+v", v)
	}

	return logrus.WithContext(ctx).WithFields(fields)
}

// InitLogger 初始化日志模块
func InitLogger(filename string) (func(), error) {
	SetLevel(5)
	SetFormatter("json")

	// 设定日志输出
	var file *os.File
	var logFullName = filepath.Join("./storage/logs", filename)
	_ = os.MkdirAll(filepath.Dir(logFullName), 0777)

	f, err := os.OpenFile(logFullName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	//writers := []io.Writer{f, os.Stdout}
	//fileAndStdoutWriter := io.MultiWriter(writers...)
	//SetOutput(fileAndStdoutWriter)
	SetOutput(f)

	file = f

	//logger.SetOutput(os.Stdout)
	return func() {
		if file != nil {
			//defer file.Close()
		}
	}, nil
}

// Define logrus alias
var (
	Tracef  = logrus.Tracef
	Debugf  = logrus.Debugf
	Info    = logrus.Info
	Infof   = logrus.Infof
	Errorf  = logrus.Errorf
	Fatalf  = logrus.Fatalf
	Fatalln = logrus.Fatalln
	Panicf  = logrus.Panicf
	Printf  = logrus.Printf
	Errorln = logrus.Errorln
)

func InitTLogger() *TLogger {
	return &TLogger{}
}

type TLogger struct {
}

func (cl *TLogger) Check(ctx context.Context, err error, backLevel ...int) bool {
	level := 1
	if len(backLevel) != 0 {
		level = backLevel[0]
	}
	if err != nil {
		_, filename, line, _ := runtime.Caller(level)
		_, name := filepath.Split(filename)
		traceID, ok := ctx.Value("trace_id").(string)
		if !ok {
			traceID = ""
		}
		fmt.Printf("%s \033[31m[C] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, err.Error())
		logrus.Errorf("[%s:%d] [trace_id:%s] %s", name, line, traceID, err.Error())
		return false
	}
	return true
}

func (cl *TLogger) TCheck(ctx context.Context, err error, backLevel ...int) bool {
	level := 1
	if len(backLevel) != 0 {
		level = backLevel[0]
	}
	if err != nil {
		_, filename, line, _ := runtime.Caller(level)
		_, name := filepath.Split(filename)
		traceID, ok := ctx.Value("trace_id").(string)
		if !ok {
			traceID = ""
		}
		fmt.Printf("%s \033[31m[C] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, err.Error())
		logrus.Errorf("[%s:%d] [trace_id:%s] %s", name, line, traceID, err.Error())
		return false
	}
	return true
}

func (cl *TLogger) WarnCheck(ctx context.Context, err error, backLevel ...int) bool {
	level := 1
	if len(backLevel) != 0 {
		level = backLevel[0]
	}
	if err != nil {
		_, filename, line, _ := runtime.Caller(level)
		_, name := filepath.Split(filename)
		traceID, ok := ctx.Value("trace_id").(string)
		if !ok {
			traceID = ""
		}
		fmt.Printf("%s \033[31m[W] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, err.Error())
		logrus.Warnf("[%s:%d] [trace_id:%s] %s", name, line, traceID, err.Error())
		return false
	}
	return true
}

func (cl *TLogger) Debug(ctx context.Context, values ...interface{}) {
	var datas []string
	for _, value := range values {
		datas = append(datas, fmt.Sprintf("%#v", value))
	}
	_, filename, line, _ := runtime.Caller(1)
	_, name := filepath.Split(filename)
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = ""
	}
	fmt.Printf("%s \033[34m[D] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, strings.Join(datas, " "))
	logrus.Debugf("[%s:%d] [trace_id:%s] %s", name, line, traceID, strings.Join(datas, " "))
}

func (cl *TLogger) TDebug(ctx context.Context, values ...interface{}) {
	var datas []string
	for _, value := range values {
		datas = append(datas, fmt.Sprintf("%#v", value))
	}
	_, filename, line, _ := runtime.Caller(1)
	_, name := filepath.Split(filename)
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = ""
	}
	fmt.Printf("%s \033[34m[D] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, strings.Join(datas, " "))
	logrus.Debugf("[%s:%d] [trace_id:%s] %s", name, line, traceID, strings.Join(datas, " "))
}

//修改颜色
func (cl *TLogger) Warn(ctx context.Context, values ...interface{}) {
	var datas []string
	for _, value := range values {
		datas = append(datas, fmt.Sprintf("%#v", value))
	}
	_, filename, line, _ := runtime.Caller(1)
	_, name := filepath.Split(filename)
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = ""
	}
	fmt.Printf("%s \033[34m[D] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, strings.Join(datas, " "))
	logrus.Debugf("[%s:%d] [trace_id:%s] %s", name, line, traceID, strings.Join(datas, " "))
}

//需要出发告警
func (cl *TLogger) TErrorAlarm(ctx context.Context, values ...interface{}) {
	var datas []string
	for _, value := range values {
		datas = append(datas, fmt.Sprintf("%#v", value))
	}
	_, filename, line, _ := runtime.Caller(1)
	_, name := filepath.Split(filename)
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = ""
	}
	fmt.Printf("%s \033[34m[D] [%s:%d]\033[32m\033[0m [trace_id:%s]\u001B[32m\u001B[0m %s\n", time.Now().Format(time.RFC3339), name, line, traceID, strings.Join(datas, " "))
	logrus.Debugf("[%s:%d] [trace_id:%s] %s", name, line, traceID, strings.Join(datas, " "))
}
