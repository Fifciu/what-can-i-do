package models

type Idea struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ProblemID   uint    `json:"problem_id"`
	UserID   uint    `json:"user_id"`
	AuthorName string `json:"author_name"`
	IsPublished   bool    `json:"-"`
	ActionDescription   string    `json:"action_description"`
	ResultsDescription   string    `json:"results_description"`
	MoneyPrice float32 `json:"money_price"`
	TimePrice int `json:"time_price"`
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
	GetDB().Table("ideas").Select("id, action_description, results_description, money_price, time_price").Where("problem_id = ? AND is_published = 1", problemId).Scan(&ideas)

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

func (idea *Idea) Save(userID uint, problemID uint, actionDescription string, resultsDescription string, moneyPrice float32, timePrice int) bool {
	existingIdea := &Idea{}
	GetDB().Table("ideas").Select("description").Where("action_description = ? AND problem_id = ? AND is_published = 1", actionDescription, problemID).First(existingIdea)
	if existingIdea.ActionDescription == actionDescription {
		return false
	}

	idea.UserID = userID
	idea.ProblemID = problemID
	idea.ActionDescription = actionDescription
	idea.ResultsDescription = resultsDescription
	idea.MoneyPrice = moneyPrice
	idea.TimePrice = timePrice
	idea.IsPublished = false
	GetDB().Create(idea)
	return true
}