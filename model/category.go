package model

// CREATE TABLE IF NOT EXISTS `category` (
//     `id` BIGINT(20) NOT NULL PRIMARY KEY AUTO_INCREMENT comment '分類 id',
//     `category_name` VARCHAR(255) NOT NULL comment '分類名字',
//     `category_no` INT(10) NOT NULL comment '分類排序',
//     `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment '發佈時間',
//     `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment '更新時間'
// ) CHARSET `utf8mb4`;

// import (
// 	"time"
// )

type Category struct {
	Id int64 `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo int32 `db:"category_no"`
}
