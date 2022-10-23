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

func TestInsertCategory(t *testing.T) {
	model := &model.Category {
		CategoryName: "javascript",
		CategoryNo: 2,
	}

	categoryID, err := InsertCategory(model)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("categoryID: %v\n", categoryID)
}


func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", category)
}


func TestGetCategorysByIds(t *testing.T) {
	categorys, err := GetCategorysByIds([] int64 { 1, 2 })
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range categorys {
		t.Logf("%#v", v)
	}

}

func TestGetCategorysAll(t *testing.T) {
	categorys, err := GetCategorysAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range categorys {
		t.Logf("%#v", v)
	}

}
