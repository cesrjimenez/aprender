package store

type MockStoreOptions struct {
	GetUserByIDFunc func(id int) (string, error)
}

type mockStore struct {
	*MockStoreOptions
}

func NewMockStore(opts *MockStoreOptions) Store {
	return &mockStore{MockStoreOptions: opts}
}

func (m *mockStore) GetUserByID(id int) (string, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(id)
	}

	return "", nil
}
