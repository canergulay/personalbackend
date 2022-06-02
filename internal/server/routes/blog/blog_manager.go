package blog

import (
	"encoding/json"
	"io"
	"net/http"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BlogManager struct {
	DB *gorm.DB
}

func NewBlogManager(db *gorm.DB) *BlogManager {
	return &BlogManager{DB: db}
}

func (h BlogManager) Handler(c echo.Context) error {

	body := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != io.EOF {
		c.String(500, "an unexpected error has occured"+err.Error()+"\n")
	}

	lastID := body["lastID"]

	lastIDParsed, ok := lastID.(int)

	var posts []Post

	if ok {
		posts = h.GetPosts(lastIDParsed)
	} else {
		posts = h.GetPosts(-1)
	}

	c.JSON(200, posts)

	return c.String(http.StatusOK, "Hello, World!")
}

func (h BlogManager) GetBlogEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/blog", h.Handler, "GET")
}

func (bm BlogManager) CreateBlog() int {
	post := Post{}
	bm.DB.Create(&post)
	return post.ID
}

func (bm BlogManager) SavePost(post *Post) *Post {
	bm.DB.Save(post)
	return post
}

func (bm BlogManager) GetPosts(lastID int) []Post {
	var posts []Post

	// IF THE LAST ID IS NOT SPECIFIED, IT WILL BE PASSED DOWN AS -1
	// WE WILL RETURN THE FIRST 10, WHICH REPRESENTS THE FIRST PAGE
	// IF IT IS SPECIFIED, IT MEANS WE ARE PAGINATING THROUGH A SPECIFIG PAGE.
	if lastID == -1 {
		bm.DB.Model(Post{}).Limit(10).Order("id desc").Find(&posts)
	} else {
		bm.DB.Model(Post{}).Where("id >", lastID).Limit(10).Order("id desc").Find(&posts)
	}
	return posts
}
