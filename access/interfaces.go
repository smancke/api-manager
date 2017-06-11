package access

//go:generate go get github.com/golang/mock/mockgen
//go:generate $GOPATH/bin/mockgen -package access -destination $GOPATH/src/github.com/smancke/api-manager/access/mocks_test.go github.com/smancke/api-manager/access UsageStore
type UsageStore interface {
	GetLimit(key, secret, ip, referer string) (remaining, limit int, reset uint64, err error)
}
