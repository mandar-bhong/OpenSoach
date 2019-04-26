-- =============================================================================
-- Filename   :  schema_1.0.0.sql

-- Version    :  1.0

-- Created on :  22-Feb-2018

-- Author     :  sanjay.sawant@opensoach.com

-- Description:
-- - db schema script to create database for HPFT domain database (multiple instances to be created of the same schema).
-- - spl_hpft_node_xxxx, the placeholder 'xxx' to be replaced with the instance number, instance number will start from 0000
-- ==========================================================================
drop database if exists spl_hpft_node_0001;
create database spl_hpft_node_0001 DEFAULT CHARACTER SET utf8;
use spl_hpft_node_0001;

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
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
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
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`dev_id_fk`),
  KEY `fk_dev_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_dev_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: dev';

--
-- Table structure for table `spl_node_sp_tbl`
--

CREATE TABLE `spl_node_sp_tbl` (
  `sp_id_fk` int(10) unsigned NOT NULL,
  `uuid` VARCHAR(50) NOT NULL,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `spc_id_fk` int(10) unsigned NOT NULL,
  `sp_name` varchar(50) NOT NULL,
  `short_desc` VARCHAR(250) NULL DEFAULT NULL,
  `sp_state` TINYINT(3) UNSIGNED NOT NULL COMMENT '1: Active, 2: Inactive, 3: Suspended etc.',
  `sp_state_since` DATETIME NOT NULL,
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `updated_by` INT(10) UNSIGNED NULL DEFAULT NULL,
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
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
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
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
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
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
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
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
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
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
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
  `complaint_state` tinyint(4) NOT NULL COMMENT '1: Open, 2: Closed, 3: Force Closed etc.',
  `closed_on` datetime DEFAULT NULL,
  `remarks` varchar(500) DEFAULT NULL,
  `created_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_on` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
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
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
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
	`report_query_params` VARCHAR(150) NOT NULL,
	`report_query` VARCHAR(1000) NOT NULL,
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
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (`dev_id_fk`),
	CONSTRAINT `fk_devstate_dev` FOREIGN KEY (`dev_id_fk`) REFERENCES `spl_node_dev_tbl` (`dev_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='Short Name for Table: devstate';


--
-- Table structure for table `spl_node_sync_config_tbl`
--

CREATE TABLE `spl_node_sync_config_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`store_name` VARCHAR(50) NOT NULL,
	`device_type` TINYINT(3) NOT NULL COMMENT '0 :shared device; 1 :user device',
	`updated_on` DATETIME NOT NULL,
	`has_qry` VARCHAR(5000) NOT NULL,
	`select_count_qry` VARCHAR(5000) NOT NULL,
	`select_qry` VARCHAR(5000) NOT NULL,
	`insert_qry` VARCHAR(1000) NOT NULL,
	`update_qry` VARCHAR(1000) NOT NULL,
	`data_source` int(10) NOT NULL DEFAULT '2' COMMENT '1: MstDataSource, 2: NodeDataSource',
	`query_data` JSON NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='Short Name for Table: sync';


--
-- Table structure for table `spl_hpft_patient_master_tbl`
--

CREATE TABLE `spl_hpft_patient_master_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`patient_reg_no` VARCHAR(50) NOT NULL,
	`uuid` VARCHAR(50) NOT NULL,
	`fname` VARCHAR(25) NOT NULL,
	`lname` VARCHAR(25) NOT NULL,
	`mob_no` VARCHAR(15) NOT NULL,
	`date_of_birth` DATE NULL DEFAULT NULL,
	`age` VARCHAR(10) NOT NULL,
	`blood_grp` VARCHAR(10) NOT NULL,
	`gender` TINYINT(3) NOT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_patient_cpm` (`cpm_id_fk`),
	CONSTRAINT `fk_patient_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='Short Name for Table: patient';


--
-- Table structure for table `spl_hpft_patient_admission_tbl`
--

CREATE TABLE `spl_hpft_patient_admission_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`patient_id_fk` INT(10) UNSIGNED NOT NULL,
	`patient_reg_no` VARCHAR(50) NOT NULL,
	`uuid` VARCHAR(50) NOT NULL,
	`sp_id_fk` INT(10) UNSIGNED NOT NULL,
	`dr_incharge` INT(10) NOT NULL,
	`bed_no` VARCHAR(10) NOT NULL,
	`status` TINYINT(3) NOT NULL COMMENT '1 :hospitalize, 2 :discharged',
	`admitted_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`discharged_on` DATETIME NULL DEFAULT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_padmsn_cpm` (`cpm_id_fk`),
	INDEX `fk_padmsn_patient` (`patient_id_fk`),
	INDEX `fk_padmsn_sp` (`sp_id_fk`),
	CONSTRAINT `fk_padmsn_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_padmsn_patient` FOREIGN KEY (`patient_id_fk`) REFERENCES `spl_hpft_patient_master_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_padmsn_sp` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_node_sp_tbl` (`sp_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='Short Name for Table: padmsn';


--
-- Table structure for table `spl_hpft_patient_personal_details_tbl`
--

CREATE TABLE `spl_hpft_patient_personal_details_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`patient_id` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`uuid` VARCHAR(50) NOT NULL,
	`age` VARCHAR(10) NOT NULL,
	`other_details` JSON NULL DEFAULT NULL,
	`person_accompanying` JSON NULL DEFAULT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `admission_id_fk` (`admission_id_fk`),
	INDEX `fk_ppd_cpm` (`cpm_id_fk`),
	CONSTRAINT `fk_ppd_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_ppd_padmsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='Short Name for Table: ppd';



--
-- Table structure for table `spl_hpft_patient_medical_details_tbl`
-- 

CREATE TABLE `spl_hpft_patient_medical_details_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`patient_id` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`present_complaints` JSON NULL DEFAULT NULL,
	`reason_for_admission` JSON NULL DEFAULT NULL,
	`history_present_illness` JSON NULL DEFAULT NULL,
	`past_history` JSON NULL DEFAULT NULL,
	`treatment_before_admission` JSON NULL DEFAULT NULL,
	`investigation_before_admission` JSON NULL DEFAULT NULL,
	`family_history` JSON NULL DEFAULT NULL,
	`allergies` JSON NULL DEFAULT NULL,
	`personal_history` JSON NULL DEFAULT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `admission_id_fk` (`admission_id_fk`),
	INDEX `fk_pmd_cpm` (`cpm_id_fk`),
	CONSTRAINT `fk_pmd_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_pmd_padmsn_id` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='Short Name for Table: pmd';



--
-- Table structure for table `spl_hpft_conf_tbl`
--

CREATE TABLE `spl_hpft_conf_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`conf_type_code` VARCHAR(25) NOT NULL,
	`conf` JSON NOT NULL,
	`short_desc` VARCHAR(50) NULL DEFAULT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_conf_cpm` (`cpm_id_fk`),
	CONSTRAINT `fk_conf_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='Short Name for Table: conf';


--
-- Table structure for table `spl_hpft_patient_conf_tbl`
--

CREATE TABLE `spl_hpft_patient_conf_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`conf_type_code` VARCHAR(50) NOT NULL,
	`conf` JSON NOT NULL,
	`start_date` DATETIME NOT NULL,
	`end_date` DATETIME NOT NULL,
	`status` TINYINT(4) NOT NULL COMMENT '0- active, 1 - cancelled',
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_pconf_cpm` (`cpm_id_fk`),
	INDEX `fk_pconf_padmsn` (`admission_id_fk`),
	CONSTRAINT `fk_pconf_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_pconf_padmsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name: pconf';



--
-- Table structure for table `spl_hpft_action_tbl`
--

CREATE TABLE `spl_hpft_action_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`patient_conf_id_fk` INT(10) UNSIGNED NOT NULL,
	`conf_type_code` VARCHAR(25) NOT NULL,
	`scheduled_time` DATETIME NULL DEFAULT NULL,
	`is_deleted` TINYINT(4) NOT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(11) NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_actn_cpm` (`cpm_id_fk`),
	INDEX `fk_actn_admsn` (`admission_id_fk`),
	INDEX `fk_actn_pconf` (`patient_conf_id_fk`),
	CONSTRAINT `fk_actn_admsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_actn_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_actn_pconf` FOREIGN KEY (`patient_conf_id_fk`) REFERENCES `spl_hpft_patient_conf_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : actn';



--
-- Table structure for table `spl_hpft_action_txn_tbl`
--

CREATE TABLE `spl_hpft_action_txn_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`patient_conf_id_fk` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`txn_data` JSON NOT NULL,
	`runtime_config_data` JSON NULL DEFAULT NULL,
	`scheduled_time` DATETIME NOT NULL,
	`txn_state` INT(11) NOT NULL,
	`conf_type_code` VARCHAR(25) NOT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_actn_txn_cpm` (`cpm_id_fk`),
	INDEX `fk_actn_txn_pconf` (`patient_conf_id_fk`),
	INDEX `fk_actn_txn_admsn` (`admission_id_fk`),
	CONSTRAINT `fk_actn_txn_admsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_actn_txn_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_actn_txn_pconf` FOREIGN KEY (`patient_conf_id_fk`) REFERENCES `spl_hpft_patient_conf_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name: actn_txn';


--
-- Table structure for table `spl_hpft_document_tbl`
--

CREATE TABLE `spl_hpft_document_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cpm_id_fk` INT(11) UNSIGNED NOT NULL,
	`uuid` VARCHAR(50) NOT NULL,
	`name` VARCHAR(50) NULL DEFAULT NULL,
	`doctype` VARCHAR(200) NULL DEFAULT NULL,
	`store_name` VARCHAR(50) NULL DEFAULT NULL,
	`location` VARCHAR(200) NULL DEFAULT NULL,
	`location_type` TINYINT(3) NULL DEFAULT NULL,
	`persisted` TINYINT(3) NOT NULL COMMENT '0- not persisted, 1 - persisted',
	`updated_by` INT(11) UNSIGNED NOT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (`id`),
	INDEX `fk_doc_cpm` (`cpm_id_fk`),
	CONSTRAINT `doc_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : doc';



--
-- Table structure for table `spl_hpft_patient_document_tbl`
--

CREATE TABLE `spl_hpft_patient_document_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`document_id_fk` INT(10) UNSIGNED NOT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (`id`),
	INDEX `fk_pdoc_cpm` (`cpm_id_fk`),
	INDEX `fk_pdoc_padmsn` (`admission_id_fk`),
	INDEX `fk_pdoc_doc` (`document_id_fk`),
	CONSTRAINT `fk_pdoc_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_pdoc_doc` FOREIGN KEY (`document_id_fk`) REFERENCES `spl_hpft_document_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_pdoc_padmsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : pdoc';


--
-- Table structure for table `spl_hpft_doctors_orders_tbl`
--

CREATE TABLE `spl_hpft_doctors_orders_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`doctor_id_fk` INT(10) UNSIGNED NOT NULL,
	`doctors_orders` VARCHAR(1500) NOT NULL,
	`comment` VARCHAR(5000) NULL DEFAULT NULL,
	`ack_by` INT(10) UNSIGNED NULL DEFAULT NULL,
	`ack_time` TIMESTAMP NULL DEFAULT NULL,
	`status` TINYINT(4) NULL DEFAULT NULL COMMENT '0 :new, 1 :acked, default :0',
	`order_created_time` TIMESTAMP NULL DEFAULT NULL,
	`order_type` VARCHAR(50) NULL DEFAULT NULL,
	`document_id_fk` INT(10) UNSIGNED NULL DEFAULT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	`updated_by` INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_doc_ordrs_cpm` (`cpm_id_fk`),
	INDEX `fk_doc_ordrs_admsn` (`admission_id_fk`),
	INDEX `fk_doc_ordrs_doc` (`document_id_fk`),
	CONSTRAINT `fk_doc_ordrs_admsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_doc_ordrs_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_doc_ordrs_doc` FOREIGN KEY (`document_id_fk`) REFERENCES `spl_hpft_document_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : doc_ordrs';


--
-- Table structure for table `spl_hpft_treatment_tbl`
--


CREATE TABLE `spl_hpft_treatment_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`treatment_done` VARCHAR(1000) NOT NULL,
	`treatment_performed_time` TIMESTAMP NULL DEFAULT NULL,
	`details` VARCHAR(1000) NULL DEFAULT NULL,
	`post_observation` VARCHAR(1000) NULL DEFAULT NULL,
	`updated_by` INT(10) UNSIGNED NOT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (`id`),
	INDEX `fk_trtmnt_cpm` (`cpm_id_fk`),
	INDEX `fk_trtmnt_padmsn` (`admission_id_fk`),
	CONSTRAINT `fk_trtmnt_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_trtmnt_padmsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : trtmnt';



--
-- Table structure for table `spl_hpft_treatment_doc_tbl`
--

CREATE TABLE `spl_hpft_treatment_doc_tbl` (
	`treatment_id_fk` INT(10) UNSIGNED NOT NULL,
	`document_id_fk` INT(10) UNSIGNED NOT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (`treatment_id_fk`, `document_id_fk`),
	INDEX `fk_tdoc_doc` (`document_id_fk`),
	CONSTRAINT `fk_tdoc_doc` FOREIGN KEY (`document_id_fk`) REFERENCES `spl_hpft_document_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_tdoc_trtmnt` FOREIGN KEY (`treatment_id_fk`) REFERENCES `spl_hpft_treatment_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : tdoc';



--
-- Table structure for table `spl_hpft_pathology_record_tbl`
--

CREATE TABLE `spl_hpft_pathology_record_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`admission_id_fk` INT(10) UNSIGNED NOT NULL,
	`test_performed` VARCHAR(5000) NOT NULL,
	`test_performed_time` TIMESTAMP NULL DEFAULT NULL,
	`test_result` VARCHAR(5000) NULL DEFAULT NULL,
	`comments` VARCHAR(1000) NULL DEFAULT NULL,
	`updated_by` INT(10) UNSIGNED NOT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (`id`),
	INDEX `fk_prec_cpm` (`cpm_id_fk`),
	INDEX `fk_prec_padmsn` (`admission_id_fk`),
	CONSTRAINT `fk_prec_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_prec_padmsn` FOREIGN KEY (`admission_id_fk`) REFERENCES `spl_hpft_patient_admission_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : prec';


--
-- Table structure for table `spl_hpft_pathology_record_doc_tbl`
--

CREATE TABLE `spl_hpft_pathology_record_doc_tbl` (
	`pathology_id_fk` INT(10) UNSIGNED NOT NULL,
	`document_id_fk` INT(10) UNSIGNED NOT NULL,
	`created_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
	`updated_on` TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	PRIMARY KEY (`pathology_id_fk`, `document_id_fk`),
	INDEX `precdoc_doc` (`document_id_fk`),
	CONSTRAINT `precdoc_doc` FOREIGN KEY (`document_id_fk`) REFERENCES `spl_hpft_document_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `precdoc_prec` FOREIGN KEY (`pathology_id_fk`) REFERENCES `spl_hpft_pathology_record_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : precdoc';


--
-- Table structure for table `spl_hpft_user_patient_monitor_mapping`
--

CREATE TABLE `spl_hpft_user_patient_monitor_mapping` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`uuid` VARCHAR(50) NOT NULL,
	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
	`usr_id_fk` INT(10) UNSIGNED NOT NULL,
	`sp_id_fk` INT(10) UNSIGNED NULL DEFAULT NULL,
	`patient_id_fk` INT(10) UNSIGNED NULL DEFAULT NULL,
	`client_updated_at` TIMESTAMP(3) NULL DEFAULT NULL,
	`created_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`updated_by` INT(10) NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `fk_upmm_cpm` (`cpm_id_fk`),
	INDEX `fk_upmm_sp` (`sp_id_fk`),
	INDEX `fk_upmm_patient` (`patient_id_fk`),
	CONSTRAINT `fk_upmm_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_node_cpm_tbl` (`cpm_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_upmm_patient` FOREIGN KEY (`patient_id_fk`) REFERENCES `spl_hpft_patient_master_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_upmm_sp` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_node_sp_tbl` (`sp_id_fk`) ON UPDATE NO ACTION ON DELETE CASCADE
)	ENGINE=InnoDB COMMENT='short name : upmm';












