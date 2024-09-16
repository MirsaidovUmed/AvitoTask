package errors

import (
	"errors"
	"net/http"
)

var (
	ErrTenderNotFound           = errors.New("tender not found")
	ErrBidNotFound              = errors.New("bid not found")
	ErrOrganizationNotFound     = errors.New("organization not found")
	ErrEmployeeNotFound         = errors.New("employee not found")
	ErrUserNotIsOwnerOrNotExist = errors.New("user is not owner or user/organization does not exist")
	ErrNotEnoughRights          = errors.New("not enough rights")
)

var ErrorStatusMapping = map[error]int{
	ErrTenderNotFound:       http.StatusNotFound,
	ErrOrganizationNotFound: http.StatusNotFound,
	ErrEmployeeNotFound:     http.StatusUnauthorized,
	ErrNotEnoughRights:      http.StatusForbidden,
}
