package repositories

// import (
// 	"log"

// 	"gopkg.in/DATA-DOG/go-sqlmock.v1"
// )

// func TestUserFetch() {
// 	db, mock, err := sqlmock.New()

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
// 		AddRow(1, "Test", "John").
// 		AddRow(2, "Test", "Ann")

// 	mock.ExpectQuery("SELECT id, first_name, last_name FROM users").WillReturnRows(rows)

// 	userRepository := NewUserRepository(db)
// }
