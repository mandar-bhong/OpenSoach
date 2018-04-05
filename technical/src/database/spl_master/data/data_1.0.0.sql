--
-- Dumping data for table `spl_master_product_tbl`
--

INSERT INTO `spl_master_product_tbl` VALUES (1,'SPL_HKT');

--
-- Dumping data for table `spl_master_database_instance_tbl`
--

INSERT INTO `spl_master_database_instance_tbl` VALUES (1,'spl_hkt_node_0001','root:welcome@tcp(localhost:3306)/spl_hkt_node_0001?parseTime=true',1);

--
-- Dumping data for table `spl_master_total_count_tbl`
--

INSERT INTO `spl_master_total_count_tbl` VALUES (1,0,0,0,0,0,0);

--
-- Dumping data for table `spl_master_user_role_tbl`
--

INSERT INTO `spl_master_user_role_tbl` (`urole_code`,`urole_name`) VALUES ('ADMIN','Administrator');

--
-- Dumping data for table `spl_master_config`
--

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address','localhost:6379');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.DB','0');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.Password','');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Web.Service.Address','localhost:80');
