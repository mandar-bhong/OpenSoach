-- =============================================================================
-- Filename   :  schema_1.0.0.sql

-- Version    :  1.0

-- Created on :  22-Feb-2018

-- Author     :  sanjay.sawant@opensoach.com

-- Description:
-- - db schema script to create database for spl_hkt_master. a master database applicable for all HKT node databases.

-- ==========================================================================
drop database if exists spl_hkt_master;
create database spl_hkt_master DEFAULT CHARACTER SET utf8;
use spl_hkt_master;

--
-- Table structure for table `spl_hkt_master_sp_category`
--

CREATE TABLE `spl_hkt_master_sp_category_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `spc_name` varchar(50) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `spc_name_UNIQUE` (`spc_name`)
) ENGINE=InnoDB COMMENT='Short Name for Table: spc';


--
-- Table structure for table `spl_hkt_master_task_lib_tbl`
--

CREATE TABLE `spl_hkt_master_task_lib_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `task_name` varchar(50) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `task_name_UNIQUE` (`task_name`)
) ENGINE=InnoDB COMMENT='Short Name for Table: mtask';

--
-- Table structure for table `spl_hkt_master_spc_task_lib_tbl`
--

CREATE TABLE `spl_hkt_master_spc_task_lib_tbl` (
  `spc_id_fk` int(10) unsigned NOT NULL,
  `mtask_id_fk` int(10) unsigned NOT NULL,
  PRIMARY KEY (`spc_id_fk`,`mtask_id_fk`),
  KEY `fk_spct_mtask_idx` (`mtask_id_fk`),
  KEY `fk_spct_spc_idx` (`spc_id_fk`),
  CONSTRAINT `fk_spct_mtask` FOREIGN KEY (`mtask_id_fk`) REFERENCES `spl_hkt_master_task_lib_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_spct_spc` FOREIGN KEY (`spc_id_fk`) REFERENCES `spl_hkt_master_sp_category_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: spct';
