package test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_Register(t *testing.T) {
	var db *sql.DB
	var err error
	var mock sqlmock.Sqlmock
	var repo repository.UserRepo

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp)) // mock sql.DB

	if err != nil {
		panic(err)
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "test",
		DriverName:           "sqlite3",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gdb, err := gorm.Open(dialector, &gorm.Config{}) // open gorm db

	if err != nil {
		panic(err)
	}

	repo = repository.NewUserRepo(gdb)

	defer db.Close()

	user := entity.User{
		Email:        "test@mail.com",
		Password:     "1234567",
		Username:     "test",
		CreationTime: time.Now(),
		UpdateTime:   time.Now(),
	}

	mock.ExpectBegin()

	const sqlInsert = `
					INSERT INTO "users" ("username","email","password")
					VALUES ($1,$2,$3) RETURNING "users"."user_id"`

	mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(user.UserID, user.Email, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(user.UserID))
	mock.ExpectCommit() // commit transaction

	res, err := repo.Create(user)

	if err != nil {
		fmt.Println(res.UserID)
		panic(err)
	}
}
