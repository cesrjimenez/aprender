package service

import (
	"fmt"
	"testear/store"
	"testing"
)

func TestGetUpperCaseUserName(t *testing.T) {
	testCases := []struct {
		name           string
		userID         int
		mockStoreOpts  *store.MockStoreOptions
		expectedResult string
		expectedErr    bool
	}{
		{
			name:   "success - user found",
			userID: 1,
			mockStoreOpts: &store.MockStoreOptions{
				GetUserByIDFunc: func(id int) (string, error) {
					return "alice", nil
				},
			},
			expectedResult: "ALICE",
			expectedErr:    false,
		},
		{
			name:   "error - user not found",
			userID: 999,
			mockStoreOpts: &store.MockStoreOptions{
				GetUserByIDFunc: func(id int) (string, error) {
					return "", fmt.Errorf("user not found")
				},
			},
			expectedResult: "",
			expectedErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := store.NewMockStore(tc.mockStoreOpts)

			s := service{
				store: mockStore,
			}

			result, err := s.GetUpperCaseUserName(tc.userID)
			if err != nil {
				t.Fatal("error")
			}

			if tc.expectedErr && err == nil {
				t.Fatalf("expected error but got nil")
			}

			if !tc.expectedErr && err != nil {
				t.Fatalf("did not expect error but got: %v", err)
			}

			if result != tc.expectedResult {
				t.Fatalf("expected %q, got %q", tc.expectedResult, result)
			}
		})
	}
}
