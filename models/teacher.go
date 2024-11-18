package models

type Teacher struct {
	TeacherID  int    `json:"teacher_id" gorm:"primary_key"` // 教师id
	Password   string `json:"-"`                             // 密码
	Name       string `json:"name"`                          // 教师姓名
	Age        int    `json:"age"`                           // 教师年龄
	Degree     string `json:"degree"`                        // 教师学历
	Employer   string `json:"employer"`                      // 教师工作单位
	Profession string `json:"profession"`                    // 教师专业
}
