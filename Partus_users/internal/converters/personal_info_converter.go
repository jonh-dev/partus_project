package converters

import (
	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/model"
)

func ToModelPersonalInfo(personalInfo *api.PersonalInfo) *model.PersonalInfo {
	return &model.PersonalInfo{
		FirstName: personalInfo.FirstName,
		LastName:  personalInfo.LastName,
		Email:     personalInfo.Email,
		BirthDate: personalInfo.BirthDate.AsTime(),
		Phone:     personalInfo.Phone,
	}
}
