package commentRepository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Higins/go_blog2/domain"
)

func TestSave(t *testing.T) {

	var mock sqlmock.Sqlmock

	const sqlInsert = `INSERT INTO "comments" ("text","blogid") VALUES ($1,$2) RETURNING "comments"."id"`
	commentInsert := domain.Comment{
		Text:   "test",
		BlogId: 1,
	}
	newId := 1
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
	WithArgs(commentInsert.Text, commentInsert.BlogId).
	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit()
}
