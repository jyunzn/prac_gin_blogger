package service

import (
	"blogger/model"
	"blogger/dao/mysql"
)


func getArticleCategoryId(articleList []*model.ArticleInfo) (ids []int64) {

	// 循環去重複 ( ps. golang 沒有 dir 實在是很麻煩 )
	for _, articleInfo := range articleList {
		flag := false
		for _, id := range ids {

			if id == articleInfo.CategoryId {
				flag = true
				break
			}
		}

		if !flag {
			ids = append(ids, articleInfo.CategoryId)
		}

	}

	return

}


func GetArticleRecords(pageNum, pageSize int)(articleRecords []*model.ArticleRecord, err error) {
	// 獲取所有 articleInfo
	var articleList []*model.ArticleInfo
	articleList, err = db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}



	// 獲取 articleInfo 對應的 category 數據
	var categorys []*model.Category
	categoryIds := getArticleCategoryId(articleList)
	categorys, err = db.GetCategorysByIds(categoryIds)
	if err != nil {
		return
	}


	for _, articleInfo := range articleList {
		articleRecord := &model.ArticleRecord {}
		articleRecord.ArticleInfo = *articleInfo
		for _, category := range categorys {
			if category.Id == articleInfo.CategoryId {
				articleRecord.Category = *category
				break
			}
		}

		articleRecords = append(articleRecords, articleRecord)
	}

	return
}


func GetArticleRecordsByCategoryId(categoryId int64, pageNum, pageSize int)(articleRecords []*model.ArticleRecord, err error) {

	// 獲取所有 articleInfo
	var articleList []*model.ArticleInfo
	articleList, err = db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}



	// 獲取 articleInfo 對應的 category 數據
	var categorys []*model.Category
	categoryIds := getArticleCategoryId(articleList)
	categorys, err = db.GetCategorysByIds(categoryIds)
	if err != nil {
		return
	}


	for _, articleInfo := range articleList {
		articleRecord := &model.ArticleRecord {}
		articleRecord.ArticleInfo = *articleInfo
		for _, category := range categorys {
			if category.Id == articleInfo.CategoryId {
				articleRecord.Category = *category
				break
			}
		}

		articleRecords = append(articleRecords, articleRecord)
	}

	return
}

