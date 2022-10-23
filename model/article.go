package model

// CREATE TABLE IF NOT EXISTS `article` (
//     `id` BIGINT(20) NOT NULL PRIMARY KEY AUTO_INCREMENT comment '文章 id',
//     `category_id` BIGINT(20) NOT NULL comment '分類 id',
//     `content` LONGTEXT NOT NULL comment '文章內容',
//     `title` VARCHAR(1024) NOT NULL comment '文章標題',
//     `view_count` INT(255) UNSIGNED NOT NULL comment '閱讀次數',
//     `comment_count` INT(255) UNSIGNED NOT NULL comment '評論次數',
//     `username` VARCHAR(128) NOT NULL comment '作者',
//     `status` INT(10) DEFAULT 1 NOT NULL comment '文章狀態; 正常為 1;',
//     `summary` VARCHAR(256) NOT NULL comment '文章摘要',
//     `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment '發佈時間',
//     `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment '更新時間'
// ) CHARSET `utf8mb4`;


import (
	"time"
)

// 列表:
// * 沒有 content，所以不要去取，以提升效率
type ArticleInfo struct {
	Id int64 `db:"id"`
	CategoryId int64 `db:"category_id"`
	Title string `db:"title"`
	ViewCount uint32 `db:"view_count"`
	CommentCount uint32 `db:"comment_count"`
	Username string `db:"username"`
	Summary string `db:"summary"`
	CreateTime time.Time `db:"create_time"`
}



// 詳情頁
// * 詳情包含了文章內容跟分類
// * 這樣寫純粹是為了偷懶，如果要更高效的話，還是只把需要的列出來就好，例如詳情頁就用不到 Summary
type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}


// 文章上下頁
// * 其實只需要標題
// * 不過老師把 Info 都塞入了
type ArticleRecord struct {
	ArticleInfo
	Category
}
