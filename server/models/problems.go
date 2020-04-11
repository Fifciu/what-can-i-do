package models

type Problem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name   string    `json:"name"`
	Description   string    `json:"description"`
	IsPublished bool `json:"is_published"`
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
	GetDB().Table("problems").Select("*").Where("is_published = 1").Scan(&problems)

	return problems
}

func GetProblem(problemId int, withIdeas bool) *Problem {
	problem := &Problem{}

	GetDB().Table("problems").Select("*").Where("id = ? AND is_published = 1", problemId).First(problem)
	if withIdeas {
		ideas := []*Idea{}
		GetDB().Table("ideas").Select("*").Where("problem_id = ? AND is_published = 1", problemId).Scan(&ideas)
		problem.Ideas = ideas
	}

	return problem
}

func GetProblemsByQuery(searchQuery string) []*Problem {
	problems := []*Problem{}
	queryAsPart := "%" + searchQuery + "%"

	GetDB().Table("problems").Select("*").Where("name LIKE ? AND is_published = 1", queryAsPart).Scan(&problems)
	return problems
}

func ProblemExists(problemId uint) bool {
	problem := &Problem{}
	GetDB().Table("problems").Select("id").Where("id = ? AND is_published = 1", problemId).First(problem)
	if problem.ID == problemId {
		return true
	}
	return false
}

func (problem *Problem) Save(name string, description string) bool {
	existingProblem := &Problem{}
	GetDB().Table("problems").Select("*").Where("name = ? AND is_published = 1", name).First(existingProblem)

	if existingProblem.Name == name {
		return false
	}

	problem.Name = name
	problem.Description = description
	GetDB().Create(problem)
	return true
}