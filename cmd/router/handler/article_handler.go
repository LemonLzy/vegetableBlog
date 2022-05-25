package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	"net/http"
	"strconv"
)

func CreateArticleHandler(c *gin.Context) {
	a := new(app.Article)

	if err := c.ShouldBindJSON(a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数绑定失败",
		})
		return
	}

	err := app.CreateArticle(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}

func ModifyArticleHandler(c *gin.Context) {

}

func ArticleDetailHandler(c *gin.Context) {
	// 获取文章id
	idStr := c.Param("id")
	articleID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文章ID获取失败",
		})
		return
	}

	articleInfo, err := app.GetArticleDetail(articleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询文章详情失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"data":    articleInfo,
	})
}

func ArticleListHandler(c *gin.Context) {

}
