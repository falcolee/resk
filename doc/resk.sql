CREATE TABLE IF NOT EXISTS `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账户ID',
  `account_no` varchar(32) NOT NULL COMMENT '账户编号,账户唯一标识 ',
  `account_name` varchar(64) NOT NULL COMMENT '账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱',
  `account_type` tinyint(2) NOT NULL COMMENT '账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户',
  `currency_code` char(3) NOT NULL DEFAULT 'CNY' COMMENT '货币类型编码：CNY人民币，EUR欧元，USD美元 。。。',
  `user_id` varchar(40) NOT NULL COMMENT '用户编号, 账户所属用户 ',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名称',
  `balance` decimal(30,6) unsigned NOT NULL DEFAULT 0.000000 COMMENT '账户可用余额',
  `status` tinyint(2) NOT NULL COMMENT '账户状态，账户状态：0账户初始化，1启用，2停用 ',
  `created_at` datetime(3) NOT NULL DEFAULT current_timestamp(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account_no_idx` (`account_no`) USING BTREE,
  KEY `id_user_idx` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32938 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;


INSERT INTO `account` (`id`, `account_no`, `account_name`, `account_type`, `currency_code`, `user_id`, `username`, `balance`, `status`, `created_at`, `updated_at`) VALUES
	(32937, '10000020190101010000000000000001', '系统红包账户', 2, 'CNY', '100001', '系统红包账户', 0.000000, 1, '2019-05-01 08:41:10.346', '2019-05-12 09:37:55.462');

CREATE TABLE IF NOT EXISTS `account_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trade_no` varchar(32) NOT NULL COMMENT '交易单号 全局不重复字符或数字，唯一性标识 ',
  `log_no` varchar(32) NOT NULL COMMENT '流水编号 全局不重复字符或数字，唯一性标识 ',
  `account_no` varchar(32) NOT NULL COMMENT '账户编号 账户ID',
  `target_account_no` varchar(32) NOT NULL COMMENT '账户编号 账户ID',
  `user_id` varchar(40) NOT NULL COMMENT '用户编号',
  `username` varchar(64) NOT NULL COMMENT '用户名称',
  `target_user_id` varchar(40) NOT NULL COMMENT '目标用户编号',
  `target_username` varchar(64) NOT NULL COMMENT '目标用户名称',
  `amount` decimal(30,6) NOT NULL DEFAULT 0.000000 COMMENT '交易金额,该交易涉及的金额 ',
  `balance` decimal(30,6) unsigned NOT NULL DEFAULT 0.000000 COMMENT '交易后余额,该交易后的余额 ',
  `change_type` tinyint(2) NOT NULL DEFAULT 0 COMMENT '流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义',
  `change_flag` tinyint(2) NOT NULL DEFAULT 0 COMMENT '交易变化标识：-1 出账 1为进账，枚举',
  `status` tinyint(2) NOT NULL DEFAULT 0 COMMENT '交易状态：',
  `decs` varchar(128) NOT NULL COMMENT '交易描述 ',
  `created_at` datetime(3) NOT NULL DEFAULT current_timestamp(3) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `id_log_no_idx` (`log_no`) USING BTREE,
  KEY `id_user_idx` (`user_id`) USING BTREE,
  KEY `id_account_idx` (`account_no`) USING BTREE,
  KEY `id_trade_idx` (`trade_no`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=43209 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

INSERT INTO `account_log` (`id`, `trade_no`, `log_no`, `account_no`, `target_account_no`, `user_id`, `username`, `target_user_id`, `target_username`, `amount`, `balance`, `change_type`, `change_flag`, `status`, `decs`, `created_at`) VALUES
	(43208, '20190501084054283000000002110000', '20190501084054283000000002110000', '10000020190101010000000000000001', '10000020190101010000000000000001', '100001', '系统红包账户', '100001', '系统红包账户', 0.000000, 0.000000, 0, 0, 0, '开户', '2019-05-01 08:41:10.371');