package hexagonal

import (
	"github.com/stretchr/testify/require"
	"testear/hexagonal/core"
	"testing"
)

type MockInventory struct{}

func (m MockInventory) Reserve(_ string) error { return nil }

type MockNotifier struct{}

func (m MockNotifier) SendConfirmation(_, _ string) error { return nil }

func TestOrderPlacement(t *testing.T) {
	svc := core.NewService(MockInventory{}, MockNotifier{})
	err := svc.PlaceOrder(core.Order{ItemID: "X", Email: "y@z.com"})
	require.NoError(t, err)
}
