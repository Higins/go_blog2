package blogRespository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Higins/go_blog2/domain"
)

func TestSave(t *testing.T) {
	var mock sqlmock.Sqlmock

	blogInsert := &domain.Blog{
		Title: "test",
		Body:  "b test",
	}

	const sqlInsert = `INSERT INTO "blogs" ("title","body") VALUES ($1,$2) RETURNING "blogs"."id"`
	const newId = 1
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(blogInsert.Title, blogInsert.Body).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit()
}

func TestFindAll(t *testing.T) {
	var mock sqlmock.Sqlmock
	const sqlSelectAll = `SELECT * FROM "blogs"`
	mock.ExpectQuery(sqlSelectAll).WillReturnRows(sqlmock.NewRows(nil))
}
