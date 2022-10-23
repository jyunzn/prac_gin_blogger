package service

import (
	"blogger/dao/mysql"
	"testing"
)

func init() {
	err := db.Init("root:youknow.@tcp(127.0.0.1:3306)/blogger?parseTime=true")
	if err != nil {
		panic(err)
		return
	}
}

func TestGetArticleRecords(t *testing.T) {
	articleRecords, err := GetArticleRecords(0, 100)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, articleRecord := range articleRecords {
		t.Logf("%#v\n", articleRecord)
	}
}



func TestGetArticleRecordsByCategoryId(t *testing.T) {
	articleRecords, err := GetArticleRecordsByCategoryId(1, 0, 100)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, articleRecord := range articleRecords {
		t.Logf("%#v", articleRecord)
	}
}
