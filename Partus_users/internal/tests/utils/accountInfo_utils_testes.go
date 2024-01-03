package utils

import (
	"strings"

	"github.com/jonh-dev/partus_users/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateValidAccountInfo() *api.AccountInfo {
	return &api.AccountInfo{
		UserId:                "valid_user_id",
		Username:              "valid_username_1",
		Password:              "ValidPassword123!",
		AccountStatus:         api.AccountStatus_ACTIVE,
		StatusReason:          "",
		CreatedAt:             timestamppb.Now(),
		UpdatedAt:             timestamppb.Now(),
		LastLogin:             timestamppb.Now(),
		FailedLoginAttempts:   0,
		LastFailedLogin:       timestamppb.Now(),
		LastFailedLoginReason: "",
		AccountLockedUntil:    timestamppb.Now(),
		AccountLockedReason:   "",
	}
}

func CreateInvalidUserIdAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.UserId = "invalid_user_id"
	return accountInfo
}

func CreateShortUsernameAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Username = "ab"
	return accountInfo
}

func CreateLongUsernameAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Username = "a123456789012345678901"
	return accountInfo
}

func CreateSpecialCharConsecutiveUsernameAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Username = "invalid__username"
	return accountInfo
}

func CreateSpecialCharStartUsernameAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Username = "_invalidusername"
	return accountInfo
}

func CreateSpecialCharEndUsernameAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Username = "invalidusername_"
	return accountInfo
}

func CreateInvalidCharUsernameAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Username = "invalidusername@"
	return accountInfo
}

func CreateNoLowercasePasswordAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Password = "VALIDPASSWORD123!"
	return accountInfo
}

func CreateNoUppercasePasswordAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Password = "validpassword123!"
	return accountInfo
}

func CreateNoNumberPasswordAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Password = "ValidPassword!"
	return accountInfo
}

func CreateNoSpecialCharPasswordAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Password = "ValidPassword123"
	return accountInfo
}

func CreateShortPasswordAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Password = "Val12!"
	return accountInfo
}

func CreateLongPasswordAccountInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.Password = "V" + strings.Repeat("a", 63) + "1!"
	return accountInfo
}

func CreateInvalidAccountStatusInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.AccountStatus = api.AccountStatus(4)
	return accountInfo
}

func CreateInvalidStatusReasonInfo() *api.AccountInfo {
	accountInfo := CreateValidAccountInfo()
	accountInfo.AccountStatus = api.AccountStatus_INACTIVE
	accountInfo.StatusReason = ""
	return accountInfo
}
