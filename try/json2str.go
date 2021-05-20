package main

import (
	"encoding/json"
	"fmt"
)

type QuestionnaireAnswer struct {
	Answer     string      `json:"answer" form:"answer"  ` //m:"type:text;comment:用户答案"` //单个题目的答案
	ID         int64       `json:"id"`
	IsMust     int64       `json:"is_must"`
	Max        int64       `json:"max"`
	Name       string      `json:"name"`
	Options    []Options   `json:"options"`
	Score      int64       `json:"score"` // 评分
	SubTitles  []SubTitles `json:"subTitles"`
	Type       string      `json:"type"`         // 单选题single 多选题,multiple , 评分：score, 问答：qa 矩阵选择matrixSelection
	TotalCount int64       `json:"total_count" ` // 作答总条数
}

type Options struct {
	Op     int64  `json:"op"`
	Option string `json:"option"`
	Count  int64  `json:"count"`
}
type SubTitles struct { // 矩阵选择题
	Title   string    `json:"title"`
	Count   int64     `json:"count"`
	Options []Options `json:"options"`
}

func main() {
	s := `[{"id":1,"type":"single","name":"1.单选题","is_must":1,"options":[{"op":"A","option":"<p>王：工会、科研<br/>卫：人事、教育、财务<br/>丁：宣传、后勤</p>"},{"op":"B","option":"<p>王：人事、工会<br/>卫：宣传、后勤<br/>丁：科研、教育、财务</p>"},{"op":"C","option":"<p>王：教育、财务<br/>卫：人事、后勤<br/>丁：科研、宣传、工会</p>"},{"op":"D","option":"<p>王：宣传、后勤<br/>卫：人事、工会<br/>丁：科研、教育、财务</p>"}]}]`
	fmt.Println(s)
	var an []QuestionnaireAnswer
	_ = json.Unmarshal([]byte(s), &an)
	fmt.Println(an)
}
