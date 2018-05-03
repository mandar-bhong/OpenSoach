--
-- Dumping data for table `spl_prod_master_sp_category_tbl`
--

INSERT INTO `spl_prod_master_sp_category_tbl` (`id`,`spc_name`) VALUES (1,'Washroom - Mens');
INSERT INTO `spl_prod_master_sp_category_tbl` (`id`,`spc_name`) VALUES (2,'Washroom - Womens');

--
-- Dumping data for table `spl_prod_master_config`
--

INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Web.Service.Address',':90');

INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Cache.Address.Host','localhost');
INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Cache.Address.Port','6379');
INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Cache.Address.Password','');
INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Cache.Address.DB','2');

INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Que.Address.Host','localhost');
INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Que.Address.Port','6379');
INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Que.Address.Password','');
INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Product.Que.Address.DB','2');

INSERT INTO `spl_prod_serv_conf_type_tbl` (`conf_type_code`) VALUES ('SERVICE_DAILY_CHART');