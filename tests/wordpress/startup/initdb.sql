#delete all users
DELETE FROM mysql.user;
DROP DATABASE test;

CREATE DATABASE db_wordpress;
#GRANT ALL PRIVILEGES ON db_wordpress.* to wp_user@'%' identified by '12345';
GRANT ALL PRIVILEGES ON db_wordpress.* to wp_user@'localhost' identified by '12345';

USE db_wordpress

-- MySQL dump 10.14  Distrib 5.5.65-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: db_wordpress
-- ------------------------------------------------------
-- Server version	5.5.65-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `wp_commentmeta`
--

DROP TABLE IF EXISTS `wp_commentmeta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_commentmeta` (
  `meta_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `comment_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `meta_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_value` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`meta_id`),
  KEY `comment_id` (`comment_id`),
  KEY `meta_key` (`meta_key`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_commentmeta`
--

LOCK TABLES `wp_commentmeta` WRITE;
/*!40000 ALTER TABLE `wp_commentmeta` DISABLE KEYS */;
/*!40000 ALTER TABLE `wp_commentmeta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_comments`
--

DROP TABLE IF EXISTS `wp_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_comments` (
  `comment_ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `comment_post_ID` bigint(20) unsigned NOT NULL DEFAULT '0',
  `comment_author` tinytext COLLATE utf8mb4_unicode_ci NOT NULL,
  `comment_author_email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `comment_author_url` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `comment_author_IP` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `comment_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `comment_date_gmt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `comment_content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `comment_karma` int(11) NOT NULL DEFAULT '0',
  `comment_approved` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1',
  `comment_agent` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `comment_type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'comment',
  `comment_parent` bigint(20) unsigned NOT NULL DEFAULT '0',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`comment_ID`),
  KEY `comment_post_ID` (`comment_post_ID`),
  KEY `comment_approved_date_gmt` (`comment_approved`,`comment_date_gmt`),
  KEY `comment_date_gmt` (`comment_date_gmt`),
  KEY `comment_parent` (`comment_parent`),
  KEY `comment_author_email` (`comment_author_email`(10))
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_comments`
--

LOCK TABLES `wp_comments` WRITE;
/*!40000 ALTER TABLE `wp_comments` DISABLE KEYS */;
INSERT INTO `wp_comments` VALUES (1,1,'A WordPress Commenter','wapuu@wordpress.example','https://wordpress.org/','','2020-10-14 18:32:03','2020-10-14 18:32:03','Hi, this is a comment.\nTo get started with moderating, editing, and deleting comments, please visit the Comments screen in the dashboard.\nCommenter avatars come from <a href=\"https://gravatar.com\">Gravatar</a>.',0,'1','','comment',0,0);
/*!40000 ALTER TABLE `wp_comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_links`
--

DROP TABLE IF EXISTS `wp_links`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_links` (
  `link_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `link_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link_image` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link_target` varchar(25) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link_description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link_visible` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'Y',
  `link_owner` bigint(20) unsigned NOT NULL DEFAULT '1',
  `link_rating` int(11) NOT NULL DEFAULT '0',
  `link_updated` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `link_rel` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link_notes` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `link_rss` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`link_id`),
  KEY `link_visible` (`link_visible`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_links`
--

LOCK TABLES `wp_links` WRITE;
/*!40000 ALTER TABLE `wp_links` DISABLE KEYS */;
/*!40000 ALTER TABLE `wp_links` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_options`
--

DROP TABLE IF EXISTS `wp_options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_options` (
  `option_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `option_name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option_value` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `autoload` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'yes',
  PRIMARY KEY (`option_id`),
  UNIQUE KEY `option_name` (`option_name`),
  KEY `autoload` (`autoload`)
) ENGINE=InnoDB AUTO_INCREMENT=136 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_options`
--

LOCK TABLES `wp_options` WRITE;
/*!40000 ALTER TABLE `wp_options` DISABLE KEYS */;
INSERT INTO `wp_options` VALUES (1,'siteurl','http://localhost:8888','yes'),(2,'home','http://localhost:8888','yes'),(3,'blogname','Test','yes'),(4,'blogdescription','Just another WordPress site','yes'),(5,'users_can_register','0','yes'),(6,'admin_email','admin@test.com','yes'),(7,'start_of_week','1','yes'),(8,'use_balanceTags','0','yes'),(9,'use_smilies','1','yes'),(10,'require_name_email','1','yes'),(11,'comments_notify','1','yes'),(12,'posts_per_rss','10','yes'),(13,'rss_use_excerpt','0','yes'),(14,'mailserver_url','mail.example.com','yes'),(15,'mailserver_login','login@example.com','yes'),(16,'mailserver_pass','password','yes'),(17,'mailserver_port','110','yes'),(18,'default_category','1','yes'),(19,'default_comment_status','open','yes'),(20,'default_ping_status','open','yes'),(21,'default_pingback_flag','1','yes'),(22,'posts_per_page','10','yes'),(23,'date_format','F j, Y','yes'),(24,'time_format','g:i a','yes'),(25,'links_updated_date_format','F j, Y g:i a','yes'),(26,'comment_moderation','0','yes'),(27,'moderation_notify','1','yes'),(28,'permalink_structure','','yes'),(29,'rewrite_rules','','yes'),(30,'hack_file','0','yes'),(31,'blog_charset','UTF-8','yes'),(32,'moderation_keys','','no'),(33,'active_plugins','a:0:{}','yes'),(34,'category_base','','yes'),(35,'ping_sites','http://rpc.pingomatic.com/','yes'),(36,'comment_max_links','2','yes'),(37,'gmt_offset','0','yes'),(38,'default_email_category','1','yes'),(39,'recently_edited','','no'),(40,'template','twentytwenty','yes'),(41,'stylesheet','twentytwenty','yes'),(42,'comment_registration','0','yes'),(43,'html_type','text/html','yes'),(44,'use_trackback','0','yes'),(45,'default_role','subscriber','yes'),(46,'db_version','48748','yes'),(47,'uploads_use_yearmonth_folders','1','yes'),(48,'upload_path','','yes'),(49,'blog_public','1','yes'),(50,'default_link_category','2','yes'),(51,'show_on_front','posts','yes'),(52,'tag_base','','yes'),(53,'show_avatars','1','yes'),(54,'avatar_rating','G','yes'),(55,'upload_url_path','','yes'),(56,'thumbnail_size_w','150','yes'),(57,'thumbnail_size_h','150','yes'),(58,'thumbnail_crop','1','yes'),(59,'medium_size_w','300','yes'),(60,'medium_size_h','300','yes'),(61,'avatar_default','mystery','yes'),(62,'large_size_w','1024','yes'),(63,'large_size_h','1024','yes'),(64,'image_default_link_type','none','yes'),(65,'image_default_size','','yes'),(66,'image_default_align','','yes'),(67,'close_comments_for_old_posts','0','yes'),(68,'close_comments_days_old','14','yes'),(69,'thread_comments','1','yes'),(70,'thread_comments_depth','5','yes'),(71,'page_comments','0','yes'),(72,'comments_per_page','50','yes'),(73,'default_comments_page','newest','yes'),(74,'comment_order','asc','yes'),(75,'sticky_posts','a:0:{}','yes'),(76,'widget_categories','a:2:{i:2;a:4:{s:5:\"title\";s:0:\"\";s:5:\"count\";i:0;s:12:\"hierarchical\";i:0;s:8:\"dropdown\";i:0;}s:12:\"_multiwidget\";i:1;}','yes'),(77,'widget_text','a:0:{}','yes'),(78,'widget_rss','a:0:{}','yes'),(79,'uninstall_plugins','a:0:{}','no'),(80,'timezone_string','','yes'),(81,'page_for_posts','0','yes'),(82,'page_on_front','0','yes'),(83,'default_post_format','0','yes'),(84,'link_manager_enabled','0','yes'),(85,'finished_splitting_shared_terms','1','yes'),(86,'site_icon','0','yes'),(87,'medium_large_size_w','768','yes'),(88,'medium_large_size_h','0','yes'),(89,'wp_page_for_privacy_policy','3','yes'),(90,'show_comments_cookies_opt_in','1','yes'),(91,'admin_email_lifespan','1618252323','yes'),(92,'disallowed_keys','','no'),(93,'comment_previously_approved','1','yes'),(94,'auto_plugin_theme_update_emails','a:0:{}','no'),(95,'initial_db_version','48748','yes'),(96,'wp_user_roles','a:5:{s:13:\"administrator\";a:2:{s:4:\"name\";s:13:\"Administrator\";s:12:\"capabilities\";a:61:{s:13:\"switch_themes\";b:1;s:11:\"edit_themes\";b:1;s:16:\"activate_plugins\";b:1;s:12:\"edit_plugins\";b:1;s:10:\"edit_users\";b:1;s:10:\"edit_files\";b:1;s:14:\"manage_options\";b:1;s:17:\"moderate_comments\";b:1;s:17:\"manage_categories\";b:1;s:12:\"manage_links\";b:1;s:12:\"upload_files\";b:1;s:6:\"import\";b:1;s:15:\"unfiltered_html\";b:1;s:10:\"edit_posts\";b:1;s:17:\"edit_others_posts\";b:1;s:20:\"edit_published_posts\";b:1;s:13:\"publish_posts\";b:1;s:10:\"edit_pages\";b:1;s:4:\"read\";b:1;s:8:\"level_10\";b:1;s:7:\"level_9\";b:1;s:7:\"level_8\";b:1;s:7:\"level_7\";b:1;s:7:\"level_6\";b:1;s:7:\"level_5\";b:1;s:7:\"level_4\";b:1;s:7:\"level_3\";b:1;s:7:\"level_2\";b:1;s:7:\"level_1\";b:1;s:7:\"level_0\";b:1;s:17:\"edit_others_pages\";b:1;s:20:\"edit_published_pages\";b:1;s:13:\"publish_pages\";b:1;s:12:\"delete_pages\";b:1;s:19:\"delete_others_pages\";b:1;s:22:\"delete_published_pages\";b:1;s:12:\"delete_posts\";b:1;s:19:\"delete_others_posts\";b:1;s:22:\"delete_published_posts\";b:1;s:20:\"delete_private_posts\";b:1;s:18:\"edit_private_posts\";b:1;s:18:\"read_private_posts\";b:1;s:20:\"delete_private_pages\";b:1;s:18:\"edit_private_pages\";b:1;s:18:\"read_private_pages\";b:1;s:12:\"delete_users\";b:1;s:12:\"create_users\";b:1;s:17:\"unfiltered_upload\";b:1;s:14:\"edit_dashboard\";b:1;s:14:\"update_plugins\";b:1;s:14:\"delete_plugins\";b:1;s:15:\"install_plugins\";b:1;s:13:\"update_themes\";b:1;s:14:\"install_themes\";b:1;s:11:\"update_core\";b:1;s:10:\"list_users\";b:1;s:12:\"remove_users\";b:1;s:13:\"promote_users\";b:1;s:18:\"edit_theme_options\";b:1;s:13:\"delete_themes\";b:1;s:6:\"export\";b:1;}}s:6:\"editor\";a:2:{s:4:\"name\";s:6:\"Editor\";s:12:\"capabilities\";a:34:{s:17:\"moderate_comments\";b:1;s:17:\"manage_categories\";b:1;s:12:\"manage_links\";b:1;s:12:\"upload_files\";b:1;s:15:\"unfiltered_html\";b:1;s:10:\"edit_posts\";b:1;s:17:\"edit_others_posts\";b:1;s:20:\"edit_published_posts\";b:1;s:13:\"publish_posts\";b:1;s:10:\"edit_pages\";b:1;s:4:\"read\";b:1;s:7:\"level_7\";b:1;s:7:\"level_6\";b:1;s:7:\"level_5\";b:1;s:7:\"level_4\";b:1;s:7:\"level_3\";b:1;s:7:\"level_2\";b:1;s:7:\"level_1\";b:1;s:7:\"level_0\";b:1;s:17:\"edit_others_pages\";b:1;s:20:\"edit_published_pages\";b:1;s:13:\"publish_pages\";b:1;s:12:\"delete_pages\";b:1;s:19:\"delete_others_pages\";b:1;s:22:\"delete_published_pages\";b:1;s:12:\"delete_posts\";b:1;s:19:\"delete_others_posts\";b:1;s:22:\"delete_published_posts\";b:1;s:20:\"delete_private_posts\";b:1;s:18:\"edit_private_posts\";b:1;s:18:\"read_private_posts\";b:1;s:20:\"delete_private_pages\";b:1;s:18:\"edit_private_pages\";b:1;s:18:\"read_private_pages\";b:1;}}s:6:\"author\";a:2:{s:4:\"name\";s:6:\"Author\";s:12:\"capabilities\";a:10:{s:12:\"upload_files\";b:1;s:10:\"edit_posts\";b:1;s:20:\"edit_published_posts\";b:1;s:13:\"publish_posts\";b:1;s:4:\"read\";b:1;s:7:\"level_2\";b:1;s:7:\"level_1\";b:1;s:7:\"level_0\";b:1;s:12:\"delete_posts\";b:1;s:22:\"delete_published_posts\";b:1;}}s:11:\"contributor\";a:2:{s:4:\"name\";s:11:\"Contributor\";s:12:\"capabilities\";a:5:{s:10:\"edit_posts\";b:1;s:4:\"read\";b:1;s:7:\"level_1\";b:1;s:7:\"level_0\";b:1;s:12:\"delete_posts\";b:1;}}s:10:\"subscriber\";a:2:{s:4:\"name\";s:10:\"Subscriber\";s:12:\"capabilities\";a:2:{s:4:\"read\";b:1;s:7:\"level_0\";b:1;}}}','yes'),(97,'fresh_site','0','yes'),(98,'widget_search','a:2:{i:2;a:1:{s:5:\"title\";s:0:\"\";}s:12:\"_multiwidget\";i:1;}','yes'),(99,'widget_recent-posts','a:2:{i:2;a:2:{s:5:\"title\";s:0:\"\";s:6:\"number\";i:5;}s:12:\"_multiwidget\";i:1;}','yes'),(100,'widget_recent-comments','a:2:{i:2;a:2:{s:5:\"title\";s:0:\"\";s:6:\"number\";i:5;}s:12:\"_multiwidget\";i:1;}','yes'),(101,'widget_archives','a:2:{i:2;a:3:{s:5:\"title\";s:0:\"\";s:5:\"count\";i:0;s:8:\"dropdown\";i:0;}s:12:\"_multiwidget\";i:1;}','yes'),(102,'widget_meta','a:2:{i:2;a:1:{s:5:\"title\";s:0:\"\";}s:12:\"_multiwidget\";i:1;}','yes'),(103,'sidebars_widgets','a:4:{s:19:\"wp_inactive_widgets\";a:0:{}s:9:\"sidebar-1\";a:3:{i:0;s:8:\"search-2\";i:1;s:14:\"recent-posts-2\";i:2;s:17:\"recent-comments-2\";}s:9:\"sidebar-2\";a:3:{i:0;s:10:\"archives-2\";i:1;s:12:\"categories-2\";i:2;s:6:\"meta-2\";}s:13:\"array_version\";i:3;}','yes'),(104,'cron','a:7:{i:1602700327;a:5:{s:32:\"recovery_mode_clean_expired_keys\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:5:\"daily\";s:4:\"args\";a:0:{}s:8:\"interval\";i:86400;}}s:34:\"wp_privacy_delete_old_export_files\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:6:\"hourly\";s:4:\"args\";a:0:{}s:8:\"interval\";i:3600;}}s:16:\"wp_version_check\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:10:\"twicedaily\";s:4:\"args\";a:0:{}s:8:\"interval\";i:43200;}}s:17:\"wp_update_plugins\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:10:\"twicedaily\";s:4:\"args\";a:0:{}s:8:\"interval\";i:43200;}}s:16:\"wp_update_themes\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:10:\"twicedaily\";s:4:\"args\";a:0:{}s:8:\"interval\";i:43200;}}}i:1602700339;a:2:{s:19:\"wp_scheduled_delete\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:5:\"daily\";s:4:\"args\";a:0:{}s:8:\"interval\";i:86400;}}s:25:\"delete_expired_transients\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:5:\"daily\";s:4:\"args\";a:0:{}s:8:\"interval\";i:86400;}}}i:1602700340;a:1:{s:30:\"wp_scheduled_auto_draft_delete\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:5:\"daily\";s:4:\"args\";a:0:{}s:8:\"interval\";i:86400;}}}i:1602700399;a:1:{s:28:\"wp_update_comment_type_batch\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:2:{s:8:\"schedule\";b:0;s:4:\"args\";a:0:{}}}}i:1602700596;a:1:{s:8:\"do_pings\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:2:{s:8:\"schedule\";b:0;s:4:\"args\";a:0:{}}}}i:1602786727;a:1:{s:30:\"wp_site_health_scheduled_check\";a:1:{s:32:\"40cd750bba9870f18aada2478b24840a\";a:3:{s:8:\"schedule\";s:6:\"weekly\";s:4:\"args\";a:0:{}s:8:\"interval\";i:604800;}}}s:7:\"version\";i:2;}','yes'),(105,'widget_pages','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(106,'widget_calendar','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(107,'widget_media_audio','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(108,'widget_media_image','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(109,'widget_media_gallery','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(110,'widget_media_video','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(111,'nonce_key','QylmV-c$maZX[(8m}V!fZm^!}s+c~j6EH+Oy3?d9;OgfE$9HNf[F12/{in@ A]f|','no'),(112,'nonce_salt','CnKG/*x){RgO3%~?#<mC{I6h])VCO$s|9]#Y=7AWw[u~-lR87/# !RZo[pO^iZ-A','no'),(113,'widget_tag_cloud','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(114,'widget_nav_menu','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(115,'widget_custom_html','a:1:{s:12:\"_multiwidget\";i:1;}','yes'),(116,'_transient_doing_cron','1602704913.5031890869140625000000','yes'),(117,'auth_key','+U@fTFY,;@X-C>@-eUh:G3%!E<|:7+,d[*p|qB,OI07:Lw?m28gYPTzfTh6.@h6!','no'),(118,'auth_salt','p@+^2$7I&%:QRyULNEH0q%,})2a%OivYnzY0D>;^ -N QqpP<247}kHpCu?(>jW.','no'),(119,'logged_in_key','eaX X,7|x-fXE@Il_I:TS?.88_[ZN[> dns9OsupBCmm=Gm7z5gYH-2~kFo)pYJO','no'),(120,'logged_in_salt','l3?*yh0m|:mlFt),>!W3XiAKvv(1j{{%kVs?=!`+,1}y0tU!7[Ch,&[li%I-yt7S','no'),(121,'_site_transient_update_core','O:8:\"stdClass\":4:{s:7:\"updates\";a:1:{i:0;O:8:\"stdClass\":10:{s:8:\"response\";s:6:\"latest\";s:8:\"download\";s:59:\"https://downloads.wordpress.org/release/wordpress-5.5.1.zip\";s:6:\"locale\";s:5:\"en_US\";s:8:\"packages\";O:8:\"stdClass\":5:{s:4:\"full\";s:59:\"https://downloads.wordpress.org/release/wordpress-5.5.1.zip\";s:10:\"no_content\";s:70:\"https://downloads.wordpress.org/release/wordpress-5.5.1-no-content.zip\";s:11:\"new_bundled\";s:71:\"https://downloads.wordpress.org/release/wordpress-5.5.1-new-bundled.zip\";s:7:\"partial\";s:0:\"\";s:8:\"rollback\";s:0:\"\";}s:7:\"current\";s:5:\"5.5.1\";s:7:\"version\";s:5:\"5.5.1\";s:11:\"php_version\";s:6:\"5.6.20\";s:13:\"mysql_version\";s:3:\"5.0\";s:11:\"new_bundled\";s:3:\"5.3\";s:15:\"partial_version\";s:0:\"\";}}s:12:\"last_checked\";i:1602700339;s:15:\"version_checked\";s:5:\"5.5.1\";s:12:\"translations\";a:0:{}}','no'),(122,'_site_transient_update_plugins','O:8:\"stdClass\":4:{s:12:\"last_checked\";i:1602700339;s:8:\"response\";a:0:{}s:12:\"translations\";a:0:{}s:9:\"no_update\";a:2:{s:19:\"akismet/akismet.php\";O:8:\"stdClass\":9:{s:2:\"id\";s:21:\"w.org/plugins/akismet\";s:4:\"slug\";s:7:\"akismet\";s:6:\"plugin\";s:19:\"akismet/akismet.php\";s:11:\"new_version\";s:5:\"4.1.6\";s:3:\"url\";s:38:\"https://wordpress.org/plugins/akismet/\";s:7:\"package\";s:56:\"https://downloads.wordpress.org/plugin/akismet.4.1.6.zip\";s:5:\"icons\";a:2:{s:2:\"2x\";s:59:\"https://ps.w.org/akismet/assets/icon-256x256.png?rev=969272\";s:2:\"1x\";s:59:\"https://ps.w.org/akismet/assets/icon-128x128.png?rev=969272\";}s:7:\"banners\";a:1:{s:2:\"1x\";s:61:\"https://ps.w.org/akismet/assets/banner-772x250.jpg?rev=479904\";}s:11:\"banners_rtl\";a:0:{}}s:9:\"hello.php\";O:8:\"stdClass\":9:{s:2:\"id\";s:25:\"w.org/plugins/hello-dolly\";s:4:\"slug\";s:11:\"hello-dolly\";s:6:\"plugin\";s:9:\"hello.php\";s:11:\"new_version\";s:5:\"1.7.2\";s:3:\"url\";s:42:\"https://wordpress.org/plugins/hello-dolly/\";s:7:\"package\";s:60:\"https://downloads.wordpress.org/plugin/hello-dolly.1.7.2.zip\";s:5:\"icons\";a:2:{s:2:\"2x\";s:64:\"https://ps.w.org/hello-dolly/assets/icon-256x256.jpg?rev=2052855\";s:2:\"1x\";s:64:\"https://ps.w.org/hello-dolly/assets/icon-128x128.jpg?rev=2052855\";}s:7:\"banners\";a:1:{s:2:\"1x\";s:66:\"https://ps.w.org/hello-dolly/assets/banner-772x250.jpg?rev=2052855\";}s:11:\"banners_rtl\";a:0:{}}}}','no'),(123,'_site_transient_timeout_theme_roots','1602702139','no'),(124,'_site_transient_theme_roots','a:3:{s:14:\"twentynineteen\";s:7:\"/themes\";s:15:\"twentyseventeen\";s:7:\"/themes\";s:12:\"twentytwenty\";s:7:\"/themes\";}','no'),(125,'_site_transient_update_themes','O:8:\"stdClass\":5:{s:12:\"last_checked\";i:1602700339;s:7:\"checked\";a:3:{s:14:\"twentynineteen\";s:3:\"1.7\";s:15:\"twentyseventeen\";s:3:\"2.4\";s:12:\"twentytwenty\";s:3:\"1.5\";}s:8:\"response\";a:0:{}s:9:\"no_update\";a:3:{s:14:\"twentynineteen\";a:6:{s:5:\"theme\";s:14:\"twentynineteen\";s:11:\"new_version\";s:3:\"1.7\";s:3:\"url\";s:44:\"https://wordpress.org/themes/twentynineteen/\";s:7:\"package\";s:60:\"https://downloads.wordpress.org/theme/twentynineteen.1.7.zip\";s:8:\"requires\";s:5:\"4.9.6\";s:12:\"requires_php\";s:5:\"5.2.4\";}s:15:\"twentyseventeen\";a:6:{s:5:\"theme\";s:15:\"twentyseventeen\";s:11:\"new_version\";s:3:\"2.4\";s:3:\"url\";s:45:\"https://wordpress.org/themes/twentyseventeen/\";s:7:\"package\";s:61:\"https://downloads.wordpress.org/theme/twentyseventeen.2.4.zip\";s:8:\"requires\";s:3:\"4.7\";s:12:\"requires_php\";s:5:\"5.2.4\";}s:12:\"twentytwenty\";a:6:{s:5:\"theme\";s:12:\"twentytwenty\";s:11:\"new_version\";s:3:\"1.5\";s:3:\"url\";s:42:\"https://wordpress.org/themes/twentytwenty/\";s:7:\"package\";s:58:\"https://downloads.wordpress.org/theme/twentytwenty.1.5.zip\";s:8:\"requires\";s:3:\"4.7\";s:12:\"requires_php\";s:5:\"5.2.4\";}}s:12:\"translations\";a:0:{}}','no'),(126,'_site_transient_timeout_browser_2e3b07d506595682929bc3b13f080001','1603305140','no'),(127,'_site_transient_browser_2e3b07d506595682929bc3b13f080001','a:10:{s:4:\"name\";s:6:\"Chrome\";s:7:\"version\";s:13:\"85.0.4183.121\";s:8:\"platform\";s:9:\"Macintosh\";s:10:\"update_url\";s:29:\"https://www.google.com/chrome\";s:7:\"img_src\";s:43:\"http://s.w.org/images/browsers/chrome.png?1\";s:11:\"img_src_ssl\";s:44:\"https://s.w.org/images/browsers/chrome.png?1\";s:15:\"current_version\";s:2:\"18\";s:7:\"upgrade\";b:0;s:8:\"insecure\";b:0;s:6:\"mobile\";b:0;}','no'),(128,'_site_transient_timeout_php_check_97f83d63b8a66f6e8c057d89a83d8845','1603305140','no'),(129,'_site_transient_php_check_97f83d63b8a66f6e8c057d89a83d8845','a:5:{s:19:\"recommended_version\";s:3:\"7.4\";s:15:\"minimum_version\";s:6:\"5.6.20\";s:12:\"is_supported\";b:0;s:9:\"is_secure\";b:0;s:13:\"is_acceptable\";b:0;}','no'),(130,'_site_transient_timeout_community-events-f119fd185e2033d6979a549e98b5e616','1602743542','no'),(131,'_site_transient_community-events-f119fd185e2033d6979a549e98b5e616','a:4:{s:9:\"sandboxed\";b:0;s:5:\"error\";N;s:8:\"location\";a:1:{s:2:\"ip\";s:10:\"172.22.0.0\";}s:6:\"events\";a:7:{i:0;a:10:{s:4:\"type\";s:6:\"meetup\";s:5:\"title\";s:59:\"Discussion Group: Intro to Publishing with the Block Editor\";s:3:\"url\";s:68:\"https://www.meetup.com/learn-wordpress-discussions/events/273756574/\";s:6:\"meetup\";s:27:\"Learn WordPress Discussions\";s:10:\"meetup_url\";s:51:\"https://www.meetup.com/learn-wordpress-discussions/\";s:4:\"date\";s:19:\"2020-10-14 09:00:00\";s:8:\"end_date\";s:19:\"2020-10-14 10:00:00\";s:20:\"start_unix_timestamp\";i:1602691200;s:18:\"end_unix_timestamp\";i:1602694800;s:8:\"location\";a:4:{s:8:\"location\";s:6:\"Online\";s:7:\"country\";s:2:\"US\";s:8:\"latitude\";d:37.779998779297;s:9:\"longitude\";d:-122.41999816895;}}i:1;a:10:{s:4:\"type\";s:8:\"wordcamp\";s:5:\"title\";s:22:\"WordCamp Italia Online\";s:3:\"url\";s:33:\"https://2020.italia.wordcamp.org/\";s:6:\"meetup\";N;s:10:\"meetup_url\";N;s:4:\"date\";s:19:\"2020-10-16 00:00:00\";s:8:\"end_date\";s:19:\"2020-10-17 00:00:00\";s:20:\"start_unix_timestamp\";i:1602799200;s:18:\"end_unix_timestamp\";i:1602885600;s:8:\"location\";a:4:{s:8:\"location\";s:6:\"Online\";s:7:\"country\";s:2:\"IT\";s:8:\"latitude\";d:41.87194;s:9:\"longitude\";d:12.56738;}}i:2;a:10:{s:4:\"type\";s:6:\"meetup\";s:5:\"title\";s:49:\"WordPress Toronto - Let\'s Fix Your WordPress Site\";s:3:\"url\";s:54:\"https://www.meetup.com/WPToronto/events/xfnwwrybcnbbc/\";s:6:\"meetup\";s:27:\"The Toronto WordPress Group\";s:10:\"meetup_url\";s:33:\"https://www.meetup.com/WPToronto/\";s:4:\"date\";s:19:\"2020-10-20 18:30:00\";s:8:\"end_date\";s:19:\"2020-10-20 20:30:00\";s:20:\"start_unix_timestamp\";i:1603233000;s:18:\"end_unix_timestamp\";i:1603240200;s:8:\"location\";a:4:{s:8:\"location\";s:6:\"Online\";s:7:\"country\";s:2:\"CA\";s:8:\"latitude\";d:43.659999847412;s:9:\"longitude\";d:-79.379997253418;}}i:3;a:10:{s:4:\"type\";s:6:\"meetup\";s:5:\"title\";s:46:\"Brampton Wordpress Meetup - Fix My Site Clinic\";s:3:\"url\";s:70:\"https://www.meetup.com/Brampton-WordPress-Meetup/events/mrgqvrybcnblc/\";s:6:\"meetup\";s:25:\"Brampton WordPress Meetup\";s:10:\"meetup_url\";s:49:\"https://www.meetup.com/Brampton-WordPress-Meetup/\";s:4:\"date\";s:19:\"2020-10-28 18:00:00\";s:8:\"end_date\";s:19:\"2020-10-28 20:00:00\";s:20:\"start_unix_timestamp\";i:1603922400;s:18:\"end_unix_timestamp\";i:1603929600;s:8:\"location\";a:4:{s:8:\"location\";s:6:\"Online\";s:7:\"country\";s:2:\"CA\";s:8:\"latitude\";d:43.790000915527;s:9:\"longitude\";d:-79.73999786377;}}i:4;a:10:{s:4:\"type\";s:6:\"meetup\";s:5:\"title\";s:49:\"WordPress Toronto - Let\'s Fix Your WordPress Site\";s:3:\"url\";s:54:\"https://www.meetup.com/WPToronto/events/xfnwwrybcpbwb/\";s:6:\"meetup\";s:27:\"The Toronto WordPress Group\";s:10:\"meetup_url\";s:33:\"https://www.meetup.com/WPToronto/\";s:4:\"date\";s:19:\"2020-11-17 18:30:00\";s:8:\"end_date\";s:19:\"2020-11-17 20:30:00\";s:20:\"start_unix_timestamp\";i:1605655800;s:18:\"end_unix_timestamp\";i:1605663000;s:8:\"location\";a:4:{s:8:\"location\";s:6:\"Online\";s:7:\"country\";s:2:\"CA\";s:8:\"latitude\";d:43.659999847412;s:9:\"longitude\";d:-79.379997253418;}}i:5;a:10:{s:4:\"type\";s:6:\"meetup\";s:5:\"title\";s:46:\"Brampton Wordpress Meetup - Fix My Site Clinic\";s:3:\"url\";s:70:\"https://www.meetup.com/Brampton-WordPress-Meetup/events/mrgqvrybcpbhc/\";s:6:\"meetup\";s:25:\"Brampton WordPress Meetup\";s:10:\"meetup_url\";s:49:\"https://www.meetup.com/Brampton-WordPress-Meetup/\";s:4:\"date\";s:19:\"2020-11-25 18:00:00\";s:8:\"end_date\";s:19:\"2020-11-25 20:00:00\";s:20:\"start_unix_timestamp\";i:1606345200;s:18:\"end_unix_timestamp\";i:1606352400;s:8:\"location\";a:4:{s:8:\"location\";s:6:\"Online\";s:7:\"country\";s:2:\"CA\";s:8:\"latitude\";d:43.790000915527;s:9:\"longitude\";d:-79.73999786377;}}i:6;a:10:{s:4:\"type\";s:6:\"meetup\";s:5:\"title\";s:49:\"WordPress Toronto - Let\'s Fix Your WordPress Site\";s:3:\"url\";s:54:\"https://www.meetup.com/WPToronto/events/xfnwwrybcqbtb/\";s:6:\"meetup\";s:27:\"The Toronto WordPress Group\";s:10:\"meetup_url\";s:33:\"https://www.meetup.com/WPToronto/\";s:4:\"date\";s:19:\"2020-12-15 18:30:00\";s:8:\"end_date\";s:19:\"2020-12-15 20:30:00\";s:20:\"start_unix_timestamp\";i:1608075000;s:18:\"end_unix_timestamp\";i:1608082200;s:8:\"location\";a:4:{s:8:\"location\";s:6:\"Online\";s:7:\"country\";s:2:\"CA\";s:8:\"latitude\";d:43.659999847412;s:9:\"longitude\";d:-79.379997253418;}}}}','no'),(132,'_transient_timeout_dash_v2_88ae138922fe95674369b1cb3d215a2b','1602743542','no'),(133,'_transient_dash_v2_88ae138922fe95674369b1cb3d215a2b','<div class=\"rss-widget\"><p><strong>RSS Error:</strong> XML or PCRE extensions not loaded!</p></div><div class=\"rss-widget\"><p><strong>RSS Error:</strong> XML or PCRE extensions not loaded!</p></div>','no'),(134,'can_compress_scripts','1','no'),(135,'theme_mods_twentytwenty','a:1:{s:18:\"custom_css_post_id\";i:-1;}','yes');
/*!40000 ALTER TABLE `wp_options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_postmeta`
--

DROP TABLE IF EXISTS `wp_postmeta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_postmeta` (
  `meta_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `meta_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_value` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`meta_id`),
  KEY `post_id` (`post_id`),
  KEY `meta_key` (`meta_key`(191))
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_postmeta`
--

LOCK TABLES `wp_postmeta` WRITE;
/*!40000 ALTER TABLE `wp_postmeta` DISABLE KEYS */;
INSERT INTO `wp_postmeta` VALUES (1,2,'_wp_page_template','default'),(2,3,'_wp_page_template','default'),(3,5,'_edit_lock','1602700454:1'),(4,5,'_pingme','1'),(5,5,'_encloseme','1'),(6,7,'_edit_lock','1602700476:1'),(7,7,'_pingme','1'),(8,7,'_encloseme','1'),(9,9,'_edit_lock','1602700499:1'),(10,9,'_pingme','1'),(11,9,'_encloseme','1'),(12,11,'_edit_lock','1602700571:1'),(13,11,'_pingme','1'),(14,11,'_encloseme','1'),(15,13,'_edit_lock','1602700610:1'),(16,13,'_pingme','1'),(17,13,'_encloseme','1');
/*!40000 ALTER TABLE `wp_postmeta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_posts`
--

DROP TABLE IF EXISTS `wp_posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_posts` (
  `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `post_author` bigint(20) unsigned NOT NULL DEFAULT '0',
  `post_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `post_date_gmt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `post_content` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `post_title` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `post_excerpt` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `post_status` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'publish',
  `comment_status` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'open',
  `ping_status` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'open',
  `post_password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `post_name` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `to_ping` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `pinged` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `post_modified` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `post_modified_gmt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `post_content_filtered` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `post_parent` bigint(20) unsigned NOT NULL DEFAULT '0',
  `guid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `menu_order` int(11) NOT NULL DEFAULT '0',
  `post_type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'post',
  `post_mime_type` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `comment_count` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ID`),
  KEY `post_name` (`post_name`(191)),
  KEY `type_status_date` (`post_type`,`post_status`,`post_date`,`ID`),
  KEY `post_parent` (`post_parent`),
  KEY `post_author` (`post_author`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_posts`
--

LOCK TABLES `wp_posts` WRITE;
/*!40000 ALTER TABLE `wp_posts` DISABLE KEYS */;
INSERT INTO `wp_posts` VALUES (1,1,'2020-10-14 18:32:03','2020-10-14 18:32:03','<!-- wp:paragraph -->\n<p>Welcome to WordPress. This is your first post. Edit or delete it, then start writing!</p>\n<!-- /wp:paragraph -->','Hello world!','','publish','open','open','','hello-world','','','2020-10-14 18:32:03','2020-10-14 18:32:03','',0,'http://localhost:8888/?p=1',0,'post','',1),(2,1,'2020-10-14 18:32:03','2020-10-14 18:32:03','<!-- wp:paragraph -->\n<p>This is an example page. It\'s different from a blog post because it will stay in one place and will show up in your site navigation (in most themes). Most people start with an About page that introduces them to potential site visitors. It might say something like this:</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:quote -->\n<blockquote class=\"wp-block-quote\"><p>Hi there! I\'m a bike messenger by day, aspiring actor by night, and this is my website. I live in Los Angeles, have a great dog named Jack, and I like pi&#241;a coladas. (And gettin\' caught in the rain.)</p></blockquote>\n<!-- /wp:quote -->\n\n<!-- wp:paragraph -->\n<p>...or something like this:</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:quote -->\n<blockquote class=\"wp-block-quote\"><p>The XYZ Doohickey Company was founded in 1971, and has been providing quality doohickeys to the public ever since. Located in Gotham City, XYZ employs over 2,000 people and does all kinds of awesome things for the Gotham community.</p></blockquote>\n<!-- /wp:quote -->\n\n<!-- wp:paragraph -->\n<p>As a new WordPress user, you should go to <a href=\"http://localhost:8888/wp-admin/\">your dashboard</a> to delete this page and create new pages for your content. Have fun!</p>\n<!-- /wp:paragraph -->','Sample Page','','publish','closed','open','','sample-page','','','2020-10-14 18:32:03','2020-10-14 18:32:03','',0,'http://localhost:8888/?page_id=2',0,'page','',0),(3,1,'2020-10-14 18:32:03','2020-10-14 18:32:03','<!-- wp:heading --><h2>Who we are</h2><!-- /wp:heading --><!-- wp:paragraph --><p>Our website address is: http://localhost:8888.</p><!-- /wp:paragraph --><!-- wp:heading --><h2>What personal data we collect and why we collect it</h2><!-- /wp:heading --><!-- wp:heading {\"level\":3} --><h3>Comments</h3><!-- /wp:heading --><!-- wp:paragraph --><p>When visitors leave comments on the site we collect the data shown in the comments form, and also the visitor&#8217;s IP address and browser user agent string to help spam detection.</p><!-- /wp:paragraph --><!-- wp:paragraph --><p>An anonymized string created from your email address (also called a hash) may be provided to the Gravatar service to see if you are using it. The Gravatar service privacy policy is available here: https://automattic.com/privacy/. After approval of your comment, your profile picture is visible to the public in the context of your comment.</p><!-- /wp:paragraph --><!-- wp:heading {\"level\":3} --><h3>Media</h3><!-- /wp:heading --><!-- wp:paragraph --><p>If you upload images to the website, you should avoid uploading images with embedded location data (EXIF GPS) included. Visitors to the website can download and extract any location data from images on the website.</p><!-- /wp:paragraph --><!-- wp:heading {\"level\":3} --><h3>Contact forms</h3><!-- /wp:heading --><!-- wp:heading {\"level\":3} --><h3>Cookies</h3><!-- /wp:heading --><!-- wp:paragraph --><p>If you leave a comment on our site you may opt-in to saving your name, email address and website in cookies. These are for your convenience so that you do not have to fill in your details again when you leave another comment. These cookies will last for one year.</p><!-- /wp:paragraph --><!-- wp:paragraph --><p>If you visit our login page, we will set a temporary cookie to determine if your browser accepts cookies. This cookie contains no personal data and is discarded when you close your browser.</p><!-- /wp:paragraph --><!-- wp:paragraph --><p>When you log in, we will also set up several cookies to save your login information and your screen display choices. Login cookies last for two days, and screen options cookies last for a year. If you select &quot;Remember Me&quot;, your login will persist for two weeks. If you log out of your account, the login cookies will be removed.</p><!-- /wp:paragraph --><!-- wp:paragraph --><p>If you edit or publish an article, an additional cookie will be saved in your browser. This cookie includes no personal data and simply indicates the post ID of the article you just edited. It expires after 1 day.</p><!-- /wp:paragraph --><!-- wp:heading {\"level\":3} --><h3>Embedded content from other websites</h3><!-- /wp:heading --><!-- wp:paragraph --><p>Articles on this site may include embedded content (e.g. videos, images, articles, etc.). Embedded content from other websites behaves in the exact same way as if the visitor has visited the other website.</p><!-- /wp:paragraph --><!-- wp:paragraph --><p>These websites may collect data about you, use cookies, embed additional third-party tracking, and monitor your interaction with that embedded content, including tracking your interaction with the embedded content if you have an account and are logged in to that website.</p><!-- /wp:paragraph --><!-- wp:heading {\"level\":3} --><h3>Analytics</h3><!-- /wp:heading --><!-- wp:heading --><h2>Who we share your data with</h2><!-- /wp:heading --><!-- wp:heading --><h2>How long we retain your data</h2><!-- /wp:heading --><!-- wp:paragraph --><p>If you leave a comment, the comment and its metadata are retained indefinitely. This is so we can recognize and approve any follow-up comments automatically instead of holding them in a moderation queue.</p><!-- /wp:paragraph --><!-- wp:paragraph --><p>For users that register on our website (if any), we also store the personal information they provide in their user profile. All users can see, edit, or delete their personal information at any time (except they cannot change their username). Website administrators can also see and edit that information.</p><!-- /wp:paragraph --><!-- wp:heading --><h2>What rights you have over your data</h2><!-- /wp:heading --><!-- wp:paragraph --><p>If you have an account on this site, or have left comments, you can request to receive an exported file of the personal data we hold about you, including any data you have provided to us. You can also request that we erase any personal data we hold about you. This does not include any data we are obliged to keep for administrative, legal, or security purposes.</p><!-- /wp:paragraph --><!-- wp:heading --><h2>Where we send your data</h2><!-- /wp:heading --><!-- wp:paragraph --><p>Visitor comments may be checked through an automated spam detection service.</p><!-- /wp:paragraph --><!-- wp:heading --><h2>Your contact information</h2><!-- /wp:heading --><!-- wp:heading --><h2>Additional information</h2><!-- /wp:heading --><!-- wp:heading {\"level\":3} --><h3>How we protect your data</h3><!-- /wp:heading --><!-- wp:heading {\"level\":3} --><h3>What data breach procedures we have in place</h3><!-- /wp:heading --><!-- wp:heading {\"level\":3} --><h3>What third parties we receive data from</h3><!-- /wp:heading --><!-- wp:heading {\"level\":3} --><h3>What automated decision making and/or profiling we do with user data</h3><!-- /wp:heading --><!-- wp:heading {\"level\":3} --><h3>Industry regulatory disclosure requirements</h3><!-- /wp:heading -->','Privacy Policy','','draft','closed','open','','privacy-policy','','','2020-10-14 18:32:03','2020-10-14 18:32:03','',0,'http://localhost:8888/?page_id=3',0,'page','',0),(4,1,'2020-10-14 18:32:20','0000-00-00 00:00:00','','Auto Draft','','auto-draft','open','open','','','','','2020-10-14 18:32:20','0000-00-00 00:00:00','',0,'http://localhost:8888/?p=4',0,'post','',0),(5,1,'2020-10-14 18:36:35','2020-10-14 18:36:35','<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur eleifend dictum justo, tempor dignissim elit tempus vel. Nam ut efficitur mi. Curabitur aliquam a erat nec efficitur. Quisque tortor diam, dapibus id gravida in, rhoncus vitae leo. Maecenas interdum sed augue sed consectetur. Curabitur eu enim elit. Aenean eu sapien sodales, malesuada mauris non, porttitor erat. Maecenas consectetur orci dolor, eget vulputate libero viverra ut. Sed ac congue diam, in porttitor leo. Pellentesque euismod, tortor eget finibus pellentesque, augue erat molestie purus, vel lacinia risus velit non massa. Praesent ligula lacus, dictum vitae volutpat blandit, hendrerit in elit. Duis orci ligula, molestie vel eros sit amet, vulputate lacinia ex.</p>\n<!-- /wp:paragraph -->','TESTING 1','','publish','open','open','','testing-1','','','2020-10-14 18:36:35','2020-10-14 18:36:35','',0,'http://localhost:8888/?p=5',0,'post','',0),(6,1,'2020-10-14 18:36:35','2020-10-14 18:36:35','<!-- wp:paragraph -->\n<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur eleifend dictum justo, tempor dignissim elit tempus vel. Nam ut efficitur mi. Curabitur aliquam a erat nec efficitur. Quisque tortor diam, dapibus id gravida in, rhoncus vitae leo. Maecenas interdum sed augue sed consectetur. Curabitur eu enim elit. Aenean eu sapien sodales, malesuada mauris non, porttitor erat. Maecenas consectetur orci dolor, eget vulputate libero viverra ut. Sed ac congue diam, in porttitor leo. Pellentesque euismod, tortor eget finibus pellentesque, augue erat molestie purus, vel lacinia risus velit non massa. Praesent ligula lacus, dictum vitae volutpat blandit, hendrerit in elit. Duis orci ligula, molestie vel eros sit amet, vulputate lacinia ex.</p>\n<!-- /wp:paragraph -->','TESTING 1','','inherit','closed','closed','','5-revision-v1','','','2020-10-14 18:36:35','2020-10-14 18:36:35','',5,'http://localhost:8888/?p=6',0,'revision','',0),(7,1,'2020-10-14 18:36:58','2020-10-14 18:36:58','<!-- wp:paragraph -->\n<p>Aliquam erat volutpat. Duis id blandit ex, nec ullamcorper est. Nunc sollicitudin vulputate tincidunt. Donec lectus massa, posuere ac nunc et, consectetur sollicitudin urna. Aliquam eu tortor sapien. Sed imperdiet, erat ut suscipit commodo, massa mi eleifend felis, quis congue mauris lectus a eros. Fusce in nibh sem. Vestibulum eros est, consectetur a dictum volutpat, fermentum at urna. Integer faucibus in purus eget volutpat. Aliquam accumsan ex nec dolor commodo, ac hendrerit ligula luctus. Nullam tincidunt sagittis suscipit. Praesent nec mauris sodales, blandit nibh ac, tristique mi. In porttitor rhoncus tincidunt. Suspendisse id tellus turpis. Quisque fringilla efficitur nisi, eleifend posuere nisl ornare ac. Duis quis enim gravida, gravida lectus ac, placerat nisi.</p>\n<!-- /wp:paragraph -->','Testing2','','publish','open','open','','testing2','','','2020-10-14 18:36:58','2020-10-14 18:36:58','',0,'http://localhost:8888/?p=7',0,'post','',0),(8,1,'2020-10-14 18:36:58','2020-10-14 18:36:58','<!-- wp:paragraph -->\n<p>Aliquam erat volutpat. Duis id blandit ex, nec ullamcorper est. Nunc sollicitudin vulputate tincidunt. Donec lectus massa, posuere ac nunc et, consectetur sollicitudin urna. Aliquam eu tortor sapien. Sed imperdiet, erat ut suscipit commodo, massa mi eleifend felis, quis congue mauris lectus a eros. Fusce in nibh sem. Vestibulum eros est, consectetur a dictum volutpat, fermentum at urna. Integer faucibus in purus eget volutpat. Aliquam accumsan ex nec dolor commodo, ac hendrerit ligula luctus. Nullam tincidunt sagittis suscipit. Praesent nec mauris sodales, blandit nibh ac, tristique mi. In porttitor rhoncus tincidunt. Suspendisse id tellus turpis. Quisque fringilla efficitur nisi, eleifend posuere nisl ornare ac. Duis quis enim gravida, gravida lectus ac, placerat nisi.</p>\n<!-- /wp:paragraph -->','Testing2','','inherit','closed','closed','','7-revision-v1','','','2020-10-14 18:36:58','2020-10-14 18:36:58','',7,'http://localhost:8888/?p=8',0,'revision','',0),(9,1,'2020-10-14 18:37:21','2020-10-14 18:37:21','<!-- wp:paragraph -->\n<p>Cras fermentum mattis nulla. Phasellus molestie dui sit amet suscipit finibus. Nulla purus felis, egestas sagittis sodales vel, consectetur eu nunc. Cras hendrerit nibh vitae mollis luctus. Nulla ultrices tincidunt ipsum non dapibus. Morbi imperdiet vehicula sollicitudin. Sed non hendrerit augue, eget malesuada nulla. Fusce vestibulum sit amet ante venenatis porttitor. Integer efficitur feugiat sem eu sagittis. Cras aliquam dignissim lacus id vehicula. Quisque at felis sapien. Aenean sit amet ipsum feugiat, condimentum quam eget, convallis metus. Quisque vehicula lobortis lorem, ac sodales justo laoreet pretium. Vivamus eget quam augue.</p>\n<!-- /wp:paragraph -->','Testing Three','','publish','open','open','','testing-three','','','2020-10-14 18:37:21','2020-10-14 18:37:21','',0,'http://localhost:8888/?p=9',0,'post','',0),(10,1,'2020-10-14 18:37:21','2020-10-14 18:37:21','<!-- wp:paragraph -->\n<p>Cras fermentum mattis nulla. Phasellus molestie dui sit amet suscipit finibus. Nulla purus felis, egestas sagittis sodales vel, consectetur eu nunc. Cras hendrerit nibh vitae mollis luctus. Nulla ultrices tincidunt ipsum non dapibus. Morbi imperdiet vehicula sollicitudin. Sed non hendrerit augue, eget malesuada nulla. Fusce vestibulum sit amet ante venenatis porttitor. Integer efficitur feugiat sem eu sagittis. Cras aliquam dignissim lacus id vehicula. Quisque at felis sapien. Aenean sit amet ipsum feugiat, condimentum quam eget, convallis metus. Quisque vehicula lobortis lorem, ac sodales justo laoreet pretium. Vivamus eget quam augue.</p>\n<!-- /wp:paragraph -->','Testing Three','','inherit','closed','closed','','9-revision-v1','','','2020-10-14 18:37:21','2020-10-14 18:37:21','',9,'http://localhost:8888/?p=10',0,'revision','',0),(11,1,'2020-10-14 18:38:32','2020-10-14 18:38:32','<!-- wp:paragraph -->\n<p>Integer sodales dapibus erat, non tincidunt metus imperdiet ac. Aenean porttitor non arcu quis convallis. Duis rhoncus eros urna, ac finibus magna rhoncus quis. Praesent dictum ante dictum venenatis eleifend. Curabitur eleifend magna imperdiet viverra faucibus. Vivamus rhoncus felis nec justo fringilla placerat. Donec lacinia justo massa, ut posuere massa ornare id. Etiam efficitur tristique neque vel egestas.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Donec egestas et metus eget vulputate. Praesent efficitur, metus nec consectetur ultricies, nulla libero faucibus ex, eu posuere lectus justo eget nunc. Nullam at orci tortor. Vestibulum bibendum neque vitae mollis rhoncus. Integer eu dapibus risus, et pharetra risus. Phasellus elementum consequat lorem. Donec pretium placerat ligula, non luctus diam porttitor cursus. Phasellus in orci bibendum, scelerisque leo a, mattis metus. Etiam metus felis, varius a pulvinar eu, tincidunt imperdiet dolor. Integer sodales nec dui non viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aliquam erat volutpat. Nullam cursus lacus mauris, sit amet fringilla lectus condimentum at. Nam maximus sapien sed tortor tincidunt, in posuere ante ornare. Nullam diam ipsum, volutpat ac diam at, posuere ornare odio. Proin eu accumsan erat.</p>\n<!-- /wp:paragraph -->','Testing Quatro','','publish','open','open','','testing-quatro','','','2020-10-14 18:38:32','2020-10-14 18:38:32','',0,'http://localhost:8888/?p=11',0,'post','',0),(12,1,'2020-10-14 18:38:32','2020-10-14 18:38:32','<!-- wp:paragraph -->\n<p>Integer sodales dapibus erat, non tincidunt metus imperdiet ac. Aenean porttitor non arcu quis convallis. Duis rhoncus eros urna, ac finibus magna rhoncus quis. Praesent dictum ante dictum venenatis eleifend. Curabitur eleifend magna imperdiet viverra faucibus. Vivamus rhoncus felis nec justo fringilla placerat. Donec lacinia justo massa, ut posuere massa ornare id. Etiam efficitur tristique neque vel egestas.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Donec egestas et metus eget vulputate. Praesent efficitur, metus nec consectetur ultricies, nulla libero faucibus ex, eu posuere lectus justo eget nunc. Nullam at orci tortor. Vestibulum bibendum neque vitae mollis rhoncus. Integer eu dapibus risus, et pharetra risus. Phasellus elementum consequat lorem. Donec pretium placerat ligula, non luctus diam porttitor cursus. Phasellus in orci bibendum, scelerisque leo a, mattis metus. Etiam metus felis, varius a pulvinar eu, tincidunt imperdiet dolor. Integer sodales nec dui non viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aliquam erat volutpat. Nullam cursus lacus mauris, sit amet fringilla lectus condimentum at. Nam maximus sapien sed tortor tincidunt, in posuere ante ornare. Nullam diam ipsum, volutpat ac diam at, posuere ornare odio. Proin eu accumsan erat.</p>\n<!-- /wp:paragraph -->','Testing Quatro','','inherit','closed','closed','','11-revision-v1','','','2020-10-14 18:38:32','2020-10-14 18:38:32','',11,'http://localhost:8888/?p=12',0,'revision','',0),(13,1,'2020-10-14 18:39:12','2020-10-14 18:39:12','<!-- wp:paragraph -->\n<p>Mauris vel lacinia sem. Ut placerat augue molestie turpis dictum vulputate. Duis malesuada interdum quam, et efficitur lectus varius non. Etiam vel urna in nulla eleifend placerat non eu nisl. Etiam pellentesque, diam sed gravida convallis, nisl tortor congue sapien, vel ornare dolor erat at sem. Nunc egestas sodales urna, quis suscipit purus convallis a. Suspendisse congue congue nulla ut tempor. Quisque odio eros, euismod sed hendrerit sit amet, luctus ac nulla. Donec in enim lobortis, luctus velit bibendum, suscipit sapien. Pellentesque pretium, massa nec suscipit faucibus, massa mi eleifend nisl, sed finibus nisi lacus eget mauris.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Duis velit ipsum, congue vitae elit commodo, facilisis tristique tellus. Phasellus id hendrerit eros, non gravida nisi. Nulla facilisi. Quisque accumsan pharetra ipsum, id elementum turpis maximus quis. Nam semper massa quis mi feugiat, non aliquet quam lobortis. Nam porta, diam finibus sodales tincidunt, erat elit lobortis purus, et semper risus nulla sit amet odio. Aliquam felis arcu, malesuada vel nulla non, semper porttitor risus. Etiam fermentum nulla non ligula ornare placerat. Ut lobortis, purus id aliquet pulvinar, libero arcu rhoncus lectus, vitae interdum risus mauris quis metus. Proin pretium nunc ultrices ante vestibulum, vitae tincidunt libero faucibus. Phasellus non purus vehicula, vehicula lacus non, aliquet ex. Quisque vestibulum ultrices dolor. Nam lobortis pulvinar volutpat. Maecenas mollis sodales erat vel lobortis.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Pellentesque vitae est eu eros rutrum dapibus. Suspendisse eu interdum odio. Nunc lobortis scelerisque volutpat. Ut feugiat imperdiet lobortis. Sed porta nec libero vel luctus. Vivamus turpis sapien, ultricies convallis nulla porta, sodales eleifend neque. Phasellus sit amet libero leo. Nullam suscipit vehicula nisi, id tristique massa imperdiet ac.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Pellentesque at blandit nulla. Donec ac dapibus urna. Nulla facilisi. Aliquam quis tincidunt ante. Vivamus interdum orci quis ipsum elementum, nec congue est euismod. Mauris diam justo, luctus vel luctus eu, rutrum a felis. Proin vel est in quam finibus rhoncus. Donec libero odio, sagittis eu metus id, eleifend iaculis enim. Nulla eget lorem vulputate, condimentum nunc a, venenatis est.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Nullam sagittis maximus turpis, pharetra malesuada tortor. Etiam ut finibus urna, eget sodales erat. Etiam laoreet mi dapibus dapibus pulvinar. In rutrum posuere sem, finibus consequat dui imperdiet vel. Pellentesque euismod lobortis lacinia. Nam sed arcu arcu. Mauris ornare enim ipsum, vel semper augue scelerisque faucibus. Maecenas venenatis facilisis lacus sit amet laoreet. Curabitur auctor massa a lectus tincidunt, non vulputate enim pretium. Aenean euismod tellus dolor, sed ultricies velit hendrerit in. Donec elementum est ut sollicitudin dignissim.</p>\n<!-- /wp:paragraph -->','Cinq Cinq Cinq Cinq Cinq','','publish','open','open','','cinq-cinq-cinq-cinq-cinq','','','2020-10-14 18:39:12','2020-10-14 18:39:12','',0,'http://localhost:8888/?p=13',0,'post','',0),(14,1,'2020-10-14 18:39:12','2020-10-14 18:39:12','<!-- wp:paragraph -->\n<p>Mauris vel lacinia sem. Ut placerat augue molestie turpis dictum vulputate. Duis malesuada interdum quam, et efficitur lectus varius non. Etiam vel urna in nulla eleifend placerat non eu nisl. Etiam pellentesque, diam sed gravida convallis, nisl tortor congue sapien, vel ornare dolor erat at sem. Nunc egestas sodales urna, quis suscipit purus convallis a. Suspendisse congue congue nulla ut tempor. Quisque odio eros, euismod sed hendrerit sit amet, luctus ac nulla. Donec in enim lobortis, luctus velit bibendum, suscipit sapien. Pellentesque pretium, massa nec suscipit faucibus, massa mi eleifend nisl, sed finibus nisi lacus eget mauris.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Duis velit ipsum, congue vitae elit commodo, facilisis tristique tellus. Phasellus id hendrerit eros, non gravida nisi. Nulla facilisi. Quisque accumsan pharetra ipsum, id elementum turpis maximus quis. Nam semper massa quis mi feugiat, non aliquet quam lobortis. Nam porta, diam finibus sodales tincidunt, erat elit lobortis purus, et semper risus nulla sit amet odio. Aliquam felis arcu, malesuada vel nulla non, semper porttitor risus. Etiam fermentum nulla non ligula ornare placerat. Ut lobortis, purus id aliquet pulvinar, libero arcu rhoncus lectus, vitae interdum risus mauris quis metus. Proin pretium nunc ultrices ante vestibulum, vitae tincidunt libero faucibus. Phasellus non purus vehicula, vehicula lacus non, aliquet ex. Quisque vestibulum ultrices dolor. Nam lobortis pulvinar volutpat. Maecenas mollis sodales erat vel lobortis.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Pellentesque vitae est eu eros rutrum dapibus. Suspendisse eu interdum odio. Nunc lobortis scelerisque volutpat. Ut feugiat imperdiet lobortis. Sed porta nec libero vel luctus. Vivamus turpis sapien, ultricies convallis nulla porta, sodales eleifend neque. Phasellus sit amet libero leo. Nullam suscipit vehicula nisi, id tristique massa imperdiet ac.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Pellentesque at blandit nulla. Donec ac dapibus urna. Nulla facilisi. Aliquam quis tincidunt ante. Vivamus interdum orci quis ipsum elementum, nec congue est euismod. Mauris diam justo, luctus vel luctus eu, rutrum a felis. Proin vel est in quam finibus rhoncus. Donec libero odio, sagittis eu metus id, eleifend iaculis enim. Nulla eget lorem vulputate, condimentum nunc a, venenatis est.</p>\n<!-- /wp:paragraph -->\n\n<!-- wp:paragraph -->\n<p>Nullam sagittis maximus turpis, pharetra malesuada tortor. Etiam ut finibus urna, eget sodales erat. Etiam laoreet mi dapibus dapibus pulvinar. In rutrum posuere sem, finibus consequat dui imperdiet vel. Pellentesque euismod lobortis lacinia. Nam sed arcu arcu. Mauris ornare enim ipsum, vel semper augue scelerisque faucibus. Maecenas venenatis facilisis lacus sit amet laoreet. Curabitur auctor massa a lectus tincidunt, non vulputate enim pretium. Aenean euismod tellus dolor, sed ultricies velit hendrerit in. Donec elementum est ut sollicitudin dignissim.</p>\n<!-- /wp:paragraph -->','Cinq Cinq Cinq Cinq Cinq','','inherit','closed','closed','','13-revision-v1','','','2020-10-14 18:39:12','2020-10-14 18:39:12','',13,'http://localhost:8888/?p=14',0,'revision','',0);
/*!40000 ALTER TABLE `wp_posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_term_relationships`
--

DROP TABLE IF EXISTS `wp_term_relationships`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_term_relationships` (
  `object_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `term_taxonomy_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `term_order` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`object_id`,`term_taxonomy_id`),
  KEY `term_taxonomy_id` (`term_taxonomy_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_term_relationships`
--

LOCK TABLES `wp_term_relationships` WRITE;
/*!40000 ALTER TABLE `wp_term_relationships` DISABLE KEYS */;
INSERT INTO `wp_term_relationships` VALUES (1,1,0),(5,1,0),(7,1,0),(9,1,0),(11,1,0),(13,1,0);
/*!40000 ALTER TABLE `wp_term_relationships` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_term_taxonomy`
--

DROP TABLE IF EXISTS `wp_term_taxonomy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_term_taxonomy` (
  `term_taxonomy_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `term_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `taxonomy` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `description` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `parent` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`term_taxonomy_id`),
  UNIQUE KEY `term_id_taxonomy` (`term_id`,`taxonomy`),
  KEY `taxonomy` (`taxonomy`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_term_taxonomy`
--

LOCK TABLES `wp_term_taxonomy` WRITE;
/*!40000 ALTER TABLE `wp_term_taxonomy` DISABLE KEYS */;
INSERT INTO `wp_term_taxonomy` VALUES (1,1,'category','',0,6);
/*!40000 ALTER TABLE `wp_term_taxonomy` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_termmeta`
--

DROP TABLE IF EXISTS `wp_termmeta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_termmeta` (
  `meta_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `term_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `meta_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_value` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`meta_id`),
  KEY `term_id` (`term_id`),
  KEY `meta_key` (`meta_key`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_termmeta`
--

LOCK TABLES `wp_termmeta` WRITE;
/*!40000 ALTER TABLE `wp_termmeta` DISABLE KEYS */;
/*!40000 ALTER TABLE `wp_termmeta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_terms`
--

DROP TABLE IF EXISTS `wp_terms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_terms` (
  `term_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `slug` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `term_group` bigint(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`term_id`),
  KEY `slug` (`slug`(191)),
  KEY `name` (`name`(191))
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_terms`
--

LOCK TABLES `wp_terms` WRITE;
/*!40000 ALTER TABLE `wp_terms` DISABLE KEYS */;
INSERT INTO `wp_terms` VALUES (1,'Uncategorized','uncategorized',0);
/*!40000 ALTER TABLE `wp_terms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_usermeta`
--

DROP TABLE IF EXISTS `wp_usermeta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_usermeta` (
  `umeta_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `meta_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_value` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`umeta_id`),
  KEY `user_id` (`user_id`),
  KEY `meta_key` (`meta_key`(191))
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_usermeta`
--

LOCK TABLES `wp_usermeta` WRITE;
/*!40000 ALTER TABLE `wp_usermeta` DISABLE KEYS */;
INSERT INTO `wp_usermeta` VALUES (1,1,'nickname','admin'),(2,1,'first_name',''),(3,1,'last_name',''),(4,1,'description',''),(5,1,'rich_editing','true'),(6,1,'syntax_highlighting','true'),(7,1,'comment_shortcuts','false'),(8,1,'admin_color','fresh'),(9,1,'use_ssl','0'),(10,1,'show_admin_bar_front','true'),(11,1,'locale',''),(12,1,'wp_capabilities','a:1:{s:13:\"administrator\";b:1;}'),(13,1,'wp_user_level','10'),(14,1,'dismissed_wp_pointers',''),(15,1,'show_welcome_panel','1'),(16,1,'session_tokens','a:1:{s:64:\"37deaf05ed8d17fdaa07e64d663215496dc3e51c7b89f977670e4e06180ffcf9\";a:4:{s:10:\"expiration\";i:1603909937;s:2:\"ip\";s:10:\"172.22.0.1\";s:2:\"ua\";s:121:\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36\";s:5:\"login\";i:1602700337;}}'),(17,1,'wp_dashboard_quick_press_last_post_id','4'),(18,1,'community-events-location','a:1:{s:2:\"ip\";s:10:\"172.22.0.0\";}');
/*!40000 ALTER TABLE `wp_usermeta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wp_users`
--

DROP TABLE IF EXISTS `wp_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wp_users` (
  `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_login` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_pass` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_nicename` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_url` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_registered` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `user_activation_key` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user_status` int(11) NOT NULL DEFAULT '0',
  `display_name` varchar(250) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`),
  KEY `user_login_key` (`user_login`),
  KEY `user_nicename` (`user_nicename`),
  KEY `user_email` (`user_email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wp_users`
--

LOCK TABLES `wp_users` WRITE;
/*!40000 ALTER TABLE `wp_users` DISABLE KEYS */;
INSERT INTO `wp_users` VALUES (1,'admin','$P$BKVMOwYYjuNym14vSaGsJ2xkCkni/k/','admin','admin@test.com','http://localhost:8888','2020-10-14 18:32:03','',0,'admin');
/*!40000 ALTER TABLE `wp_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-10-14 20:00:00