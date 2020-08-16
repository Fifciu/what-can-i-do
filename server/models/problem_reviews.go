package models

import (
	"time"
)

type ProblemReview struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ReviewerID        uint      `json:"user_id"`
	ProblemID uint `json:"problem" gorm:"foreignkey:ProblemID"`
	Message   string    `json:"message"`
	Date   time.Time    `json:"date"`
}

func (problemReview ProblemReview) TableName() string {
	return "problem_reviews"
}

func (problemReview *ProblemReview) SingularName() string {
	return "problem_review"
}

func (problemReview *ProblemReview) PluralName() string {
	return "problem_reviews"
}

func (problemReview *ProblemReview) GetReviews(problemId uint, userId uint, isModerator bool) []*ProblemReview {
	reviews := []*ProblemReview{}
	if isModerator {
		GetDB().Table(problemReview.TableName()).Select("*").Where("problem_id = ? AND reviewer_id = ?", problemId, userId).Scan(&reviews)
	} else {
		GetDB().
			Table(problemReview.TableName()).
			Select(problemReview.TableName() + ".*").
			Joins("INNER JOIN problems ON problems.id = problem_reviews.problem_id").
			Where("problem_reviews.problem_id = ? AND problems.user_id = ?", problemId, userId).
			Scan(&reviews)
	}

	return reviews
}

func (r *ProblemReview) Save(problemId uint, reviewerId uint, message string) error {
	review := &ProblemReview{}
	review.ProblemID = problemId
	review.ReviewerID = reviewerId
	review.Message = message
	d := GetDB().Create(review)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func GetProblemReviews(problemId uint, userId uint) []*ProblemReview {
	reviews := []*ProblemReview{}
	GetDB().
		Table("problem_reviews").
		Select("problem_reviews.*").
		Joins("INNER JOIN problems ON problems.id = problem_reviews.problem_id").
		Where("problem_reviews.problem_id = ? AND problems.user_id = ?", problemId, userId).
		Scan(&reviews)

	return reviews
}