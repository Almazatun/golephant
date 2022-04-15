package test

import (
	"database/sql"
	"regexp"
	"time"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pg", func() {
	var repo repository.UserRepo
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp)) // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err := gorm.Open("postgres", db) // open gorm db
		Expect(err).ShouldNot(HaveOccurred())

		repo = repository.NewUserRepo(gdb)
	})

	Context("Register", func() {
		var userDB entity.User

		rows := sqlmock.NewRows([]string{
			"username", "email", "password",
		}).AddRow(userDB.Username, userDB.Email, userDB.Password)

		BeforeEach(func() {
			userDB = entity.User{
				Email:        "test@mail.com",
				Password:     "1234567",
				Username:     "test",
				CreationTime: time.Now(),
				UpdateTime:   time.Now(),
			}
		})

		It("Create", func() {
			mock.ExpectBegin()

			const sqlInsert = `
					INSERT INTO "users" ("username","email","password")
					VALUES ($1,$2,$3) RETURNING "users"."user_id"`
			const user_id = 1
			mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
				WithArgs(userDB.UserID, userDB.Email, userDB.Password).
				WillReturnRows(rows)
			mock.ExpectCommit() // commit transaction
			Expect(userDB.Email).Should(Equal("test@mail.com"))

			_, err := repo.Create(userDB)
			Expect(err).ShouldNot(HaveOccurred())

		})

	})

})
