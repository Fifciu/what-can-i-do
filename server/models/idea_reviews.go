package models

import (
	"time"
)

type IdeaReview struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ReviewerID        uint      `json:"user_id"`
	IdeaID   uint    `json:"idea" gorm:"foreignkey:IdeaID"`
	Message   string    `json:"message"`
	Date   time.Time    `json:"date"`
}

func (ideaReview IdeaReview) TableName() string {
	return "idea_reviews"
}

func (ideaReview *IdeaReview) SingularName() string {
	return "idea_review"
}

func (ideaReview *IdeaReview) PluralName() string {
	return "idea_reviews"
}

func (ideaReview *IdeaReview) GetIdea(ideaId uint, userId uint, isModerator bool) []*IdeaReview {
	reviews := []*IdeaReview{}
	if isModerator {
		GetDB().Table(ideaReview.TableName()).Select("*").Where("idea_id = ? AND reviewer_id = ?", ideaId, userId).Scan(&reviews)
	} else {
		GetDB().
			Table(ideaReview.TableName()).
			Select(ideaReview.TableName() + ".*").
			Joins("INNER JOIN ideas ON ideas.id = idea_reviews.idea_id").
			Where("idea_reviews.idea_id = ? AND ideas.user_id = ?", ideaId, userId).
			Scan(&reviews)
	}

	return reviews
}

func (r *IdeaReview) Save(ideaId uint, reviewerId uint, message string) error {
	review := &IdeaReview{}
	review.IdeaID = ideaId
	review.ReviewerID = reviewerId
	review.Message = message
	d := GetDB().Create(review)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func GetIdeaReviews(ideaId uint, userId uint) []*IdeaReview {
	reviews := []*IdeaReview{}
	GetDB().
		Table("idea_reviews").
		Select("idea_reviews.*").
		Joins("INNER JOIN ideas ON ideas.id = idea_reviews.idea_id").
		Where("idea_reviews.idea_id = ? AND ideas.user_id = ?", ideaId, userId).
		Scan(&reviews)

	return reviews
}