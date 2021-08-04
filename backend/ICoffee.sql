-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 30, 2021 at 10:43 AM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 8.0.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ICoffee`
--

-- --------------------------------------------------------

--
-- Table structure for table `ico_admin_meta`
--

CREATE TABLE `ico_admin_meta` (
  `id` int(11) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `meta_key` varchar(255) DEFAULT NULL,
  `meta_value` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_admin_meta`
--

INSERT INTO `ico_admin_meta` (`id`, `admin_id`, `meta_key`, `meta_value`) VALUES
(1, 25, 'birth_date', '1370/08/23'),
(2, 25, 'address', 'test address'),
(3, 25, 'avatar', ''),
(4, 25, 'phone', '05138474541'),
(5, 25, 'mobile', '09352399329'),
(11, 27, 'birth_date', '1370/08/23'),
(12, 27, 'address', 'test address'),
(13, 27, 'avatar', ''),
(14, 27, 'phone', '05138474541'),
(15, 27, 'mobile', '09352399329');

-- --------------------------------------------------------

--
-- Table structure for table `ico_card`
--

CREATE TABLE `ico_card` (
  `id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `card` text NOT NULL,
  `created_at` varchar(30) NOT NULL DEFAULT current_timestamp(),
  `status` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `ico_category`
--

CREATE TABLE `ico_category` (
  `id` int(11) NOT NULL,
  `type` varchar(255) NOT NULL,
  `category_name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `parent` int(11) NOT NULL,
  `status` tinyint(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_category`
--

INSERT INTO `ico_category` (`id`, `type`, `category_name`, `slug`, `parent`, `status`) VALUES
(4, 'food-industry', 'coffee shop', 'coffee-shop', 0, 0),
(5, 'food-industry', 'restaurant', 'restaurant', 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `ico_jwt_token`
--

CREATE TABLE `ico_jwt_token` (
  `id` int(11) NOT NULL,
  `token` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_jwt_token`
--

INSERT INTO `ico_jwt_token` (`id`, `token`) VALUES
(22, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ3NDc2Nn0.omCvPorI_TDc-xhrxGvT6NGRAl-geVH8EoguLeztFb8'),
(23, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ3NjU4Mn0.x6Qyae3Vvj1KatdJUCgZ5rKC_OH28Nj9KjXRhgH9qOo'),
(24, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ3NzE0N30.OFJLpisOJJXF7rCvPIMIh-eBmYE8ckpuCDCscAeh64o'),
(25, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ3NzMxM30.NApuJeSPehjX7RFFZ2_RJiLBrIYHxyt3col-gK5m2U8'),
(26, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ5MTE4MX0.dbCTmWC5akssvQ4L0WUVy8PJlFrM7qH6ViEaiNXBjHE'),
(27, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ5MjEyN30.TQ_lsCmzEZD9Dlgfv0SlYxVre4EXILXcoFxGnC-v6FM'),
(28, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ5NTg0NH0.x68_qf7S4_znEtMBvyWPQJVuS8E5fKwYj-17mNfpWcg'),
(29, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ5NjU0OH0.aKLo70x70PPvoFtG9iAvT9hF1Rwf38cT5D2uRZfVook'),
(30, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzQ5NzYxMn0.3J3-VCkxrN-LE-LDAq2_iu614VKtCMvsNFECQSotlVc'),
(31, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUwNTgxM30.wsCAvgrIC94ToJkVLLX-vZrFPywQiPmmRug4vUGXnRI'),
(32, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUwNjA1N30.LThcz5JCkdYnF9j-d-Zrn3qZXdbDYDpAVwUhCZwWaSU'),
(33, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUwNzM5OX0.uMYBwnHcJuTguZRDUzjrkrWCUpULdsjZUn9zDzcVQ3Q'),
(34, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUwOTg2NH0.aPkX-Is99AkGRxyw_HNY_HZUbYeyy-C4i3a474bXNAw'),
(35, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUxMTI5N30.gRFeCPlc9tJmLJzr1YoUBhtxEM-vsNdBkXEhTrAcBWI'),
(36, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUxMTY2MX0.vjEvrCkLfy1tW8xIdb6JKOsr96S4tsNWMxAVc_rnblg'),
(37, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUxMTkzNH0.wcvOp4NaTS-HWbpBfnAVH3d6MLpr6_bu9s0kYwAE9VE'),
(38, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUxMjQ1MX0.iNGSu1Pilf7-yvhiNZQ029G2TO1BlCiNENjTBrAGi50'),
(39, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUxMzcxMn0.7YSsOnXwDAXT5NcgNEbPMEiC8-ps3LPZjps85EXwoNg'),
(40, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzUxOTc3OX0.E7M040ycCMcVyFHTYsUipZtccrmJrFMfaf83_YxSYTQ'),
(41, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzU3NjA3OX0.llTisU5hxwJBodI2hF5aCei0cSf2AiOhobZie2zT74E'),
(42, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzU4NTI4M30.ZiRTqVPTqI8RwLH7_QnvfiDjGVNWblVGLPdjIX0gTBE'),
(43, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzU5MDA4NX0.yORk3v19n-nmDN4tlHMIDkg10LJ7nKK0fzCzLW_vWaY'),
(44, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzYzNTUwNn0.UYTnc7osh7FWm1kOnNOpoYd1JQaoK4H-7j4E20fZcVc'),
(45, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsInVzZXJuYW1lIjoia2Fuc2FpIiwiZmlyc3RuYW1lIjoiYWxpcmV6YSIsImxhc3RuYW1lIjoic2FmZmFyIiwiYmlydGhkYXRlIjoiMTM3MC8wOC8yMyIsInJvbGUiOjIsImV4cCI6MTYyNzYzNTk0NH0.7VnWSmfRhrDJ4oXPgGkm1JCsJo4mGjMKtLugpTWU4ng');

-- --------------------------------------------------------

--
-- Table structure for table `ico_order`
--

CREATE TABLE `ico_order` (
  `id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `products` text NOT NULL,
  `ref_id` varchar(255) NOT NULL,
  `order_token` varchar(255) NOT NULL,
  `authority` varchar(255) NOT NULL,
  `transport_price` varchar(255) NOT NULL,
  `total_price` varchar(255) NOT NULL,
  `discount_amount` varchar(255) NOT NULL,
  `discount_coupon` varchar(255) NOT NULL,
  `wallet_reduction` varchar(255) NOT NULL,
  `created_at` varchar(255) NOT NULL,
  `date_paid` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `status` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `ico_product`
--

CREATE TABLE `ico_product` (
  `id` int(11) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `status` tinyint(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_product`
--

INSERT INTO `ico_product` (`id`, `product_name`, `slug`, `status`) VALUES
(1, 'میلک شیک شکلات', 'mylkh-shykh-shkhlt', 0),
(5, 'test product', 'test-product', 0);

-- --------------------------------------------------------

--
-- Table structure for table `ico_product_meta`
--

CREATE TABLE `ico_product_meta` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `meta_key` varchar(255) NOT NULL,
  `meta_value` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_product_meta`
--

INSERT INTO `ico_product_meta` (`id`, `product_id`, `meta_key`, `meta_value`) VALUES
(3, 5, 'avatar', 'https://example.test/image-url.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `ico_roles`
--

CREATE TABLE `ico_roles` (
  `id` int(11) NOT NULL,
  `role_name` varchar(255) NOT NULL,
  `status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_roles`
--

INSERT INTO `ico_roles` (`id`, `role_name`, `status`) VALUES
(1, 'admin', 1),
(2, 'super-admin', 1),
(4, 'employer', 1);

-- --------------------------------------------------------

--
-- Table structure for table `ico_shop`
--

CREATE TABLE `ico_shop` (
  `id` int(11) NOT NULL,
  `owner_id` int(11) NOT NULL,
  `cat_id` int(11) NOT NULL,
  `shop_name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `status` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_shop`
--

INSERT INTO `ico_shop` (`id`, `owner_id`, `cat_id`, `shop_name`, `slug`, `status`) VALUES
(7, 27, 4, 'کافه وال', 'khfh-wl', 0);

-- --------------------------------------------------------

--
-- Table structure for table `ico_shop_meta`
--

CREATE TABLE `ico_shop_meta` (
  `id` int(11) NOT NULL,
  `shop_id` int(11) NOT NULL,
  `meta_key` varchar(255) NOT NULL,
  `meta_value` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `ico_taxonomy_products_cat`
--

CREATE TABLE `ico_taxonomy_products_cat` (
  `id` int(11) NOT NULL,
  `cat_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `ico_taxonomy_products_shop`
--

CREATE TABLE `ico_taxonomy_products_shop` (
  `id` int(11) NOT NULL,
  `shop_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_taxonomy_products_shop`
--

INSERT INTO `ico_taxonomy_products_shop` (`id`, `shop_id`, `product_id`) VALUES
(1, 7, 1),
(5, 7, 5);

-- --------------------------------------------------------

--
-- Table structure for table `ico_usermeta`
--

CREATE TABLE `ico_usermeta` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `meta_key` varchar(255) DEFAULT NULL,
  `meta_value` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `ico_users`
--

CREATE TABLE `ico_users` (
  `id` int(11) NOT NULL,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` varchar(255) DEFAULT NULL,
  `status` tinyint(2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `ico_user_admin`
--

CREATE TABLE `ico_user_admin` (
  `id` int(11) NOT NULL,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `created_at` varchar(255) DEFAULT NULL,
  `status` tinyint(2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ico_user_admin`
--

INSERT INTO `ico_user_admin` (`id`, `firstname`, `lastname`, `username`, `password`, `email`, `role_id`, `created_at`, `status`) VALUES
(25, 'alireza', 'saffar', 'kansai', '1f78c629944e103fe7db671208bfc8038dee47b9', 'a.khorasany@gmail.com', 2, '2021-07-26 00:00:02.724701195 +0430 +0430 m=+221.862843278', 1),
(27, 'alireza', 'saffar', 'quintin', '1f78c629944e103fe7db671208bfc8038dee47b9', 'a.khorasany@gmail.com', 1, '2021-07-26 15:51:48.182930761 +0430 +0430 m=+47.485192166', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `ico_admin_meta`
--
ALTER TABLE `ico_admin_meta`
  ADD PRIMARY KEY (`id`),
  ADD KEY `admin_id` (`admin_id`);

--
-- Indexes for table `ico_card`
--
ALTER TABLE `ico_card`
  ADD PRIMARY KEY (`id`),
  ADD KEY `customer_id` (`customer_id`);

--
-- Indexes for table `ico_category`
--
ALTER TABLE `ico_category`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ico_jwt_token`
--
ALTER TABLE `ico_jwt_token`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ico_order`
--
ALTER TABLE `ico_order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `customer_id` (`customer_id`);

--
-- Indexes for table `ico_product`
--
ALTER TABLE `ico_product`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ico_product_meta`
--
ALTER TABLE `ico_product_meta`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `ico_roles`
--
ALTER TABLE `ico_roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ico_shop`
--
ALTER TABLE `ico_shop`
  ADD PRIMARY KEY (`id`),
  ADD KEY `cat_id` (`cat_id`),
  ADD KEY `owner_id` (`owner_id`);

--
-- Indexes for table `ico_shop_meta`
--
ALTER TABLE `ico_shop_meta`
  ADD PRIMARY KEY (`id`),
  ADD KEY `shop_id` (`shop_id`);

--
-- Indexes for table `ico_taxonomy_products_cat`
--
ALTER TABLE `ico_taxonomy_products_cat`
  ADD PRIMARY KEY (`id`),
  ADD KEY `cat_id` (`cat_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `ico_taxonomy_products_shop`
--
ALTER TABLE `ico_taxonomy_products_shop`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `shop_id` (`shop_id`);

--
-- Indexes for table `ico_usermeta`
--
ALTER TABLE `ico_usermeta`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `ico_users`
--
ALTER TABLE `ico_users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ico_user_admin`
--
ALTER TABLE `ico_user_admin`
  ADD PRIMARY KEY (`id`),
  ADD KEY `role_id` (`role_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `ico_admin_meta`
--
ALTER TABLE `ico_admin_meta`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `ico_card`
--
ALTER TABLE `ico_card`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ico_category`
--
ALTER TABLE `ico_category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `ico_jwt_token`
--
ALTER TABLE `ico_jwt_token`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=46;

--
-- AUTO_INCREMENT for table `ico_order`
--
ALTER TABLE `ico_order`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ico_product`
--
ALTER TABLE `ico_product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `ico_product_meta`
--
ALTER TABLE `ico_product_meta`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `ico_roles`
--
ALTER TABLE `ico_roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `ico_shop`
--
ALTER TABLE `ico_shop`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `ico_shop_meta`
--
ALTER TABLE `ico_shop_meta`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ico_taxonomy_products_cat`
--
ALTER TABLE `ico_taxonomy_products_cat`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ico_taxonomy_products_shop`
--
ALTER TABLE `ico_taxonomy_products_shop`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `ico_usermeta`
--
ALTER TABLE `ico_usermeta`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ico_users`
--
ALTER TABLE `ico_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ico_user_admin`
--
ALTER TABLE `ico_user_admin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `ico_admin_meta`
--
ALTER TABLE `ico_admin_meta`
  ADD CONSTRAINT `ico_admin_meta_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `ico_user_admin` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_card`
--
ALTER TABLE `ico_card`
  ADD CONSTRAINT `ico_card_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `ico_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_order`
--
ALTER TABLE `ico_order`
  ADD CONSTRAINT `ico_order_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `ico_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_product_meta`
--
ALTER TABLE `ico_product_meta`
  ADD CONSTRAINT `ico_product_meta_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `ico_product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_shop`
--
ALTER TABLE `ico_shop`
  ADD CONSTRAINT `ico_shop_ibfk_1` FOREIGN KEY (`cat_id`) REFERENCES `ico_category` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_shop_meta`
--
ALTER TABLE `ico_shop_meta`
  ADD CONSTRAINT `ico_shop_meta_ibfk_1` FOREIGN KEY (`shop_id`) REFERENCES `ico_shop` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_taxonomy_products_cat`
--
ALTER TABLE `ico_taxonomy_products_cat`
  ADD CONSTRAINT `ico_taxonomy_products_cat_ibfk_1` FOREIGN KEY (`cat_id`) REFERENCES `ico_category` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `ico_taxonomy_products_cat_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `ico_product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_taxonomy_products_shop`
--
ALTER TABLE `ico_taxonomy_products_shop`
  ADD CONSTRAINT `ico_taxonomy_products_shop_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `ico_product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `ico_taxonomy_products_shop_ibfk_2` FOREIGN KEY (`shop_id`) REFERENCES `ico_shop` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_usermeta`
--
ALTER TABLE `ico_usermeta`
  ADD CONSTRAINT `ico_usermeta_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `ico_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ico_user_admin`
--
ALTER TABLE `ico_user_admin`
  ADD CONSTRAINT `ico_user_admin_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `ico_roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
