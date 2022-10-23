package db

import (
	"blogger/model"
	"github.com/jmoiron/sqlx"
	"fmt"
)

func InsertCategory(m *model.Category) (int64, error) {
	sql := `
		INSERT INTO
			category(category_name, category_no)
			VALUES(?,?)
	`

	ret, err := DB.Exec(sql, m.CategoryName, m.CategoryNo)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%T\n", ret)

	return ret.LastInsertId()
}


func GetCategoryById(id int64) (category *model.Category, err error) {
	category =  &model.Category {}
	sql := `
		SELECT
			id, category_name, category_no
		FROM
			category
		WHERE
			id = ?
	`

	err = DB.Get(category, sql, id)
	return
}

func GetCategorysByIds(ids []int64)(categorys []*model.Category, err error)  {
	sqlstr := `
		SELECT
			id, category_name, category_no
		FROM
			category
		WHERE
			id IN (?)
	`

	var args []interface{}
	sqlstr, args, err = sqlx.In(sqlstr, ids)
	if err != nil {
		return
	}

	err = DB.Select(&categorys, sqlstr, args...)
	return
}


func GetCategorysAll()(categorys []*model.Category, err error) {
	sqlstr := `
		SELECT
			id, category_name, category_no
		FROM
			category
		ORDER BY
			category_no asc
	`

	err = DB.Select(&categorys, sqlstr)
	return

}
