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
INSERT INTO `spl_prod_master_config` (`config_key`,`config_value`) VALUES ('Spl.Prod.Base.Url','http://172.105.232.148');


--
-- Dumping data for table `spl_prod_master_sp_category_tbl`
--

INSERT INTO `spl_prod_master_sp_category_tbl` (`id`,`spc_name`) VALUES (1,'General Ward');
INSERT INTO `spl_prod_master_sp_category_tbl` (`id`,`spc_name`) VALUES (2,'Private Room');
INSERT INTO `spl_prod_master_sp_category_tbl` (`id`,`spc_name`) VALUES (3,'Semi Private');
INSERT INTO `spl_prod_master_sp_category_tbl` (`id`,`spc_name`) VALUES (4,'ICU');