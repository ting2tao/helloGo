package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type IndexCtl struct {
}

// 首页数据
type IndexParam struct {
	Types      string `json:"types,omitempty" form:"types" `
	SourceType string `json:"source_type,omitempty" form:"source_type" `
	Limit      uint   `json:"limit,omitempty" form:"limit"`
	Offset     uint   `json:"offset,omitempty" form:"offset"`
	SortBy     string `json:"sort_by,omitempty"`
	IsSelling  int    `json:"is_selling,omitempty"`
	ColumnID   int    `json:"column_id,omitempty" form:"column_id" `

	Enable int `json:"enable" form:"enable"`

	// 打卡營参数
	IsNeedUserTodayStats     int `json:"is_need_user_today_stats,omitempty" form:"is_need_user_today_stats"`
	IsNeedCampClassTypeStats int `json:"is_need_camp_class_type_stats,omitempty" form:"is_need_camp_class_type_stats"`
	IsNeedCampTeacher        int `json:"is_need_camp_teacher,omitempty" form:"is_need_camp_teacher"`
	IsExcludeCourseMode      int `json:"is_exclude_course_mode,omitempty" form:":"is_exclude_course_mode"`
	IsNoviceCamp             int `json:"is_novice_camp,omitempty"`
	IsNeedCampStats          int `json:"is_need_camp_stats,omitempty"`

	IsShow          int    `json:"is_show,omitempty" form:"is_show"`
	SortByStartTime string `json:"sort_by_start_time"`
}

func main() {
	var ctx *gin.Context
	var ctl IndexCtl
	ctl.IndexData(ctx)
}
func (ctl IndexCtl) IndexData(ctx *gin.Context) {
	var lp IndexParam

	lp.Enable = 1
	var (
		liveList         []LiveDataList
		studyRoomInfo    StudyRoomInfo
		publicCourse     PublicCourseInfo
		campDataList     []CampDataList
		examEncyclopedia []ExamEncyclopediaInfo
		secondsDataList  []SecondsDataList
		hotNews          HotNews
		publicCourseList []PublicCourseList
		search           Search
		moduleList       []string
	)
	ch := make(chan map[string]interface{}, 7)

	go func(ctx *gin.Context, ch *chan map[string]interface{}) {
		ctl.liveList(ctx, lp)
		*ch <- gin.H{"key": "liveList", "value": liveList}
	}(ctx, &ch)

	go func(ctx *gin.Context, ch *chan map[string]interface{}) {
		ctl.studyRoom(ctx, lp)
		*ch <- gin.H{"key": "studyRoomInfo", "value": studyRoomInfo}
	}(ctx, &ch)

	go func(ctx *gin.Context, ch *chan map[string]interface{}) {
		ctl.camp(ctx, lp)
		*ch <- gin.H{"key": "campDataList", "value": campDataList}
	}(ctx, &ch)

	go func(ctx *gin.Context, ch *chan map[string]interface{}) {
		ctl.examEncyclopedia(ctx, lp)
		*ch <- gin.H{"key": "examEncyclopedia", "value": examEncyclopedia}
	}(ctx, &ch)

	go func(ctx *gin.Context, ch *chan map[string]interface{}) {
		ctl.secondsDataList(ctx, lp)
		*ch <- gin.H{"key": "secondsDataList", "value": secondsDataList}
	}(ctx, &ch)

	go func(ctx *gin.Context, ch *chan map[string]interface{}) {
		ctl.hotNews(ctx, lp)
		*ch <- gin.H{"key": "hotNews", "value": hotNews}
	}(ctx, &ch)

	go func(ctx *gin.Context, ch *chan map[string]interface{}) {
		ctl.publicCourseList(ctx, &lp)
		*ch <- gin.H{"key": "publicCourseList", "value": gin.H{"publicCourse": publicCourse, "publicCourseList": publicCourseList}}
	}(ctx, &ch)

	for i := 0; i < 7; i++ {
		res := <-ch
		switch res["key"] {
		case "liveList":
			liveList = res["value"].([]LiveDataList)
			break
		case "studyRoomInfo":
			studyRoomInfo = res["value"].(StudyRoomInfo)
			break
		case "campDataList":
			campDataList = res["value"].([]CampDataList)
			break
		case "examEncyclopedia":
			examEncyclopedia = res["value"].([]ExamEncyclopediaInfo)
			break
		case "secondsDataList":
			secondsDataList = res["value"].([]SecondsDataList)
			break
		case "hotNews":
			hotNews = res["value"].(HotNews)
			break
		case "publicCourseList":
			publicCourse = res["value"].(gin.H)["publicCourse"].(PublicCourseInfo)
			publicCourseList = res["value"].(gin.H)["publicCourseList"].([]PublicCourseList)
			break

		}

	}

	mp := map[string]interface{}{
		"live_list":          liveList,
		"study_room":         studyRoomInfo,
		"public_course":      publicCourse,
		"camp_data_list":     campDataList,
		"exam_encyclopedia":  examEncyclopedia,
		"seconds_data_list":  secondsDataList,
		"hot_news":           hotNews,
		"public_course_list": publicCourseList,
		"search_data":        search,
		"module_list":        moduleList,
	}

	fmt.Println(mp)
}

// 直播数据

type LiveDataList struct {
	CoverImage     string `json:"cover_image"`
	ID             int64  `json:"id"`
	IsReservation  int64  `json:"is_reservation"`
	LiveName       string `json:"live_name"`
	LiveStartTime  string `json:"live_start_time"`
	LiveEndTime    string `json:"live_end_time"`
	LiveStatus     int64  `json:"live_status"` //1-等待 2-直播中 3-已结束
	Type           string `json:"type"`        //Type    comment:直播类型:live 直播，liveRecord回放"`
	LiveID         int64  `json:"live_id"`
	VID            int64  `json:"vid"`
	PlayCount      int64  `json:"play_count"`
	ReservationNum int64  `json:"reservation_num"`
	StartDate      string `json:"start_date"`
	TeacherName    string `json:"teacher_name"`
	TeacherAvatar  string `json:"teacher_avatar"`
	BackImage      string `json:"back_image"`
}

func (ctl IndexCtl) liveList(ctx *gin.Context, lp IndexParam) ([]LiveDataList, []error) {

	liveList := []LiveDataList{}
	time.Sleep(time.Second * 2)
	fmt.Println("live-1")
	return liveList, nil
}

// 自犀室
type StudyRoomInfo struct {
	RoomTitle  string `json:"room_title"`
	StudyNum   int64  `json:"study_num"`
	LiveStatus string `json:"live_status"` //live-直播中，playback-回放中，end-已结束，waiting-未开始
}

func (ctl IndexCtl) studyRoom(ctx *gin.Context, lp IndexParam) (StudyRoomInfo, []error) {

	var data StudyRoomInfo
	time.Sleep(time.Second * 2)
	fmt.Println("StudyRoomInfo-2")
	return data, nil
}

// 打卡营
type CampDataList struct {
	ClassTypeID       int64    `json:"class_type_id"`
	IsBuy             int      `json:"is_buy"`
	CourseID          int64    `json:"course_id"`
	CountDown         int64    `json:"count_down"`
	CourseCoverImage  string   `json:"course_cover_image"`
	GuideCopy         string   `json:"guide_copy"`
	ID                int64    `json:"id"`
	Intro             string   `json:"intro"`
	JoinAvatar        []string `json:"join_avatar"`
	JoinNum           int64    `json:"join_num"`
	ModelType         string   `json:"model_type"`
	CampName          string   `json:"camp_name"`
	CampClassName     string   `json:"camp_class_name"`
	PlayRule          string   `json:"play_rule"`
	PlayRuleText      string   `json:"play_rule_text"`
	Prices            int64    `json:"prices"`
	PrimaryCoverImage string   `json:"primary_cover_image"`
	Subtitle          string   `json:"subtitle"`
	TeacherID         int64    `json:"teacher_id"`
	TeacherName       string   `json:"teacher_name"`
	TeacherNameArr    []string `json:"teacher_name_arr"`
	UIType            int      `json:"ui_type"`
	Sort              int64    `json:"sort"`
	LabelName         string   `json:"label_name"`
}

func (ctl IndexCtl) camp(ctx *gin.Context, lp IndexParam) ([]CampDataList, []error) {
	var dataList []CampDataList
	time.Sleep(time.Second * 2)
	fmt.Println("CampDataList-3")
	return dataList, nil
}

// 公考百科
type ExamEncyclopediaInfo struct {
	ID    int64  `json:"id"`
	Tag   string `json:"tag"`
	Title string `json:"title"`
}

func (ctl IndexCtl) examEncyclopedia(ctx *gin.Context, lp IndexParam) ([]ExamEncyclopediaInfo, []error) {

	var data []ExamEncyclopediaInfo
	time.Sleep(time.Second * 2)
	fmt.Println("ExamEncyclopediaInfo-4")
	return data, nil
}

// 秒懂课堂
type SecondsDataList struct {
	CoverImage    string `json:"cover_image"`
	ID            int64  `json:"id"`
	LabelID       int64  `json:"label_id"`
	LabelName     string `json:"label_name"`
	Name          string `json:"name"`
	TeacherName   string `json:"teacher_name"`
	VideoDuration int64  `json:"video_duration"`
	VideoURL      string `json:"video_url"`
}

func (ctl IndexCtl) secondsDataList(ctx *gin.Context, lp IndexParam) ([]SecondsDataList, []error) {
	var dataList []SecondsDataList
	time.Sleep(time.Second * 2)
	fmt.Println("SecondsDataList-5")
	return dataList, nil
}

// 热点早餐
type HotList struct {
	AudioTime int64  `json:"audio_time"`
	AudioURL  string `json:"audio_url"`
	ID        int64  `json:"id"`
	IsPlay    int64  `json:"is_play"`
	Title     string `json:"title"`
}
type HotNews struct {
	HotList    []HotList `json:"hot_list"`
	NewestInfo struct {
		ID          int64  `json:"id"`
		NewestImage string `json:"newest_image"`
		NewestNo    int64  `json:"newest_no"`
		NewestURL   string `json:"newest_url"`
	} `json:"newest_info"`
}

type Search struct {
	Recommend string `json:"recommend"`
}

func (ctl IndexCtl) hotNews(ctx *gin.Context, lp IndexParam) (HotNews, []error) {
	var dataList HotNews
	time.Sleep(time.Second * 3)
	fmt.Println("HotNews-6")
	return dataList, nil
}

// 缓存 struct
type CachePublicCourse struct {
	PublicCourseInfo PublicCourseInfo   `json:"public_course_info"`
	PublicCourseList []PublicCourseList `json:"public_course_list"`
}

// 公开课
type PublicCourseInfo struct {
	CourseTitle    string `json:"course_title"`
	TotalCourseNum int64  `json:"total_course_num"`
}

// 公开课列表
type PublicCourseList struct {
	CoverImage string   `json:"cover_image"`
	ID         int64    `json:"id"`
	Intro      string   `json:"intro"`
	Labels     []string `json:"labels"`
	Name       string   `json:"name"`
	ShowType   int64    `json:"show_type"`
	Subtitle   string   `json:"subtitle"`
	Type       int64    `json:"type"`
	ViewNum    int64    `json:"view_num"`
	Sort       int64    `json:"sort"`
	VideoType  string   `json:"video_type"` // 直播live、录播video、音频audio 回放replay
}

func (ctl IndexCtl) publicCourseList(ctx *gin.Context, lp *IndexParam) (PublicCourseInfo, []PublicCourseList, []error) {
	var dataList []PublicCourseList
	var publicInfo PublicCourseInfo
	time.Sleep(time.Second * 2)
	fmt.Println("publicCourseList-7")
	return publicInfo, dataList, nil
}
