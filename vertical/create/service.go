package create

import "context"

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) Exec(ctx context.Context, req Request) (Response, error) {
	// Logica de negocio para este caso de uso
	id, err := s.store.InsertOrder(ctx, req.CustomerID, req.ItemID, req.Quantity)
	if err != nil {
		return Response{}, err
	}

	return Response{
		OrderID: id,
		Status:  "created",
	}, nil
}
