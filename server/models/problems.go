package models

import (
	"github.com/gosimple/slug"
)

type Problem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID        uint      `json:"user_id"`
	Name   string    `json:"name"`
	Slug string `json:"slug"`
	Description   string    `json:"description"`
	IsPublished bool `json:"is_published"`
	Views uint `json:"views"`
	Ideas []*Idea `json:"ideas" gorm:"foreignkey:ProblemID"`
}

func (problem Problem) TableName() string {
	return "problems"
}

func (problem *Problem) SingularName() string {
	return "problem"
}

func (problem *Problem) PluralName() string {
	return "problems"
}

func GetAllProblems() []*Problem {
	problems := []*Problem{}
	GetDB().Table("problems").Select("*").Where("is_published = 1").Scan(&problems)
	return problems
}

func (problem *Problem) GetByUserId(userId uint) []UserCreatedEntity {
	problems := []*Problem{}
	GetDB().Table("problems").Select("*").Where("user_id = ?", userId).Scan(&problems)

	uces := make([]UserCreatedEntity, len(problems))
	for i, problem := range problems {
		uces[i] = problem
	}
	return uces
}

// UserID to fetch votes if user's voted
func GetProblem(problemSlug string, withIdeas bool, userId uint) *Problem {
	problem := &Problem{}

	GetDB().Table("problems").Select("*").Where("slug = ? AND is_published = 1", problemSlug).First(problem)

	if withIdeas {
		problemId := problem.ID
		// Add a view
		if problem.ID > 0 {
			GetDB().Model(&problem).Update("views", problem.Views + 1)
		}

		ideas := []*Idea{}
		if userId > 0 {
			//SELECT IFNULL(myVotes.delta,25) as my_vote, IFNULL(SUM(allVotes.delta),0) as score, ideas.*, users.fullname as author_name
			//FROM ideas
			//INNER JOIN users ON ideas.user_id = users.id
			//JOIN votes allVotes ON ideas.id = allVotes.idea_id
			//JOIN votes myVotes ON ideas.id = myVotes.idea_id AND myVotes.user_id = 1
			//WHERE ideas.problem_id = 7 AND ideas.is_published = 1

			GetDB().
				Table("ideas").
				Select("IFNULL(myVotes.delta,25) as my_vote, IFNULL(SUM(allVotes.delta),0) as score, ideas.*, users.fullname as author_name").
				Joins("INNER JOIN users ON ideas.user_id = users.id").
				Joins("JOIN votes allVotes ON ideas.id = allVotes.idea_id").
				Joins("JOIN votes myVotes ON ideas.id = myVotes.idea_id AND myVotes.user_id = ?", userId).
				Where("ideas.problem_id = ? AND ideas.is_published = 1", problemId).
				Scan(&ideas)
		} else {
			GetDB().
				Table("ideas").
				Select("IFNULL(SUM(votes.delta),0) as score, ideas.*, users.fullname as author_name").
				Joins("INNER JOIN users ON ideas.user_id = users.id").
				Joins("LEFT JOIN votes ON ideas.id = votes.idea_id").
				Where("problem_id = ? AND is_published = 1", problemId).
				Group("id").
				Scan(&ideas)
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

func GetMostPopular() []*Problem {
	problems := []*Problem{}

	GetDB().Table("problems").Select("*").Where("is_published = 1").Order("views desc").Limit(10).Scan(&problems)
	return problems
}

func (problem *Problem) Save(userId uint, name string, description string) bool {
	// TODO Max requests per time for user
	existingUnpublishedProblem := &Problem{}
	// Antyspam
	GetDB().Table("problems").Select("*").Where("name = ? AND description = ? AND is_published = 0", name, description).First(existingUnpublishedProblem)
	if existingUnpublishedProblem.ID > 0 {
		return false
	}

	problem.UserID = userId
	problem.Name = name
	problem.Description = description
	problem.Slug = slug.Make(name)
	d := GetDB().Create(problem)
	if d.Error != nil {
		return false
	}
	return true
}
