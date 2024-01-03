package validation

type OperationType int

const (
	Create OperationType = iota
	Update
	UpdateProfile
)
