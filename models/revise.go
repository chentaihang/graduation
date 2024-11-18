package models

type TopicResult struct {
	TopicID     int    `db:"topic_id"`
	Name        string `db:"name"`
	Type        string `db:"type"`
	Description string `db:"description"`
	CreateDate  string `db:"create_date"`
	UpdateDate  string `db:"update_date"`
}

func Topicrevise(student *Student, result *TopicResult) {
	student.Topic = Topic{
		ID:          result.TopicID, // 注意这里映射到 ID 字段
		Name:        result.Name,
		Type:        result.Type,
		Description: result.Description,
		CreateDate:  result.CreateDate,
		UpdateDate:  result.UpdateDate,
	}
	if student.Topic.Name == "" {
		student.Topic.Name = "未选题"
		return
	}
}
