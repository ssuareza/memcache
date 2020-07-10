package memcache

import "testing"

func (m *mockClient) FlushAll() error {
	return nil
}

func TestFlushAll(t *testing.T) {
	mockClient := &mockClient{}

	err := mockClient.FlushAll()
	if err != nil {
		t.Error(err)
	}
}
