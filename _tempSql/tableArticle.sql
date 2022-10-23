CREATE TABLE IF NOT EXISTS `article` (
    `id` BIGINT(20) NOT NULL PRIMARY KEY AUTO_INCREMENT comment '文章 id',
    `category_id` BIGINT(20) NOT NULL comment '分類 id',
    `content` LONGTEXT NOT NULL comment '文章內容',
    `title` VARCHAR(1024) NOT NULL comment '文章標題',
    `view_count` INT(255) UNSIGNED NOT NULL comment '閱讀次數',
    `comment_count` INT(255) UNSIGNED NOT NULL comment '評論次數',
    `username` VARCHAR(128) NOT NULL comment '作者',
    `status` INT(10) DEFAULT 1 NOT NULL comment '文章狀態; 正常為 1;',
    `summary` VARCHAR(256) NOT NULL comment '文章摘要',
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment '發佈時間',
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment '更新時間'
) CHARSET `utf8mb4`;
