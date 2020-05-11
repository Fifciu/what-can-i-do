package models

import (
	//"sort"
	"errors"
)

type Vote struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	IdeaID   uint    `json:"idea_id"`
	UserID   uint    `json:"user_id"`
	Delta   int    `json:"delta"`
}

func (vote Vote) TableName() string {
	return "votes"
}

func (vote *Vote) SingularName() string {
	return "vote"
}

func (vote *Vote) PluralName() string {
	return "votes"
}

func (vote *Vote) Validate() error {
	if vote.Delta != 1 && vote.Delta != -1 {
		return errors.New("Bad vote's delta")
	}

	if vote.UserID < 1 {
		return errors.New("Bad user's ID")
	}

	if vote.IdeaID < 1 {
		return errors.New("Bad idea's ID")
	}

	if !UserExists(vote.UserID) {
		return errors.New("User does not exist")
	}

	if !IdeaExistsAndPublished(vote.IdeaID) {
		return errors.New("Idea does not exist")
	}

	return nil
}

//func (vote *Vote) GetByUserId(userId uint) []UserCreatedEntity {
//	ideas := []*Vote{}
//	GetDB().
//		Table("ideas").
//		Select("problems.name AS problem_name, ideas.id, ideas.problem_id, ideas.action_description, ideas.results_description, ideas.money_price, ideas.time_price, ideas.is_published").
//		Joins("INNER JOIN problems ON ideas.problem_id = problems.id").
//		Where("ideas.user_id = ?", userId).
//		Scan(&ideas)
//
//	uces := make([]UserCreatedEntity, len(ideas))
//	for i, idea := range ideas {
//		uces[i] = idea
//	}
//	return uces
//}

func (vote *Vote) Save() error {
	existingVote := &Vote{}
	GetDB().
		Table("votes").
		Select("id, delta").
		Where("idea_id = ? AND user_id = ?", vote.IdeaID, vote.UserID).
		First(existingVote)

	// Existing vote with same delta
	if existingVote.ID > 0 && vote.Delta == existingVote.Delta {
		return errors.New("This vote already exist")
	}

	// Existing vote with different delta
	if existingVote.ID > 0 && vote.Delta != existingVote.Delta {
		// Update
		GetDB().Model(&existingVote).Update("delta", vote.Delta)
	} else {
		GetDB().Create(vote)
	}

	return nil
}

func (vote *Vote) SetUserId(userId uint) {
	vote.UserID = userId
}

func (vote *Vote) GetNewInstance() DatabaseType {
	return &Vote{}
}