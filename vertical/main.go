package vertical

import (
	"database/sql"
	"log"
	"net/http"
	"testear/vertical/create"

	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "POSTGRES_URL")

	// crea el slice
	store := create.NewStore(db)
	svc := create.NewService(store)
	handler := create.NewHandler(svc)

	http.Handle("/orders/create", handler)

	log.Println("listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
