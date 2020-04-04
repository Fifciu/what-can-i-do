package models

type Idea struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ProblemID   uint    `json:"problem_id"`
	Description   string    `json:"description"`
	Price float32 `json:"price"`
}

func (idea Idea) TableName() string {
	return "ideas"
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

func GetProblemIdeas(problemId int) []*Idea {
	ideas := []*Idea{}
	GetDB().Table("ideas").Select("id, description, price").Where("problem_id = ?", problemId).Scan(&ideas)

	return ideas
}

//func GetProblem(problemId int, withIdeas bool) *Problem {
//	problem := &Problem{}
//	if !withIdeas {
//		GetDB().Table("problems").Select("*").Where("id = ?", problemId).First(problem)
//	}
//	//GetDB().Table("message").Select("message.*, user.name").Joins("INNER JOIN user ON user.id = message.user_id").Scan(&messages)
//
//	return problem
//}