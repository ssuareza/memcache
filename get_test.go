package memcache

import (
	"testing"
)

func (m *mockClient) Get(key string) (*Item, error) {
	return &Item{Key: "key1", Value: []byte("value1")}, nil
}

// mock GetMulti response
func (m *mockClient) GetMulti(keys []string) (map[string]*Item, error) {
	// create items to return
	items := make(map[string]*Item, 0)
	for _, v := range keys {
		items[v] = &Item{
			Key:   v,
			Value: []byte("random text"),
		}
	}

	return items, nil
}

func TestGet(t *testing.T) {
	mockClient := &mockClient{}
	item, err := mockClient.Get("key1")
	if err != nil {
		t.Error(err)
	}

	expected := "value1"
	if string(item.Value) != expected {
		t.Errorf("Value should be %s and is %s", expected, string(item.Value))
	}
}

func TestGetMulti(t *testing.T) {
	mockClient := &mockClient{}
	keys := []string{"key1", "key2"}

	// and get the keys!
	items, err := mockClient.GetMulti(keys)
	if err != nil {
		t.Error(err)
	}

	expected := 2
	if len(items) != expected {
		t.Errorf("Key number is %v and expected is %v", len(items), expected)
	}
}
