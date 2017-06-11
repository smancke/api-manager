package model

import (
	"time"
)

type GetApiKeyInfo func(key string) (info ApiKeyInfo, exist bool, err error)
type CalculateAccessCount func(ip string, ipLimitTimeframe time.Duration) (int, error)

type Store interface {
	GetApiKeyInfo(key string) (info ApiKeyInfo, exist bool, err error)
	CalculateAccessCount(ip string, ipLimitTimeframe time.Duration) (int, error)
}
