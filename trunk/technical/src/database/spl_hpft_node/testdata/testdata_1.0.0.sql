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

INSERT INTO `spl_node_service_conf_tbl` (`cpm_id_fk`, `spc_id_fk`, `conf_type_code`, `serv_conf_name`, `serv_conf`) VALUES ('3', '3', 'SERVICE_DAILY_CHART', 'Patient File Template 1', '{"timeconf":{"starttime":480, "interval":30, "endtime": 1020}, "taskconf":{"tasks":[{"taskname": "Urinal"}, {"taskname": "Wash Basin"}, {"taskname": "Hand Wash"},{"taskname": "Tissue Roll"},{"taskname": "Sanitary Ware"},{"taskname": "Dustbins"},{"taskname": "Mopping"},{"taskname": "Air Freshner"},{"taskname": "Doors"}, {"taskname": "Wash Platform"}]}}');


--
-- Dumping data for table `spl_node_service_instance_tbl`
--

INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');
INSERT INTO `spl_node_service_instance_tbl` (`cpm_id_fk`, `serv_conf_id_fk`, `sp_id_fk`) VALUES ('3', '1', '3');

--
-- Dumping data for table `spl_hpft_patient_master_tbl`
--

INSERT INTO `spl_hpft_patient_master_tbl` (`id`, `cpm_id_fk`, `patient_details`, `medical_details`, `patient_file_template`, `sp_id_fk`, `serv_in_id_fk`, `status`) VALUES ('1', '3', '{"allergies": "dust", "treatmentdone": "tablets", "reasonadmission": "serious", "patientmedicalhistory": "no history"}', '{"age": "34", "ward": "ward 1", "bedno": "12", "patientname": "Patient1", "admissiondate": "2018-08-02T18:30:00.000Z", "dischargedate": "2018-08-08T18:30:00.000Z", "emergencycontactno": "2334534435", "patientregistrationno": "2334124324"}', '1', '3', '1', '1');

INSERT INTO `spl_hpft_patient_master_tbl` (`id`, `cpm_id_fk`, `patient_details`, `medical_details`, `patient_file_template`, `sp_id_fk`, `serv_in_id_fk`, `status`) VALUES ('2', '3', '{"allergies": "dust", "treatmentdone": "tablets", "reasonadmission": "serious", "patientmedicalhistory": "no history"}', '{"age": "34", "ward": "ward 1", "bedno": "12", "patientname": "Patient1", "admissiondate": "2018-08-02T18:30:00.000Z", "dischargedate": "2018-08-08T18:30:00.000Z", "emergencycontactno": "2334534435", "patientregistrationno": "2334124324"}', '1', '3', '2', '1');




