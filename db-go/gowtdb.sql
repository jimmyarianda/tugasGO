-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 09, 2022 at 04:43 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 7.4.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gowtdb`
--

-- --------------------------------------------------------

--
-- Table structure for table `pegawai`
--

CREATE TABLE `pegawai` (
  `id` int(11) NOT NULL,
  `task` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `assignee` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `deadline` date DEFAULT NULL,
  `status` enum('True','False') COLLATE utf8_unicode_ci NOT NULL DEFAULT 'False'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `pegawai`
--

INSERT INTO `pegawai` (`id`, `task`, `assignee`, `deadline`, `status`) VALUES
(1, 'Tugas-GO1', 'Jimmy', '2022-08-08', 'True'),
(2, 'Tugas-GO2', 'Yudi', '2022-08-08', 'True'),
(3, 'Tugas-GO3', 'Ricky', '2022-08-08', 'True'),
(4, 'Tugas-GO4', 'Kiting', '2022-08-08', 'True'),
(6, 'Tugas-GO5', 'Candra', '2022-08-09', 'False'),
(7, 'Tugas-GO6', 'Agus', '2022-08-09', 'False'),
(8, 'Tugas-GO6', 'Yuli', '2022-08-09', 'False'),
(9, 'Tugas-GO7', 'Vico', '2022-08-09', 'False'),
(10, 'Tugas-GO8', 'Gayu', '2022-08-09', 'False'),
(11, 'Tugas-GO9', 'Rani', '2022-08-09', 'False'),
(12, 'Tugas-GO10', 'Budi', '2022-08-09', 'False'),
(13, 'Tugas-GO11', 'Ucup', '2022-08-09', 'False');

-- --------------------------------------------------------

--
-- Table structure for table `tools`
--

CREATE TABLE `tools` (
  `id` int(11) NOT NULL,
  `name` varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  `category` varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  `url` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `rating` int(11) DEFAULT NULL,
  `notes` text COLLATE utf8_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `pegawai`
--
ALTER TABLE `pegawai`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tools`
--
ALTER TABLE `tools`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `pegawai`
--
ALTER TABLE `pegawai`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `tools`
--
ALTER TABLE `tools`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
