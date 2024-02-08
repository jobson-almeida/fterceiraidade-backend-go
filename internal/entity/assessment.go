package entity

import (
	"github.com/lib/pq"
)

type Assessment struct {
	Base
	Description string         `json:"description" validate:"required" gorm:"type:varchar(255)"`
	Courses     pq.StringArray `json:"image" validate:"required" gorm:"type:text[]"`
	Classrooms  pq.StringArray `json:"classrooms" validate:"required" gorm:"type:text[]"`
	StartDate   string         `json:"startdate" validate:"required" gorm:"type:varchar"`
	EndDate     string         `json:"enddate" validate:"required" gorm:"type:varchar"`
	Quiz        []*Quiz        `json:"quiz" gorm:"serializer:json"`
}

func init() {}

func NewAssessment() *Assessment {
	return &Assessment{}
}

func UpdateAssessment() *Assessment {
	return &Assessment{}
}

/*
func (assessment *Assessment) Prepare() *validator.ErrResponse {
	res, err := regexp.MatchString("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", assessment.ID)
	if err != nil {
		return validator.ToErrResponse(err)
	}
	if !res {
		return validator.ToErrResponse(err)
	}

	err = validator.New().Struct(assessment)
	if err != nil {
		return validator.ToErrResponse(err)
	}
	return nil
}
*/
