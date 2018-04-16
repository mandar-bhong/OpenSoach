--
-- Dumping data for table `spl_master_product_tbl`
--

INSERT INTO `spl_master_product_tbl` (`id`,`prod_code`) VALUES (1,'SPL_HKT');

--
-- Dumping data for table `spl_master_database_instance_tbl`
--

INSERT INTO `spl_master_database_instance_tbl` (`id`,`dbi_name`,`connection_string`,`prod_id_fk`) VALUES (1,'spl_hkt_node_0001','root:welcome@tcp(localhost:3306)/spl_hkt_node_0001?parseTime=true',1);

--
-- Dumping data for table `spl_master_total_count_tbl`
--

INSERT INTO `spl_master_total_count_tbl` (`id`,`cust_cnt`,`usr_cnt`,`dev_cnt`,`sp_cnt`,`dev_unallocated_cnt`,`dev_active_cnt`) VALUES (1,0,0,0,0,0,0);

--
-- Dumping data for table `spl_master_user_role_tbl`
--

INSERT INTO `spl_master_user_role_tbl` (`id`,`urole_code`,`urole_name`) VALUES (1,'ADMIN','Administrator');

--
-- Dumping data for table `spl_master_config`
--

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address','localhost:6379');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.DB','0');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.Password','');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Web.Service.Address',':80');

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Server.Win.BaseDir','');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Server.Lin.BaseDir','/opt/build/spl/SPLBuild/');
