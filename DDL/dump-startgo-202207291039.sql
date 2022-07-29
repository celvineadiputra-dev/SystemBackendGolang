-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: localhost    Database: startgo
-- ------------------------------------------------------
-- Server version	5.5.5-10.4.22-MariaDB

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
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `created_at` date NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES (1,'AC','2022-07-29','2022-07-29 08:28:30');
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `email` varchar(60) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (17,'HELLO','Hello@mail.com','IX','2022-07-29 09:42:25','2022-07-29 09:42:25'),(18,'HELLO','Hello@mail.com','IX','2022-07-29 10:00:00','2022-07-29 10:00:00'),(19,'HELLO','Hello@mail.com','IX','2022-07-29 10:02:03','2022-07-29 10:02:03'),(20,'HELLO','Hello@mail.com','IX','2022-07-29 10:03:17','2022-07-29 10:03:17'),(21,'HELLO','Hello@mail.com','IX','2022-07-29 10:06:07','2022-07-29 10:06:07'),(22,'HELLO','Hello@mail.com','IX','2022-07-29 10:06:55','2022-07-29 10:06:55'),(23,'HELLO','Hello@mail.com','IX','2022-07-29 10:07:57','2022-07-29 10:07:57'),(24,'HELLO','Hello@mail.com','IX','2022-07-29 10:08:23','2022-07-29 10:08:23'),(25,'HELLO','Hello@mail.com','IX','2022-07-29 10:11:54','2022-07-29 10:11:54'),(26,'HELLO','Hello@mail.com','IX','2022-07-29 10:16:24','2022-07-29 10:16:24'),(27,'HELLO','Hello@mail.com','IX','2022-07-29 10:17:06','2022-07-29 10:17:06'),(28,'HELLO','Hello@mail.com','IX','2022-07-29 10:17:47','2022-07-29 10:17:47'),(29,'HELLO','Hello@mail.com','IX','2022-07-29 10:19:04','2022-07-29 10:19:04'),(30,'HELLO','Hello@mail.com','IX','2022-07-29 10:20:07','2022-07-29 10:20:07'),(31,'HELLO','Hello@mail.com','IX','2022-07-29 10:20:26','2022-07-29 10:20:26'),(32,'HELLO','Hello@mail.com','IX','2022-07-29 10:20:52','2022-07-29 10:20:52'),(33,'HELLO','Hello@mail.com','IX','2022-07-29 10:23:10','2022-07-29 10:23:10'),(34,'HELLO','Hello@mail.com','IX','2022-07-29 10:23:38','2022-07-29 10:23:38'),(35,'HELLO','Hello@mail.com','IX','2022-07-29 10:29:36','2022-07-29 10:29:36'),(36,'HELLO','Hello@mail.com','IX','2022-07-29 10:29:55','2022-07-29 10:29:55');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'startgo'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-29 10:39:38
