package es

import (
	"fmt"
	"github.com/smancke/go-ht"
)

// Store is a Elasticsearch based store implementation
type Store struct {
	esBaseUrl    string
	accountIndex string
}

func (s *Store) GetApiKeyInfo(key string) (info ApiKeyInfo, exist bool, err error) {
	url := fmt.Sprintf("%v/%v/apiKeys/%v", s.esBaseUrl, s.accountIndex, sanitizeKey(key))
	err = ht.FetchJson()
	// TODO: handle the case, where the key does not exist
	return info, true, error
}

// TODO: Don't trust the key! It may be userinput
func sanitizeKey(userInputKey string) string {
	//TODO: implement
	return userInputKey
}
