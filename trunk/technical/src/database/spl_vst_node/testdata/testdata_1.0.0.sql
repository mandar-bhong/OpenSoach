--
-- Dumping data for table `spl_node_cpm_tbl`
--

INSERT INTO `spl_node_cpm_tbl` (`cpm_id_fk`) VALUES ('4');

--
-- Dumping data for table `spl_node_dev_tbl`
--

INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('1', '4','device 1','1234567890123456');
INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('2', '4','device 2','1345494544733456');
INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('3', '4','device 3','1155623421323222');
--
-- Dumping data for table `spl_node_sp_category_tbl`
--

INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (1,4, 'Token Generation');
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (2,4, 'Job Creation');
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (3,4, 'Job Execution');

--
-- Dumping data for table `spl_node_sp_tbl`
--

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('4', '4', '1', 'Service Point 4', '1', UTC_TIMESTAMP);
INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('5', '4', '2', 'Service Point 5', '1', UTC_TIMESTAMP);
INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('6', '4', '3', 'Service Point 6', '1', UTC_TIMESTAMP);

--
-- Dumping data for table `spl_node_dev_sp_mapping`
--

INSERT INTO `spl_node_dev_sp_mapping` (`dev_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('1', '4', '4');
INSERT INTO `spl_node_dev_sp_mapping` (`dev_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('2', '5', '4');
INSERT INTO `spl_node_dev_sp_mapping` (`dev_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('3', '6', '4');


--
-- Dumping data for table `spl_node_service_conf_tbl`
--

INSERT INTO `spl_node_service_conf_tbl` (`cpm_id_fk`, `spc_id_fk`, `conf_type_code`, `serv_conf_name`, `serv_conf`) VALUES 
('4', '1', 'TOKEN_GENERATION', 'Chart 1', '{}'),
('4', '2', 'JOB_CREATION', 'Chart 2', '{}'),
('4', '3', 'JOB_EXECUTION', 'Chart 3', '{}');

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
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('4', '2', '5');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('4', '3', '6');


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


--
-- Dumping data for table `spl_node_sp_complaint_tbl`
--

INSERT INTO `spl_node_sp_complaint_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `complaint_title`, `complaint_by`, `severity`, `raised_on`, `complaint_state`, `closed_on`, `created_on`, `updated_on`) VALUES 
('1', '4', '4', 'Complaint 1', 'Asda', '1', '2018-01-04 19:20:00', '3', '2018-01-06 19:20:01', '2018-01-04 19:20:00', '2018-01-06 19:20:01'),
('2', '4', '4', 'Complaint 2', 'Asda', '1', '2018-01-22 19:20:00', '3', '2018-01-23 19:20:01', '18-01-22 19:20:00', '2018-01-23 19:20:01'),
('3', '4', '4', 'Complaint 3', 'Asda', '1', '2018-02-11 19:20:00', '3', '2018-02-14 19:20:00', '2018-02-11 19:20:00', '2018-02-14 19:20:00'),
('4', '4', '4', 'Complaint 4', 'Asda', '1', '2018-02-16 19:20:00', '3', '2018-02-17 19:20:01', '2018-02-16 19:20:00', '2018-02-17 19:20:01'),
('5', '4', '4', 'Complaint 5', 'Asda', '1', '2018-02-21 19:20:00', '3', '2018-02-22 12:45:11', '2018-02-21 19:20:00', '2018-02-22 12:45:11'),
('6', '4', '4', 'Complaint 6', 'Asda', '2', '2018-02-28 19:20:00', '3', '2018-03-01 19:20:01', '2018-02-28 19:20:00', '2018-03-01 19:20:01'),
('7', '4', '4', 'Complaint 7', 'Asda', '4', '2018-03-01 19:20:00', '3', '2018-03-02 19:20:01', '2018-03-01 19:20:00', '2018-03-02 19:20:01'),
('8', '4', '4', 'Complaint 8', 'Asda', '3', '2018-03-25 19:20:00', '3', '2018-03-26 19:20:01', '2018-03-25 19:20:00', '2018-03-26 19:20:01'),
('9', '4', '4', 'Complaint 9', 'Asda', '2', '2018-04-14 19:20:00', '3', '2018-04-15 19:20:01', '2018-04-14 19:20:00', '2018-04-15 19:20:01'),
('10', '4', '4', 'Complaint 10', 'Asda', '1', '2018-05-08 19:20:00', '3', '2018-05-09 12:45:11', '2018-05-08 19:20:00', '2018-05-09 12:45:11'),
('11', '4', '4', 'Complaint 11', 'Asda', '2', '2018-05-13 19:20:00', '3', '2018-05-14 19:20:01', '2018-05-13 19:20:00', '2018-05-14 19:20:01'),
('12', '4', '4', 'Complaint 12', 'Asda', '1', '2018-05-17 19:20:00', '3', '2018-05-18 19:20:01', '2018-05-17 19:20:00', '2018-05-18 19:20:01'),
('13', '4', '4', 'Complaint 12', 'Asda', '1', '2018-05-23 19:20:00', '3', '2018-05-24 19:20:01', '2018-05-23 19:20:00', '2018-05-24 19:20:01'),
('14', '4', '4', 'Complaint 12', 'Asda', '1', '2018-06-20 19:20:00', '3', '2018-06-21 19:20:01', '2018-06-20 19:20:00', '2018-06-21 19:20:01'),
('15', '4', '4', 'Complaint 12', 'Asda', '1', '2018-06-21 19:20:00', '3', '2018-06-23 19:20:01', '2018-06-21 19:20:00', '2018-06-23 19:20:01'),
('16', '4', '4', 'Complaint 12', 'Asda', '1', '2018-07-17 19:20:00', '3', '2018-07-18 19:20:01', '2018-07-17 19:20:00', '2018-07-18 19:20:01'),
('17', '4', '4', 'Complaint 12', 'Asda', '1', '2018-07-19 19:20:00', '2', NULL, '2018-07-19 19:20:00', '2018-07-19 20:20:02'),
('18', '4', '4', 'Complaint 12', 'Asda', '1', '2018-07-29 19:20:00', '2', NULL, '2018-07-29 19:20:00', '2018-07-29 20:20:02'),
('19', '4', '4', 'Complaint 12', 'Asda', '1', '2018-08-02 19:20:00', '3', '2018-08-02 19:20:01', '2018-08-02 19:20:00', '2018-08-02 19:20:01'),
('20', '4', '4', 'Complaint 12', 'Asda', '1', '2018-08-12 19:20:00', '2', NULL, '2018-08-12 19:20:00', '2018-08-12 20:20:02'),
('21', '4', '4', 'Complaint 12', 'Asda', '1', '2018-08-15 19:20:00', '2', NULL, '2018-08-15 19:20:00', '2018-08-15 20:20:02'),
('22', '4', '4', 'Complaint 12', 'Asda', '1', '2018-09-03 19:20:00', '1', NULL, '2018-09-03 19:20:00', '2018-09-03 19:20:00'),
('23', '4', '4', 'Complaint 12', 'Asda', '1', '2018-09-09 19:20:00', '1', NULL, '2018-09-09 19:20:00', '2018-09-09 19:20:00');


--
-- Dumping data for table `spl_node_feedback_tbl`
--
	
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('1', '4', '4', '5', '2018-01-10 11:04:05', '2018-01-10 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('2', '4', '4', '5', '2018-01-12 11:04:05', '2018-01-12 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('3', '4', '4', '5', '2018-01-25 11:04:05', '2018-01-25 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('4', '4', '4', '4', '2018-02-08 11:04:05', '2018-02-08 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('5', '4', '4', '4', '2018-02-15 11:04:05', '2018-02-15 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('6', '4', '4', '5', '2018-02-18 11:04:05', '2018-02-18 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('7', '4', '4', '4', '2018-03-05 11:04:05', '2018-03-05 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('8', '4', '4', '5', '2018-03-11 11:04:05', '2018-03-11 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('9', '4', '4', '4', '2018-03-13 11:04:05', '2018-03-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('10', '4', '4', '5', '2018-03-14 11:04:05', '2018-03-14 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('11', '4', '4', '5', '2018-03-18 11:04:05', '2018-03-18 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('12', '4', '4', '5', '2018-03-24 11:04:05', '2018-03-24 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('13', '4', '4', '5', '2018-04-01 11:04:05', '2018-04-01 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('14', '4', '4', '5', '2018-04-14 11:04:05', '2018-04-14 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('15', '4', '4', '5', '2018-04-19 11:04:05', '2018-04-19 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('16', '4', '4', '5', '2018-04-21 11:04:05', '2018-04-21 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('17', '4', '4', '5', '2018-04-27 11:04:05', '2018-04-27 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('18', '4', '4', '3', '2018-05-02 11:04:05', '2018-05-02 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('19', '4', '4', '5', '2018-06-07 11:04:05', '2018-06-07 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('20', '4', '4', '5', '2018-06-14 11:04:05', '2018-06-14 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('21', '4', '4', '5', '2018-06-15 11:04:05', '2018-06-15 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('22', '4', '4', '4', '2018-07-22 11:04:05', '2018-07-22 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('23', '4', '4', '5', '2018-07-23 11:04:05', '2018-07-23 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('24', '4', '4', '4', '2018-08-01 11:04:05', '2018-08-01 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('25', '4', '4', '5', '2018-08-17 11:04:05', '2018-08-17 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('26', '4', '4', '5', '2018-08-24 11:04:05', '2018-08-24 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('27', '4', '4', '5', '2018-08-29 11:04:05', '2018-08-29 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('28', '4', '4', '5', '2018-09-13 11:04:05', '2018-09-13 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('29', '4', '4', '5', '2018-09-16 11:04:05', '2018-09-16 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('30', '4', '4', '5', '2018-09-22 11:04:05', '2018-09-22 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('31', '4', '4', '4', '2018-09-23 11:04:05', '2018-09-23 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('32', '4', '4', '2', '2018-09-24 11:04:05', '2018-09-24 11:04:05');
INSERT INTO `spl_node_feedback_tbl` (`id`, `cpm_id_fk`, `sp_id_fk`, `feedback`, `raised_on`, `created_on`) VALUES ('33', '4', '4', '1', '2018-10-08 11:04:05', '2018-10-08 11:04:05');
