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

INSERT INTO `spl_node_dev_status_tbl` (`dev_id_fk`,`connection_state`,`connection_state_since`,`sync_state`,`sync_state_since`,`battery_level`,`battery_level_since`) VALUES ('1', '0', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP);

--
-- Dumping data for table `spl_hkt_sp_complaint_tbl`
--

INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `closed_on`, `created_on`, `updated_on`) VALUES ('1', '1', '1', 'Complaint 1', 'Asda', '1', '2018-06-01 19:20:00', '1', '2018-06-01 19:20:01', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `closed_on`, `created_on`, `updated_on`) VALUES ('2', '1', '1', 'Complaint 2', 'Asda', '1', '2018-06-01 19:20:00', '1', '2018-06-01 19:20:01', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `closed_on`, `created_on`, `updated_on`) VALUES ('3', '1', '1', 'Complaint 3', 'Asda', '1', '2018-06-01 19:20:00', '1', '2018-06-01 19:20:01', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('4', '1', '1', 'Complaint 4', 'Asda', '1', '2018-05-01 19:20:00', '1', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('5', '1', '1', 'Complaint 5', 'Asda', '1', '2018-04-01 19:20:00', '2', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('6', '1', '1', 'Complaint 6', 'Asda', '2', '2018-04-01 19:20:00', '1', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('7', '1', '1', 'Complaint 7', 'Asda', '4', '2018-04-01 19:20:00', '1', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('8', '1', '1', 'Complaint 8', 'Asda', '3', '2018-01-01 19:20:00', '3', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('9', '1', '1', 'Complaint 9', 'Asda', '2', '2018-02-01 19:20:00', '1', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('10', '1', '1', 'Complaint 10', 'Asda', '1', '2018-03-01 19:20:00', '2', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('11', '1', '1', 'Complaint 11', 'Asda', '2', '2018-02-01 19:20:00', '1', '2018-06-01 19:20:02', '2018-06-01 19:20:02');
INSERT INTO `spl_hkt_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `created_on`, `updated_on`) VALUES ('12', '1', '1', 'Complaint 12', 'Asda', '1', '2018-03-01 19:20:00', '3', '2018-06-01 19:20:02', '2018-06-01 19:20:02');

--
-- Dumping data for table `spl_node_service_in_txn_tbl`
--

INSERT INTO `spl_node_service_in_txn_tbl` (`id`, `cpm_id_fk`, `serv_in_id_fk`, `fopcode`, `status`, `txn_data`, `txn_date`, `created_on`, `updated_on`) VALUES
	(1, 1, 1, '11', 1, '{"taskname": "Urinal", "slotendtime": 630, "slotstarttime": 600}', '2018-06-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(2, 1, 1, '23', 2, '{"taskname": "Hand Wash", "slotendtime": 510, "slotstarttime": 480}', '2018-06-05 14:59:44', '2018-06-05 15:00:07', '2018-06-05 15:00:07'),
	(3, 1, 1, '33', 1, '{"taskname": "Urinal", "slotendtime": 630, "slotstarttime": 600}', '2018-03-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(4, 1, 1, '44', 2, '{"taskname": "Urinal", "slotendtime": 510, "slotstarttime": 480}', '2018-03-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(5, 1, 1, '55', 1, '{"taskname": "Sanitary Ware", "slotendtime": 630, "slotstarttime": 600}', '2018-04-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(6, 1, 1, '66', 1, '{"taskname": "Dustbins", "slotendtime": 630, "slotstarttime": 600}', '2018-04-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(7, 1, 1, '77', 2, '{"taskname": "Mopping",  "slotendtime": 510, "slotstarttime": 480}', '2018-01-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(8, 1, 1, '88', 2, '{"taskname": "Air Freshner",  "slotendtime": 510, "slotstarttime": 480}', '2018-01-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(9, 1, 1, '88', 1, '{"taskname": "Doors", "slotendtime": 630, "slotstarttime": 600}', '2018-01-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(10, 1, 1, '77', 1, '{"taskname": "Sanitary Ware", "slotendtime": 630, "slotstarttime": 600}', '2018-02-05 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(11, 1, 1, '88', 1, '{"taskname": "Doors", "slotendtime": 630, "slotstarttime": 600}', '2018-06-22 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(12, 1, 1, '77', 1, '{"taskname": "Air Freshner", "slotendtime": 630, "slotstarttime": 600}', '2018-06-22 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(13, 1, 1, '66', 1, '{"taskname": "Dustbins", "slotendtime": 630, "slotstarttime": 600}', '2018-06-22 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(14, 1, 1, '66', 1, '{"taskname": "Mopping", "slotendtime": 630, "slotstarttime": 600}', '2018-06-22 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20'),
	(15, 1, 1, '77', 1, '{"taskname": "Urinal", "slotendtime": 630, "slotstarttime": 600}', '2018-06-22 13:00:17', '2018-06-05 13:00:18', '2018-06-05 14:58:20');
	
--
-- Dumping data for table `spl_node_feedback_tbl`
--
	
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('1', '1', '1', '1', '2018-06-14 11:04:05', '2018-06-14 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('2', '1', '1', '2', '2018-05-13 11:04:05', '2018-05-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('3', '1', '1', '1', '2018-05-13 11:04:05', '2018-05-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('4', '1', '1', '3', '2018-04-13 11:04:05', '2018-04-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('5', '1', '1', '1', '2018-03-13 11:04:05', '2018-03-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('6', '1', '2', '2', '2018-02-13 11:04:05', '2018-02-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('7', '1', '1', '1', '2018-02-13 11:04:05', '2018-02-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('8', '1', '2', '3', '2018-01-13 11:04:05', '2018-01-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('9', '1', '1', '4', '2018-01-13 11:04:05', '2018-01-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('10', '1', '1', '5', '2018-02-13 11:04:05', '2018-02-13 11:04:05');
	