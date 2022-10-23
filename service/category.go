package service


import (
	"blogger/model"
	"blogger/dao/mysql"
)

func GetAllCagoryList()(category []*model.Category, err error) {

	category, err = db.GetCategorysAll()
	return
}
