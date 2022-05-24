package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/app"
)

func CreateArticleHandler(c *gin.Context) {
	err := app.CreateArticle()
	if err != nil {

	}
}

func ArticleDetailHandler(c *gin.Context) {

}

func ArticleListHandler(c *gin.Context) {

}
