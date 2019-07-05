-- =============================================================================
-- Filename   :  schema_1.0.0.sql

-- Version    :  1.0

-- Created on :  22-Feb-2018

-- Author     :  sanjay.sawant@opensoach.com

-- Description:
-- - db schema script to create database for spl_hpft_master. a master database applicable for all HPFT node databases.

-- ==========================================================================
drop database if exists spl_hpft_master;
create database spl_hpft_master DEFAULT CHARACTER SET utf8;
use spl_hpft_master;

--
-- Table structure for table `spl_hkt_prod_config`
--

CREATE TABLE `spl_prod_master_config` (
	`config_key` VARCHAR(50) NOT NULL,
	`config_value` VARCHAR(500) NOT NULL DEFAULT '',
	`created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`config_key`)
) ENGINE=InnoDB COMMENT='Short Name for Table: config\r\nThis table will contain configuration for hkt product';

--
-- Table structure for table `spl_prod_sp_category_tbl`
--

CREATE TABLE `spl_prod_master_sp_category_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `spc_name` varchar(50) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `spc_name_UNIQUE` (`spc_name`)
) ENGINE=InnoDB COMMENT='Short Name for Table: spc';

--
-- Table structure for table `spl_hpft_master_task_lib_tbl`
--

CREATE TABLE `spl_hpft_master_task_lib_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `spc_id_fk` INT(10) UNSIGNED NOT NULL,
  `task_name` varchar(50) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `task_name_UNIQUE` (`task_name`),
  KEY `fk_mtask_spc` (`spc_id_fk`),
  CONSTRAINT `fk_mtask_spc` FOREIGN KEY (`spc_id_fk`) REFERENCES `spl_prod_master_sp_category_tbl` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: mtask';


--
-- Table structure for table `spl_prod_master_serv_conf_type_tbl`
--

CREATE TABLE `spl_prod_master_serv_conf_type_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `conf_type_code` varchar(20) NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `conf_type_code_UNIQUE` (`conf_type_code`)
) ENGINE=InnoDB COMMENT='Short Name for Table: serv_conf_type';

