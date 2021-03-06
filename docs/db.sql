-- 交易流水表
CREATE TABLE strategy_logs (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name` varchar(255) NOT NULL COMMENT '策略名称',
    `number` decimal(20,10) NOT NULL COMMENT '币数',
    `amount` decimal(20,10) NOT NULL COMMENT '总金额',
    `percentage` decimal(4,2) NOT NULL COMMENT '每次补仓的金额比',
    `percentage_increase` decimal(6,2) NOT NULL COMMENT '价格加仓百分比',
    `percentage_drop` decimal(6,2) NOT NULL COMMENT '价格减仓百分比',
    `set_price` decimal(20,10) NOT NULL COMMENT '设置价格',
    `final_price` decimal(20,10) NOT NULL COMMENT '最终成交价',
    `final_type` tinyint(1) NOT NULL COMMENT '成交方式 1:加仓 2:减仓',
    `final_date` varchar(32) NOT NULL COMMENT '添加时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;