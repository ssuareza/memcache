package memcache

// ClientIface is an interface for mocking.
type ClientIface interface {
	Get(key string) (*Item, error)
	GetMulti(keys []string) (map[string]*Item, error)
	Set(item *Item) error
	Replace(item *Item) error
	FlushAll() error
}

type mockClient struct {
	ClientIface
}
