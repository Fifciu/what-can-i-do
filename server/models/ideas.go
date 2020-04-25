package models

import (
	"sort"
	"errors"
)

type Idea struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ProblemID   uint    `json:"problem_id"`
	Name string `json:"name" gorm:"-" `
	ProblemName string `json:"problem_name" gorm:"-" `
	UserID   uint    `json:"user_id"`
	AuthorName string `gorm:"-" json:"author_name"`
	IsPublished   bool    `json:"is_published"`
	ActionDescription   string    `json:"action_description"`
	ResultsDescription   string    `json:"results_description"`
	MoneyPrice float32 `json:"money_price"`
	TimePrice int `json:"time_price"`
}

func (idea Idea) TableName() string {
	return "ideas"
}

func (idea *Idea) SingularName() string {
	return "idea"
}

func (idea *Idea) PluralName() string {
	return "ideas"
}

func (idea *Idea) Validate() error {
	// TODO TimePrice validator with special values
	if idea.ProblemID < 1 {
		return errors.New("Bad Problem ID")
	}

	if !ProblemExists(idea.ProblemID) {
		return errors.New("Problem does not exist")
	}

	if len(idea.ActionDescription) < 15 {
		return errors.New("Idea's action description must have at least 15 characters")
	}

	if len(idea.ResultsDescription) < 15 {
		return errors.New("Idea's results description must have at least 15 characters")
	}

	if idea.MoneyPrice < 0 {
		return errors.New("Idea's price be bigger or equal $0")
	}

	return nil
}

type IdeasMapper func ([]*Idea)

func MapperAddProblemsName (ideas []*Idea) {
	problemIdsSet := make(map[uint]bool)
	for _, idea := range ideas {
		problemIdsSet[idea.ProblemID] = true
	}
	problemIds := make([]int, len(problemIdsSet))
	counter := 0
	for id, _ := range problemIdsSet {
		problemIds[counter] = int(id)
		counter++
	}
	sort.Ints(problemIds)

	problems := []*Problem{}
	GetDB().
		Table("problems").
		Select("id, name").
		Where("id IN (?)", problemIds).
		Scan(&problems)
	problemsMap := make(map[uint]*Problem)
	for _, problem := range problems {
		problemsMap[problem.ID] = problem
	}
	for _, idea := range ideas {
		if problem, ok := problemsMap[idea.ProblemID]; ok {
			idea.ProblemName = problem.Name
		} else {
			// I SHOULD LOG IT
			idea.ProblemName = "Not found"
		}
	}
}

func GetProblemIdeas(problemId uint) []*Idea {
	ideas := []*Idea{}
	GetDB().
		Table("ideas").
		Select("id, action_description, results_description, money_price, time_price").
		Where("problem_id = ? AND is_published = 1", problemId).
		Scan(&ideas)
	return ideas
}

func (idea *Idea) GetByUserId(userId uint) []*Idea {
	ideas := []*Idea{}
	GetDB().
		Table("ideas").
		Select("problems.name, ideas.id, ideas.problem_id, ideas.action_description, ideas.results_description, ideas.money_price, ideas.time_price, ideas.is_published").
		Joins("INNER JOIN problems ON ideas.problem_id = problems.id").
		Where("user_id = ?", userId).
		Scan(&ideas)

	//GetDB().Table("message").Select("message.*, user.name").Joins("INNER JOIN user ON user.id = message.user_id").Scan(&messages)


	return ideas
}

func (idea *Idea) Save() error {
	existingIdea := &Idea{}
	GetDB().
		Table("ideas").
		Select("action_description").
		Where("action_description = ? AND problem_id = ? AND is_published = 1", idea.ActionDescription, idea.ProblemID).
		First(existingIdea)
	if existingIdea.ActionDescription == idea.ActionDescription {
		return errors.New("This action already exist for selected problem")
	}

	idea.IsPublished = false
	GetDB().Create(idea)
	return nil
}

func (idea *Idea) SetUserId(userId uint) {
	idea.UserID = userId
}

func (idea *Idea) GetNewInstance() DatabaseType {
	return &Idea{}
}