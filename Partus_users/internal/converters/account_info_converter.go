package converters

import (
	"errors"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToModelAccountInfo(objectId primitive.ObjectID, accountInfo *api.AccountInfo) (*model.AccountInfo, error) {
	if accountInfo == nil {
		return nil, errors.New("accountInfo não pode ser nil")
	}

	return &model.AccountInfo{
		UserId:                objectId,
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
	}, nil
}
