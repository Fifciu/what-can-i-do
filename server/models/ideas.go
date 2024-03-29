package models

import (
	"errors"
)

type Idea struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ProblemID   uint    `json:"problem_id"`
	Score int `json:"score" gorm:"-" `
	MyVote int `json:"my_vote" gorm:"-" `
	ProblemName string `json:"problem_name" gorm:"-" `
	Reviews int `json:"reviews" gorm:"-" `
	UserID   uint    `json:"user_id"`
	AuthorName string `gorm:"-" json:"author_name"`
	IsPublished   bool    `json:"is_published"`
	IsReviewed   bool    `json:"is_reviewed" gorm:"-" `
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

//func MapperAddProblemsName (ideas []*Idea) {
//	problemIdsSet := make(map[uint]bool)
//	for _, idea := range ideas {
//		problemIdsSet[idea.ProblemID] = true
//	}
//	problemIds := make([]int, len(problemIdsSet))
//	counter := 0
//	for id, _ := range problemIdsSet {
//		problemIds[counter] = int(id)
//		counter++
//	}
//	sort.Ints(problemIds)
//
//	problems := []*Problem{}
//	GetDB().
//		Table("problems").
//		Select("id, name").
//		Where("id IN (?)", problemIds).
//		Scan(&problems)
//	problemsMap := make(map[uint]*Problem)
//	for _, problem := range problems {
//		problemsMap[problem.ID] = problem
//	}
//	for _, idea := range ideas {
//		if problem, ok := problemsMap[idea.ProblemID]; ok {
//			idea.ProblemName = problem.Name
//		} else {
//			// I SHOULD LOG IT
//			idea.ProblemName = "Not found"
//		}
//	}
//}

func GetProblemIdeas(problemId uint) []*Idea {
	ideas := []*Idea{}
	GetDB().
		Table("ideas").
		Select("id, action_description, results_description, money_price, time_price").
		Where("problem_id = ? AND is_published = 1", problemId).
		Scan(&ideas)
	return ideas
}

func IdeaExistsAndPublished(ideaId uint) bool {
	idea := &Idea{}

	GetDB().Table("ideas").Select("id").Where("id = ? AND is_published = 1", ideaId).First(idea)

	if idea.ID > 0 {
		return true
	}
	return false
}

func (idea *Idea) GetByUserId(userId uint) []UserCreatedEntity {
	ideas := []*Idea{}
	GetDB().
		Table("ideas").
		Select("idea_reviews.id AS reviews, problems.name AS problem_name, ideas.id, ideas.problem_id, ideas.action_description, ideas.results_description, ideas.money_price, ideas.time_price, ideas.is_published").
		Joins("INNER JOIN problems ON ideas.problem_id = problems.id").
		Joins("LEFT JOIN idea_reviews ON ideas.id = idea_reviews.idea_id").
		Where("ideas.user_id = ?", userId).
		Group("idea_reviews.idea_id").
		Scan(&ideas)

	uces := make([]UserCreatedEntity, len(ideas))
	for i, idea := range ideas {
		if idea.Reviews > 0 {
			idea.IsReviewed = true
		}
		uces[i] = idea
	}
	return uces
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

func (idea *Idea) Resolve() error {
	existingIdea := &Idea{}
	GetDB().
		Table("ideas").
		Select("*").
		Where("id = ?", idea.ID).
		First(existingIdea)
	if existingIdea.ID < 1 {
		return errors.New("This idea does not exist")
	}
	if existingIdea.IsPublished {
		return errors.New("This idea is already published")
	}

	existingIdea.IsPublished = true
	d := GetDB().Save(existingIdea)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func GetIdeasToReview() []*Idea {
	ideas := []*Idea{}

	GetDB().Table("ideas").Select("*").Where("is_published = 0").Order("id desc").Scan(&ideas)
	return ideas
}
