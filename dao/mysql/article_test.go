package db

import (
	"testing"
	"blogger/model"
)


func init() {
	err := Init("root:youknow.@tcp(127.0.0.1:3306)/blogger?parseTime=true")
	if err != nil {
		panic(err)
		return
	}
}

// category_id, title, view_count, comment_count, username, summary, content
func TestInsertArticle(t *testing.T) {


	article := &model.ArticleDetail {
		ArticleInfo: model.ArticleInfo {
			CategoryId: 2,
			Title: "aaaa",
			ViewCount: 0,
			CommentCount: 0,
			Username: "bebe",
			Summary: "aaaa",
		},
		Content: "aaaaaaaaa",
	}

	articleId, err := InsertArticle(article)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", articleId)
}


func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(0, 2)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range articleList {
		t.Logf("%#v", v)
	}
}

func TestGetArticleById(t *testing.T) {
	article, err := GetArticleById(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", article)
}


func TestGetArticleListByCategoryId(t *testing.T) {
	articleList, err := GetArticleListByCategoryId(2, 0, 2)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range articleList {
		t.Logf("%#v", v)
	}
}
