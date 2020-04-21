package models

import (
	"sort"
)

type Idea struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ProblemID   uint    `json:"problem_id"`
	ProblemName string `gorm:"-" json:"problem_name"`
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

func GetUserIdeas(userId uint, ideasMappers []IdeasMapper) []*Idea {
	ideas := []*Idea{}
	GetDB().
		Table("ideas").
		Select("id, problem_id, action_description, results_description, money_price, time_price, is_published").
		Where("user_id = ?", userId).
		Scan(&ideas)
	if len(ideasMappers) > 0 {
		for _, mapper := range ideasMappers {
			mapper(ideas)
		}
	}
	return ideas
}

func (idea *Idea) Save(userID uint, problemID uint, actionDescription string, resultsDescription string, moneyPrice float32, timePrice int) bool {
	existingIdea := &Idea{}
	GetDB().
		Table("ideas").
		Select("action_description").
		Where("action_description = ? AND problem_id = ? AND is_published = 1", actionDescription, problemID).
		First(existingIdea)
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