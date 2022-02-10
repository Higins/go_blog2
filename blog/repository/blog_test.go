package blogRespository

import (
	"fmt"
	"testing"

	"github.com/Higins/go_blog2/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type blogRepositoryTestSuite struct {
	suite.Suite
	db    *gorm.DB
	badDB *gorm.DB
}

func (s *blogRepositoryTestSuite) SetupTest() {
	var err error
	s.db, err = gorm.Open(sqlite.Open("../../blog.db"), &gorm.Config{})
	err = s.db.AutoMigrate(domain.Blog{})
	defer s.db.Migrator().DropTable(domain.Blog{})

	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

	s.badDB, err = gorm.Open(sqlite.Open("../../blogBAD.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

}
func (s *blogRepositoryTestSuite) TestSaveBadDB(t *testing.T) {
	blogRepo := NewBlogRepository(s.badDB)
	post := domain.Blog{
		Title: "test",
		Body:  "body",
	}
	blog, err := blogRepo.Save(post)

	require.Nil(t, err)
	require.Empty(t, blog)

}
func (s *blogRepositoryTestSuite) TestSaveGoodDB(t *testing.T) {
	blogRepo := NewBlogRepository(s.db)
	post := domain.Blog{
		Title: "test",
		Body:  "body",
	}
	blog, err := blogRepo.Save(post)

	require.NoError(t, err)
	assert.Equal(t, "test", blog.Title)
	assert.Equal(t, "body", blog.Body)
	assert.NotZero(t, blog.ID)

}

func (s *blogRepositoryTestSuite) TestFindAll(t *testing.T) {
	blogRepo := NewBlogRepository(s.db)
	posts := []domain.Blog{
		{Title: "Test1", Body: "Body1"},
		{Title: "Test2", Body: "Body2"},
		{Title: "Test3", Body: "Body3"},
	}

	for _, post := range posts {
		s.db.Create(&post)
	}

	blog, err := blogRepo.FindAll()
	require.NoError(t, err)
	assert.Equal(t, blog, posts)

}

func (s *blogRepositoryTestSuite) TestGetBlogById(t *testing.T) {
	blogRepo := NewBlogRepository(s.db)
	post := domain.Blog{
		Title: "test getblogidby",
		Body:  "body",
	}
	blog, err := blogRepo.Save(post)
	require.NoError(t, err)
	assert.Equal(t, "test getblogidby", blog.Title)
	assert.Equal(t, "body", blog.Body)
	assert.NotZero(t, blog.ID)
}
