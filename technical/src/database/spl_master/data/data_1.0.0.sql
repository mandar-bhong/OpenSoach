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
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (2,1,'ADMIN','Administrator');

--
-- Dumping data for table `spl_master_config`
--

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.Host','localhost');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.Port','6379');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.DB','0');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Cache.Address.Password','');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Web.Service.Address',':80');

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('HKT.Master.DB.Connection','root:welcome@tcp(localhost:3306)/spl_hkt_master?parseTime=true');

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Server.Win.BaseDir','');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Server.Lin.BaseDir','/opt/build/spl/SPLBuild/');

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Que.Address.Host','localhost');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Que.Address.Port','6379');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Que.Address.Password','');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Que.Address.DB','2');

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Logger.Log.Level','Debug');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Logger.LogType','Std');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Logger.Fluent.Host','http://172.105.232.148:9999

--
-- Dumping data for table `spl_master_user_tbl`
--
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (1,'admin@servicepoint.live','admin',1,1,1,UTC_TIMESTAMP);
