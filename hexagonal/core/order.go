package core

type Order struct {
	ID       int
	Customer string
	ItemID   string
	Email    string
}

type Service struct {
	inv Inventory
	not Notifier
}

func NewService(inv Inventory, not Notifier) *Service {
	return &Service{inv: inv, not: not}
}

func (s *Service) PlaceOrder(order Order) error {
	// Regla de negocio #1: reserve inventory
	if err := s.inv.Reserve(order.ItemID); err != nil {
		return err
	}

	// Regla de negocio 2: send confirmation
	if err := s.not.SendConfirmation(order.Email, "Order confirmed!"); err != nil {
		return err
	}

	return nil
}
