--
-- Dumping data for table `spl_node_cpm_tbl`
--

INSERT INTO `spl_node_cpm_tbl` (`cpm_id_fk`) VALUES ('4');

--
-- Dumping data for table `spl_node_dev_tbl`
--

INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('1', '4','device 1','1234567890123456');

--
-- Dumping data for table `spl_node_sp_category_tbl`
--

INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (1,4, 'Washroom - Mens');
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (2,4, 'Washroom - Womens');

--
-- Dumping data for table `spl_node_sp_tbl`
--

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('4', '4', '1', 'Service Point 4', '1', UTC_TIMESTAMP);

--
-- Dumping data for table `spl_node_dev_sp_mapping`
--

INSERT INTO `spl_node_dev_sp_mapping` (`dev_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('1', '4', '4');


--
-- Dumping data for table `spl_node_service_conf_tbl`
--

INSERT INTO `spl_node_service_conf_tbl` (`cpm_id_fk`, `spc_id_fk`, `conf_type_code`, `serv_conf_name`, `serv_conf`) VALUES 
('4', '1', 'Conf_Type_1', 'Chart 1', '{"timeconf":{"starttime":480, "interval":30, "endtime": 1020}, "taskconf":{"tasks":[{"taskname": "Urinal"}, {"taskname": "Wash Basin"}, {"taskname": "Hand Wash"},{"taskname": "Tissue Roll"},{"taskname": "Sanitary Ware"},{"taskname": "Dustbins"},{"taskname": "Mopping"},{"taskname": "Air Freshner"},{"taskname": "Doors"}, {"taskname": "Wash Platform"}]}}'),
('4', '1', 'Conf_Type_2', 'Chart 2', '{"timeconf":{"starttime":480, "interval":30, "endtime": 1020}, "taskconf":{"tasks":[{"taskname": "Urinal"}, {"taskname": "Wash Basin"}, {"taskname": "Hand Wash"},{"taskname": "Tissue Roll"},{"taskname": "Sanitary Ware"},{"taskname": "Dustbins"},{"taskname": "Mopping"},{"taskname": "Air Freshner"},{"taskname": "Doors"}, {"taskname": "Wash Platform"}]}}'),
('4', '1', 'Conf_Type_3', 'Chart 3', '{"timeconf":{"starttime":480, "interval":30, "endtime": 1020}, "taskconf":{"tasks":[{"taskname": "Urinal"}, {"taskname": "Wash Basin"}, {"taskname": "Hand Wash"},{"taskname": "Tissue Roll"},{"taskname": "Sanitary Ware"},{"taskname": "Dustbins"},{"taskname": "Mopping"},{"taskname": "Air Freshner"},{"taskname": "Doors"}, {"taskname": "Wash Platform"}]}}');

--
-- Dumping data for table `spl_node_field_operator_tbl`
--

INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Rohini Thakre','rohini.thakre@gmail.com','4', '1111', '9911223344', '1', '1');
INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Pooja Dessai','pooja.dessai@gmail.com','4', '2222', '9811223344', '1', '2');


--
-- Dumping data for table `spl_node_fop_sp_tbl`
--

INSERT INTO `spl_node_fop_sp_tbl` (`fop_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('2', '4', '4');


--
-- Dumping data for table `spl_node_service_instance_tbl`
--

INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('4', '1', '4');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('4', '2', '4');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('4', '3', '4');


--
-- Dumping data for table `spl_node_dev_status_tbl`
--

INSERT INTO `spl_node_dev_status_tbl` (`dev_id_fk`,`connection_state`,`connection_state_since`,`sync_state`,`sync_state_since`,`battery_level`,`battery_level_since`) VALUES ('1', '0', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP);


--
-- Dumping data for table `spl_vst_vehicle_master_tbl`
--

INSERT INTO `spl_vst_vehicle_master_tbl` (`cpm_id_fk`, `vehicle_no`, `details`, `created_on`, `updated_on`) VALUES ('4', 'MH 14 AB 1234', '{"owner": {"Name": "Sarang Patil"}, "vehicle": {"make": "honda", "model": "cbz"}}', UTC_TIMESTAMP, UTC_TIMESTAMP);


--
-- Dumping data for table `spl_vst_token`
--

INSERT INTO `spl_vst_token` (`token`, `vhl_id_fk`, `mapping_details`, `state`, `generated_on`, `created_on`, `updated_on`) VALUES ('1', '1', '{"jobexeid": 0, "jobcreationid": 0, "tokenconfigid": 0}', '1', '2018-08-21 15:58:35', UTC_TIMESTAMP, UTC_TIMESTAMP);