mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 9.6.0, for macos26.4 (arm64)
--
-- Host: localhost    Database: gva
-- ------------------------------------------------------
-- Server version	9.6.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=303 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (9,'p','888','/api/createApi','POST','','',''),(12,'p','888','/api/deleteApi','POST','','',''),(15,'p','888','/api/deleteApisByIds','DELETE','','',''),(18,'p','888','/api/enterSyncApi','POST','','',''),(296,'p','888','/api/freshCasbin','GET',NULL,NULL,NULL),(14,'p','888','/api/getAllApis','POST','','',''),(11,'p','888','/api/getApiById','POST','','',''),(17,'p','888','/api/getApiGroups','GET','','',''),(10,'p','888','/api/getApiList','POST','','',''),(20,'p','888','/api/getApiRoles','GET','','',''),(19,'p','888','/api/ignoreApi','POST','','',''),(21,'p','888','/api/setApiRoles','POST','','',''),(16,'p','888','/api/syncApi','GET','','',''),(13,'p','888','/api/updateApi','POST','','',''),(183,'p','888','/attachmentCategory/addCategory','POST','','',''),(184,'p','888','/attachmentCategory/deleteCategory','POST','','',''),(182,'p','888','/attachmentCategory/getCategoryList','GET','','',''),(22,'p','888','/authority/copyAuthority','POST','','',''),(24,'p','888','/authority/createAuthority','POST','','',''),(25,'p','888','/authority/deleteAuthority','POST','','',''),(26,'p','888','/authority/getAuthorityList','POST','','',''),(28,'p','888','/authority/getUsersByAuthority','GET','','',''),(27,'p','888','/authority/setDataAuthority','POST','','',''),(29,'p','888','/authority/setRoleUsers','POST','','',''),(23,'p','888','/authority/updateAuthority','PUT','','',''),(151,'p','888','/authorityBtn/canRemoveAuthorityBtn','POST','','',''),(150,'p','888','/authorityBtn/getAuthorityBtn','POST','','',''),(149,'p','888','/authorityBtn/setAuthorityBtn','POST','','',''),(61,'p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),(60,'p','888','/casbin/updateCasbin','POST','','',''),(144,'p','888','/email/emailTest','POST','','',''),(145,'p','888','/email/sendEmail','POST','','',''),(56,'p','888','/fileUploadAndDownload/deleteFile','POST','','',''),(57,'p','888','/fileUploadAndDownload/editFileName','POST','','',''),(58,'p','888','/fileUploadAndDownload/getFileList','POST','','',''),(59,'p','888','/fileUploadAndDownload/importURL','POST','','',''),(55,'p','888','/fileUploadAndDownload/upload','POST','','',''),(62,'p','888','/jwt/jsonInBlacklist','POST','','',''),(293,'p','888','/mall/member/detail','POST',NULL,NULL,NULL),(292,'p','888','/mall/member/list','POST',NULL,NULL,NULL),(32,'p','888','/menu/addBaseMenu','POST','','',''),(34,'p','888','/menu/addMenuAuthority','POST','','',''),(38,'p','888','/menu/deleteBaseMenu','POST','','',''),(40,'p','888','/menu/getBaseMenuById','POST','','',''),(33,'p','888','/menu/getBaseMenuTree','POST','','',''),(30,'p','888','/menu/getMenu','POST','','',''),(35,'p','888','/menu/getMenuAuthority','POST','','',''),(31,'p','888','/menu/getMenuList','POST','','',''),(36,'p','888','/menu/getMenuRoles','GET','','',''),(37,'p','888','/menu/setMenuRoles','POST','','',''),(39,'p','888','/menu/updateBaseMenu','POST','','',''),(147,'p','888','/simpleUploader/checkFileMd5','GET','','',''),(148,'p','888','/simpleUploader/mergeFileMd5','GET','','',''),(146,'p','888','/simpleUploader/upload','POST','','',''),(6,'p','888','/sysApiToken/createApiToken','POST','','',''),(8,'p','888','/sysApiToken/deleteApiToken','POST','','',''),(7,'p','888','/sysApiToken/getApiTokenList','POST','','',''),(134,'p','888','/sysDictionary/createSysDictionary','POST','','',''),(135,'p','888','/sysDictionary/deleteSysDictionary','DELETE','','',''),(137,'p','888','/sysDictionary/exportSysDictionary','GET','','',''),(131,'p','888','/sysDictionary/findSysDictionary','GET','','',''),(133,'p','888','/sysDictionary/getSysDictionaryList','GET','','',''),(136,'p','888','/sysDictionary/importSysDictionary','POST','','',''),(132,'p','888','/sysDictionary/updateSysDictionary','PUT','','',''),(124,'p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),(126,'p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),(122,'p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),(129,'p','888','/sysDictionaryDetail/getDictionaryDetailsByParent','GET','','',''),(130,'p','888','/sysDictionaryDetail/getDictionaryPath','GET','','',''),(127,'p','888','/sysDictionaryDetail/getDictionaryTreeList','GET','','',''),(128,'p','888','/sysDictionaryDetail/getDictionaryTreeListByType','GET','','',''),(125,'p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),(123,'p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),(162,'p','888','/sysError/createSysError','POST','','',''),(163,'p','888','/sysError/deleteSysError','DELETE','','',''),(164,'p','888','/sysError/deleteSysErrorByIds','DELETE','','',''),(166,'p','888','/sysError/findSysError','GET','','',''),(167,'p','888','/sysError/getSysErrorList','GET','','',''),(168,'p','888','/sysError/getSysErrorSolution','GET','','',''),(165,'p','888','/sysError/updateSysError','PUT','','',''),(2,'p','888','/sysLoginLog/deleteLoginLog','DELETE','','',''),(3,'p','888','/sysLoginLog/deleteLoginLogByIds','DELETE','','',''),(4,'p','888','/sysLoginLog/findLoginLog','GET','','',''),(5,'p','888','/sysLoginLog/getLoginLogList','GET','','',''),(140,'p','888','/sysOperationRecord/createSysOperationRecord','POST','','',''),(142,'p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),(143,'p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),(138,'p','888','/sysOperationRecord/findSysOperationRecord','GET','','',''),(141,'p','888','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(139,'p','888','/sysOperationRecord/updateSysOperationRecord','PUT','','',''),(175,'p','888','/sysParams/createSysParams','POST','','',''),(176,'p','888','/sysParams/deleteSysParams','DELETE','','',''),(177,'p','888','/sysParams/deleteSysParamsByIds','DELETE','','',''),(179,'p','888','/sysParams/findSysParams','GET','','',''),(181,'p','888','/sysParams/getSysParam','GET','','',''),(180,'p','888','/sysParams/getSysParamsList','GET','','',''),(178,'p','888','/sysParams/updateSysParams','PUT','','',''),(65,'p','888','/system/getServerInfo','POST','','',''),(63,'p','888','/system/getSystemConfig','POST','','',''),(64,'p','888','/system/setSystemConfig','POST','','',''),(190,'p','888','/sysVersion/deleteSysVersion','DELETE','','',''),(191,'p','888','/sysVersion/deleteSysVersionByIds','DELETE','','',''),(187,'p','888','/sysVersion/downloadVersionJson','GET','','',''),(188,'p','888','/sysVersion/exportVersion','POST','','',''),(185,'p','888','/sysVersion/findSysVersion','GET','','',''),(186,'p','888','/sysVersion/getSysVersionList','GET','','',''),(189,'p','888','/sysVersion/importVersion','POST','','',''),(1,'p','888','/user/admin_register','POST','','',''),(46,'p','888','/user/changePassword','POST','','',''),(45,'p','888','/user/deleteUser','DELETE','','',''),(41,'p','888','/user/getUserInfo','GET','','',''),(44,'p','888','/user/getUserList','POST','','',''),(49,'p','888','/user/resetPassword','POST','','',''),(43,'p','888','/user/setSelfInfo','PUT','','',''),(50,'p','888','/user/setSelfSetting','PUT','','',''),(48,'p','888','/user/setUserAuthorities','POST','','',''),(47,'p','888','/user/setUserAuthority','POST','','',''),(42,'p','888','/user/setUserInfo','PUT','','',''),(193,'p','8881','/api/createApi','POST','','',''),(196,'p','8881','/api/deleteApi','POST','','',''),(297,'p','8881','/api/freshCasbin','GET',NULL,NULL,NULL),(198,'p','8881','/api/getAllApis','POST','','',''),(195,'p','8881','/api/getApiById','POST','','',''),(194,'p','8881','/api/getApiList','POST','','',''),(199,'p','8881','/api/getApiRoles','GET','','',''),(200,'p','8881','/api/setApiRoles','POST','','',''),(197,'p','8881','/api/updateApi','POST','','',''),(201,'p','8881','/authority/createAuthority','POST','','',''),(202,'p','8881','/authority/deleteAuthority','POST','','',''),(203,'p','8881','/authority/getAuthorityList','POST','','',''),(205,'p','8881','/authority/getUsersByAuthority','GET','','',''),(204,'p','8881','/authority/setDataAuthority','POST','','',''),(206,'p','8881','/authority/setRoleUsers','POST','','',''),(227,'p','8881','/casbin/getPolicyPathByAuthorityId','POST','','',''),(226,'p','8881','/casbin/updateCasbin','POST','','',''),(223,'p','8881','/fileUploadAndDownload/deleteFile','POST','','',''),(224,'p','8881','/fileUploadAndDownload/editFileName','POST','','',''),(222,'p','8881','/fileUploadAndDownload/getFileList','POST','','',''),(225,'p','8881','/fileUploadAndDownload/importURL','POST','','',''),(221,'p','8881','/fileUploadAndDownload/upload','POST','','',''),(228,'p','8881','/jwt/jsonInBlacklist','POST','','',''),(295,'p','8881','/mall/member/detail','POST',NULL,NULL,NULL),(294,'p','8881','/mall/member/list','POST',NULL,NULL,NULL),(209,'p','8881','/menu/addBaseMenu','POST','','',''),(211,'p','8881','/menu/addMenuAuthority','POST','','',''),(215,'p','8881','/menu/deleteBaseMenu','POST','','',''),(217,'p','8881','/menu/getBaseMenuById','POST','','',''),(210,'p','8881','/menu/getBaseMenuTree','POST','','',''),(207,'p','8881','/menu/getMenu','POST','','',''),(212,'p','8881','/menu/getMenuAuthority','POST','','',''),(208,'p','8881','/menu/getMenuList','POST','','',''),(213,'p','8881','/menu/getMenuRoles','GET','','',''),(214,'p','8881','/menu/setMenuRoles','POST','','',''),(216,'p','8881','/menu/updateBaseMenu','POST','','',''),(229,'p','8881','/system/getSystemConfig','POST','','',''),(230,'p','8881','/system/setSystemConfig','POST','','',''),(192,'p','8881','/user/admin_register','POST','','',''),(218,'p','8881','/user/changePassword','POST','','',''),(236,'p','8881','/user/getUserInfo','GET','','',''),(219,'p','8881','/user/getUserList','POST','','',''),(220,'p','8881','/user/setUserAuthority','POST','','',''),(238,'p','9528','/api/createApi','POST','','',''),(241,'p','9528','/api/deleteApi','POST','','',''),(298,'p','9528','/api/freshCasbin','GET',NULL,NULL,NULL),(243,'p','9528','/api/getAllApis','POST','','',''),(240,'p','9528','/api/getApiById','POST','','',''),(239,'p','9528','/api/getApiList','POST','','',''),(244,'p','9528','/api/getApiRoles','GET','','',''),(245,'p','9528','/api/setApiRoles','POST','','',''),(242,'p','9528','/api/updateApi','POST','','',''),(246,'p','9528','/authority/createAuthority','POST','','',''),(247,'p','9528','/authority/deleteAuthority','POST','','',''),(248,'p','9528','/authority/getAuthorityList','POST','','',''),(250,'p','9528','/authority/getUsersByAuthority','GET','','',''),(249,'p','9528','/authority/setDataAuthority','POST','','',''),(251,'p','9528','/authority/setRoleUsers','POST','','',''),(272,'p','9528','/casbin/getPolicyPathByAuthorityId','POST','','',''),(271,'p','9528','/casbin/updateCasbin','POST','','',''),(268,'p','9528','/fileUploadAndDownload/deleteFile','POST','','',''),(269,'p','9528','/fileUploadAndDownload/editFileName','POST','','',''),(267,'p','9528','/fileUploadAndDownload/getFileList','POST','','',''),(270,'p','9528','/fileUploadAndDownload/importURL','POST','','',''),(266,'p','9528','/fileUploadAndDownload/upload','POST','','',''),(273,'p','9528','/jwt/jsonInBlacklist','POST','','',''),(299,'p','9528','/mall/member/detail','POST',NULL,NULL,NULL),(300,'p','9528','/mall/member/list','POST',NULL,NULL,NULL),(254,'p','9528','/menu/addBaseMenu','POST','','',''),(256,'p','9528','/menu/addMenuAuthority','POST','','',''),(260,'p','9528','/menu/deleteBaseMenu','POST','','',''),(262,'p','9528','/menu/getBaseMenuById','POST','','',''),(255,'p','9528','/menu/getBaseMenuTree','POST','','',''),(252,'p','9528','/menu/getMenu','POST','','',''),(257,'p','9528','/menu/getMenuAuthority','POST','','',''),(253,'p','9528','/menu/getMenuList','POST','','',''),(258,'p','9528','/menu/getMenuRoles','GET','','',''),(259,'p','9528','/menu/setMenuRoles','POST','','',''),(261,'p','9528','/menu/updateBaseMenu','POST','','',''),(274,'p','9528','/system/getSystemConfig','POST','','',''),(275,'p','9528','/system/setSystemConfig','POST','','',''),(237,'p','9528','/user/admin_register','POST','','',''),(263,'p','9528','/user/changePassword','POST','','',''),(291,'p','9528','/user/getUserInfo','GET','','',''),(264,'p','9528','/user/getUserList','POST','','',''),(265,'p','9528','/user/setUserAuthority','POST','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_attachment_category`
--

DROP TABLE IF EXISTS `exa_attachment_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_attachment_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '分类名称',
  `pid` bigint DEFAULT '0' COMMENT '父节点ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_attachment_category_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_attachment_category`
--

LOCK TABLES `exa_attachment_category` WRITE;
/*!40000 ALTER TABLE `exa_attachment_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_attachment_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_upload_and_downloads`
--

DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `class_id` bigint DEFAULT '0' COMMENT '分类id',
  `url` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_upload_and_downloads`
--

LOCK TABLES `exa_file_upload_and_downloads` WRITE;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
INSERT INTO `exa_file_upload_and_downloads` VALUES (1,'2026-06-06 03:33:51.822','2026-06-06 03:33:51.822',NULL,'10.png',0,'https://qmplusimg.henrongyi.top/gvalogo.png','png','158787308910.png'),(2,'2026-06-06 03:33:51.822','2026-06-06 03:33:51.822',NULL,'logo.png',0,'https://qmplusimg.henrongyi.top/1576554439myAvatar.png','png','1587973709logo.png');
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jwt_blacklists`
--

DROP TABLE IF EXISTS `jwt_blacklists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jwt_blacklists`
--

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mall_member_levels`
--

DROP TABLE IF EXISTS `mall_member_levels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mall_member_levels` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '等级名称',
  `level` bigint DEFAULT NULL COMMENT '等级序号 数值越大等级越高',
  `min_growth` bigint DEFAULT '0' COMMENT '升级所需成长值',
  `discount_rate` bigint DEFAULT '100' COMMENT '折扣率 100=无折扣 95=95折',
  `icon` varchar(512) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '等级图标',
  `status` bigint DEFAULT '1' COMMENT '状态 1启用 2禁用',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mall_member_levels_level` (`level`),
  KEY `idx_mall_member_levels_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mall_member_levels`
--

LOCK TABLES `mall_member_levels` WRITE;
/*!40000 ALTER TABLE `mall_member_levels` DISABLE KEYS */;
INSERT INTO `mall_member_levels` VALUES (1,'2026-06-06 06:23:22.978','2026-06-06 06:23:22.978',NULL,'普通会员',1,0,100,'',1,''),(2,'2026-06-06 06:23:22.978','2026-06-06 06:23:22.978',NULL,'银卡会员',2,1000,98,'',1,''),(3,'2026-06-06 06:23:22.978','2026-06-06 06:23:22.978',NULL,'金卡会员',3,5000,95,'',1,'');
/*!40000 ALTER TABLE `mall_member_levels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mall_member_login_logs`
--

DROP TABLE IF EXISTS `mall_member_login_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mall_member_login_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `member_id` bigint unsigned DEFAULT NULL COMMENT '会员ID',
  `mobile` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `ip` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '登录IP',
  `user_agent` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'UserAgent',
  `status` tinyint(1) DEFAULT NULL COMMENT '是否成功',
  `error_message` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '失败原因',
  PRIMARY KEY (`id`),
  KEY `idx_mall_member_login_logs_deleted_at` (`deleted_at`),
  KEY `idx_mall_member_login_logs_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mall_member_login_logs`
--

LOCK TABLES `mall_member_login_logs` WRITE;
/*!40000 ALTER TABLE `mall_member_login_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `mall_member_login_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mall_member_oauths`
--

DROP TABLE IF EXISTS `mall_member_oauths`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mall_member_oauths` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `member_id` bigint unsigned DEFAULT NULL COMMENT '会员ID',
  `provider` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '平台 wechat/miniprogram/app',
  `open_id` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '平台OpenID',
  `union_id` varchar(128) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'UnionID',
  `extra` text COLLATE utf8mb4_general_ci COMMENT '扩展信息JSON',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mall_oauth_provider_openid` (`provider`,`open_id`),
  KEY `idx_mall_member_oauths_deleted_at` (`deleted_at`),
  KEY `idx_mall_member_oauths_member_id` (`member_id`),
  KEY `idx_mall_member_oauths_union_id` (`union_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mall_member_oauths`
--

LOCK TABLES `mall_member_oauths` WRITE;
/*!40000 ALTER TABLE `mall_member_oauths` DISABLE KEYS */;
/*!40000 ALTER TABLE `mall_member_oauths` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mall_member_profiles`
--

DROP TABLE IF EXISTS `mall_member_profiles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mall_member_profiles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `member_id` bigint unsigned DEFAULT NULL COMMENT '会员ID',
  `gender` tinyint DEFAULT '0' COMMENT '性别 0未知 1男 2女',
  `birthday` datetime(3) DEFAULT NULL COMMENT '生日',
  `level_id` bigint unsigned DEFAULT '1' COMMENT '会员等级ID',
  `points` bigint DEFAULT '0' COMMENT '可用积分',
  `growth_value` bigint DEFAULT '0' COMMENT '成长值',
  `total_spent` bigint DEFAULT '0' COMMENT '累计消费金额(分)',
  `order_count` bigint DEFAULT '0' COMMENT '累计订单数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mall_member_profiles_member_id` (`member_id`),
  KEY `idx_mall_member_profiles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mall_member_profiles`
--

LOCK TABLES `mall_member_profiles` WRITE;
/*!40000 ALTER TABLE `mall_member_profiles` DISABLE KEYS */;
/*!40000 ALTER TABLE `mall_member_profiles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mall_members`
--

DROP TABLE IF EXISTS `mall_members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mall_members` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '会员UUID',
  `mobile` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '邮箱',
  `password` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '登录密码(bcrypt)',
  `nickname` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(512) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '头像URL',
  `status` bigint DEFAULT '1' COMMENT '状态 1正常 2冻结',
  `register_source` varchar(32) COLLATE utf8mb4_general_ci DEFAULT 'h5' COMMENT '注册来源 h5/app/wechat/miniprogram',
  `register_ip` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '注册IP',
  `last_login_at` datetime(3) DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '最后登录IP',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mall_members_mobile` (`mobile`),
  KEY `idx_mall_members_deleted_at` (`deleted_at`),
  KEY `idx_mall_members_uuid` (`uuid`),
  KEY `idx_mall_members_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mall_members`
--

LOCK TABLES `mall_members` WRITE;
/*!40000 ALTER TABLE `mall_members` DISABLE KEYS */;
/*!40000 ALTER TABLE `mall_members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_api_tokens`
--

DROP TABLE IF EXISTS `sys_api_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_api_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `token` text COLLATE utf8mb4_general_ci COMMENT 'Token',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `expires_at` datetime(3) DEFAULT NULL COMMENT '过期时间',
  `remark` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_tokens_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_api_tokens`
--

LOCK TABLES `sys_api_tokens` WRITE;
/*!40000 ALTER TABLE `sys_api_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_api_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_apis`
--

DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=194 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_apis`
--

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/jwt/jsonInBlacklist','jwt加入黑名单(退出，必选)','jwt','POST'),(2,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysLoginLog/deleteLoginLog','删除登录日志','登录日志','DELETE'),(3,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysLoginLog/deleteLoginLogByIds','批量删除登录日志','登录日志','DELETE'),(4,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysLoginLog/findLoginLog','根据ID获取登录日志','登录日志','GET'),(5,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysLoginLog/getLoginLogList','获取登录日志列表','登录日志','GET'),(6,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysApiToken/createApiToken','签发API Token','API Token','POST'),(7,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysApiToken/getApiTokenList','获取API Token列表','API Token','POST'),(8,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysApiToken/deleteApiToken','作废API Token','API Token','POST'),(9,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/deleteUser','删除用户','系统用户','DELETE'),(10,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/admin_register','用户注册','系统用户','POST'),(11,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/getUserList','获取用户列表','系统用户','POST'),(12,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/setUserInfo','设置用户信息','系统用户','PUT'),(13,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/setSelfInfo','设置自身信息(必选)','系统用户','PUT'),(14,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/getUserInfo','获取自身信息(必选)','系统用户','GET'),(15,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/setUserAuthorities','设置权限组','系统用户','POST'),(16,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/changePassword','修改密码（建议选择)','系统用户','POST'),(17,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/setUserAuthority','修改用户角色(必选)','系统用户','POST'),(18,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/resetPassword','重置用户密码','系统用户','POST'),(19,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/user/setSelfSetting','用户界面配置','系统用户','PUT'),(20,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/createApi','创建api','api','POST'),(21,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/deleteApi','删除Api','api','POST'),(22,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/updateApi','更新Api','api','POST'),(23,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/getApiList','获取api列表','api','POST'),(24,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/getAllApis','获取所有api','api','POST'),(25,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/getApiById','获取api详细信息','api','POST'),(26,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/deleteApisByIds','批量删除api','api','DELETE'),(27,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/syncApi','获取待同步API','api','GET'),(28,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/getApiGroups','获取路由组','api','GET'),(29,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/enterSyncApi','确认同步API','api','POST'),(30,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/ignoreApi','忽略API','api','POST'),(31,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/getApiRoles','获取指定API关联角色列表','api','GET'),(32,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/api/setApiRoles','全量覆盖API关联角色列表','api','POST'),(33,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/copyAuthority','拷贝角色','角色','POST'),(34,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/createAuthority','创建角色','角色','POST'),(35,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/deleteAuthority','删除角色','角色','POST'),(36,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/updateAuthority','更新角色信息','角色','PUT'),(37,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/getAuthorityList','获取角色列表','角色','POST'),(38,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/setDataAuthority','设置角色资源权限','角色','POST'),(39,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/getUsersByAuthority','获取角色关联用户ID列表','角色','GET'),(40,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authority/setRoleUsers','全量覆盖角色关联用户','角色','POST'),(41,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST'),(42,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST'),(43,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/addBaseMenu','新增菜单','菜单','POST'),(44,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/getMenu','获取菜单树(必选)','菜单','POST'),(45,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/deleteBaseMenu','删除菜单','菜单','POST'),(46,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/updateBaseMenu','更新菜单','菜单','POST'),(47,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/getBaseMenuById','根据id获取菜单','菜单','POST'),(48,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/getMenuList','分页获取基础menu列表','菜单','POST'),(49,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/getBaseMenuTree','获取用户动态路由','菜单','POST'),(50,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/getMenuAuthority','获取指定角色menu','菜单','POST'),(51,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','菜单','POST'),(52,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/getMenuRoles','获取菜单关联角色列表','菜单','GET'),(53,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/menu/setMenuRoles','全量覆盖菜单关联角色列表','菜单','POST'),(58,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/fileUploadAndDownload/upload','文件上传（建议选择）','文件上传与下载','POST'),(59,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/fileUploadAndDownload/deleteFile','删除文件','文件上传与下载','POST'),(60,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/fileUploadAndDownload/editFileName','文件名或者备注编辑','文件上传与下载','POST'),(61,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','文件上传与下载','POST'),(62,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/fileUploadAndDownload/importURL','导入URL','文件上传与下载','POST'),(63,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/system/getServerInfo','获取服务器信息','系统服务','POST'),(64,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/system/getSystemConfig','获取配置文件内容','系统服务','POST'),(65,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/system/setSystemConfig','设置配置文件内容','系统服务','POST'),(121,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/updateSysDictionaryDetail','更新字典内容','系统字典详情','PUT'),(122,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/createSysDictionaryDetail','新增字典内容','系统字典详情','POST'),(123,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/deleteSysDictionaryDetail','删除字典内容','系统字典详情','DELETE'),(124,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/findSysDictionaryDetail','根据ID获取字典内容','系统字典详情','GET'),(125,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/getSysDictionaryDetailList','获取字典内容列表','系统字典详情','GET'),(126,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/getDictionaryTreeList','获取字典数列表','系统字典详情','GET'),(127,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/getDictionaryTreeListByType','根据分类获取字典数列表','系统字典详情','GET'),(128,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/getDictionaryDetailsByParent','根据父级ID获取字典详情','系统字典详情','GET'),(129,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionaryDetail/getDictionaryPath','获取字典详情的完整路径','系统字典详情','GET'),(130,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionary/createSysDictionary','新增字典','系统字典','POST'),(131,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionary/deleteSysDictionary','删除字典','系统字典','DELETE'),(132,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionary/updateSysDictionary','更新字典','系统字典','PUT'),(133,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionary/findSysDictionary','根据ID获取字典（建议选择）','系统字典','GET'),(134,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionary/getSysDictionaryList','获取字典列表','系统字典','GET'),(135,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionary/importSysDictionary','导入字典JSON','系统字典','POST'),(136,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysDictionary/exportSysDictionary','导出字典JSON','系统字典','GET'),(137,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysOperationRecord/createSysOperationRecord','新增操作记录','操作记录','POST'),(138,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysOperationRecord/findSysOperationRecord','根据ID获取操作记录','操作记录','GET'),(139,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysOperationRecord/getSysOperationRecordList','获取操作记录列表','操作记录','GET'),(140,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysOperationRecord/deleteSysOperationRecord','删除操作记录','操作记录','DELETE'),(141,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysOperationRecord/deleteSysOperationRecordByIds','批量删除操作历史','操作记录','DELETE'),(142,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/simpleUploader/upload','插件版分片上传','断点续传(插件版)','POST'),(143,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/simpleUploader/checkFileMd5','文件完整度验证','断点续传(插件版)','GET'),(144,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/simpleUploader/mergeFileMd5','上传完成合并文件','断点续传(插件版)','GET'),(145,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/email/emailTest','发送测试邮件','email','POST'),(146,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/email/sendEmail','发送邮件','email','POST'),(147,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authorityBtn/setAuthorityBtn','设置按钮权限','按钮权限','POST'),(148,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authorityBtn/getAuthorityBtn','获取已有按钮权限','按钮权限','POST'),(149,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/authorityBtn/canRemoveAuthorityBtn','删除按钮','按钮权限','POST'),(160,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysError/createSysError','新建错误日志','错误日志','POST'),(161,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysError/deleteSysError','删除错误日志','错误日志','DELETE'),(162,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysError/deleteSysErrorByIds','批量删除错误日志','错误日志','DELETE'),(163,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysError/updateSysError','更新错误日志','错误日志','PUT'),(164,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysError/findSysError','根据ID获取错误日志','错误日志','GET'),(165,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysError/getSysErrorList','获取错误日志列表','错误日志','GET'),(166,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysError/getSysErrorSolution','触发错误处理(异步)','错误日志','GET'),(173,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysParams/createSysParams','新建参数','参数管理','POST'),(174,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysParams/deleteSysParams','删除参数','参数管理','DELETE'),(175,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysParams/deleteSysParamsByIds','批量删除参数','参数管理','DELETE'),(176,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysParams/updateSysParams','更新参数','参数管理','PUT'),(177,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysParams/findSysParams','根据ID获取参数','参数管理','GET'),(178,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysParams/getSysParamsList','获取参数列表','参数管理','GET'),(179,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysParams/getSysParam','获取参数列表','参数管理','GET'),(180,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/attachmentCategory/getCategoryList','分类列表','媒体库分类','GET'),(181,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/attachmentCategory/addCategory','添加/编辑分类','媒体库分类','POST'),(182,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/attachmentCategory/deleteCategory','删除分类','媒体库分类','POST'),(183,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysVersion/findSysVersion','获取单一版本','版本控制','GET'),(184,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysVersion/getSysVersionList','获取版本列表','版本控制','GET'),(185,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysVersion/downloadVersionJson','下载版本json','版本控制','GET'),(186,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysVersion/exportVersion','创建版本','版本控制','POST'),(187,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysVersion/importVersion','同步版本','版本控制','POST'),(188,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysVersion/deleteSysVersion','删除版本','版本控制','DELETE'),(189,'2026-06-06 03:33:51.633','2026-06-06 03:33:51.633',NULL,'/sysVersion/deleteSysVersionByIds','批量删除版本','版本控制','DELETE'),(191,'2026-06-06 05:38:57.000','2026-06-06 05:38:57.000',NULL,'/mall/member/list','会员列表','商城会员','POST'),(192,'2026-06-06 05:38:57.000','2026-06-06 05:38:57.000',NULL,'/mall/member/detail','会员详情','商城会员','POST'),(193,'2026-06-07 01:03:55.963','2026-06-07 01:03:55.963',NULL,'/api/freshCasbin','刷新casbin权限缓存','api','GET');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authorities`
--

DROP TABLE IF EXISTS `sys_authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `uni_sys_authorities_authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authorities`
--

LOCK TABLES `sys_authorities` WRITE;
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` VALUES ('2026-06-06 03:33:51.645','2026-06-06 03:33:51.813',NULL,888,'普通用户',0,'dashboard'),('2026-06-06 03:33:51.645','2026-06-06 03:33:51.819',NULL,8881,'普通用户子角色',888,'dashboard'),('2026-06-06 03:33:51.645','2026-06-06 03:33:51.817',NULL,9528,'测试角色',0,'dashboard');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_btns`
--

DROP TABLE IF EXISTS `sys_authority_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_btns`
--

LOCK TABLES `sys_authority_btns` WRITE;
/*!40000 ALTER TABLE `sys_authority_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_authority_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_menus`
--

DROP TABLE IF EXISTS `sys_authority_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_menus`
--

LOCK TABLES `sys_authority_menus` WRITE;
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` VALUES (1,888),(1,8881),(1,9528),(3,888),(3,8881),(4,888),(4,8881),(4,9528),(6,888),(6,8881),(8,888),(8,8881),(8,9528),(9,888),(9,8881),(10,888),(11,888),(12,888),(13,888),(14,888),(15,888),(16,888),(17,888),(18,888),(19,888),(20,888),(21,888),(39,888),(44,888),(44,8881),(44,9528),(45,888),(45,8881),(45,9528);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_btns`
--

DROP TABLE IF EXISTS `sys_base_menu_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_btns`
--

LOCK TABLES `sys_base_menu_btns` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_parameters`
--

DROP TABLE IF EXISTS `sys_base_menu_parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_parameters`
--

LOCK TABLES `sys_base_menu_parameters` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_parameters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menus`
--

DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '高亮菜单',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '是否是基础路由（开发中）',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单名',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单图标',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '自动关闭tab',
  `transition_type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由切换动画',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menus`
--

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (1,'2026-06-06 03:33:51.661','2026-06-06 03:33:51.661',NULL,0,0,'dashboard','dashboard',0,'view/dashboard/index.vue',1,'',0,0,'仪表盘','odometer',0,''),(3,'2026-06-06 03:33:51.661','2026-06-06 03:33:51.661',NULL,0,0,'admin','superAdmin',0,'view/superAdmin/index.vue',3,'',0,0,'超级管理员','user',0,''),(4,'2026-06-06 03:33:51.661','2026-06-06 03:33:51.661',NULL,0,0,'person','person',1,'view/person/person.vue',4,'',0,0,'个人信息','message',0,''),(6,'2026-06-06 03:33:51.661','2026-06-06 03:33:51.661',NULL,0,0,'systemTools','systemTools',0,'view/systemTools/index.vue',5,'',0,0,'系统工具','tools',0,''),(8,'2026-06-06 03:33:51.661','2026-06-06 03:33:51.661',NULL,0,0,'state','state',0,'view/system/state.vue',8,'',0,0,'服务器状态','cloudy',0,''),(9,'2026-06-06 03:33:51.661','2026-06-06 03:33:51.661',NULL,0,0,'plugin','plugin',0,'view/routerHolder.vue',6,'',0,0,'插件系统','cherry',0,''),(10,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'authority','authority',0,'view/superAdmin/authority/authority.vue',1,'',0,0,'角色管理','avatar',0,''),(11,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'menu','menu',0,'view/superAdmin/menu/menu.vue',2,'',1,0,'菜单管理','tickets',0,''),(12,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'api','api',0,'view/superAdmin/api/api.vue',3,'',1,0,'api管理','platform',0,''),(13,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'user','user',0,'view/superAdmin/user/user.vue',4,'',0,0,'用户管理','coordinate',0,''),(14,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue',5,'',0,0,'字典管理','notebook',0,''),(15,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue',6,'',0,0,'操作历史','pie-chart',0,''),(16,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'sysParams','sysParams',0,'view/superAdmin/params/sysParams.vue',7,'',0,0,'参数管理','compass',0,''),(17,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'system','system',0,'view/systemTools/system/system.vue',8,'',0,0,'系统配置','operation',0,''),(18,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'apiToken','apiToken',0,'view/systemTools/apiToken/index.vue',9,'',0,0,'API Token','key',0,''),(19,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'loginLog','loginLog',0,'view/systemTools/loginLog/index.vue',10,'',0,0,'登录日志','monitor',0,''),(20,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'sysVersion','sysVersion',0,'view/systemTools/version/version.vue',11,'',0,0,'版本管理','server',0,''),(21,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,3,'sysError','sysError',0,'view/systemTools/sysError/sysError.vue',12,'',0,0,'错误日志','warn',0,''),(39,'2026-06-06 03:33:51.662','2026-06-06 03:33:51.662',NULL,1,9,'plugin-email','plugin-email',0,'plugin/email/view/index.vue',4,'',0,0,'邮件插件','message',0,''),(41,'2026-06-06 03:33:51.850','2026-06-06 03:33:51.850',NULL,0,0,'programmingAssistant','AutoRoot',0,'view/routerHolder.vue',91,'',0,0,'编程辅助','cpu',0,''),(44,'2026-06-06 05:50:54.000','2026-06-06 05:50:54.000',NULL,0,0,'mall','mall',0,'view/mall/index.vue',2,NULL,NULL,NULL,'商城管理','shopping-cart',NULL,NULL),(45,'2026-06-06 05:50:54.000','2026-06-06 05:50:54.000',NULL,1,44,'member','mallMember',0,'view/mall/member/index.vue',1,NULL,NULL,NULL,'会员管理','user',NULL,NULL);
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_data_authority_id`
--

DROP TABLE IF EXISTS `sys_data_authority_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_data_authority_id`
--

LOCK TABLES `sys_data_authority_id` WRITE;
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` VALUES (888,888),(888,8881),(888,9528),(9528,8881),(9528,9528);
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionaries`
--

DROP TABLE IF EXISTS `sys_dictionaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父级字典ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionaries`
--

LOCK TABLES `sys_dictionaries` WRITE;
/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
INSERT INTO `sys_dictionaries` VALUES (1,'2026-06-06 03:33:51.650','2026-06-06 03:33:51.651',NULL,'性别','gender',1,'性别字典',NULL),(2,'2026-06-06 03:33:51.650','2026-06-06 03:33:51.653',NULL,'数据库int类型','int',1,'int类型对应的数据库类型',NULL),(3,'2026-06-06 03:33:51.650','2026-06-06 03:33:51.654',NULL,'数据库时间日期类型','time.Time',1,'数据库时间日期类型',NULL),(4,'2026-06-06 03:33:51.650','2026-06-06 03:33:51.656',NULL,'数据库浮点型','float64',1,'数据库浮点型',NULL),(5,'2026-06-06 03:33:51.650','2026-06-06 03:33:51.658',NULL,'数据库字符串','string',1,'数据库字符串',NULL),(6,'2026-06-06 03:33:51.650','2026-06-06 03:33:51.660',NULL,'数据库bool类型','bool',1,'数据库bool类型',NULL);
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionary_details`
--

DROP TABLE IF EXISTS `sys_dictionary_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父级字典详情ID',
  `level` bigint DEFAULT NULL COMMENT '层级深度',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '层级路径',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionary_details`
--

LOCK TABLES `sys_dictionary_details` WRITE;
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
INSERT INTO `sys_dictionary_details` VALUES (1,'2026-06-06 03:33:51.652','2026-06-06 03:33:51.652',NULL,'男','1','',1,1,1,NULL,0,''),(2,'2026-06-06 03:33:51.652','2026-06-06 03:33:51.652',NULL,'女','2','',1,2,1,NULL,0,''),(3,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'smallint','1','mysql',1,1,2,NULL,0,''),(4,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'mediumint','2','mysql',1,2,2,NULL,0,''),(5,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'int','3','mysql',1,3,2,NULL,0,''),(6,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'bigint','4','mysql',1,4,2,NULL,0,''),(7,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'int2','5','pgsql',1,5,2,NULL,0,''),(8,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'int4','6','pgsql',1,6,2,NULL,0,''),(9,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'int6','7','pgsql',1,7,2,NULL,0,''),(10,'2026-06-06 03:33:51.653','2026-06-06 03:33:51.653',NULL,'int8','8','pgsql',1,8,2,NULL,0,''),(11,'2026-06-06 03:33:51.655','2026-06-06 03:33:51.655',NULL,'date','0','mysql',1,0,3,NULL,0,''),(12,'2026-06-06 03:33:51.655','2026-06-06 03:33:51.655',NULL,'time','1','mysql',1,1,3,NULL,0,''),(13,'2026-06-06 03:33:51.655','2026-06-06 03:33:51.655',NULL,'year','2','mysql',1,2,3,NULL,0,''),(14,'2026-06-06 03:33:51.655','2026-06-06 03:33:51.655',NULL,'datetime','3','mysql',1,3,3,NULL,0,''),(15,'2026-06-06 03:33:51.655','2026-06-06 03:33:51.655',NULL,'timestamp','5','mysql',1,5,3,NULL,0,''),(16,'2026-06-06 03:33:51.655','2026-06-06 03:33:51.655',NULL,'timestamptz','6','pgsql',1,5,3,NULL,0,''),(17,'2026-06-06 03:33:51.656','2026-06-06 03:33:51.656',NULL,'float','0','mysql',1,0,4,NULL,0,''),(18,'2026-06-06 03:33:51.656','2026-06-06 03:33:51.656',NULL,'double','1','mysql',1,1,4,NULL,0,''),(19,'2026-06-06 03:33:51.656','2026-06-06 03:33:51.656',NULL,'decimal','2','mysql',1,2,4,NULL,0,''),(20,'2026-06-06 03:33:51.656','2026-06-06 03:33:51.656',NULL,'numeric','3','pgsql',1,3,4,NULL,0,''),(21,'2026-06-06 03:33:51.656','2026-06-06 03:33:51.656',NULL,'smallserial','4','pgsql',1,4,4,NULL,0,''),(22,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'char','0','mysql',1,0,5,NULL,0,''),(23,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'varchar','1','mysql',1,1,5,NULL,0,''),(24,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'tinyblob','2','mysql',1,2,5,NULL,0,''),(25,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'tinytext','3','mysql',1,3,5,NULL,0,''),(26,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'text','4','mysql',1,4,5,NULL,0,''),(27,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'blob','5','mysql',1,5,5,NULL,0,''),(28,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'mediumblob','6','mysql',1,6,5,NULL,0,''),(29,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'mediumtext','7','mysql',1,7,5,NULL,0,''),(30,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'longblob','8','mysql',1,8,5,NULL,0,''),(31,'2026-06-06 03:33:51.658','2026-06-06 03:33:51.658',NULL,'longtext','9','mysql',1,9,5,NULL,0,''),(32,'2026-06-06 03:33:51.660','2026-06-06 03:33:51.660',NULL,'tinyint','1','mysql',1,0,6,NULL,0,''),(33,'2026-06-06 03:33:51.660','2026-06-06 03:33:51.660',NULL,'bool','2','pgsql',1,0,6,NULL,0,'');
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_error`
--

DROP TABLE IF EXISTS `sys_error`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_error` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `form` text COLLATE utf8mb4_general_ci COMMENT '错误来源',
  `info` text COLLATE utf8mb4_general_ci COMMENT '错误内容',
  `level` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '日志等级',
  `solution` text COLLATE utf8mb4_general_ci COMMENT '解决方案',
  `status` varchar(20) COLLATE utf8mb4_general_ci DEFAULT '未处理' COMMENT '处理状态',
  PRIMARY KEY (`id`),
  KEY `idx_sys_error_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_error`
--

LOCK TABLES `sys_error` WRITE;
/*!40000 ALTER TABLE `sys_error` DISABLE KEYS */;
INSERT INTO `sys_error` VALUES (1,'2026-06-06 05:29:56.031','2026-06-06 05:29:56.031',NULL,'前端','错误信息: ReferenceError: reactive is not defined\nStack: 调用栈: ReferenceError: reactive is not defined\n    at setup (http://localhost:8080/src/view/layout/index.vue?t=1780703995926:43:16)\n    at callWithErrorHandling (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2083:17)\n    at setupStatefulComponent (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:6085:23)\n    at setupComponent (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:6061:35)\n    at mountComponent (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:4931:3)\n    at processComponent (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:4921:8)\n    at patch (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:4716:28)\n    at ReactiveEffect.componentUpdateFn [as fn] (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:5036:5)\n    at ReactiveEffect.run (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:1474:17)\n    at http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2313:21','error',NULL,'未处理'),(2,'2026-06-06 05:29:56.300','2026-06-06 05:29:56.300',NULL,'前端','错误信息: TypeError: Cannot read properties of null (reading \'flags\')\nStack: 调用栈: TypeError: Cannot read properties of null (reading \'flags\')\n    at http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2311:23\n    at callWithErrorHandling (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2083:31)\n    at flushJobs (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2223:5)','error',NULL,'未处理'),(3,'2026-06-06 05:29:56.670','2026-06-06 05:29:56.670',NULL,'前端','错误信息: TypeError: Cannot read properties of null (reading \'flags\')\nStack: 调用栈: TypeError: Cannot read properties of null (reading \'flags\')\n    at http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2311:23\n    at callWithErrorHandling (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2083:31)\n    at flushJobs (http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-DXgHk3a4.js?v=50675631:2223:5)','error',NULL,'未处理'),(4,'2026-06-06 05:36:52.165','2026-06-06 05:36:52.165',NULL,'前端','错误信息: TypeError: Failed to fetch dynamically imported module: http://localhost:8080/src/view/dashboard/index.vue?t=1780704410935\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://localhost:8080/src/view/dashboard/index.vue?t=1780704410935','error',NULL,'未处理'),(5,'2026-06-07 01:02:27.721','2026-06-07 01:02:27.721',NULL,'后端','server启动失败 | 错误: listen tcp :8888: bind: address already in use \n 源文件:/Users/admin/Documents/xuexi/mall-admin/server/core/server_run.go:37 \n 调用栈：mall-admin/server/core.initServer.func1\n	/Users/admin/Documents/xuexi/mall-admin/server/core/server_run.go:37','error',NULL,'未处理');
/*!40000 ALTER TABLE `sys_error` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_ignore_apis`
--

DROP TABLE IF EXISTS `sys_ignore_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_ignore_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_ignore_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_ignore_apis`
--

LOCK TABLES `sys_ignore_apis` WRITE;
/*!40000 ALTER TABLE `sys_ignore_apis` DISABLE KEYS */;
INSERT INTO `sys_ignore_apis` VALUES (1,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/swagger/*any','GET'),(3,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/uploads/file/*filepath','GET'),(4,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/health','GET'),(5,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/uploads/file/*filepath','HEAD'),(8,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/system/reloadSystem','POST'),(9,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/base/login','POST'),(10,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/base/captcha','POST'),(11,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/init/initdb','POST'),(12,'2026-06-06 03:33:51.639','2026-06-06 03:33:51.639',NULL,'/init/checkdb','POST'),(15,'2026-06-07 01:03:55.961','2026-06-07 01:03:55.961',NULL,'/site/auth/register','POST'),(16,'2026-06-07 01:03:55.961','2026-06-07 01:03:55.961',NULL,'/site/auth/login','POST'),(17,'2026-06-07 01:03:55.961','2026-06-07 01:03:55.961',NULL,'/site/auth/logout','POST'),(18,'2026-06-07 01:03:55.961','2026-06-07 01:03:55.961',NULL,'/site/auth/profile','GET'),(19,'2026-06-07 01:03:55.961','2026-06-07 01:03:55.961',NULL,'/base/uploadConfig','GET');
/*!40000 ALTER TABLE `sys_ignore_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_login_logs`
--

DROP TABLE IF EXISTS `sys_login_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_login_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名',
  `ip` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `status` tinyint(1) DEFAULT NULL COMMENT '登录状态',
  `error_message` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `agent` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_login_logs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_login_logs`
--

LOCK TABLES `sys_login_logs` WRITE;
/*!40000 ALTER TABLE `sys_login_logs` DISABLE KEYS */;
INSERT INTO `sys_login_logs` VALUES (1,'2026-06-06 03:33:59.112','2026-06-06 03:33:59.112',NULL,'admin','127.0.0.1',0,'验证码错误','curl/8.7.1',0),(2,'2026-06-06 03:34:30.427','2026-06-06 03:34:30.427',NULL,'admin','127.0.0.1',1,'登录成功','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36',1),(3,'2026-06-06 04:30:11.461','2026-06-06 04:30:11.461',NULL,'admin','127.0.0.1',1,'登录成功','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36',1),(4,'2026-06-06 06:03:50.008','2026-06-06 06:03:50.008',NULL,'admin','::1',0,'验证码错误','curl/8.7.1',0),(5,'2026-06-06 06:23:31.694','2026-06-06 06:23:31.694',NULL,'admin','127.0.0.1',0,'验证码错误','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36',0),(6,'2026-06-06 06:23:38.251','2026-06-06 06:23:38.251',NULL,'admin','127.0.0.1',1,'登录成功','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36',1);
/*!40000 ALTER TABLE `sys_login_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_operation_records`
--

DROP TABLE IF EXISTS `sys_operation_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` text COLLATE utf8mb4_general_ci COMMENT '代理',
  `error_message` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` text COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_operation_records`
--

LOCK TABLES `sys_operation_records` WRITE;
/*!40000 ALTER TABLE `sys_operation_records` DISABLE KEYS */;
INSERT INTO `sys_operation_records` VALUES (2,'2026-06-06 05:57:15.518','2026-06-06 05:57:15.518',NULL,'127.0.0.1','GET','/api/getApiGroups',200,1232916,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"mall\":\"商城会员\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"商城会员\"]},\"msg\":\"成功\"}',1),(3,'2026-06-06 05:58:25.155','2026-06-06 05:58:25.155',NULL,'127.0.0.1','GET','/api/getApiGroups',200,5713584,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"mall\":\"商城会员\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"商城会员\"]},\"msg\":\"成功\"}',1),(4,'2026-06-06 05:58:26.655','2026-06-06 05:58:26.655',NULL,'127.0.0.1','GET','/api/getApiGroups',200,816333,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"mall\":\"商城会员\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"商城会员\"]},\"msg\":\"成功\"}',1),(5,'2026-06-06 05:58:27.236','2026-06-06 05:58:27.236',NULL,'127.0.0.1','GET','/api/getApiGroups',200,1470583,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"mall\":\"商城会员\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"商城会员\"]},\"msg\":\"成功\"}',1),(6,'2026-06-06 05:58:28.094','2026-06-06 05:58:28.094',NULL,'127.0.0.1','GET','/api/getApiGroups',200,1628625,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"mall\":\"商城会员\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"商城会员\"]},\"msg\":\"成功\"}',1),(7,'2026-06-07 00:45:08.857','2026-06-07 00:45:08.857',NULL,'127.0.0.1','GET','/api/getApiGroups',200,821667,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"mall\":\"商城会员\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"商城会员\"]},\"msg\":\"成功\"}',1);
/*!40000 ALTER TABLE `sys_operation_records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_params`
--

DROP TABLE IF EXISTS `sys_params`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_params` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数名称',
  `key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数键',
  `value` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数值',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数说明',
  PRIMARY KEY (`id`),
  KEY `idx_sys_params_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_params`
--

LOCK TABLES `sys_params` WRITE;
/*!40000 ALTER TABLE `sys_params` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_params` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_authority`
--

DROP TABLE IF EXISTS `sys_user_authority`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_authority`
--

LOCK TABLES `sys_user_authority` WRITE;
/*!40000 ALTER TABLE `sys_user_authority` DISABLE KEYS */;
INSERT INTO `sys_user_authority` VALUES (1,888),(1,8881),(1,9528),(2,888);
/*!40000 ALTER TABLE `sys_user_authority` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  `origin_setting` json DEFAULT NULL COMMENT '配置',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'2026-06-06 03:33:51.806','2026-06-06 03:33:51.807',NULL,'bbba2744-7dce-4345-a7c8-6ca8d89c3a71','admin','$2a$10$35E/FEgSp3W5JOpr4Le7jOwP.QFQUfnNCuxMYHuT2DujdzyM0o2Kq','Mr.奇淼','https://qmplusimg.henrongyi.top/gva_header.jpg',888,'17611111111','333333333@qq.com',1,NULL),(2,'2026-06-06 03:33:51.806','2026-06-06 03:33:51.809',NULL,'53da3c96-21f6-4db7-8a7c-1693c22960f6','a303176530','$2a$10$ouJZah0mU3DfgqLzNzrm5e.yVHKgAJEsUxxuQAtST/VVxCSWRdO9m','用户1','https://qmplusimg.henrongyi.top/1572075907logo.png',9528,'17611111111','333333333@qq.com',1,NULL);
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_versions`
--

DROP TABLE IF EXISTS `sys_versions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_versions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `version_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本名称',
  `version_code` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本号',
  `description` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本描述',
  `version_data` text COLLATE utf8mb4_general_ci COMMENT '版本数据JSON',
  PRIMARY KEY (`id`),
  KEY `idx_sys_versions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_versions`
--

LOCK TABLES `sys_versions` WRITE;
/*!40000 ALTER TABLE `sys_versions` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_versions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-06-07  1:15:06
