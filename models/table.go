package models

import "time"

type Table struct {
	LeaderTeachers []Teacher     `json:"leaderteachers"` // 指导教师信息
	Topic          Topic         `json:"topic"`          // 课题信息
	Student        Student       `json:"student"`        // 学生信息
	Tasks          Task          `json:"tasks"`          // 任务信息
	References     []Reference   `json:"references"`     // 参考文献
	Arrangements   []Arrangement `json:"arrangement"`    // 任务安排
	Commit         Commit        `json:"commit"`         // 审核评价
}

type Topic struct {
	ID          int    `json:"topic_id"`    // 课题id
	Name        string `json:"name"`        // 课题名称
	Type        string `json:"type"`        // 课题类型
	Description string `json:"description"` // 课题描述
	CreateDate  string `json:"create_date"` // 课题创建时间
	UpdateDate  string `json:"update_date"` // 课题更新时间
}

type Task struct {
	TaskInstruction      string `json:"task_instruction"`       // 任务描述
	TaskRequirement      string `json:"task_requirement"`       // 任务要求
	IsIndicative         bool   `json:"is_indicative"`          // 任务指标是否明确
	IsSuitable           bool   `json:"is_suitable"`            // 任务内容是否合适
	ScheduleIsReasonable bool   `json:"schedule_is_reasonable"` // 安排是否合理
}

// TableName 指定表名
func (Topic) TableName() string {
	return "topics"
}

type Reference struct {
	Name      string `json:"name"`      // 参考文献名称
	Author    string `json:"author"`    // 作者
	Year      int    `json:"year"`      // 出版年份
	Publisher string `json:"publisher"` // 出版社
	ISBN      string `json:"isbn"`      // ISBN
	Link      string `json:"link"`      // 链接
}

type Arrangement struct {
	ArrangeID      int       `json:"arrange_id"`      // 任务安排id
	ArrangeDate    time.Time `json:"arrange_date"`    // 任务安排时间
	ArrangeName    string    `json:"arrange_name"`    // 任务名称
	ArrangeContent string    `json:"arrange_content"` // 任务详细描述
}

type Commit struct {
	TeacherAdvice       string    `json:"teacher_advice"`        // 教师建议
	InstituteAdvice     string    `json:"institute_advice"`      // 学院建议
	TeacherPass         bool      `json:"teacher_pass"`          // 教师是否允许通过
	InstitutePass       bool      `json:"institute_pass"`        // 学院是否允许通过
	TeacherReviewDate   time.Time `json:"teacher_review_date"`   // 教师审核时间
	InstituteReviewDate time.Time `json:"institute_review_date"` // 学院审核时间
	FinalResult         bool      `json:"final_result"`          // 最终审核结果
}
