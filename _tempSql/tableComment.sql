CREATE TABLE IF NOT EXISTS `comment` (
    `id` BIGINT(20) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '分類 id',
    `content` TEXT NOT NULL COMMENT '評論內容',
    `username` VARCHAR(64) NOT NULL COMMENT '評論者',
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '發佈時間',
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '更新時間',
    `status` INT(255) DEFAULT 1 NOT NULL COMMENT '評論狀態: 0; 刪除: 1',
    `article_id` BIGINT(20) NOT NULL comment '被評論文章 id'
) CHARSET `utf8mb4`;
