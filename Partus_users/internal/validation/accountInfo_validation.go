package validation

import (
	"errors"
	"regexp"
	"time"

	"github.com/jonh-dev/partus_users/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OperationType int

const (
	Create OperationType = iota
	Update
)

var (
	ErrInvalidUsername             = errors.New("o nome de usuário deve começar e terminar com um caractere alfanumérico, pode conter letras, números, pontos, hifens e sublinhados, não pode conter caracteres especiais consecutivos, e deve ter entre 3 e 20 caracteres")
	ErrInvalidPassword             = errors.New("a senha deve ter entre 8 e 64 caracteres, conter pelo menos uma letra maiúscula, uma letra minúscula, um número e um caractere especial")
	ErrInvalidAccountStatus        = errors.New("o status da conta deve ser ACTIVE, INACTIVE, PENDING ou SUSPENDED")
	ErrInvalidStatusReason         = errors.New("a razão do status deve ser uma string não vazia se o status da conta não for ACTIVE")
	ErrCreatedAtCannotBeUpdated    = errors.New("o campo createdAt não pode ser atualizado")
	ErrLastFailedLoginInFuture     = errors.New("a última tentativa de login falhada não pode estar no futuro")
	ErrFailedLoginAttemptsNegative = errors.New("o número de tentativas de login falhadas não pode ser negativo")
	ErrLastFailedLoginReasonEmpty  = errors.New("a razão da última tentativa de login falhada não pode estar vazia se houve uma tentativa de login falhada")
)

func ValidateAccountInfo(accountInfo *api.AccountInfo, operation OperationType, originalAccountInfo *api.AccountInfo) error {
	if !isValidUsername(accountInfo.Username) {
		return ErrInvalidUsername
	}

	if !isValidPassword(accountInfo.Password) {
		return ErrInvalidPassword
	}

	if !isValidAccountStatus(accountInfo.AccountStatus) {
		return ErrInvalidAccountStatus
	}

	if !isValidStatusReason(accountInfo.AccountStatus, accountInfo.StatusReason) {
		return ErrInvalidStatusReason
	}

	if operation == Update {
		if !isCreatedAtUnchanged(originalAccountInfo.CreatedAt, accountInfo.CreatedAt) {
			return ErrCreatedAtCannotBeUpdated
		}
	}

	// Adicione aqui as outras validações do AccountInfo

	return nil
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^(?i)[a-z0-9]+([._-]?[a-z0-9]+)*$`)
	return len(username) >= 3 && len(username) <= 20 && re.MatchString(username)
}

func isValidPassword(password string) bool {
	hasMin := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasMaj := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNum := regexp.MustCompile(`\d`).MatchString(password)
	hasSpec := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)
	length := len(password)

	if hasMin && hasMaj && hasNum && hasSpec && length >= 8 && length <= 64 {
		return true
	}

	return false
}

func isValidAccountStatus(accountStatus api.AccountStatus) bool {
	switch accountStatus {
	case api.AccountStatus_ACTIVE, api.AccountStatus_INACTIVE, api.AccountStatus_PENDING, api.AccountStatus_SUSPENDED:
		return true
	default:
		return false
	}
}

func isValidStatusReason(accountStatus api.AccountStatus, statusReason string) bool {
	if accountStatus != api.AccountStatus_ACTIVE && statusReason == "" {
		return false
	}
	return true
}

func isCreatedAtUnchanged(originalCreatedAt *timestamppb.Timestamp, updatedCreatedAt *timestamppb.Timestamp) bool {
	return originalCreatedAt.AsTime().Equal(updatedCreatedAt.AsTime())
}

func ValidateLoginFields(accountInfo *api.AccountInfo) error {
	if !isLastFailedLoginNotInFuture(accountInfo.LastFailedLogin) {
		return ErrLastFailedLoginInFuture
	}

	if accountInfo.FailedLoginAttempts < 0 {
		return ErrFailedLoginAttemptsNegative
	}

	if accountInfo.FailedLoginAttempts > 0 && accountInfo.LastFailedLoginReason == "" {
		return ErrLastFailedLoginReasonEmpty
	}

	return nil
}

func isLastFailedLoginNotInFuture(lastFailedLogin *timestamppb.Timestamp) bool {
	return lastFailedLogin.AsTime().Before(time.Now())
}
