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

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('3', '3', '1', 'General Ward 3', '1', UTC_TIMESTAMP);


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
('3', '{"age": "35", "bedno": "3A/312", "patientname": "Sanjay Sawant", "admissiondate": "2018-08-05T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "7798847950", "patientregistrationno": "RHC-2018-3456","bloodgroup":"A+","weight":"64","drinst":"Arun Tripathi","gender":"Male"}', '{"allergies": "Allergy with Peanuts", "treatmentdone": "Under Diagnosis", "reasonadmission": "Uneasyness in the chest", "patientmedicalhistory": "Undergone treatment for bone fracture in right leg"}', '1', '3', '1', '1'),
('3', '{"age": "32", "bedno": "3B/313", "patientname": "Praveen Pandey", "admissiondate": "2018-08-06T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "7028841950", "patientregistrationno": "RHC-2018-3457","bloodgroup":"A-","weight":"70","drinst":"Arun Tripathi","gender":"Male"}', '{"allergies": "Allergy with Shellfish", "treatmentdone": "Surgery for right ear", "reasonadmission": "Swelling in Ear", "patientmedicalhistory": "None"}', '1', '3', '2', '1'),
('3', '{"age": "22", "bedno": "2A/314", "patientname": "Mandar bhong", "admissiondate": "2018-08-07T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "7938841950", "patientregistrationno": "RHC-2018-3458","bloodgroup":"AB+","weight":"55","drinst":"Arun Tripathi","gender":"Male"}', '{"allergies": "None", "treatmentdone": "traetment3", "reasonadmission": "Iritation in eye from last 15 days", "patientmedicalhistory": "None"}', '1', '3', '3', '1'),
('3', '{"age": "60", "bedno": "4A/315", "patientname": "Amol Patil", "admissiondate": "2018-08-13T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "7648841950", "patientregistrationno": "RHC-2018-3459","bloodgroup":"O+","weight":"65","drinst":"Arun Tripathi","gender":"Male"}', '{"allergies": "None", "treatmentdone": "traetment4", "reasonadmission": "Uneasyness in the chest", "patientmedicalhistory": "None"}', '1', '3', '4', '1'),
('3', '{"age": "67", "bedno": "5A/316", "patientname": "Sumeet Karande", "admissiondate": "2018-08-13T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9563242432", "patientregistrationno": "RHC-2018-3460","bloodgroup":"0-","weight":"80","drinst":"Arun Tripathi","gender":"Male"}', '{"allergies": "None", "treatmentdone": "traetment5", "reasonadmission": "reason5", "patientmedicalhistory": "None"}', '1', '3', '5', '1'),
('3', '{"age": "45", "bedno": "6A/317", "patientname": "Arun Tripathi", "admissiondate": "2018-08-13T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "8793447950", "patientregistrationno": "RHC-2018-3461","bloodgroup":"B+","weight":"74","drinst":"Arun Tripathi","gender":"Male"}', '{"allergies": "None", "treatmentdone": "traetment6", "reasonadmission": "reason6", "patientmedicalhistory": "None"}', '1', '3', '6', '1'),
('3', '{"age": "50", "bedno": "7A/318", "patientname": "Om Kumar", "admissiondate": "2018-08-13T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9068242932", "patientregistrationno": "RHC-2018-3462","bloodgroup":"AB-","weight":"93","drinst":"Arun Tripathi","gender":"Male"}', '{"allergies": "None", "treatmentdone": "traetment7", "reasonadmission": "reason7", "patientmedicalhistory": "None"}', '1', '3', '7', '1'),
('3', '{"age": "42", "bedno": "7A/319", "patientname": "Mayuri Jain", "admissiondate": "2018-08-14T18:30:00.000Z", "dischargedate": "2018-08-09T18:30:00.000Z", "emergencycontactno": "9053241932", "patientregistrationno": "RHC-2018-3463","bloodgroup":"AB-","weight":"55","drinst":"Arun Tripathi","gender":"Female"}', '{"allergies": "None", "treatmentdone": "traetment7", "reasonadmission": "reason7", "patientmedicalhistory": "None"}', '1', '3', '7', '1');


--
-- Dumping data for table `spl_node_field_operator_tbl`
--

INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Rohini Thakre','operator1@cust1.com','3', '1111', '1222', '1', '1');
INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Pooja Dessai','operator2@cust1.com','3', '2222', '222', '1', '2');


--
-- Dumping data for table `spl_node_fop_sp_tbl`
--

INSERT INTO `spl_node_fop_sp_tbl` (`fop_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('2', '3', '3');


--
-- Dumping data for table `spl_node_service_in_txn_tbl`
--


INSERT INTO `spl_node_service_in_txn_tbl` (`cpm_id_fk`, `serv_in_id_fk`, `fopcode`, `status`, `txn_data`, `txn_date`, `created_on`, `updated_on`) VALUES
	(3, 1, '1111', 1, '{"value": 140, "comment": "High Blood Pressure", "taskname": "Monitor Blood Pressure", "slotendtime": 630, "slotstarttime": 600}', '2018-08-13 19:13:19', '2018-08-07 16:13:19', '2018-08-07 16:13:19'),
	(3, 1, '2222', 1, '{"value": 140, "comment": "High Blood Pressure", "taskname": "Monitor Blood Pressure", "slotendtime": 630, "slotstarttime": 600}', '2018-08-14 16:13:39', '2018-08-07 16:13:19', '2018-08-07 16:13:19'),
	(3, 1, '1111', 1, '{"comment": "Saline", "taskname": "Saline 250ML", "slotendtime": 630, "slotstarttime": 600}', '2018-08-14 16:13:39', '2018-08-07 16:13:19', '2018-08-07 16:13:19'),
	(3, 1, '1111', 1, '{"value": 104, "comment": "High fever, 1 Crocin tablet provided", "taskname": "Monitor Temperature", "slotendtime": 730, "slotstarttime": 700}', '2018-08-14 16:13:39', '2018-08-07 16:13:19', '2018-08-07 16:13:19');


