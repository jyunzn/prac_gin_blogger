package service

import (
	"testing"
	"blogger/dao/mysql"
)

func init() {
	err := db.Init("root:youknow.@tcp(127.0.0.1:3306)/blogger?parseTime=true")
	if err != nil {
		panic(err)
		return
	}
}

func TestGetAllCagoryList(t *testing.T) {
	categorys, err := GetAllCagoryList()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range categorys {
		t.Logf("%#v", v)
	}
}
