--
-- Dumping data for table `spl_master_product_tbl`
--

INSERT INTO `spl_master_product_tbl` (`id`,`prod_code`) VALUES (1,'SPL_HKT'),
(2,'SPL_HPFT'),
(3,'SPL_VST');

--
-- Dumping data for table `spl_master_database_instance_tbl`
--

INSERT INTO `spl_master_database_instance_tbl` (`id`,`dbi_name`,`connection_string`,`prod_id_fk`) VALUES (1,'spl_hkt_node_0001','root:welcome@tcp(localhost:3306)/spl_hkt_node_0001?parseTime=true',1),
(2,'spl_hpft_node_0001','root:welcome@tcp(localhost:3306)/spl_hpft_node_0001?parseTime=true',2),
(3,'spl_vst_node_0001','root:welcome@tcp(localhost:3306)/spl_vst_node_0001?parseTime=true',3);


--
-- Dumping data for table `spl_master_total_count_tbl`
--

INSERT INTO `spl_master_total_count_tbl` (`id`,`cust_cnt`,`usr_cnt`,`dev_cnt`,`sp_cnt`,`dev_unallocated_cnt`,`dev_active_cnt`) VALUES (1,0,0,0,0,0,0);

--
-- Dumping data for table `spl_master_user_role_tbl`
--

INSERT INTO `spl_master_user_role_tbl` (`id`,`urole_code`,`urole_name`) VALUES (1,'ADMIN','Administrator');
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (2,1,'ADMIN','Administrator');
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (3,2,'ADMIN','Administrator');
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (4,3,'ADMIN','Administrator');
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (5,2,'EX_DOC','External Doctor');
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (6,2,'IN_DOC','Internal Doctor');
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (7,2,'NURSE','Nurse');
INSERT INTO `spl_master_user_role_tbl` (`id`,`prod_id_fk`,`urole_code`,`urole_name`) VALUES (8,2,'LAB','Laboratory');

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
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Logger.LogType','InfluxDB');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Logger.Fluent.Host','http://172.105.232.148:9999');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('Logger.InfluxDB.Host','http://172.105.232.148:8086');

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('SMTP.Address','send.one.com');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('SMTP.From','support@opensoach.com');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('SMTP.Password','opensoach.support@123');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('SMTP.Port','25');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('SMTP.UserName','support@opensoach.com');

INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('HPFT.Master.DB.Connection','root:welcome@tcp(localhost:3306)/spl_hpft_master?parseTime=true');
INSERT INTO `spl_master_config` (`config_key`,`config_value`) VALUES ('VST.Master.DB.Connection','root:welcome@tcp(localhost:3306)/spl_vst_master?parseTime=true');


--
-- Dumping data for table `spl_master_user_tbl`
--
INSERT INTO `spl_master_user_tbl` (`id`,`usr_name`,`usr_password`,`usr_category`,`urole_id_fk`,`usr_state`,`usr_state_since`) VALUES (1,'admin@servicepoint.live','admin',1,1,1,UTC_TIMESTAMP);

INSERT INTO `spl_master_server_register` (`id`, `server_type_code`, `server_address`, `prod_id_fk`, `server_state`, `server_state_since`) VALUES (1, 'SPL', 'ws://172.105.232.148:8081/ws', 1, 1, '2018-05-05 20:08:55'),
(2, 'SPL', 'ws://172.105.232.148:8090/ws', 2, 1, '2018-05-05 20:08:55'),
(3, 'SPL', 'ws://172.105.232.148:8091/ws', 3, 1, '2018-05-05 20:08:55');

INSERT INTO `spl_master_email_template_tbl` (`id`, `code`, `subject`, `body`, `bcc`, `maxretry`) VALUES
	(1, 'USER_ASSOCIATED', 'User account activation', '<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head><meta name="viewport" content="width=device-width" /><meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /><title></title><style>*{margin:0;font-family:"Helvetica Neue",Helvetica,Arial,sans-serif;box-sizing:border-box;font-size:14px}img{max-width:100%}body{-webkit-font-smoothing:antialiased;-webkit-text-size-adjust:none;width:100% !important;height:100%;line-height:1.6em}table td{vertical-align:top}body{background-color:#f6f6f6}.body-wrap{background-color:#f6f6f6;width:100%}.container{display:block !important;max-width:600px !important;margin:0 auto !important;clear:both !important}.content{max-width:600px;margin:0 auto;display:block;padding:20px}.main{background-color:#fff;border:1px solid #e9e9e9;border-radius:3px}.content-wrap{padding:20px}.content-block{padding:0 0 20px}.header{width:100%;margin-bottom:20px}.footer{width:100%;clear:both;color:#999;padding:20px}.footer p, .footer a, .footer td{color:#999;font-size:12px}h1,h2,h3{font-family:"Helvetica Neue",Helvetica,Arial,"Lucida Grande",sans-serif;color:#000;margin:40px 0 0;line-height:1.2em;font-weight:400}h1{font-size:32px;font-weight:500}h2{font-size:24px}h3{font-size:18px}h4{font-size:14px;font-weight:600}p,ul,ol{margin-bottom:10px;font-weight:normal}p li, ul li, ol li{margin-left:5px;list-style-position:inside}a{color:#348eda;text-decoration:underline}.btn-primary{text-decoration:none;color:#FFF;background-color:#348eda;border:solid #348eda;border-width:10px 20px;line-height:2em;font-weight:bold;text-align:center;cursor:pointer;display:inline-block;border-radius:5px;text-transform:capitalize}.last{margin-bottom:0}.first{margin-top:0}.aligncenter{text-align:center}.alignright{text-align:right}.alignleft{text-align:left}.clear{clear:both}.alert{font-size:16px;color:#fff;font-weight:500;padding:20px;text-align:center;border-radius:3px 3px 0 0}.alert a{color:#fff;text-decoration:none;font-weight:500;font-size:16px}.alert.alert-warning{background-color:#FF9F00}.alert.alert-bad{background-color:#D0021B}.alert.alert-good{background-color:#68B90F}.invoice{margin:40px auto;text-align:left;width:80%}.invoice td{padding:5px 0}.invoice .invoice-items{width:100%}.invoice .invoice-items td{border-top:#eee 1px solid}.invoice .invoice-items .total td{border-top:2px solid #333;border-bottom:2px solid #333;font-weight:700}@media only screen and (max-width: 640px){body{padding:0 !important}h1,h2,h3,h4{font-weight:800 !important;margin:20px 0 5px !important}h1{font-size:22px !important}h2{font-size:18px !important}h3{font-size:16px !important}.container{padding:0 !important;width:100% !important}.content{padding:0 !important}.content-wrap{padding:10px !important}.invoice{width:100% !important}}</style></head><body itemscope itemtype="http://schema.org/EmailMessage"><table class="body-wrap"><tr><td></td><td class="container" width="600"><div class="content"><table class="main" width="100%" cellpadding="0" cellspacing="0" itemprop="action" itemscope itemtype="http://schema.org/ConfirmAction"><tr><td class="content-wrap"><meta itemprop="name" content="Confirm Email"/><table width="100%" cellpadding="0" cellspacing="0"><tr><td class="content-block"> Dear Customer, <img alt="Logo" title="" align="right" height="auto" src="http://172.105.232.148:91/shared/assets/images/opensoach_logo_black.png"></td></tr><tr><td class="content-block"><p>Your user account has been created.</p><p>Please click the button below to activate your account.</p></td></tr><tr><td class="content-block" itemprop="handler" itemscope itemtype="http://schema.org/HttpActionHandler"> <a href="http://172.105.232.148:91/auth/change-password?code=$Code$" class="btn-primary">Activate</a></td></tr><tr><td class="content-block"><p>Alternatively, you may copy and paste the link given below in your internet browser and to activate your account.</p><p><a href="http://172.105.232.148:91/auth/change-password?code=$ReplacableCode$">http://172.105.232.148:91/auth/change-password?code=$ReplacableCode$</a></p></td></tr><tr><td class="content-block"><div>Thank you,</div><div>Team OpenSoach</div><div><a href="www.opensoach.com">www.opensoach.com</a></div><hr style="margin-top: 56px"></td></tr><tr><td class="content-block" style="color:#808080"> This message was sent from a notification-only email address that does not accept incoming email. Please do not reply to this message.</td></tr></table></td></tr></table><div class="footer"><table width="100%"><tr><td class="aligncenter content-block"></td></tr></table></div></div></td><td></td></tr></table></body></html>', NULL, 3),
	(2, 'USER_OTP', 'PatientCare.live | Reset Password', '<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head><meta name="viewport" content="width=device-width" /><meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /><title></title><style>*{margin:0;font-family:"Helvetica Neue",Helvetica,Arial,sans-serif;box-sizing:border-box;font-size:14px}img{max-width:100%}body{-webkit-font-smoothing:antialiased;-webkit-text-size-adjust:none;width:100% !important;height:100%;line-height:1.6em}table td{vertical-align:top}body{background-color:#f6f6f6}.body-wrap{background-color:#f6f6f6;width:100%}.container{display:block !important;max-width:600px !important;margin:0 auto !important;clear:both !important}.content{max-width:600px;margin:0 auto;display:block;padding:20px}.main{background-color:#fff;border:1px solid #e9e9e9;border-radius:3px}.content-wrap{padding:20px}.content-block{padding:0 0 20px}.header{width:100%;margin-bottom:20px}.footer{width:100%;clear:both;color:#999;padding:20px}.footer p, .footer a, .footer td{color:#999;font-size:12px}h1,h2,h3{font-family:"Helvetica Neue",Helvetica,Arial,"Lucida Grande",sans-serif;color:#000;margin:40px 0 0;line-height:1.2em;font-weight:400}h1{font-size:32px;font-weight:500}h2{font-size:24px}h3{font-size:18px}h4{font-size:14px;font-weight:600}p,ul,ol{margin-bottom:10px;font-weight:normal}p li, ul li, ol li{margin-left:5px;list-style-position:inside}a{color:#348eda;text-decoration:underline}.btn-primary{text-decoration:none;color:#FFF;background-color:#348eda;border:solid #348eda;border-width:10px 20px;line-height:2em;font-weight:bold;text-align:center;cursor:pointer;display:inline-block;border-radius:5px;text-transform:capitalize}.last{margin-bottom:0}.first{margin-top:0}.aligncenter{text-align:center}.alignright{text-align:right}.alignleft{text-align:left}.clear{clear:both}.alert{font-size:16px;color:#fff;font-weight:500;padding:20px;text-align:center;border-radius:3px 3px 0 0}.alert a{color:#fff;text-decoration:none;font-weight:500;font-size:16px}.alert.alert-warning{background-color:#FF9F00}.alert.alert-bad{background-color:#D0021B}.alert.alert-good{background-color:#68B90F}.invoice{margin:40px auto;text-align:left;width:80%}.invoice td{padding:5px 0}.invoice .invoice-items{width:100%}.invoice .invoice-items td{border-top:#eee 1px solid}.invoice .invoice-items .total td{border-top:2px solid #333;border-bottom:2px solid #333;font-weight:700}@media only screen and (max-width: 640px){body{padding:0 !important}h1,h2,h3,h4{font-weight:800 !important;margin:20px 0 5px !important}h1{font-size:22px !important}h2{font-size:18px !important}h3{font-size:16px !important}.container{padding:0 !important;width:100% !important}.content{padding:0 !important}.content-wrap{padding:10px !important}.invoice{width:100% !important}}</style></head><body itemscope itemtype="http://schema.org/EmailMessage"><table class="body-wrap"><tr><td></td><td class="container" width="600"><div class="content"><table class="main" width="100%" cellpadding="0" cellspacing="0" itemprop="action" itemscope itemtype="http://schema.org/ConfirmAction"><tr><td class="content-wrap"><meta itemprop="name" content="Confirm Email"/><table width="100%" cellpadding="0" cellspacing="0"><tr><td class="content-block"> Dear Customer, <img alt="Logo" title="" align="right" height="auto" src="http://172.105.232.148:91/shared/assets/images/opensoach_logo_black.png"></td></tr><tr><td class="content-block"><p>Please use the OTP code $Code$ to reset your password.</p><p>This OTP is confidential. For security reasons, DO NOT share the OTP with anyone.</p></td></tr><tr><td class="content-block"><div>Thank you,</div><div>Team OpenSoach</div><div><a href="www.opensoach.com">www.opensoach.com</a></div><hr style="margin-top: 56px"></td></tr><tr><td class="content-block" style="color:#808080"> This message was sent from a notification-only email address that does not accept incoming email. Please do not reply to this message.</td></tr></table></td></tr></table><div class="footer"><table width="100%"><tr><td class="aligncenter content-block"></td></tr></table></div></div></td><td></td></tr></table></body></html>', NULL, 3);