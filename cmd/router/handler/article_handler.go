package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	"net/http"
)

func CreateArticleHandler(c *gin.Context) {
	a := new(app.Article)

	if err := c.ShouldBindJSON(a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数绑定失败",
		})
	}

	err := app.CreateArticle(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}

func ArticleDetailHandler(c *gin.Context) {

}

func ArticleListHandler(c *gin.Context) {

}
