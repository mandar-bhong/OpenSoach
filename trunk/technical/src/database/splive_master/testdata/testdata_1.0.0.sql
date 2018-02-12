--
-- Dumping data for table `spl_master_customer_tbl`
--

INSERT INTO `spl_master_customer_tbl` VALUES (1,'Customer 1',1);

--
-- Dumping data for table `spl_master_cust_prod_mapping_tbl`
--

INSERT INTO `spl_master_cust_prod_mapping_tbl` VALUES (1,1,1,1,1);

--
-- Dumping data for table `spl_master_device_tbl`
--

INSERT INTO `spl_master_device_tbl` VALUES (1,'1234567890123456',1);

--
-- Dumping data for table `spl_master_cpm_dev_mapping_tbl`
--

INSERT INTO `spl_master_cpm_dev_mapping_tbl` VALUES (1,1);

--
-- Dumping data for table `spl_master_servicepoint_tbl`
--

INSERT INTO `spl_master_servicepoint_tbl` VALUES (1,1);

--
-- Dumping data for table `spl_master_cpm_sp_mapping_tbl`
--

INSERT INTO `spl_master_cpm_sp_mapping_tbl` VALUES (1,1);

--
-- Dumping data for table `spl_master_total_count_tbl`
--

UPDATE `splive_master`.`spl_master_total_count_tbl` SET `cust_cnt`='1', `usr_cnt`='1', `dev_cnt`='1', `sp_cnt`='1', `dev_active_cnt`='1' WHERE `id`='1';

--
-- Dumping data for table `spl_master_user_tbl`
--

INSERT INTO `spl_master_user_tbl` VALUES (1,'admin@customer1.com','admin',1);

--
-- Dumping data for table `spl_master_usr_cpm_tbl`
--

INSERT INTO `spl_master_usr_cpm_tbl` VALUES (1,1,1);

--
-- Dumping data for table `spl_master_cust_prod_count_tbl`
--

INSERT INTO `spl_master_cust_prod_count_tbl` VALUES (1,1,1,1,1);