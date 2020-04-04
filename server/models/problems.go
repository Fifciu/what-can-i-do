package models

type Problem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title   string    `json:"title"`
	Description   string    `json:"description"`
	IsAccepted bool `json:"is_accepted"`
}

func (problem Problem) TableName() string {
	return "problems"
}

//func (problem *Problem) Save(userId uint, userName string) {
//	if len(message.Message) < 1 {
//		return
//	}
//
//	message.UserID = userId
//	message.Name = userName
//	message.CreatedAt = time.Now().UTC()
//
//	GetDB().Create(message)
//}

func GetAllProblems() []*Problem {
	problems := []*Problem{}
	GetDB().Table("problems").Select("*").Scan(&problems)

	return problems
}

func GetProblem(problemId int, withIdeas bool) *Problem {
	problem := &Problem{}
	if !withIdeas {
		GetDB().Table("problems").Select("*").Where("id = ?", problemId).First(problem)
	}
	//GetDB().Table("message").Select("message.*, user.name").Joins("INNER JOIN user ON user.id = message.user_id").Scan(&messages)

	return problem
}