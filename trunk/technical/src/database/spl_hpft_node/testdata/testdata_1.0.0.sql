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

INSERT INTO `spl_node_service_conf_tbl` (`cpm_id_fk`, `spc_id_fk`, `conf_type_code`, `serv_conf_name`, `serv_conf`) VALUES ('3', '3', 'SERVICE_DAILY_CHART', 'Patient File Template 1', '{"taskconf":{"tasks":[{"taskname":"Monitor Temperature","fields":["Value","Comments"]},{"taskname":"Monitor Pressure","fields":["Value","Comments"]},{"taskname":"Saline 250ML","fields":["Comments"]},{"taskname":"Monitor Heart Rate","fields":["Value","Comments"]},{"taskname":"Monitor Blood Pressure","fields":["Value","Comments"]},{"taskname":"Physiotherapy","fields":["Comments"]},{"taskname":"Dressing","fields":["Comments"]}]},"timeconf":{"endtime":1440,"interval":240,"starttime":0}}');


--
-- Dumping data for table `spl_node_service_instance_tbl`
--

INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');


--
-- Dumping data for table `spl_hpft_patient_master_tbl`
--

INSERT INTO `spl_hpft_patient_master_tbl` (`cpm_id_fk`, `patient_details`, `medical_details`, `patient_file_template`, `sp_id_fk`, `serv_in_id_fk`, `status`) VALUES 
('3', '{"age": "67", "bedno": "12", "patientname": "Patient1", "admissiondate": "2018-08-05T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "4532234346", "patientregistrationno": "233966567657","bloodgroup":"A+","weight":"64kg","drinst":"doctor instance1"}', '{"allergies": "allergy1", "treatmentdone": "traetment1", "reasonadmission": "reason1", "patientmedicalhistory": "history1"}', '1', '3', '1', '1'),
('3', '{"age": "76", "bedno": "13", "patientname": "Patient2", "admissiondate": "2018-08-05T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9843242382", "patientregistrationno": "5423477234236","bloodgroup":"A-","weight":"70kg","drinst":"doctor instance2"}', '{"allergies": "allergy2", "treatmentdone": "traetment2", "reasonadmission": "reason2", "patientmedicalhistory": "history2"}', '1', '3', '2', '1'),
('3', '{"age": "19", "bedno": "14", "patientname": "Patient3", "admissiondate": "2018-08-06T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "34433242382", "patientregistrationno": "5255345234236","bloodgroup":"AB+","weight":"55kg","drinst":"doctor instance3"}', '{"allergies": "allergy3", "treatmentdone": "traetment3", "reasonadmission": "reason3", "patientmedicalhistory": "history2"}', '1', '3', '3', '1'),
('3', '{"age": "42", "bedno": "15", "patientname": "Patient4", "admissiondate": "2018-08-07T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "8673242382", "patientregistrationno": "892234234236","bloodgroup":"O+","weight":"65kg","drinst":"doctor instance4"}', '{"allergies": "allergy4", "treatmentdone": "traetment4", "reasonadmission": "reason4", "patientmedicalhistory": "history2"}', '1', '3', '4', '1'),
('3', '{"age": "55", "bedno": "17", "patientname": "Patient5", "admissiondate": "2018-08-08T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9563242432", "patientregistrationno": "13423477234236","bloodgroup":"0-","weight":"80kg","drinst":"doctor instance5"}', '{"allergies": "allergy5", "treatmentdone": "traetment5", "reasonadmission": "reason5", "patientmedicalhistory": "history2"}', '1', '3', '5', '1'),
('3', '{"age": "15", "bedno": "18", "patientname": "Patient6", "admissiondate": "2018-08-08T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9123342432", "patientregistrationno": "12423477234236","bloodgroup":"B+","weight":"74kg","drinst":"doctor instance6"}', '{"allergies": "allergy6", "treatmentdone": "traetment6", "reasonadmission": "reason6", "patientmedicalhistory": "history2"}', '1', '3', '6', '1'),
('3', '{"age": "25", "bedno": "19", "patientname": "Patient7", "admissiondate": "2018-08-08T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9063248932", "patientregistrationno": "14423477234236","bloodgroup":"AB-","weight":"93kg","drinst":"doctor instance7"}', '{"allergies": "allergy7", "treatmentdone": "traetment7", "reasonadmission": "reason7", "patientmedicalhistory": "history2"}', '1', '3', '7', '1');


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


