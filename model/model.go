package model

import (
	"time"
)

type User struct {
	// The unique id for the user
	Id string `json:"id"`

	// The subject or username, this is only unique in combination with the origin
	Sub string `json:"sub"`

	// The URL for an avatar or picture of the user
	Picture string `json:"picture,omitempty"`

	// The full Name
	Name string `json:"name,omitempty"`

	// The Email address
	Email string `json:"email,omitempty"`

	// The origin or the user e.g. github, google, ...
	Origin string `json:"origin,omitempty"`

	// The id of the user within its origin
	OriginId string `json:"originId,omitempty"`
}

type Plan struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	MonthlyFee       string `json:"monthlyFee"`
	ApiCallsPerMonth int    `json:"apiCallsPerMonth"`
}

type Account struct {
	// The id of the account
	Id string `json:"id"`

	// A list of members for the account
	Members []struct {
		UserId  string `json:"userId"`
		IsOwner bool
	}

	// The selected plan for this account
	PlanId string `json:"planId"`
}

type ApiKey struct {
	// The account, the key belongs to
	AccountId       string `json:"id"`
	Key             string `json:"key"`
	Secret          string `json:"secret"`
	UnsecuredAccess Limit  `json:"unsecuredAccess"`
	SecuredAccess   Limit  `json:"securedAccess"`
}

type Limit struct {
	Active           bool          `json:"active"`
	Unlimited        bool          `json:"unlimited"`
	Limit            int           `json:"limit"`
	LimitTimeframe   time.Duration `json:"limitTimeframe"`
	RefererWhitelist []string      `json:"refererWhitelist"`
}

type ApiUsage struct {
	// The key which was used
	Key string `json:"key"`

	// The account, where the usage was
	AccountId string `json:"id"`

	// the moment, when the usage occurs
	Time time.Time `json:"time"`

	// The ip of the client
	Ip string `json:"ip"`

	// Application specific usage attributes
	Attributes map[string]string `json:"attributes"`
}
