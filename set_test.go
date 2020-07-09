package memcache

import "testing"

func (m *mockClient) Set(item *Item) error {
	return nil
}

func TestSet(t *testing.T) {
	mockClient := &mockClient{}
	item := &Item{Key: "key1", Value: []byte("value1")}

	err := mockClient.Set(item)
	if err != nil {
		t.Error(err)
	}
}
