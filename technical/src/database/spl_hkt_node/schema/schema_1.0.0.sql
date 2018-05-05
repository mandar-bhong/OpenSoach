-- =============================================================================
-- Filename   :  schema_1.0.0.sql

-- Version    :  1.0

-- Created on :  22-Feb-2018

-- Author     :  sanjay.sawant@opensoach.com

-- Description:
-- - db schema script to create database for HKT domain database (multiple instances to be created of the same schema).
-- - spl_hkt_node_xxxx, the placeholder 'xxx' to be replaced with the instance number, instance number will start from 0000
-- ==========================================================================
drop database if exists spl_hkt_node_0001;
create database spl_hkt_node_0001 DEFAULT CHARACTER SET utf8;
use spl_hkt_node_0001;

--
-- Table structure for table `spl_node_cpm_tbl`
--

CREATE TABLE `spl_node_cpm_tbl` (
  `cpm_id_fk` int(10) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`cpm_id_fk`)
) ENGINE=InnoDB COMMENT='Short Name for Table: cpm';

--
-- Table structure for table `spl_node_sp_category_tbl`
--

CREATE TABLE `spl_node_sp_category_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `spc_name` varchar(50) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `spc_name_UNIQUE` (`spc_name`),
  KEY `fk_spc_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_spc_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: spc';

--
-- Table structure for table `spl_node_dev_tbl`
--

CREATE TABLE `spl_node_dev_tbl` (
  `dev_id_fk` int(10) unsigned NOT NULL,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`dev_id_fk`),
  KEY `fk_dev_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_dev_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: dev';

--
-- Table structure for table `spl_node_sp_tbl`
--

CREATE TABLE `spl_node_sp_tbl` (
  `sp_id_fk` int(10) unsigned NOT NULL,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `spc_id_fk` int(10) unsigned NOT NULL,
  `sp_name` varchar(50) NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`sp_id_fk`),
  CONSTRAINT `fk_sp_spc` FOREIGN KEY (`spc_id_fk`) REFERENCES `spl_node_sp_category_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_sp_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_sp_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: sp';

--
-- Table structure for table `spl_node_dev_sp_mapping`
--

CREATE TABLE `spl_node_dev_sp_mapping` (
  `dev_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`dev_id_fk`,`sp_id_fk`),
  KEY `fk_devsp_dev_idx` (`dev_id_fk`),
  KEY `fk_devsp_sp_idx` (`sp_id_fk`),
  CONSTRAINT `fk_devsp_dev` FOREIGN KEY (`dev_id_fk`) REFERENCES `spl_node_dev_tbl` (`dev_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_devsp_sp` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_node_sp_tbl` (`sp_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_devsp_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_devsp_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: devsp';

--
-- Table structure for table `spl_node_service_conf_tbl`
--

CREATE TABLE `spl_node_service_conf_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `spc_id_fk` int(10) unsigned NOT NULL,
  `conf_type_code` varchar(20) NOT NULL,
  `serv_conf_name` varchar(50) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `serv_conf` json NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_serv_conf_spc_idx` (`spc_id_fk`),  
  CONSTRAINT `fk_serv_conf_spc` FOREIGN KEY (`spc_id_fk`) REFERENCES `spl_node_sp_category_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_serv_conf_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_serv_conf_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION  
) ENGINE=InnoDB COMMENT='Short Name for Table: serv_conf';

--
-- Table structure for table `spl_node_service_instance_tbl`
--

CREATE TABLE `spl_node_service_instance_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `serv_conf_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_serv_conf_in_serv_conf` (`serv_conf_id_fk`),
  KEY `fk_serv_conf_in_sp_idx` (`sp_id_fk`),
  CONSTRAINT `fk_serv_conf_in_serv_conf` FOREIGN KEY (`serv_conf_id_fk`) REFERENCES `spl_node_service_conf_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_serv_conf_in_sp` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_node_sp_tbl` (`sp_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_serv_conf_in_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_serv_conf_in_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: serv_conf_in';

--
-- Table structure for table `spl_node_service_in_txn_tbl`
--

CREATE TABLE `spl_node_service_in_txn_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `serv_in_id_fk` int(10) unsigned NOT NULL,
  `txn_data` json NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_serv_in_txn_serv_in_idx` (`serv_in_id_fk`),
  CONSTRAINT `fk_serv_in_txn_serv_in` FOREIGN KEY (`serv_in_id_fk`) REFERENCES `spl_node_service_instance_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_serv_in_txn_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_serv_in_txn_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: serv_in_txn';

--
-- Table structure for table `spl_node_field_operator_tbl`
--

CREATE TABLE `spl_node_field_operator_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `fopcode` varchar(4) NOT NULL,
  `fop_name` varchar(50) DEFAULT NULL,
  `mobile_no` varchar(15) NOT NULL,
  `email_id` varchar(254) DEFAULT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `fop_state` tinyint(4) NOT NULL COMMENT '1: Active, 2: InActive etc.',
  `fop_area` tinyint(4) NOT NULL COMMENT '1: open, 2: restricted.',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_fop_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_fop_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: fop';

--
-- Table structure for table `spl_node_fop_sp_tbl`
--

CREATE TABLE `spl_node_fop_sp_tbl` (
  `fop_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`fop_id_fk`,`sp_id_fk`),
  KEY `fk_fopsp_fop_idx` (`fop_id_fk`),
  KEY `fk_fopsp_sp_idx` (`sp_id_fk`),
  CONSTRAINT `fk_fopsp_fop` FOREIGN KEY (`fop_id_fk`) REFERENCES `spl_node_field_operator_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_fopsp_sp` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_node_sp_tbl` (`sp_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_fopsp_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_fopsp_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: fopsp';

--
-- Table structure for table `spl_hkt_task_lib_tbl`
--

CREATE TABLE `spl_hkt_task_lib_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `spc_id_fk` int(10) unsigned DEFAULT NULL,
  `task_name` varchar(50) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_task_spc_idx` (`spc_id_fk`),
  UNIQUE KEY `cpm_spc_task_name_UNIQUE` (`cpm_id_fk`,`spc_id_fk`,`task_name`),
  CONSTRAINT `fk_task_spc` FOREIGN KEY (`spc_id_fk`) REFERENCES `spl_node_sp_category_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_task_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_task_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: task';

--
-- Table structure for table `spl_hkt_sp_complaint_tbl`
--

CREATE TABLE `spl_hkt_sp_complaint_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `complaint_title` varchar(250) NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `complaint_by` varchar(50) NOT NULL,
  `mobile_no` varchar(15) DEFAULT NULL,
  `email_id` varchar(254) DEFAULT NULL,
  `employee_id` varchar(16) DEFAULT NULL,
  `raised_on` datetime NOT NULL,
  `complaint_state` tinyint(4) NOT NULL COMMENT '1: Open, 2: Closed, 3: Force Closed etc.',
  `closed_on` datetime DEFAULT NULL,
  `remarks` varchar(500) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_spcomplaint_sp_idx` (`sp_id_fk`),
  CONSTRAINT `fk_spcomplaint_sp` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_node_sp_tbl` (`sp_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION,
  KEY `fk_spcomplaint_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_spcomplaint_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: spcomplaint';