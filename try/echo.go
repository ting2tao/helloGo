package main

import (
	"fmt"
	"strconv"
)

func main() {
	echoJson()
}

type CityShort struct {
	DataName  string `json:"data_name"`
	DataValue string `json:"data_value"`
}

func echoJson() {
	var CityShortWordMap = map[string]string{
		"成都":  "CD",
		"昆明":  "KM",
		"贵阳":  "GY",
		"重庆":  "CQ",
		"佛山":  "FS",
		"武汉":  "WH",
		"郑州":  "ZZ",
		"西安":  "XA",
		"长春":  "CC",
		"哈尔滨": "HB",
		"沈阳":  "SY",
		"大连":  "DL",
		"广州":  "GZ",
		"长沙":  "CS",
		"苏州":  "SU",
		"无锡":  "WX",
		"泉州":  "QZ",
		"厦门":  "XM",
		"福州":  "FZ",
		"上海":  "SH",
		"青岛":  "QD",
		"济南":  "JN",
		"烟台":  "YT",
		"南京":  "NJ",
		"合肥":  "HF",
		"深圳":  "SZ",
		"海南":  "HN",
		"东莞":  "DG",
		"杭州":  "HZ",
		"南昌":  "NC",
		"宁波":  "NB",
		"北京":  "BJ",
		"天津":  "TJ",
		"太原":  "TY"}
	var jcs []CityShort
	i := 6
	for kcs, v := range CityShortWordMap {
		jcs = append(jcs, CityShort{
			DataName:  kcs,
			DataValue: v,
		})
		i++
		sql := `INSERT INTO "public"."pride_dictionary"("code", "category", "name", "pinyin", "version", "created_at", "created_source", "updated_at", "updated_source", "status", "basic_attrs", "extend_attrs", "code_mapping", "tags", "geometry") 
VALUES (`
		sql += strconv.Itoa(i)
		sql += `, 'dictionary', '店铺城市缩写', 'dianpuchengshisuoxie', 1, 1668765146445, '房屋字典后台', 
        1668765146445, '房屋字典后台', 'ACTIVE', '{"data_name":"`
		sql += kcs
		sql += `", "data_value": "`
		sql += v
		sql += `"}', '{}', '{}', '{}', NULL);`
		fmt.Println(sql)
	}
	//marshal, err := json.Marshal(jcs)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(marshal))
}
