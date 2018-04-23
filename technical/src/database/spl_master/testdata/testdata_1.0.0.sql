--
-- Dumping data for table `spl_master_corp_tbl`
--
   
INSERT INTO `spl_master_corp_tbl` (`id`,`corp_name`,`corp_mobile_no`,`corp_email_id`,`corp_landline_no`) VALUES (1,'Corporation 1','1234568','corp@gmail.com','01919191');



--
-- Dumping data for table `spl_master_customer_tbl`
--

INSERT INTO `spl_master_customer_tbl` (`id`,`corp_id_fk`,`cust_name`,`cust_state`,`cust_state_since`) VALUES (1,1,'Customer 1',1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_cust_prod_mapping_tbl`
--

INSERT INTO `spl_master_cust_prod_mapping_tbl` (`id`,`cust_id_fk`,`prod_id_fk`,`dbi_id_fk`,`cpm_state`,`cpm_state_since`) VALUES (1,1,1,1,1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_device_tbl`
--

INSERT INTO `spl_master_device_tbl` (`id`,`serialno`,`dev_state`,`dev_state_since`) VALUES (1,'1234567890123456',1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_dev_status_tbl`
--

INSERT INTO `spl_master_dev_status_tbl` (`dev_id_fk`,`connection_state`,`connection_state_since`,`sync_state`,`sync_state_since`,`battery_level`,`battery_level_since`) VALUES ('1', '2', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_cpm_dev_mapping_tbl`
--

INSERT INTO `spl_master_cpm_dev_mapping_tbl` (`cpm_id_fk`,`dev_id_fk`) VALUES (1,1);

--
-- Dumping data for table `spl_master_servicepoint_tbl`
--

INSERT INTO `spl_master_servicepoint_tbl` (`id`,`sp_state`,`sp_state_since`) VALUES (1,1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_cpm_sp_mapping_tbl`
--

INSERT INTO `spl_master_cpm_sp_mapping_tbl` (`cpm_id_fk`,`sp_id_fk`) VALUES (1,1);

--
-- Dumping data for table `spl_master_total_count_tbl`
--

UPDATE `spl_master_total_count_tbl` SET `cust_cnt`='1', `usr_cnt`='1', `dev_cnt`='1', `sp_cnt`='1', `dev_active_cnt`='1' WHERE `id`='1';

--
-- Dumping data for table `spl_master_user_tbl`
--
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (1,'admin@servicepoint.live','admin',1,1,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (2,'admin@customer1.com','admin',2,null,1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_usr_cpm_tbl`
--

INSERT INTO `spl_master_usr_cpm_tbl` (`id`,`user_id_fk`,`cpm_id_fk`,`urole_id_fk`) VALUES (1,1,1,1);

--
-- Dumping data for table `spl_master_cust_prod_count_tbl`
--

INSERT INTO `spl_master_cust_prod_count_tbl` (`id`,`cpm_id_fk`,`dev_cnt`,`sp_cnt`,`usr_cnt`) VALUES (1,1,1,1,1);

--
-- Dumping data for table `spl_master_customer_tbl`
--

INSERT INTO `spl_master_customer_tbl` (`id`, `corp_id_fk`, `cust_name`, `cust_state`, `cust_state_since`) VALUES (2, 1, 'Customer 2', 1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_cust_details_tbl`
--

INSERT INTO `spl_master_cust_details_tbl` (`cust_id_fk`, `poc1_name`, `poc1_email_id`, `poc1_mobile_no`) VALUES (1, 'poc1', 'poc1@email.com', '1244324');



