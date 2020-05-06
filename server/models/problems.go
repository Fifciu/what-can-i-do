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

func GetProblem(problemSlug string, withIdeas bool, userId uint) *Problem {
	problem := &Problem{}

	GetDB().Table("problems").Select("*").Where("slug = ? AND is_published = 1", problemSlug).First(problem)

	if withIdeas {
		problemId := problem.ID
		ideas := []*Idea{}
		if userId > 0 {
			GetDB().
				Table("ideas").
				Select("SUM(myVotes.delta) as my_vote, SUM(votes.delta) as score, ideas.*, users.fullname as author_name").
				Joins("INNER JOIN users ON ideas.user_id = users.id").
				Joins("INNER JOIN votes ON ideas.id = votes.idea_id").
				Joins("INNER JOIN votes myVotes ON ideas.id = myVotes.idea_id AND myVotes.user_id = ?", userId).
				Where("problem_id = ? AND is_published = 1", problemId).
				Scan(&ideas)
		} else {
			GetDB().
				Table("ideas").
				Select("SUM(votes.delta) as score, ideas.*, users.fullname as author_name").
				Joins("INNER JOIN users ON ideas.user_id = users.id").
				Joins("INNER JOIN votes ON ideas.id = votes.idea_id").
				Where("problem_id = ? AND is_published = 1", problemId).
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