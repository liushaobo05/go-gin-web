
# 用户表
CREATE TABLE `tb_main_user` (
  `seq` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `id` char(65) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `username` varchar(64) NOT NULL,
  `passwordHash` text,
  `name` varchar(256) DEFAULT NULL,
  `mobile` varchar(32) DEFAULT NULL,
  `markers` text,
  `roles` text,
  `privileges` text,
  `isDisabled` tinyint(1) NOT NULL DEFAULT '0',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`seq`),
  UNIQUE KEY `ID` (`id`),
  UNIQUE KEY `USERNAME` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

# 初始化数据
LOCK TABLES `tb_main_user` WRITE;
/*!40000 ALTER TABLE `tb_main_user` DISABLE KEYS */;

INSERT INTO `tb_main_user` (`seq`, `id`, `username`, `passwordHash`, `name`, `mobile`, `markers`, `roles`, `customPrivileges`, `isDisabled`, `createTime`, `updateTime`)
VALUES
  (1,X'000001','admin','admin','系统管理员',NULL,NULL,'sa','*',0,'2019-04-20 10:08:03','2019-04-20 16:10:09'),
/*!40000 ALTER TABLE `tb_main_user` ENABLE KEYS */;
UNLOCK TABLES;

# ak
DROP TABLE IF EXISTS `tb_main_app_key`;

CREATE TABLE `tb_main_app_key` (
  `seq` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `id` char(65) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `userId` char(65) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `name` varchar(256) NOT NULL,
  `secret` varchar(64) NOT NULL,
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`seq`),
  UNIQUE KEY `ID` (`id`),
  KEY `USER_ID` (`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

# 钉钉发送日志
DROP TABLE IF EXISTS `biz_log_result_ding_talk`;

CREATE TABLE `biz_log_result_ding_talk` (
  `seq` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `id` char(65) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `task` varchar(128) NOT NULL DEFAULT '' COMMENT '通知类型',
  `startTime` int(11) DEFAULT NULL COMMENT '任务开始时间(秒级UNIX时间戳)',
  `endTime` int(11) DEFAULT NULL COMMENT '任务结束时间(秒级UNIX时间戳)',
  `msgType` varchar(20) DEFAULT NULL COMMENT '消息类型',
  `argsJSON` text COMMENT '数组参数JSON',
  `paramsJSON` text COMMENT 'map参数JSON',
  `retvalJSON` text COMMENT '执行结果JSON',
  `status` varchar(64) DEFAULT '' COMMENT '任务状态: SUCCESS|FAILURE',
  `einfoTEXT` text COMMENT '错误信息TEXT',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`seq`),
  UNIQUE KEY `ID` (`id`),
  KEY `task` (`task`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='钉钉机器人任务结果';