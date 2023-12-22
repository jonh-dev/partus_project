package converters

import (
	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/model"
)

func ToModelAccountInfo(accountInfo *api.AccountInfo) *model.AccountInfo {
	return &model.AccountInfo{
		Username:              accountInfo.Username,
		Password:              accountInfo.Password,
		AccountStatus:         model.AccountStatus(accountInfo.AccountStatus),
		StatusReason:          accountInfo.StatusReason,
		CreatedAt:             accountInfo.CreatedAt.AsTime(),
		UpdatedAt:             accountInfo.UpdatedAt.AsTime(),
		LastLogin:             accountInfo.LastLogin.AsTime(),
		FailedLoginAttempts:   accountInfo.FailedLoginAttempts,
		LastFailedLogin:       accountInfo.LastFailedLogin.AsTime(),
		LastFailedLoginReason: accountInfo.LastFailedLoginReason,
		AccountLockedUntil:    accountInfo.AccountLockedUntil.AsTime(),
		AccountLockedReason:   accountInfo.AccountLockedReason,
	}
}
