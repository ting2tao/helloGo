package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	sliceA := [][]string{{"V0001ZHZ8", "坂田_SHZ001"},
		{"V0001ZH6C", "长阳半岛_BJ001"},
		{"V0001ZHX0", "龙城_SHZ002"},
		{"V0001ZHZ8", "坂田_SHZ001"},
		{"V0001ZHSK", "光谷南_WH003"},
		{"V0001ZI73", "东火_ZS001"},
		{"V0001ZI0C", "布吉_SHZ005"},
		{"V0001ZH8K", "长白_SY001"},
		{"V0001ZHX0", "龙城_SHZ002"},
		{"V0001ZHQC", "金银湖南_WH001"},
		{"V0001ZHD0", "秣陵北_NJ001"},
		{"V0001ZHP8", "双山_QD002"},
		{"V0001ZI88", "人和_CQ002"},
		{"V0001ZI2K", "桂城南_FS002"},
		{"V0001ZH7G", "中北镇_TJ001"},
		{"V0001ZHHG", "金鸡湖东_SUZ001"},
		{"V0001ZHF8", "太湖_WX001"},
		{"V0001ZHBW", "七宝镇_SH003"},
		{"V0001ZHQC", "金银湖南_WH001"},
		{"V0001ZHHG", "金鸡湖东_SUZ001"},
		{"V0001ZHVW", "金沙_GZ002"},
		{"V0001ZHVW", "金沙_GZ002"},
		{"V0001ZHY4", "香蜜湖_SHZ004"},
		{"V0001ZHSK", "光谷南_WH003"},
		{"V0001ZHLW", "新店镇_FZ001"},
		{"V0001ZI5Y", "石歧南_ZS002"},
		{"V0001ZHKS", "钟公庙_NB001"},
		{"V0001ZHIK", "良渚_HZ001"},
		{"V0001ZHE4", "兴隆_NJ002"},
		{"V0001ZI3O", "灯湖_FS001"},
		{"V0001ZHIK", "良渚_HZ001"},
		{"V0001ZHJO", "良东_HZ002"},
		{"V0008IQOF", "沟赵乡_41010001"},
		{"V0001ZHAS", "宜川路_SH001"},
		{"V0001ZH9O", "莘梅_SH002"},
		{"V0001ZI9D", "化龙桥_CQ001"},
		{"V0001ZI4T", "南城_DG001"},
		{"V0001ZHY4", "香蜜湖_SHZ004"},
		{"V0001ZIAI", "保和_CD001"},
		{"V0001ZHUS", "林和_GZ001"},
		{"V0001ZIBN", "成龙路_CD002"},
		{"V003JLRLK", "官陡_WUH001"},
		{"V0001ZHRG", "唐家墩_WH002"},
		{"V0001ZHN0", "鳌峰_FZ002"},
		{"V0001ZIBN", "成龙路_CD002"},
		{"V0001ZI1G", "西丽_SHZ003"},
		{"V0001ZHO4", "城阳_QD001"},
		{"V0001ZHTO", "黎托_CS001"},
		{"V0001ZHGC", "金鸡湖西_SUZ002"}}
	fileStore, _ := os.OpenFile("D:/workSpace/goworkspace/src/hello/try/store.txt", os.O_CREATE|os.O_RDWR, 0775)
	for _, cStr := range sliceA {
		sql := `update pride_store  set basic_attrs = jsonb_set (
  	basic_attrs,
 	'{associated_entity,@onewo_town}',
  	`
		sql += "'\"" + cStr[0] + `"') where basic_attrs#>>'{onewo_name}'='` + cStr[1] + `';`
		fmt.Println(sql)
		_, err := fileStore.Write([]byte(sql + "\n"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	time.Sleep(time.Second * 5)
}
