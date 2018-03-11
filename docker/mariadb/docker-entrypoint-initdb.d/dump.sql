CREATE DATABASE  IF NOT EXISTS `prismapptest` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `prismapptest`;
-- MySQL dump 10.13  Distrib 5.7.17, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: prismapptest
-- ------------------------------------------------------
-- Server version	5.5.5-10.1.30-MariaDB-1~jessie

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
-- Table structure for table `room`
--

DROP TABLE IF EXISTS `room`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `room` (
  `id` varchar(36) NOT NULL,
  `message` longtext NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `room`
--

LOCK TABLES `room` WRITE;
/*!40000 ALTER TABLE `room` DISABLE KEYS */;
INSERT INTO `room` VALUES ('02b18ff2-0ba8-4e70-a12d-a3860911acc0','haii guys','2018-03-11 04:38:32'),('06f5e2f2-e244-46b2-a92a-a540663a3479','pada kerja dimana sekarang?','2018-03-10 21:46:06'),('1635fee7-9d6f-4d32-8df3-c9e26b5935c0','btw sekarang kerja dimana?','2018-03-10 11:43:43'),('17e45703-3cc5-4628-b738-cdd9232004cb','payah dah, kumpul lah','2018-03-10 21:44:42'),('21244cfa-1629-46d1-8d45-6cf6e48e674e','pada kerja dimana sekarang?','2018-03-10 21:45:23'),('2447601e-e3cf-4703-9036-bb45bd4be5cc','payah dah, kumpul lah','2018-03-10 21:40:21'),('26bc3abb-d5b5-47db-b4f5-4cc6564a17a8','pada kerja dimana sekarang?','2018-03-10 21:46:17'),('26d156a5-e373-4f10-99e9-0a339783a28b','broo','2018-03-10 21:30:06'),('28faba6d-7c8b-46b7-ae44-2e470d52f67f','haii guys','2018-03-11 04:39:20'),('2b0ffc67-186b-4f52-a5a1-8edfea31733c','cek di database','2018-03-11 11:31:34'),('33d7e913-e0c7-487a-a309-4f0e699fbb24','broo','2018-03-10 21:28:24'),('3f516c18-af8b-4ea1-a954-8192d32924cd','broo','2018-03-10 21:34:14'),('436b24fc-3fba-4bc6-a5a3-efefc6908354','haii guys','2018-03-10 22:08:24'),('488ae547-7641-446f-b928-39a3a947426b','emana aje','2018-03-10 21:39:46'),('5031549f-83ae-44fb-9be6-3af15c400aa3','haii guys','2018-03-10 22:09:22'),('50e98598-54cd-4a19-a44a-66e2770dd41d','haii guys','2018-03-11 04:38:24'),('5bbc2bb8-efd6-4630-8453-e152d22c688a','pada kerja dimana sekarang?','2018-03-10 21:46:08'),('5ca4b660-c4a2-42a5-a9fb-4bd81baa8a04','haii guys','2018-03-10 22:08:21'),('5f886e77-2e80-42e8-9a08-f28be94afcfe','Apa kabar?','2018-03-10 22:20:00'),('627c9343-16cc-4e21-ac11-38fcf52dea7b','pada kerja dimana sekarang?','2018-03-10 21:46:08'),('6724ac5b-75b9-48b5-9d31-c00fe9905353','pada kerja dimana sekarang?','2018-03-10 21:46:18'),('767e0d17-d6f9-4ba4-a820-becfbf3eb78b','haii guys','2018-03-10 22:09:23'),('778beac9-a484-476e-9095-a384e49c95c5','pada kerja dimana sekarang?','2018-03-10 21:46:10'),('77e9595d-4bba-44f4-99a6-1b4fe7415922','hai','2018-03-10 08:09:32'),('78574690-e136-4234-8f80-786c8f83ea12','haii guys','2018-03-10 22:09:22'),('7d0b6940-efdc-4dc4-912c-9ef98ec52adb','haii guys','2018-03-10 22:09:26'),('853efe44-ec81-4a24-a58d-eec80847c0a8','hai realtime  2','2018-03-10 20:28:49'),('896a6697-23ad-407f-93e5-0ea118267753','Apa kabarr, abis ini di commit aja yah','2018-03-11 05:27:43'),('8b07a058-c6c9-4baa-97cb-dfcf912c3019','pada kerja dimana sekarang?','2018-03-10 21:46:09'),('8c77c5a5-867a-48bd-8bff-7515a0e4b503','kabar baik baik saja','2018-03-10 11:21:45'),('8ef24b3f-a564-4972-ab72-a6ffddf3e8b3','kabar baik baik saja','2018-03-10 11:20:41'),('90c0f23b-c29b-4a8c-abc0-2efab01f0eab','apa kabar','2018-03-10 08:09:37'),('95546af2-ec22-492c-bf07-41c29404252a','apa kabar?','2018-03-10 21:39:19'),('96481681-68f0-45cb-b545-4806008bd632','pada kerja dimana sekarang?','2018-03-10 21:46:18'),('9d95dcf3-60ce-45e7-ad38-3f73dc63f962','hai realtime  2','2018-03-10 19:28:02'),('a351d9f5-85e8-433f-a074-3f76838197c4','pada kerja dimana sekarang?','2018-03-10 21:46:16'),('a523726d-924a-4264-9eaf-960f7d349220','apa kabar','2018-03-10 11:10:53'),('a5c081be-96d9-437d-8162-669692cb4b85','haii guys','2018-03-10 22:09:25'),('a7bba843-a87e-442e-a3f2-4ff63cadbf03','haii guys','2018-03-10 22:08:23'),('aadbbcba-274d-48b2-8f6b-95e6119aa6ce','hai realtime ','2018-03-10 19:27:48'),('b3f0de00-6fc6-4938-b73a-45a7a5c0c03d','haii guys','2018-03-10 22:08:27'),('b7f68053-d301-475d-a702-4a5984b28ce2','haii guys','2018-03-10 22:08:25'),('c085972c-6854-46ff-aa50-ffc80a6885a2','pada kerja dimana sekarang?','2018-03-10 21:46:07'),('c0a376fe-018a-4006-b081-088218dda058','pada kerja dimana sekarang?','2018-03-10 21:46:19'),('c129de63-4a5f-4a2a-b534-07a64819babf','haii guys','2018-03-10 22:09:19'),('c754d77e-0173-4545-b179-ca0bec5fa684','check after refactor','2018-03-11 11:56:54'),('d382991c-9b71-4e48-a34a-2f9ecd65ba2a','haii guys','2018-03-10 22:09:27'),('d66e62a6-5a17-4ec1-b135-d963277cb558','haii guys','2018-03-10 22:09:24'),('d816bf44-590c-423a-a245-8e32c43f7e5a',' test test','2018-03-10 22:21:08'),('db409290-2fdb-4ac5-8c7b-e40606f7a994','haii guys','2018-03-10 22:09:21'),('dda7767a-1cb0-42bf-9c03-effcf5eac0f6','broo','2018-03-10 21:37:56'),('de95f276-8460-4567-828d-bb58b5ad6ea8','pada kerja dimana sekarang?','2018-03-10 21:46:17'),('e082ac5d-df06-47db-a6bb-a5b46668cc7e','haii guys','2018-03-11 04:40:12'),('e18cc8e2-2ae0-465e-8c28-eefd73ce4fbd','hai realtime','2018-03-10 19:25:19'),('e2f97ac9-1573-4aba-9021-c317cf7a0020','haii guys','2018-03-10 22:04:02'),('e3525c85-75ab-4718-be71-2fcfab435a99','Sibuk apa Kalian?','2018-03-10 22:20:18'),('e889c750-3e27-4267-9cf9-64202a724eda','pada kerja dimana sekarang?','2018-03-10 21:46:16'),('e9d67a42-e02a-4593-9a95-fb3cc667348a','haii guys','2018-03-10 22:07:40'),('ea602003-6219-49a1-9e03-1cf09199a32c','haii guys','2018-03-10 22:09:20'),('f67d9bb8-1019-462c-9ecc-0418e5e48a52','hai realtime  2','2018-03-10 20:28:43'),('f792e540-15c3-4654-8961-25a6be13b46f','hai','2018-03-11 11:30:48'),('f993c1b5-b7ca-4e0b-987b-2fc9240e02d0','  nomer lu pada ganti kah?','2018-03-10 21:45:05'),('fb678fcd-a859-4cae-b309-2e8b30e3e1c9','Apa kabarr, abis ini di commit aja yah','2018-03-11 05:14:59');
/*!40000 ALTER TABLE `room` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-03-11 19:53:25
