package hexagonal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"testear/hexagonal/adapters"
	"testear/hexagonal/core"
)

/*
/hexagonal-example
  main.go
  /core
    order.go
    ports.go
  /adapters
    postgres_inventory.go
    http_notifier.go

*/

func main() {
	db, _ := sql.Open("postgres", "postgres://user:pass@localhost:5432/shop?sslmode=disable")

	// Create adapters
	inv := adapters.NewPostgresInventory(db)
	not := adapters.NewHTTPNotifier("https://email-service.local")

	// Inject dependencies into the core
	service := core.NewService(inv, not)

	order := core.Order{
		ID:       1,
		Customer: "Alice",
		ItemID:   "SKU-123",
		Email:    "alice@example.com",
	}

	if err := service.PlaceOrder(order); err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅ Order placed successfully!")
	}
}
