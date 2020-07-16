package memcache

import "testing"

func (m *mockClient) Replace(item *Item) error {
	return nil
}

func TestReplace(t *testing.T) {
	mockClient := &mockClient{}
	item := &Item{Key: "key1", Value: []byte("value1")}

	err := mockClient.Replace(item)
	if err != nil {
		t.Error(err)
	}
}
