package models

type Problem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID        uint      `json:"user_id"`
	Name   string    `json:"name"`
	Description   string    `json:"description"`
	IsPublished bool `json:"is_published"`
	Ideas []*Idea `json:"ideas" gorm:"foreignkey:ProblemID"`
}

func (problem Problem) TableName() string {
	return "problems"
}

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
		userIdsSet := make(map[uint]bool)
		for _, idea := range ideas {
			userIdsSet[idea.UserID] = true
		}
		userIds := make([]uint, len(userIdsSet))
		counter := 0
		for id, _ := range userIdsSet {
			userIds[counter] = id
			counter++
		}
		users := []*User{}
		GetDB().Table("users").Select("id, fullname").Where("id IN (?)", userIds).Scan(&users)

		userIdUserMap := make(map[uint]*User)
		for _, user := range users {
			userIdUserMap[user.ID] = user
		}

		for i, idea := range ideas {
			ideas[i].AuthorName = userIdUserMap[idea.UserID].Fullname

		}
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

func (problem *Problem) Save(userId uint, name string, description string) bool {
	existingProblem := &Problem{}
	GetDB().Table("problems").Select("*").Where("name = ? AND is_published = 1", name).First(existingProblem)

	if existingProblem.Name == name {
		return false
	}

	existingUnpublishedProblem := &Problem{}
	GetDB().Table("problems").Select("*").Where("name = ? AND description = ? AND is_published = 0", name, description).First(existingUnpublishedProblem)
	if existingUnpublishedProblem.ID > 0 {
		return false
	}

	problem.UserID = userId
	problem.Name = name
	problem.Description = description
	d := GetDB().Create(problem)
	if d.Error != nil {
		return false
	}
	return true
}