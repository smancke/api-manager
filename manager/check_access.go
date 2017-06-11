package manager

import (
	"github.com/smancke/api-manager/model"
	"math"
	"time"
)

type Access struct {
	Key     string
	Secret  string
	IP      string
	Referer string
}

type CheckAccessResult struct {
	Remaining int
	Limit     int
	Reset     time.Time
}

// CheckAccess checks the secret, the ip and the referer
// and returns the limits for an api key
func CheckAccess(getApiKeyInfo model.GetApiKeyInfo, calculateAccessCount model.CalculateAccessCount, access Access) (CheckAccessResult, error) {
	limit := CheckAccessResult{}

	keyInfo, exist, err := getApiKeyInfo(access.Key)
	if err != nil {
		return limit, err
	}
	if !exist {
		return limit, model.ErrorInvalidCredentials
	}

	authenticated := keyInfo.Secret == access.Secret

	if access.Secret != "" && !authenticated {
		return limit, model.ErrorInvalidCredentials
	}

	var r model.Restriction
	if authenticated {
		r = keyInfo.AuthenticatedAccess
	} else {
		r = keyInfo.UnauthenticatedAccess
	}

	if !r.Active {
		return limit, model.ErrorAccessDenied
	}

	if !matchReferer(r, access.Referer) {

	}

	if r.IPLimit > 0 {
		n, err := calculateAccessCount(access.IP, r.IPLimitTimeframe)
		if err != nil {
			return limit, err
		}

		return CheckAccessResult{
			Remaining: r.IPLimit - n,
			Limit:     r.IPLimit,
			Reset:     time.Now().Add(r.IPLimitTimeframe),
		}, nil
	}

	return CheckAccessResult{
		Remaining: math.MaxInt32,
		Limit:     math.MaxInt32,
		Reset:     time.Now(),
	}, nil
}

func matchReferer(r model.Restriction, item string) bool {
	if len(r.RefererWhitelist) == 0 {
		return true
	}

	for _, v := range r.RefererWhitelist {
		if v == item {
			return true
		}
	}
	return false
}
