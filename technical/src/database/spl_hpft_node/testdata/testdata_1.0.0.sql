--
-- Dumping data for table `spl_node_cpm_tbl`
--

INSERT INTO `spl_node_cpm_tbl` (`cpm_id_fk`) VALUES ('3');


--
-- Dumping data for table `spl_node_dev_tbl`
--

INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('1', '3','device 1','1234567890123456');


--
-- Dumping data for table `spl_node_sp_category_tbl`
--

INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (1,3, 'General Ward');
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (2,3, 'Private Room');
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (3,3, 'Semi Private');
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (4,3, 'ICU');


--
-- Dumping data for table `spl_node_sp_tbl`
--

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('3', '3', '1', 'Service Point 1', '1', UTC_TIMESTAMP);


--
-- Dumping data for table `spl_node_dev_sp_mapping`
--

INSERT INTO `spl_node_dev_sp_mapping` (`dev_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('1', '3', '3');


--
-- Dumping data for table `spl_node_dev_status_tbl`
--

INSERT INTO `spl_node_dev_status_tbl` (`dev_id_fk`,`connection_state`,`connection_state_since`,`sync_state`,`sync_state_since`,`battery_level`,`battery_level_since`) VALUES ('1', '0', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP);

--
-- Dumping data for table `spl_node_service_conf_tbl`
--

INSERT INTO `spl_node_service_conf_tbl` (`cpm_id_fk`, `spc_id_fk`, `conf_type_code`, `serv_conf_name`, `serv_conf`) VALUES ('3', '3', 'SERVICE_DAILY_CHART', 'Patient File Template 1', '{"taskconf":{"tasks":[{"taskname":"Monitor Temperature","fields":["Value","Comments"]},{"taskname":"Monitor Pressure","fields":["Value","Comments"]},{"taskname":"Saline 250ML","fields":["Comments"]},{"taskname":"Monitor Heart Rate","fields":["Value","Comments"]},{"taskname":"Monitor Blood Pressure","fields":["Value","Comments"]},{"taskname":"Physiotherapy","fields":["Comments"]},{"taskname":"Dressing","fields":["Comments"]}]},"timeconf":{"endtime":1020,"interval":30,"starttime":480}}');


--
-- Dumping data for table `spl_node_service_instance_tbl`
--

INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');


--
-- Dumping data for table `spl_hpft_patient_master_tbl`
--

INSERT INTO `spl_hpft_patient_master_tbl` (`cpm_id_fk`, `patient_details`, `medical_details`, `patient_file_template`, `sp_id_fk`, `serv_in_id_fk`, `status`) VALUES 
('3', '{"age": "67", "bedno": "12", "patientname": "Patient1", "admissiondate": "2018-08-05T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "4532234346", "patientregistrationno": "233966567657"}', '{"allergies": "allergy1", "treatmentdone": "traetment1", "reasonadmission": "reason1", "patientmedicalhistory": "history1"}', '1', '3', '1', '1'),
('3', '{"age": "76", "bedno": "13", "patientname": "Patient2", "admissiondate": "2018-08-05T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9843242382", "patientregistrationno": "5423477234236"}', '{"allergies": "allergy2", "treatmentdone": "traetment2", "reasonadmission": "reason2", "patientmedicalhistory": "history2"}', '1', '3', '2', '1');


--
-- Dumping data for table `spl_node_field_operator_tbl`
--

INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Operator 1','operator1@cust1.com','3', '1234', '1222', '1', '1');
INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Operator 2','operator2@cust1.com','3', '445', '222', '1', '2');


--
-- Dumping data for table `spl_node_fop_sp_tbl`
--

INSERT INTO `spl_node_fop_sp_tbl` (`fop_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('2', '3', '3');


--
-- Dumping data for table `spl_node_service_in_txn_tbl`
--


INSERT INTO `spl_node_service_in_txn_tbl` (`cpm_id_fk`, `serv_in_id_fk`, `fopcode`, `status`, `txn_data`, `txn_date`, `created_on`, `updated_on`) VALUES
	(3, 1, '11', 1, '{"value": 140, "comment": "High Blood Pressure", "taskname": "Monitor Blood Pressure", "slotendtime": 630, "slotstarttime": 600}', '2018-08-05 19:13:19', '2018-08-07 16:13:19', '2018-08-07 16:13:19'),
	(3, 1, '11', 1, '{"value": 140, "comment": "High Blood Pressure", "taskname": "Monitor Blood Pressure", "slotendtime": 630, "slotstarttime": 600}', '2018-08-07 16:13:39', '2018-08-07 16:13:19', '2018-08-07 16:13:19'),
	(3, 1, '11', 1, '{"comment": "Saline", "taskname": "Saline 250ML", "slotendtime": 630, "slotstarttime": 600}', '2018-08-07 16:13:39', '2018-08-07 16:13:19', '2018-08-07 16:13:19'),
	(3, 1, '11', 1, '{"value": 150, "comment": "High Blood Pressure", "taskname": "Monitor Blood Pressure", "slotendtime": 730, "slotstarttime": 700}', '2018-08-07 16:13:39', '2018-08-07 16:13:19', '2018-08-07 16:13:19');



