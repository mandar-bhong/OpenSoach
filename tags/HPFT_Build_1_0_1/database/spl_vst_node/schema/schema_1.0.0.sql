-- =============================================================================
-- Filename   :  schema_1.0.0.sql

-- Version    :  1.0

-- Created on :  22-Feb-2018

-- Author     :  sanjay.sawant@opensoach.com

-- Description:
-- - db schema script to create database for VST domain database (multiple instances to be created of the same schema).
-- - spl_vst_node_xxxx, the placeholder 'xxx' to be replaced with the instance number, instance number will start from 0000
-- ==========================================================================
drop database if exists spl_vst_node_0001;
create database spl_vst_node_0001 DEFAULT CHARACTER SET utf8;
use spl_vst_node_0001;

--
-- Table structure for table `spl_node_cpm_tbl`
--

CREATE TABLE `spl_node_cpm_tbl` (
  `cpm_id_fk` int(10) unsigned NOT NULL,
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
  KEY `fk_spc_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_spc_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: spc';

--
-- Table structure for table `spl_node_dev_tbl`
--

CREATE TABLE `spl_node_dev_tbl` (
  `dev_id_fk` int(10) unsigned NOT NULL,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `serialno` VARCHAR(16) NOT NULL,
  `dev_name` VARCHAR(30) NOT NULL,
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
  `short_desc` VARCHAR(250) NULL DEFAULT NULL,
  `sp_state` TINYINT(3) UNSIGNED NOT NULL COMMENT '1: Active, 2: Inactive, 3: Suspended etc.',
  `sp_state_since` DATETIME NOT NULL,
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
  `fopcode` VARCHAR(20) NOT NULL,
  `status` TINYINT(4) NOT NULL,
  `txn_data` json NOT NULL,
  `txn_date` DATETIME NOT NULL,
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
  UNIQUE INDEX `fopcode` (`fopcode`),
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
-- Table structure for table `spl_node_task_lib_tbl`
--

CREATE TABLE `spl_node_task_lib_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `spc_id_fk` int(10) unsigned NOT NULL,
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
-- Table structure for table `spl_node_sp_complaint_tbl`
--

CREATE TABLE `spl_node_sp_complaint_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `complaint_title` varchar(250) NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `complaint_by` varchar(50) NOT NULL,
  `mobile_no` varchar(15) DEFAULT NULL,
  `email_id` varchar(254) DEFAULT NULL,
  `employee_id` varchar(16) DEFAULT NULL,
  `severity` TINYINT(4) UNSIGNED NULL DEFAULT NULL COMMENT '1: Low, 2: Medium, 3: High,4: Critical etc.',
  `raised_on` datetime NOT NULL,
  `complaint_state` tinyint(4) NOT NULL COMMENT '1: Open, 2: In Progress, 3: Closed etc.',
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

--
-- Table structure for table `spl_node_feedback_tbl`
--

CREATE TABLE `spl_node_feedback_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`sp_id_fk` INT(10) UNSIGNED NOT NULL,
	`feedback` TINYINT(4) UNSIGNED NOT NULL,
	`feedback_comment` VARCHAR(150) NULL DEFAULT NULL,
	`raised_on` DATETIME NOT NULL,
	`created_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	INDEX `fk_feedback_cpm` (`cpm_id_fk`),
	INDEX `fk_feedback_sp` (`sp_id_fk`),
	CONSTRAINT `fk_feedback_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_feedback_sp` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_node_sp_tbl` (`sp_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
) 	ENGINE=InnoDB COMMENT='Short Name for Table: feedback';

--
-- Table structure for table `spl_node_report_template_tbl`
--

CREATE TABLE `spl_node_report_template_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`report_code` VARCHAR(100) NOT NULL,
	`report_desc` VARCHAR(150) NOT NULL,
	`report_header` JSON NOT NULL,
	`report_format` JSON NOT NULL,
	`report_query_params` VARCHAR(150) NOT NULL,
	`report_query` VARCHAR(3000) NOT NULL,
	PRIMARY KEY (`id`)
)   ENGINE=InnoDB COMMENT='Short Name for Table: report';

--
-- Table structure for table `spl_node_dev_status_tbl`
--

CREATE TABLE `spl_node_dev_status_tbl` (
	`dev_id_fk` INT(10) UNSIGNED NOT NULL,
	`connection_state` TINYINT(3) UNSIGNED NOT NULL COMMENT '0: Unknown, 1: Connected, 2: Disconnected.',
	`connection_state_since` DATETIME NOT NULL,
	`sync_state` TINYINT(3) UNSIGNED NOT NULL COMMENT '0: Connected., 1: Disconnected.2: Unknown',
	`sync_state_since` DATETIME NOT NULL,
	`battery_level` TINYINT(4) NOT NULL COMMENT 'In Percentage',
	`battery_level_since` DATETIME NOT NULL,
	`created_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`dev_id_fk`),
	CONSTRAINT `fk_devstate_dev` FOREIGN KEY (`dev_id_fk`) REFERENCES `spl_node_dev_tbl` (`dev_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='Short Name for Table: devstate';


--
-- Table structure for table `spl_vst_vehicle_master_tbl`
--

CREATE TABLE `spl_vst_vehicle_master_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`vehicle_no` VARCHAR(25) NOT NULL,
	`details` JSON NOT NULL,
	`created_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `vehicle_no` (`vehicle_no`),
	INDEX `fk_vehicle_cpm` (`cpm_id_fk`),
	CONSTRAINT `fk_vehicle_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
)   ENGINE=InnoDB COMMENT='Short Name for Table: vehicle';


--
-- Table structure for table `spl_vst_token`
--

CREATE TABLE `spl_vst_token` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`token` INT(10) UNSIGNED NOT NULL,
	`vhl_id_fk` INT(10) UNSIGNED NOT NULL,
	`mapping_details` JSON NOT NULL,
	`state` INT(10) UNSIGNED NOT NULL,
	`generated_on` DATETIME NOT NULL,
	`created_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	INDEX `fk_token_vehicle` (`vhl_id_fk`),
	CONSTRAINT `fk_token_vehicle` FOREIGN KEY (`vhl_id_fk`) REFERENCES `spl_vst_vehicle_master_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)   ENGINE=InnoDB COMMENT='Short Name for Table: token';

