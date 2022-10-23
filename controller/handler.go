package controller

import (
	"github.com/gin-gonic/gin"
	"blogger/service"
	"net/http"
	"fmt"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	records, err := service.GetArticleRecords(0, 10)
	if err != nil {
		fmt.Println("獲取文章數據失敗: ", err)

		c.HTML(http.StatusInternalServerError, "views/500.html", gin.H {
			"error": err,
		})
		return
	}

	categoryLists, err := service.GetAllCagoryList()
	if err != nil {
		fmt.Println("獲取分類數據失敗: ", err)

		c.HTML(http.StatusInternalServerError, "views/500.html", gin.H {
			"error": err,
		})
		return
	}


	c.HTML(http.StatusOK, "views/index.html", gin.H {
		"records": records,
		"categoryLists": categoryLists,
	})
}


func CategoryHandler(c *gin.Context) {
	queryStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(queryStr, 10, 64)
	if err != nil {
		fmt.Println("query 轉 int64 失敗: ", err)

		c.HTML(http.StatusInternalServerError, "views/500.html", gin.H {
			"error": err,
		})
		return
	}



	records, err := service.GetArticleRecordsByCategoryId(categoryId, 0, 10)
	if err != nil {
		fmt.Println("獲取文章數據失敗: ", err)

		c.HTML(http.StatusInternalServerError, "views/500.html", gin.H {
			"error": err,
		})
		return
	}

	categoryLists, err := service.GetAllCagoryList()
	if err != nil {
		fmt.Println("獲取分類數據失敗: ", err)

		c.HTML(http.StatusInternalServerError, "views/500.html", gin.H {
			"error": err,
		})
		return
	}


	c.HTML(http.StatusOK, "views/index.html", gin.H {
		"records": records,
		"categoryLists": categoryLists,
	})






}
