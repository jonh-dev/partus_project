package model

type AccountStatus int32

const (
	AccountStatus_ACTIVE    AccountStatus = 0
	AccountStatus_INACTIVE  AccountStatus = 1
	AccountStatus_PENDING   AccountStatus = 2
	AccountStatus_SUSPENDED AccountStatus = 3
)
