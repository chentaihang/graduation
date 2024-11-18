package models

type Student struct {
	StudentNumber int     `json:"student_number" gorm:"primary_key"` // 学号，设置为主键
	Name          string  `json:"name"`
	Password      string  `json:"-"` // 密码
	Age           int     `json:"age"`
	Degress       string  `json:"degress"`
	Academy       string  `json:"academy"`
	Major         string  `json:"major"`
	ClassNumber   int     `json:"class_number"`
	StudentCredit float64 `json:"student_credit"`
	TopicID       int     `json:"topic_id"`
	Topic         Topic   `json:"topic"`
}
