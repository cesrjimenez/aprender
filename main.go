package testear

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

/*
/order_processor
  main.go
  order.go             ← modelo mezclado con detalles de la base de datos
  service.go           ← contiene ambos logica de negocio + SQL + llamadas de HTTP
  repository.go        ← use do SQL directo (used directly in service)
  email_client.go      ← uso de HTTP directo (used directly in service)
*/

type Order struct {
	ID        int
	Customer  string
	ItemID    string
	Email     string
	Confirmed bool
}

func PlaceOrder(db *sql.DB, order Order) error {
	// Regla de negocio: reserve inventory (direct SQL)
	_, err := db.Exec(`UPDATE inventory SET reserved = TRUE WHERE item_id = $1`, order.ItemID)
	if err != nil {
		return fmt.Errorf("failed to reserve inventory: %w", err)
	}

	// Regla de negocio: send confirmation email (direct HTTP call)
	resp, err := http.Post("https://email-service.local/send",
		"application/json",
		strings.NewReader(fmt.Sprintf(`{"to": "%s", "message": "Order confirmed!"}`, order.Email)),
	)
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send confirmation: %w", err)
	}

	fmt.Println("✅ Order placed successfully!")
	return nil
}

func main() {
	db, _ := sql.Open("postgres", "postgres://user:pass@localhost:5432/shop?sslmode=disable")

	order := Order{
		ID:       1,
		Customer: "Alice",
		ItemID:   "SKU-123",
		Email:    "alice@example.com",
	}

	if err := PlaceOrder(db, order); err != nil {
		fmt.Println("❌", err)
	}
}
