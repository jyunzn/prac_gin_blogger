package db

import (
	"blogger/model"
	// "github.com/jmoiron/sqlx"
	"fmt"
	"errors"
)

func InsertArticle(article *model.ArticleDetail) (int64, error) {

	if article == nil {
		return 0, errors.New("article 不得為空")
	}

	sql := `
		INSERT INTO
			article(category_id, title, view_count, comment_count, username, summary, content)
			VALUES(?, ?, ?, ?, ?, ?, ?)
	`

	ret, err := DB.Exec(sql,
		article.CategoryId,
		article.Title,
		article.ViewCount,
		article.CommentCount,
		article.Username,
		article.Summary,
		article.Content,
	)

	fmt.Printf("%T\n", ret)

	if err != nil {
		return 0, err
	}

	return ret.LastInsertId()
}


func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	sql := `
		SELECT
			id, category_id, title, view_count, comment_count, username, summary, create_time
		FROM
			article
		WHERE
			status = 1
		ORDER BY
			create_time desc
		LIMIT
			?,?
	`

	err = DB.Select(&articleList, sql, pageNum, pageSize)
	return
}


func GetArticleById(id int64) (article *model.ArticleDetail, err error) {
	sql := `
		SELECT
			id, category_id, title, view_count, comment_count, username, summary, create_time, content
		FROM
			article
		WHERE
			id = ?
		AND
			status = 1
	`


	err = DB.Get(&article, sql, id)
	return
}


// 點擊分類時，查詢該分類的文章列表
func GetArticleListByCategoryId(categoryId int64, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	sql := `
		SELECT
			id, category_id, title, view_count, comment_count, username, summary, create_time
		FROM
			article
		WHERE
			status = 1
		AND
			category_id = ?
		ORDER BY
			create_time desc
		LIMIT
			?,?
	`

	err = DB.Select(&articleList, sql, categoryId, pageNum, pageSize)
	return
}
