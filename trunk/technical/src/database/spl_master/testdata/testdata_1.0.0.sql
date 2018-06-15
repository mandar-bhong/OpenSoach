--
-- Dumping data for table `spl_master_corp_tbl`
--
   
INSERT INTO `spl_master_corp_tbl` (`id`,`corp_name`,`corp_mobile_no`,`corp_email_id`,`corp_landline_no`) VALUES (1,'Corporate 1','1234568','corp@gmail.com','01919191');
INSERT INTO `spl_master_corp_tbl` (`id`,`corp_name`,`corp_mobile_no`,`corp_email_id`,`corp_landline_no`) VALUES (2,'Corporate 2','1324343','corp2@gmail.com','04342123');
INSERT INTO `spl_master_corp_tbl` (`id`,`corp_name`,`corp_mobile_no`,`corp_email_id`,`corp_landline_no`) VALUES (3,'Corporate 3','4353432','corp3@gmail.com','01132323');
INSERT INTO `spl_master_corp_tbl` (`id`,`corp_name`,`corp_mobile_no`,`corp_email_id`,`corp_landline_no`) VALUES (4,'Corporate 4','3423454','corp4@gmail.com','04322434');



--
-- Dumping data for table `spl_master_customer_tbl`
--

INSERT INTO `spl_master_customer_tbl` (`id`, `corp_id_fk`, `cust_name`, `cust_state`, `cust_state_since`) VALUES (1, 1, 'Customer 1', 1,UTC_TIMESTAMP);
INSERT INTO `spl_master_customer_tbl` (`id`, `corp_id_fk`, `cust_name`, `cust_state`, `cust_state_since`) VALUES (2, 1, 'Customer 2', 1,UTC_TIMESTAMP);
INSERT INTO `spl_master_customer_tbl` (`id`, `corp_id_fk`, `cust_name`, `cust_state`, `cust_state_since`) VALUES (3, 1, 'Customer 3', 1,UTC_TIMESTAMP);
INSERT INTO `spl_master_customer_tbl` (`id`, `corp_id_fk`, `cust_name`, `cust_state`, `cust_state_since`) VALUES (4, 2, 'Customer 4', 1,UTC_TIMESTAMP);
INSERT INTO `spl_master_customer_tbl` (`id`, `corp_id_fk`, `cust_name`, `cust_state`, `cust_state_since`) VALUES (5, 2, 'Customer 5', 2,UTC_TIMESTAMP);
INSERT INTO `spl_master_customer_tbl` (`id`, `corp_id_fk`, `cust_name`, `cust_state`, `cust_state_since`) VALUES (6, 3, 'Customer 6', 2,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_cust_prod_mapping_tbl`
--

INSERT INTO `spl_master_cust_prod_mapping_tbl` (`id`,`cust_id_fk`,`prod_id_fk`,`dbi_id_fk`,`cpm_state`,`cpm_state_since`) VALUES (1,1,1,1,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_cust_prod_mapping_tbl` (`id`,`cust_id_fk`,`prod_id_fk`,`dbi_id_fk`,`cpm_state`,`cpm_state_since`) VALUES (2,2,1,1,1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_device_tbl`
--

INSERT INTO `spl_master_device_tbl` (`id`,`cust_id_fk`,`serialno`,`dev_state`,`dev_state_since`) VALUES (1,1,'1234567890123456',1,UTC_TIMESTAMP);
INSERT INTO `spl_master_device_tbl` (`id`,`serialno`,`dev_state`,`dev_state_since`) VALUES (2,'1345494544733456',1,UTC_TIMESTAMP);
INSERT INTO `spl_master_device_tbl` (`id`,`serialno`,`dev_state`,`dev_state_since`) VALUES (3,'1155623421323222',1,UTC_TIMESTAMP);
INSERT INTO `spl_master_device_tbl` (`id`,`serialno`,`dev_state`,`dev_state_since`) VALUES (4,'1235234322122343',2,UTC_TIMESTAMP);
INSERT INTO `spl_master_device_tbl` (`id`,`serialno`,`dev_state`,`dev_state_since`) VALUES (5,'1819243234324322',1,UTC_TIMESTAMP);
INSERT INTO `spl_master_device_tbl` (`id`,`serialno`,`dev_state`,`dev_state_since`) VALUES (6,'1234523445223332',2,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_dev_details_tbl`
--

INSERT INTO `spl_master_dev_details_tbl` (`dev_id_fk`) VALUES (1);
INSERT INTO `spl_master_dev_details_tbl` (`dev_id_fk`, `make`, `technology`) VALUES (2, 'make2', 'tech2');
INSERT INTO `spl_master_dev_details_tbl` (`dev_id_fk`, `make`) VALUES (3, 'make3');
INSERT INTO `spl_master_dev_details_tbl` (`dev_id_fk`, `make`, `technology`, `tech_version`, `short_desc`) VALUES (6, 'make4', 'tech6', 'techver6', 'desc6');

--
-- Dumping data for table `spl_master_cpm_dev_mapping_tbl`
--

INSERT INTO `spl_master_cpm_dev_mapping_tbl` (`cpm_id_fk`,`dev_id_fk`) VALUES (1,1);

--
-- Dumping data for table `spl_master_servicepoint_tbl`
--

INSERT INTO `spl_master_servicepoint_tbl` (`id`,`cpm_id_fk` ,`sp_state`,`sp_state_since`) VALUES (1,1,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_servicepoint_tbl` (`id`,`cpm_id_fk` ,`sp_state`,`sp_state_since`) VALUES (2,1,1,UTC_TIMESTAMP);

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

INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (2,'admin@customer1.com','admin',2,null,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (3,'admin@servicepoint2.live','admin',1,1,2,UTC_TIMESTAMP);
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (4,'admin@customer2.com','admin',2,null,2,UTC_TIMESTAMP);
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (5,'cust@customer3.com','admin',2,null,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (6,'cust@customer4.com','admin',2,null,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (7,'cust@customer5.com','admin',2,null,1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_usr_cpm_tbl`
--

INSERT INTO `spl_master_usr_cpm_tbl` (`id`,`user_id_fk`,`cpm_id_fk`,`urole_id_fk`,`ucpm_state`,`ucpm_state_since`) VALUES (1,2,1,2,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_usr_cpm_tbl` (`id`,`user_id_fk`,`cpm_id_fk`,`urole_id_fk`,`ucpm_state`,`ucpm_state_since`) VALUES (2,6,1,2,1,UTC_TIMESTAMP);
INSERT INTO `spl_master_usr_cpm_tbl` (`id`,`user_id_fk`,`cpm_id_fk`,`urole_id_fk`,`ucpm_state`,`ucpm_state_since`) VALUES (3,7,1,2,1,UTC_TIMESTAMP);

--
-- Dumping data for table `spl_master_cust_prod_count_tbl`
--

INSERT INTO `spl_master_cust_prod_count_tbl` (`id`,`cpm_id_fk`,`dev_cnt`,`sp_cnt`,`usr_cnt`) VALUES (1,1,1,1,1);


--
-- Dumping data for table `spl_master_cust_details_tbl`
--
INSERT INTO `spl_master_cust_details_tbl` (`cust_id_fk`, `poc1_name`, `poc1_email_id`, `poc1_mobile_no`) VALUES (1, 'poc1', 'poc1@email.com', '12443241');
INSERT INTO `spl_master_cust_details_tbl` (`cust_id_fk`, `poc1_name`, `poc1_email_id`, `poc1_mobile_no`) VALUES (3, 'poc3', 'poc3@email.com', '12453434');
INSERT INTO `spl_master_cust_details_tbl` (`cust_id_fk`, `poc1_name`, `poc1_email_id`, `poc1_mobile_no`) VALUES (4, 'poc4', 'poc4@email.com', '15633434');
INSERT INTO `spl_master_cust_details_tbl` (`cust_id_fk`, `poc1_name`, `poc1_email_id`, `poc1_mobile_no`) VALUES (6, 'poc6', 'poc6@email.com', '15622434');

--
-- Dumping data for table `spl_master_usr_details_tbl`
--

INSERT INTO `spl_master_usr_details_tbl` (`usr_id_fk`, `fname`, `lname`, `mobile_no`, `alternate_contact_no`) VALUES (1, 'fname1', 'lname1', '9813123121', '9713131334');
INSERT INTO `spl_master_usr_details_tbl` (`usr_id_fk`, `fname`, `lname`, `mobile_no`, `alternate_contact_no`) VALUES (2, 'fname2', 'lname2', '9672123121', '9123131334');



INSERT INTO `spl_master_server_register` (`id`, `server_type_code`, `server_address`, `prod_id_fk`, `server_state`, `server_state_since`) VALUES (1, 'SPL', 'ws://172.105.232.148:8080/ws', 1, 1, '2018-05-05 20:08:55');