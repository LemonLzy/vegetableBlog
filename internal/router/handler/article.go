package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	errCode "github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/resp"
	"github.com/lemonlzy/vegetableBlog/internal/service"
	"github.com/lemonlzy/vegetableBlog/pkg"
	"net/http"
)

func ArticleCreateHandler(c *gin.Context) {
	a := &app.Article{}
	if err := c.ShouldBindJSON(a); err != nil {
		resp.ResponseError(c, errCode.NewClientError(errCode.ClientReqInvalid))
		return
	}

	err := service.CreateArticle(a)
	if err != nil {
		resp.ResponseError(c, err)
		return
	}

	resp.ResponseSuccess(c, nil)
}

func ArticleUpdateHandler(c *gin.Context) {
	// 获取文章id
	articleID := c.Param("id")
	a := new(app.Article)

	if err := c.ShouldBindJSON(a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数绑定失败",
		})
		return
	}

	err := app.UpdateArticleByID(articleID, a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}

func ArticleDetailHandler(c *gin.Context) {
	// 获取文章id
	articleID := c.Param("articleID")

	articleInfo, err := app.GetArticleByID(articleID)
	if err != nil {
		resp.ResponseError(c, errCode.NewClientError(errCode.ArticleInvalid))
		return
	}
	resp.ResponseSuccess(c, articleInfo)
}

func ArticleListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := pkg.GetPageInfo(c)
	// 获取列表数据
	articleList, total, err := app.GetArticleList(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取文章列表失败",
		})
		return
	}

	data := make(map[string]interface{}, 2)
	data["list"] = articleList
	data["total"] = total

	resp.ResponseSuccess(c, data)
}

func ArticlePubHandler(c *gin.Context) {
	a := &app.Article{}
	if err := c.ShouldBindJSON(a); err != nil {
		resp.ResponseError(c, errCode.NewClientError(errCode.ClientReqInvalid))
		return
	}

	err := service.PubArticle(a.ArticleID)
	if err != nil {
		resp.ResponseError(c, err)
		return
	}

	resp.ResponseSuccess(c, nil)
}
