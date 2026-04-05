package store

import (
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/jackc/pgx/v5/stdlib"

	"testing"
)

const connURI = "postgres://user:password@localhost:5555/mydb?sslmode=disable"

func init() {
	txdb.Register("txdb", "pgx", connURI)
}

func TestGetUserByID(t *testing.T) {
	// 1. Conectar a la base de datos
	db, err := sql.Open("txdb", connURI)
	if err != nil {
		t.Fatal("no abri la conexion")
	}
	defer db.Close()

	// 2. Insertar informacion para test
	if _, err := db.Exec(`INSERT INTO users (name) VALUES ('Alice'), ('Bob')`); err != nil {
		t.Fatal("could not insert test data:", err)
	}

	// 3. Conectar a store
	store := New(db)
	got, err := store.GetUserByID(1)
	if err != nil {
		t.Fatal("no encontre el user", err)
	}

	// 4. Hacer test
	want := "Alice"
	if got != want {
		t.Fatalf("espero a Alice, pero got %s", got)
	}
}
