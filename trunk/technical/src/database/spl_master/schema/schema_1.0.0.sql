-- =============================================================================
-- Filename   :  schema_1.0.0.sql

-- Version    :  1.0

-- Created on :  26-Jan-2018

-- Author     :  sanjay.sawant@opensoach.com

-- Description:
-- - db schema script to create database for servicepoint.live master.

-- ==========================================================================
drop database if exists spl_master;
create database spl_master DEFAULT CHARACTER SET utf8;
use spl_master;

--
-- Table structure for table `spl_master_product_tbl`
--

CREATE TABLE `spl_master_product_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `prod_code` varchar(7) NOT NULL COMMENT 'All product code starts with prefix "SPL" and 3 character individual product suffix separated by underscore',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `prod_code_UNIQUE` (`prod_code`)
) ENGINE=InnoDB COMMENT='Short Name for Table: prod';

--
-- Table structure for table `spl_master_database_instance_tbl`
--

CREATE TABLE `spl_master_database_instance_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dbi_name` varchar(20) NOT NULL COMMENT 'Database name: splive_product code_node_4digitnumber\ne.g.\nsplive_hkt_node_1001',
  `connection_string` varchar(500) NOT NULL,
  `prod_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `dbi_name_UNIQUE` (`dbi_name`),  KEY `fk_dbi_prod_idx` (`prod_id_fk`),
  CONSTRAINT `fk_dbi_prod` FOREIGN KEY (`prod_id_fk`) REFERENCES `spl_master_product_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: dbi';


--
-- Table structure for table `spl_master_corp_tbl`
--

CREATE TABLE `spl_master_corp_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `corp_name` varchar(50) NOT NULL,
  `corp_mobile_no` varchar(15) DEFAULT NULL,
  `corp_email_id` varchar(254) DEFAULT NULL,
  `corp_landline_no` varchar(15) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='Short Name for Table: corp';

--
-- Table structure for table `spl_master_customer_tbl`
--

CREATE TABLE `spl_master_customer_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `corp_id_fk` int(10) unsigned NOT NULL,
  `cust_name` varchar(50) NOT NULL,
  `cust_state` tinyint(3) unsigned NOT NULL COMMENT '1: Active, 2: Inactive, 3: Suspended etc.',
  `cust_state_since` datetime NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `cust_name_UNIQUE` (`cust_name`),
  KEY `fk_cust_corp_idx` (`corp_id_fk`),
  CONSTRAINT `fk_cust_corp` FOREIGN KEY (`corp_id_fk`) REFERENCES `spl_master_corp_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: cust';

--
-- Table structure for table `spl_master_cust_details_tbl`
--

CREATE TABLE `spl_master_cust_details_tbl` (
  `cust_id_fk` int(10) unsigned NOT NULL,  
  `poc1_name` varchar(50) NOT NULL,
  `poc1_email_id` varchar(254) NOT NULL,
  `poc1_mobile_no` varchar(15) NOT NULL,
  `poc2_name` varchar(50) DEFAULT NULL,
  `poc2_email_id` varchar(254) DEFAULT NULL,
  `poc2_mobile_no` varchar(50) DEFAULT NULL,
  `address` varchar(250) DEFAULT NULL,
  `address_state` varchar(50) DEFAULT NULL,
  `address_city` varchar(50) DEFAULT NULL,
  `address_pincode` varchar(6) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cust_id_fk`),
  KEY `fk_custd_cust_idx` (`cust_id_fk`),
  CONSTRAINT `fk_custd_cust` FOREIGN KEY (`cust_id_fk`) REFERENCES `spl_master_customer_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: custd';

--
-- Table structure for table `spl_master_cust_prod_mapping_tbl`
--

CREATE TABLE `spl_master_cust_prod_mapping_tbl` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cust_id_fk` INT(10) UNSIGNED NOT NULL,
	`prod_id_fk` INT(11) UNSIGNED NOT NULL,
	`dbi_id_fk` INT(10) UNSIGNED NOT NULL,
	`cpm_state` TINYINT(3) UNSIGNED NOT NULL COMMENT '1: Active, 2: Inactive, 3: Suspended etc.',
	`cpm_state_since` DATETIME NOT NULL,
	`created_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `cust_id_fk_prod_id_fk` (`cust_id_fk`, `prod_id_fk`),
	INDEX `fk_cpm_prod_idx` (`prod_id_fk`),
	INDEX `fk_cpm_cust_idx` (`cust_id_fk`),
	INDEX `fk_cpm_dbi_idx` (`dbi_id_fk`),
	CONSTRAINT `fk_cpm_cust` FOREIGN KEY (`cust_id_fk`) REFERENCES `spl_master_customer_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_cpm_dbi` FOREIGN KEY (`dbi_id_fk`) REFERENCES `spl_master_database_instance_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT `fk_cpm_prod` FOREIGN KEY (`prod_id_fk`) REFERENCES `spl_master_product_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) ENGINE=InnoDB COMMENT='Short Name for Table: cpm';

--
-- Table structure for table `spl_master_user_role_tbl`
--

CREATE TABLE `spl_master_user_role_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `prod_id_fk` int(10) unsigned DEFAULT NULL,
  `urole_code` varchar(10) NOT NULL,
  `urole_name` varchar(20) NOT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_urole_prod_idx` (`prod_id_fk`),
  CONSTRAINT `fk_urole_prod` FOREIGN KEY (`prod_id_fk`) REFERENCES `spl_master_product_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: urole';

--
-- Table structure for table `spl_master_user_tbl`
--

CREATE TABLE `spl_master_user_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `usr_name` varchar(254) NOT NULL,
  `usr_password` varchar(20) NOT NULL,
  `usr_category` tinyint(3) unsigned NOT NULL COMMENT '1: OpenSoach users.\n2: Customer users.',
  `urole_id_fk` int(10) unsigned DEFAULT NULL COMMENT 'this field is applicable only if usr_category==1(OpenSoach)',
  `usr_state` tinyint(3) unsigned NOT NULL COMMENT '1: Active, 2: Inactive, 3: Suspended etc.',
  `usr_state_since` datetime NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`usr_name`),
  KEY `fk_usr_urole_idx` (`urole_id_fk`),
  CONSTRAINT `fk_usr_urole` FOREIGN KEY (`urole_id_fk`) REFERENCES `spl_master_user_role_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: usr';

--
-- Table structure for table `spl_master_usr_details_tbl`
--

CREATE TABLE `spl_master_usr_details_tbl` (
  `usr_id_fk` int(10) unsigned NOT NULL,
  `fname` varchar(25) DEFAULT NULL,
  `lname` varchar(25) DEFAULT NULL,
  `gender` TINYINT(3) UNSIGNED NULL DEFAULT NULL COMMENT '0: Not Selected, 1: Male , 2: Female ',
  `mobile_no` varchar(15) DEFAULT NULL,
  `alternate_contact_no` varchar(15) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`usr_id_fk`),
  CONSTRAINT `fk_usrd_usr` FOREIGN KEY (`usr_id_fk`) REFERENCES `spl_master_user_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: usrd';

--
-- Table structure for table `spl_master_usr_cust_prod_mapping_tbl`
--

CREATE TABLE `spl_master_usr_cpm_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id_fk` int(10) unsigned NOT NULL,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `urole_id_fk` int(10) unsigned NOT NULL,
  `ucpm_state` TINYINT(3) UNSIGNED NOT NULL COMMENT '1: Active, 2: Inactive, 3: Suspended etc.',
  `ucpm_state_since` DATETIME NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_cpm_idx` (`user_id_fk`,`cpm_id_fk`),
  KEY `fk_ucpm_usr_idx` (`user_id_fk`),
  KEY `fk_ucpm_cpm_idx` (`cpm_id_fk`),
  KEY `fk_ucpm_urole_idx` (`urole_id_fk`),
  CONSTRAINT `fk_ucpm_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_master_cust_prod_mapping_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_ucpm_urole` FOREIGN KEY (`urole_id_fk`) REFERENCES `spl_master_user_role_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_ucpm_usr` FOREIGN KEY (`user_id_fk`) REFERENCES `spl_master_user_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: ucpm';

--
-- Table structure for table `spl_master_device_tbl`
--

CREATE TABLE `spl_master_device_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cust_id_fk` int(10) unsigned DEFAULT NULL,
  `serialno` varchar(16) NOT NULL,
  `dev_state` tinyint(3) unsigned NOT NULL COMMENT '0:Unallocated, 1: Active, 2: Inactive, 3: Suspended etc.',
  `dev_state_since` datetime NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `serialno_UNIQUE` (`serialno`),
  KEY `fk_dev_cust_idx` (`cust_id_fk`),
  CONSTRAINT `fk_dev_cust` FOREIGN KEY (`cust_id_fk`) REFERENCES `spl_master_customer_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: dev';


--
-- Table structure for table `spl_master_dev_details_tbl`
--

CREATE TABLE `spl_master_dev_details_tbl` (
  `dev_id_fk` int(10) unsigned NOT NULL,
  `make` varchar(30) DEFAULT NULL,
  `technology` varchar(30) DEFAULT NULL,
  `tech_version` varchar(30) DEFAULT NULL,
  `short_desc` varchar(250) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`dev_id_fk`),
  CONSTRAINT `fk_devd_dev` FOREIGN KEY (`dev_id_fk`) REFERENCES `spl_master_device_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: devd';

--
-- Table structure for table `spl_master_servicepoint_tbl`
--

CREATE TABLE `spl_master_servicepoint_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  	`cpm_id_fk` INT(10) UNSIGNED NOT NULL,
  `sp_state` tinyint(3) unsigned NOT NULL COMMENT '1: Active, 2: Inactive, 3: Suspended etc.',
  `sp_state_since` datetime NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  	INDEX `sp_cpm` (`cpm_id_fk`),
	CONSTRAINT `sp_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_master_cust_prod_mapping_tbl` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) ENGINE=InnoDB COMMENT='Short Name for Table: sp';

--
-- Table structure for table `spl_master_cpm_dev_mapping_tbl`
--

CREATE TABLE `spl_master_cpm_dev_mapping_tbl` (
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `dev_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cpm_id_fk`,`dev_id_fk`),
  KEY `fk_cpdm_cpm` (`cpm_id_fk`),
  KEY `fk_cpdm_dev_idx` (`dev_id_fk`),
  CONSTRAINT `fk_cpdm_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_master_cust_prod_mapping_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_cpdm_dev` FOREIGN KEY (`dev_id_fk`) REFERENCES `spl_master_device_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: cpdm';

--
-- Table structure for table `spl_master_cpm_sp_mapping_tbl`
--

CREATE TABLE `spl_master_cpm_sp_mapping_tbl` (
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `sp_id_fk` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cpm_id_fk`,`sp_id_fk`),
  KEY `fk_cpsm_cpm_idx` (`cpm_id_fk`),
  KEY `fk_cpsm_sp_idx` (`sp_id_fk`),
  CONSTRAINT `fk_cpsm_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_master_cust_prod_mapping_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_cpsm_dev` FOREIGN KEY (`sp_id_fk`) REFERENCES `spl_master_servicepoint_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: cpsm';


--
-- Table structure for table `spl_master_summary_tbl`
--

CREATE TABLE `spl_master_total_count_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cust_cnt` int(10) unsigned NOT NULL,
  `usr_cnt` int(10) unsigned NOT NULL,
  `dev_cnt` int(10) unsigned NOT NULL,
  `sp_cnt` int(10) unsigned NOT NULL,
  `dev_unallocated_cnt` int(10) unsigned NOT NULL,
  `dev_active_cnt` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB COMMENT='Short Name for Table: tcnt';

--
-- Table structure for table `spl_master_cust_prod_count_tbl`
--

CREATE TABLE `spl_master_cust_prod_count_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cpm_id_fk` int(10) unsigned NOT NULL,
  `dev_cnt` int(10) unsigned NOT NULL,
  `sp_cnt` int(10) unsigned NOT NULL,
  `usr_cnt` int(10) unsigned NOT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_cpcnt_cpm_idx` (`cpm_id_fk`),
  CONSTRAINT `fk_cpcnt_cpm` FOREIGN KEY (`cpm_id_fk`) REFERENCES `spl_master_cust_prod_mapping_tbl` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT='Short Name for Table: cpcnt';

--
-- Table structure for table `spl_master_config`
--

CREATE TABLE `spl_master_config` (
	`config_key` VARCHAR(50) NOT NULL,
	`config_value` VARCHAR(500) NOT NULL DEFAULT '',
	`created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`config_key`)
) ENGINE=InnoDB COMMENT='Short Name for Table: config\r\nThis table will contain configuration for spl  product';

CREATE TABLE `spl_master`.`spl_master_server_register` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `server_type_code` VARCHAR(50) NOT NULL,
  `server_address` VARCHAR(500) NULL,
  `prod_id_fk` INT UNSIGNED NULL,
  `server_state` TINYINT UNSIGNED NOT NULL COMMENT '0:Unallocated, 1: Running, 2: Stopped, 3: In Error',
  `server_state_since` DATETIME NOT NULL,
  `created_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_sreg_prod_idx` (`prod_id_fk` ASC),
  CONSTRAINT `fk_sreg_prod` FOREIGN KEY (`prod_id_fk`) REFERENCES `spl_master`.`spl_master_product_tbl` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB COMMENT = 'Short Name for Table: sreg';


CREATE TABLE `spl_master_email_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email_tml_id_fk` int(10) NOT NULL,
  `subject` varchar(150) NOT NULL,
  `body` longtext NOT NULL,
  `bcc` varchar(500) DEFAULT '0',
  `retrycount` int(10) NOT NULL DEFAULT '0' COMMENT 'No of retries completed',
  `status` int(10) NOT NULL COMMENT '0:unknown, 1:success, 2:failed',
  `comment` varchar(5000) DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_spl_master_email_tbl_spl_master_email_template_tbl` (`email_tml_id_fk`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='Short Name for Table: email';

CREATE TABLE `spl_master_email_template_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(50) NOT NULL,
  `subject` varchar(150) NOT NULL,
  `body` longtext NOT NULL,
  `bcc` varchar(500) DEFAULT NULL,
  `maxretry` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='Short Name for Table: emiltml';

CREATE TABLE `spl_master_usr_activation_tbl` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `usr_id_fk` int(10) unsigned NOT NULL,
  `code` varchar(150) NOT NULL,
  `password_changed` bit(1) NOT NULL DEFAULT b'0' COMMENT 'false: Not updated, true: Updated',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_spl_master_usr_activation_tbl_spl_master_user_tbl` (`usr_id_fk`),
  CONSTRAINT `FK_spl_master_usr_activation_tbl_spl_master_user_tbl` FOREIGN KEY (`usr_id_fk`) REFERENCES `spl_master_user_tbl` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Short Name for Table: usract';

