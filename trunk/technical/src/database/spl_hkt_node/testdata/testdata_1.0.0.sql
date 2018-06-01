--
-- Dumping data for table `spl_node_cpm_tbl`
--

INSERT INTO `spl_node_cpm_tbl` (`cpm_id_fk`) VALUES ('1');

--
-- Dumping data for table `spl_node_dev_tbl`
--

INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('1', '1','device 1','1234567890123456');


--
-- Dumping data for table `spl_node_sp_category_tbl`
--

INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (1,1, 'Washroom - Mens');
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (2,1, 'Washroom - Womens');


INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('1', '1', '1', 'Service Point 1', '1', '2018-05-22 21:03:08');

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('2', '1', '1', 'Service Point 2', '1', '2018-05-22 21:03:08');

INSERT INTO `spl_node_dev_sp_mapping` (`dev_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('1', '1', '1');


INSERT INTO `spl_node_service_conf_tbl` (`cpm_id_fk`, `spc_id_fk`, `conf_type_code`, `serv_conf_name`, `serv_conf`) VALUES ('1', '1', 'SERVICE_DAILY_CHART', 'Chart 1', '{"timeconf":{"starttime":480, "interval":30, "endtime": 1020}, "taskconf":{"tasks":[{"taskname": "Urinal"}, {"taskname": "Wash Basin"}, {"taskname": "Hand Wash"},{"taskname": "Tissue Roll"},{"taskname": "Sanitary Ware"},{"taskname": "Dustbins"},{"taskname": "Mopping"},{"taskname": "Air Freshner"},{"taskname": "Doors"}, {"taskname": "Wash Platform"}]}}');

INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('1', '1', '1');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('1', '1', '2');


INSERT INTO `spl_node_field_operator_tbl` (`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('1', '1234', '1222', '1', '1');
INSERT INTO `spl_node_field_operator_tbl` (`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('1', '445', '222', '1', '2');

INSERT INTO `spl_node_fop_sp_tbl` (`fop_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('2', '1', '1');

--
-- Dumping data for table `spl_node_dev_status_tbl`
--

INSERT INTO `spl_node_dev_status_tbl` (`dev_id_fk`,`connection_state`,`connection_state_since`,`sync_state`,`sync_state_since`,`battery_level`,`battery_level_since`) VALUES ('1', '2', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP);

--
-- Dumping data for table `spl_hkt_sp_complaint_tbl`
--

INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `closed_on`, `created_on`, `updated_on`) VALUES ('1', '1', '1', 'Complaint 1', 'Asda', '1', '2018-06-01 19:20:00', '1', '2018-06-01 19:20:01', '2018-06-01 19:20:02', '2018-06-01 19:20:02');