package cmd

import (
	"fmt"
	"git.vankeservice.com/deerse/pride/utils"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
	"time"

	_ "git.vankeservice.com/deerse/pride/domain"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var wwCmd = &cobra.Command{
	Use:   "ww",
	Short: "ww",
	Long:  `ww,刷入房屋字典`,
	Run: func(cmd *cobra.Command, args []string) {
		config := viper.New()
		config.AutomaticEnv()
		ctx := NewCmdAppContext(config)

		begin := time.Now()
		var success, fail int64
		defer func() {
			ctx.GetLogger().Info(
				ctx, zap.Duration("一共用时", time.Since(begin)),
				zap.Int64("成功数量", success), zap.Int64("失败数量", fail),
			)
		}()

		s := []string{
			"SHHXJ-2022-097",
			"WWY-SHHXY-2022-051",
			"VS-ZYSH-PY-2022-073",
			"VS-ZYSH-PY-2022-050",
			"SHHXJ-2022-062",
			"SHHXJ-2022-054",
			"SHHXJ-2022-022",
			"WWY-SHHXJ-2022-10",
			"WWY-SHHXJ-2022-11",
			"WWY- SHHXJ-2022-047",
			"WWYH-JZ-2022-SH-073",
			"WWY-SHHXJ-2022-096",
			"WWY-SHHXJ-2022-339",
			"WWYX-ZZ-2022-SH-345",
			"WWYX-JZ-2022-SH-529",
			"WWYX-JZ-2022-SH-446",
			"WWYX-JZ-2022-SH-700",
			"WWYX-ZZ-2022-SH-715",
			"WWYX-ZZ-2022-SH-739",
			"WWYX-ZZ-2022-SH-746",
			"WWYX-ZZ-2022-SH-741",
			"WWYX-ZZ-2022-SH-744",
			"WWYX-ZZ-2022-SH-742",
			"WWYX-ZZ-2022-SH-745",
			"WWYX-ZZ-2022-SH-747",
			"WWYX-ZZ-2022-SH-784",
			"WWYX-ZZ-2022-SH-748",
			"WWYX-ZZ-2022-SH-786",
			"WWYX-ZZ-2022-SH-738",
			"WWYX-ZZ-2022-SH-734",
			"WWYX-XS-2022-SH-652",
			"WWYX-XS-2022-SH-080",
			"WWYX-XS-2022-SH-190",
		}
		f, err := excelize.OpenFile("D:/workSpace/pride/cmd/ww.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			// Close the spreadsheet.
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		rows, err := f.GetRows("最新")
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStore, _ := os.OpenFile("D:/workSpace/pride/cmd/ww.txt", os.O_CREATE|os.O_RDWR, 0666)

		for _, row := range rows {
			conno := row[4]

			//fmt.Println(row, conno)
			if utils.Contains(s, conno) {
				success++
				o := strings.Split(row[8], "-")
				ft := fmt.Sprintf("20%s-%s-%s", o[2], o[0], o[1])
				sql := `update zeus_contract  set final_status=2,updated_source='230913-上海刷单2022年线下结转',pending_cost=0,pending_income=0,updated_time=NOW(),
                          confirmed_income= ` + row[6] + `,finish_time='` + ft + `' where final_status =1 and contract_no='` + conno + `';`
				fmt.Println(sql)
				_, err := fileStore.Write([]byte(sql + "\n"))
				if err != nil {
					fmt.Println(err.Error())
					return
				}
			}
		}

		time.Sleep(time.Second * 5)
	},
}
