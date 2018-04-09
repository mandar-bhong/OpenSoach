-- =============================================================================
-- Filename   :  schema_1.0.0.sql

-- Version    :  1.0

-- Created on :  22-Feb-2018

-- Author     :  sanjay.sawant@opensoach.com

-- Description:
-- - db schema script to create database for HKT domain database (multiple instances to be created of the same schema).
-- - spl_hkt_node_xxxx, the placeholder 'xxx' to be replaced with the instance number, instance number will start from 0000
-- ==========================================================================
drop database if exists spl_hkt_node_xxxx;
create database spl_hkt_node_xxxx DEFAULT CHARACTER SET utf8;
use spl_hkt_node_xxxx;

--
-- Table structure for table `spl_hkt_dev_sp_mapping`
--

CREATE TABLE `spl_hkt_dev_sp_mapping` (
  `dev_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`dev_id_fk`,`sp_id_fk`)
) ENGINE=InnoDB COMMENT='Short Name for Table: devsp';

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
  UNIQUE KEY `cpm_spc_task_name_UNIQUE` (`cpm_id_fk`,`spc_id_fk`,`task_name`)
) ENGINE=InnoDB COMMENT='Short Name for Table: task';

--
-- Table structure for table `spl_hkt_chart_tbl`
--

CREATE TABLE `spl_hkt_chart_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `spc_id_fk` int(11) NOT NULL,
  `chart_name` varchar(50) NOT NULL,
  `chart_type` tinyint(3) unsigned NOT NULL COMMENT '1: Daily\n2: Weekly',
  `chart_config` json NOT NULL COMMENT 'For chart_type : 1(daily). following is the json structure\n{starttime,endtime,interval}',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `cpm_spc_chart_name_UNIQUE` (`cpm_id_fk`,`spc_id_fk`,`chart_name`)
) ENGINE=InnoDB COMMENT='Short Name for Table: chart';

--
-- Table structure for table `spl_hkt_chart_tasks_tbl`
--

CREATE TABLE `spl_hkt_chart_tasks_tbl` (
  `chart_id_fk` int(10) unsigned NOT NULL,
  `task_id_fk` int(10) unsigned NOT NULL,
  `task_order` smallint(5) unsigned DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`chart_id_fk`,`task_id_fk`),
  KEY `fk_ct_task_idx` (`task_id_fk`),
  CONSTRAINT `fk_ctasks_chart` FOREIGN KEY (`chart_id_fk`) REFERENCES `spl_hkt_chart_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_ctasks_task` FOREIGN KEY (`task_id_fk`) REFERENCES `spl_hkt_task_lib_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: ctasks';

--
-- Table structure for table `spl_hkt_sp_charts_tbl`
--

CREATE TABLE `spl_hkt_sp_charts_tbl` (
  `chart_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`chart_id_fk`,`sp_id_fk`),
  CONSTRAINT `fk_spcharts_chart` FOREIGN KEY (`chart_id_fk`) REFERENCES `spl_hkt_chart_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: spcharts';

--
-- Table structure for table `spl_hkt_chart_txn_tbl`
--

CREATE TABLE `spl_hkt_chart_txn_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `chart_id_fk` int(10) unsigned NOT NULL,
  `task_id_fk` int(10) unsigned NOT NULL,
  `slot` tinyint(3) unsigned NOT NULL,
  `task_state` tinyint(3) unsigned NOT NULL,
  `entry_time` datetime NOT NULL,
  `task_txn_day` date NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_charttxn_chart_idx` (`chart_id_fk`),
  KEY `fk_charttxn_task_idx` (`task_id_fk`),
  CONSTRAINT `fk_charttxn_chart` FOREIGN KEY (`chart_id_fk`) REFERENCES `spl_hkt_chart_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_charttxn_task` FOREIGN KEY (`task_id_fk`) REFERENCES `spl_hkt_task_lib_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: charttxn';

--
-- Table structure for table `spl_hkt_field_operator_tbl`
--

CREATE TABLE `spl_hkt_field_operator_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cmp_id_fk` int(10) unsigned NOT NULL,
  `fopcode` varchar(4) NOT NULL,
  `fop_name` varchar(50) DEFAULT NULL,
  `mobile_no` varchar(15) NOT NULL,
  `email_id` varchar(254) DEFAULT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `fop_state` tinyint(4) NOT NULL COMMENT '1: Active, 2: InActive etc.',
  `fop_area` tinyint(4) NOT NULL COMMENT '1: open, 2: restricted.',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='Short Name for Table: fop';


--
-- Table structure for table `spl_hkt_fop_sp_tbl`
--

CREATE TABLE `spl_hkt_fop_sp_tbl` (
  `fop_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`fop_id_fk`,`sp_id_fk`),
  CONSTRAINT `fk_fopsp_fop` FOREIGN KEY (`fop_id_fk`) REFERENCES `spl_hkt_field_operator_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: fopsp';

--
-- Table structure for table `spl_hkt_sp_complaint_tbl`
--

CREATE TABLE `spl_hkt_sp_complaint_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='Short Name for Table: spcomplaint';