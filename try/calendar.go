package main

import (
	"encoding/json"
	"fmt"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"math/rand"
	"strings"
	"time"
)

type LunarCalendar struct {
	Ganzhi struct {
		Animal     string `json:"animal"`
		Day        string `json:"day"`
		DayOrder   int    `json:"day_order"`
		Hour       string `json:"hour"`
		HourOrder  int    `json:"hour_order"`
		Month      string `json:"month"`
		MonthOrder int    `json:"month_order"`
		Year       string `json:"year"`
		YearOrder  int    `json:"year_order"`
	} `json:"ganzhi"`
	Lunar struct {
		Animal      string `json:"animal"`
		Day         int    `json:"day"`
		DayAlias    string `json:"day_alias"`
		IsLeap      bool   `json:"is_leap"`
		IsLeapMonth bool   `json:"is_leap_month"`
		LeapMonth   int    `json:"leap_month"`
		Month       int    `json:"month"`
		MonthAlias  string `json:"month_alias"`
		Year        int    `json:"year"`
		YearAlias   string `json:"year_alias"`
	} `json:"lunar"`
	Solar struct {
		Animal        string `json:"animal"`
		Constellation string `json:"constellation"`
		Day           int    `json:"day"`
		Hour          int    `json:"hour"`
		IsLeep        bool   `json:"is_leep"`
		Minute        int    `json:"minute"`
		Month         int    `json:"month"`
		Nanosecond    int    `json:"nanosecond"`
		Second        int    `json:"second"`
		WeekAlias     string `json:"week_alias"`
		WeekNumber    int    `json:"week_number"`
		Year          int    `json:"year"`
	} `json:"solar"`
}

// 宜
var Should = []string{
	"多喝水", "睡八个小时", "吃火锅", "好好听课", "吃点好的", "吃早餐", "释放压力", "静下心来", "上班摸鱼", "吃超辣辣条",
	"穿衬衫", "坐地铁", "随身带笔记本", "吃泡面", "换手机密码", "休假", "喝肥宅快乐水", "点外卖", "给自己打气", "追剧",
	"勇敢告白", "跟爱的人亲亲", "手机开门", "跟管家聊天", "打扫卫生", "种一盆花", "出门带伞", "开车兜风", "吃点零食",
	"好好准备考试", "打听八卦", "喝杯奶茶", "给爸妈打电话", "给老板画饼", "算了算了", "吃玉米", "提前完成工作", "哈哈大笑",
	"健身房撸铁", "送礼物", "爬山", "旅行", "和朋友开黑", "穿两件", "晒太阳", "买新包包", "做高数", "六点起床", "白日做梦",
	"情绪稳定", "大白兔奶糖", "喷香水", "逛菜市场", "自言自语", "保持快乐", "戴好口罩", "酸菜鱼", "背书包", "乐于助人",
	"诚实守信", "做广播体操", "洗热水澡", "不要生气", "远离讨厌的人", "努力搬砖", "发朋友圈", "看电影", "发誓减肥", "发红包",
	"请喝奶茶", "手机充电", "约tony剪头", "背锅", "撸猫", "和爱人共进晚餐", "早点回家", "打个喷嚏", "随身带面巾纸", "换手机壳",
	"一展歌喉", "吃点坚果", "来杯咖啡", "出门检查", "多去wc", "和甲方battle", "亲亲妈妈", "泡个脚", "对自己好一点",
	"吃辣火锅", "换新电脑", "骑自行车", "去公园", "放风筝", "和小朋友在一起", "耍帅", "戴着耳机", "抬头看天", "穿帆布鞋",
	"塞满冰箱", "照镜子欣赏自己",
}

// 忌
var Avoid = []string{
	"吹牛皮儿", "打架斗殴", "违法犯罪", "偷工减料", "顶撞上司", "不喝水", "憋尿", "不带充电器", "不讲道理", "瞧不起人",
	"穿拖鞋上班", "裸考", "地铁里打啵", "乱丢垃圾", "通宵熬夜", "冲动消费", "忘带手机", "打麻将", "蹭吃蹭喝", "头疼",
	"皮笑肉不笑", "买ipad", "骗人", "没有时间观念", "炫富", "抬杠", "有话不好好说", "忘带钥匙", "打碎杯子",
	"沉迷于自己的美貌无法自拔", "画大饼", "跟甲方吵架", "骂小孩", "好吃懒做", "没心没肺", "不洗澡", "绊倒别人",
	"戴金链子", "生闷气", "不写作业", "拖拖拉拉", "借钱不还", "贴小广告", "吃香蕉", "插队", "吃完饭不洗碗", "拈花惹草",
	"写完文件不保存", "骂同事", "上厕所不冲水", "只吃一顿饭", "不出门", "熬夜加班", "上班迟到", "烫伤", "暴饮暴食",
	"开车不礼让行人", "闯红灯", "剃光头", "摔手机", "朋友圈点赞", "不穿袜子", "喝得烂醉", "自我否定", "得过且过",
	"P图只P自己", "丢失身份证", "偷东西", "哇哇大哭", "走路玩手机", "被蚊子咬", "放别人鸽子", "弄脏白衣服", "不倒垃圾",
	"担心考试", "衣服连穿三天", "觉得自己天下第一", "喉咙痛", "临时抱佛脚", "坐第一排上课", "偷看美女", "质疑", "内卷",
	"垮个批脸", "不离开被窝", "排位赛", "装腔作势", "抢功劳", "拍领导马屁", "玩物丧志", "阴阳怪气", "双标", "伤别人的心",
	"背后说坏话", "卖惨", "独吞", "唉声叹气", "小气吧啦", "占小便宜", "破口大骂",
}

// 幸运物
var Goods = []string{
	"保温杯", "雨伞", "书包", "围巾", "矿泉水", "眼镜", "耳机", "电脑", "戒指", "袜子", "充电器", "游戏机", "鞋带", "纸袋子", "充电宝", "口罩", "自行车", "篮球", "镜子", "书", "茶叶", "铅笔", "皮筋", "冰箱贴", "面巾纸", "拖鞋", "橙子", "鼠标", "发夹", "牙刷", "签字笔", "马克笔", "耳塞", "车钥匙", "工牌", "鼠标垫", "地铁卡", "单肩包", "护手霜", "围巾", "盲盒", "粉底液", "香蕉", "苹果", "大头贴", "拍立得", "身份证", "糖", "茶叶蛋", "发票", "湿纸巾", "跳绳", "棉签", "唇膏", "毛绒公仔", "手镯", "项链", "金链子", "大皮鞋", "小金表", "耳环", "帽子", "脑子", "知识", "小风扇", "充电线", "牙线", "漱口水", "现金", "遛狗绳", "气质", "手套", "饼干", "运动服", "手机", "卡包", "瓜子", "kindle", "胸针", "花", "牛奶", "购物袋", "橡皮擦", "口罩", "外套", "烧卖", "眼药水", "便签纸", "手机壳", "笔记本", "防晒霜", "口香糖", "眼镜盒", "手表", "钱包", "酸奶", "相机", "录音笔", "坚果", "驾照",
}

// 幸运词
var LuckyWords = []string{
	"烦死了", "算了算了", "中国人不骗中国人", "好吧好吧", "行", "搞不懂", "好啊", "阿弥陀佛", "别别别", "不愧是我", "哈哈哈",
	"起飞", "打工人", "饶了我吧", "我的天呐", "没问题", "路过不要错过", "笑死", "你是真的狗", "打疫苗", "拿来吧你", "针不戳",
	"内卷", "格局打开", "你懂不懂", "yyds", "绝绝子", "躺平", "伤害性不高，侮辱性极强", "甩锅", "996", "我太难了", "扎心了",
	"可以啊", "真好看", "行吧", "老凡尔赛了", "佛了佛了", "请开始你的表演", "什么鬼", "懒得理你", "吃啥", "约吗", "受不了了",
	"这天儿真好", "睡了睡了", "拜拜", "早上好", "别呀", "别抬杠", "后浪", "哦", "休息一会", "你好", "请问厕所在哪", "来一份",
	"干饭人", "防不胜防", "加油加油", "胜利就在眼前", "别忘了", "搞钱", "回家", "想妈妈", "别吹了", "走", "成功人士",
	"隐形富豪", "有房有车", "发家致富", "疫情结束", "戴口罩", "是不是", "提个建议", "小心点", "惊呆了", "无语", "醉了",
	"又来", "奖金", "就今天", "厉害了", "你怎么看", "饱了", "不吃了", "怎么回事", "笑死我了", "你在搞笑吗", "都这样",
	"差不多得了", "这怎么办", "我陪你", "你猜", "没想到吧", "哈哈哈", "ok", "居然是真的", "服了", "我就不", "管不着",
}

// GetPeriodTime
// 获取区间段的时间点
func GetPeriodTime(now time.Time, hour int, span int) (time.Time, time.Time) {
	end := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local)
	start := end.AddDate(0, 0, -span)
	return start, end
}

func main() {

	//a := 11459235 % 5
	//fmt.Println(a)
	////now2 := time.Now()
	////fmt.Println(time.Date(now2.Year(), now2.Month(), now2.Day(), 21, 0, 0, 0, time.Local).Unix())
	////start, end := GetPeriodTime(now2, 21, 1)
	////fmt.Println(start, end)
	//FillJi()
	now := time.Now()
	for i := 0; i < 30000; i++ {
		fmt.Println(Combine(i))
		//fmt.Println(GetPushContent(i))
	}
	fmt.Println(time.Since(now))
}

func GetPushContent(i int) string {
	t := time.Now().AddDate(0, 0, 2)
	// 时间戳
	c := calendar.ByTimestamp(t.Unix())
	bytes, _ := c.ToJSON()
	var lunar LunarCalendar
	json.Unmarshal(bytes, &lunar)
	str := fmt.Sprintf("明日是%s，星期%s，农历%s%s。\n", t.Format("2006年01月02日"), lunar.Solar.WeekAlias, lunar.Lunar.MonthAlias, lunar.Lunar.DayAlias)
	yi := strings.Join(RandSlice(Should, 5), "，")
	//fmt.Println(yi)
	ji := strings.Join(RandSlice(Avoid, 5), "，")
	//fmt.Println(ji)
	goods := strings.Join(RandSlice(Goods, 1), "，")
	//fmt.Println(goods)
	lucky := strings.Join(RandSlice(LuckyWords, 1), "，")

	str += fmt.Sprintf("可爱的您，明天%s；忌%s。\n明日幸运随身物：%s\n明日幸运词：%s", yi, ji, goods, lucky)

	return str
}

func Combine(i int) string {
	now := time.Now()
	//str :=
	//yi := strings.Join(GetRandomArr(Yi, 5, i), "，")
	////fmt.Println(yi)
	//ji := strings.Join(GetRandomArr(Ji, 5, i), "，")
	//fmt.Println(ji)
	goods := strings.Join(GetRandomArr(Goods, 1, int64(i)), "，")
	//fmt.Println(goods)
	lucky := strings.Join(GetRandomArr(LuckyWords, 1, int64(i)), "，")
	//fmt.Println(lucky)
	fmt.Println(time.Since(now))

	str := fmt.Sprintf("可爱的您，明天%s；忌%s。\n明日幸运随身物：%s\n明日幸运词：%s", "yi", "ji", goods, lucky)
	return str
}
func GetRandomArr(arr []string, needLen int, userID int64) []string {
	lenArr := len(arr)
	if lenArr <= needLen {
		return arr
	}
	var temp []string
	r := rand.New(rand.NewSource(time.Now().UnixNano() + userID))
	for {
		//hash.Hash64().

		randNum := r.Intn(lenArr)

		temp = append(temp, arr[randNum])
		if len(temp) == needLen {
			return temp
		}
	}
}

func RandSlice(origin []string, count int) []string {
	tmpOrigin := make([]string, len(origin))
	copy(tmpOrigin, origin)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tmpOrigin), func(i int, j int) {
		tmpOrigin[i], tmpOrigin[j] = tmpOrigin[j], tmpOrigin[i]
	})
	//result := make([]string, 0, count)
	//for index, value := range tmpOrigin {
	//	if index == count {
	//		break
	//	}
	//	result = append(result, value)
	//}
	return tmpOrigin[:count]
}
