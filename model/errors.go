package model

import (
	"errors"
)

var ErrorInvalidCredentials = errors.New("invalid credentials")
var ErrorAccessDenied = errors.New("access denied")
