package models

type Problem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title   string    `json:"title"`
	Description   string    `json:"description"`
	IsAccepted bool `json:"is_accepted"`
	Ideas []*Idea `json:"ideas" gorm:"foreignkey:ProblemID"`
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

	GetDB().Table("problems").Select("*").Where("id = ?", problemId).First(problem)
	if withIdeas {
		ideas := []*Idea{}
		GetDB().Table("ideas").Select("*").Where("problem_id = ?", problemId).Scan(&ideas)
		problem.Ideas = ideas
	}

	return problem
}

func GetProblemByQuery(searchQuery string) *Problem {
	problem := &Problem{}
	queryAsPart := "%" + searchQuery + "%"

	GetDB().Table("problems").Select("*").Where("title LIKE ?", queryAsPart).First(problem)
	return problem
}

func ProblemExists(problemId uint) bool {
	problem := &Problem{}
	GetDB().Table("problems").Select("id").Where("id = ?", problemId).First(problem)
	if problem.ID == problemId {
		return true
	}
	return false
}

func (problem *Problem) Save(title string, description string) bool {
	existingProblem := &Problem{}
	GetDB().Table("problems").Select("*").Where("title = ?", title).First(existingProblem)

	if existingProblem.Title == title {
		return false
	}

	problem.Title = title
	problem.Description = description
	GetDB().Create(problem)
	return true
}