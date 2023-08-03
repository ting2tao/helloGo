package main

import (
	"fmt"
	"git.vankeservice.com/deerse/pride/adapters"
	"git.vankeservice.com/deerse/pride/application/dtos"
	"git.vankeservice.com/deerse/pride/domain"
	"git.vankeservice.com/deerse/pride/utils"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var syncZhantuCmd = &cobra.Command{
	Use:   "sync_zhantu",
	Short: `1`,
	Long:  `1`,
	Run: func(cmd *cobra.Command, args []string) {

		config := viper.New()
		config.AutomaticEnv()
		ctx := NewCmdAppContext(config)

		dbConn, dbErr := sqlx.Connect("postgres", config.GetString("ZAHNTU_URI"))
		if dbErr != nil {
			log.Fatalf("无法连接数据库: %s", dbErr.Error())
		}
		defer dbConn.Close()
		//var data []*adapters.EntityDAO
		//sql := `SELECT * FROM pride_project WHERE code IN ('CN33020001007X')`
		sql := `SELECT * FROM pride_project WHERE code IN ('CN2101000100X9','CN2101000100Y7','CN2101000100P3','CN2101000101B5','CN210100010148','CN2101000100T6','CN210100010121','CN210100010105','CN2101000101D1','CN21010001013X','CN210100010156','CN2101000101C3','CN2101000100V2','CN2101000101G6','CN440300010074','CN130100010068','CN130200010073','CN6101000100A3','CN321000010046','CN64010001005X','CN3706000100AX','CN640100010068','CN4101000100F8','CN5101000100G4','CN2102000100E0','CN2102000100L8','CN2102000100F9','CN2102000100N4','CN2102000100M6','CN440100010061','CN210300010038','CN5101000100F6','CN5101000100F6','CN5101000100I0','CN5101000100G4','CN6101000100A3','CN2102000100Q9','CN210200010102','CN4406000100GX','CN4406000100E3','CN2101000101F8','CN520100010029','CN441300010079','CN210100010172','CN2101000100L0','CN2101000101H4','CN2101000100S8','CN2101000100N7','CN210100010113','CN2101000101H4','CN210100010199','CN2101000101A7','CN2101000100M9','CN2101000100O5','CN2101000100U4','CN2101000100Q1','CN210100010092','CN210100010164','CN320200010089','CN320200010097','CN320200010097','CN370600010059','CN640100010041','CN1101200100B8','CN5101000100E8','CN2102000100Z2','CN2102000100R7','CN2102000100T3','CN2102000100G7','CN2102000100C4','CN2102000100I3','CN2102000100O2','CN4406000100D5','CN210400010019','CN210100010113','CN330100010C87','CN441300010052' );`
		rows, _ := dbConn.QueryxContext(ctx, sql)
		var e []*domain.Entity
		for rows.Next() {
			dao := &adapters.EntityDAO{}
			if err := rows.StructScan(dao); err != nil {
				return
			}
			entity, err := dao.ToDomain()
			if err != nil {
				fmt.Println(err.Error())
			}
			e = append(e, entity)
		}

		var d []*dtos.ProjectDTO
		for _, datum := range e {
			d1 := &dtos.ProjectDTO{}
			err := d1.Scan(datum)
			if err != nil {
				return
			}
			d = append(d, d1)
		}
		for _, datum := range d {
			tmpParam := &PullZhantuParam{Body: B{Code: datum.Code, CreatedSource: datum.CreatedSource,
				CodeMapping: C{Yxs: datum.CodeMapping.Yxs}}}
			_, body, errs := utils.Post("https://weijia-x.4009515151.com/hd/v2/mgt/mq/consumer/project").Send(&tmpParam).End()
			if errs != nil {
				fmt.Println("errs" + errs[0].Error())
				return
			}
			fmt.Println("body" + body)
		}
		return
	},
}

type PullZhantuParam struct {
	Body B `json:"body"`
}

type B struct {
	Code          string `json:"code"`
	CreatedSource string `json:"created_source"`
	CodeMapping   C      `json:"code_mapping"`
}

type C struct {
	Yxs string `json:"@yxs"`
}
