--
-- Dumping data for table `spl_node_dev_tbl`
--

INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('1', '3','device 1','1234567890123456');
INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('2', '3','device 2','1345494544733456');
INSERT INTO `spl_node_dev_tbl` (`dev_id_fk`, `cpm_id_fk`,`dev_name`, `serialno`) VALUES ('3', '3','device 3','1155623421323222');


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

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `uuid`,`cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('3', 'SP001','3', '1', 'General Ward 1', '1', UTC_TIMESTAMP);
INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `uuid`,`cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('7', 'SP002','3', '1', 'General Ward 2', '1', UTC_TIMESTAMP);
INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `uuid`,`cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('8', 'SP003','3', '1', 'General Ward 3', '1', UTC_TIMESTAMP);


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
-- Dumping data for table `spl_node_field_operator_tbl`
--

INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Rohini Thakre','rohini.thakre@noblehospital.com','3', '1111', '9911223344', '1', '1');
INSERT INTO `spl_node_field_operator_tbl` (`fop_name`,`email_id`,`cpm_id_fk`, `fopcode`, `mobile_no`, `fop_state`, `fop_area`) VALUES ('Pooja Dessai','pooja.dessai@noblehospital.com','3', '2222', '9811223344', '1', '2');


--
-- Dumping data for table `spl_node_fop_sp_tbl`
--

INSERT INTO `spl_node_fop_sp_tbl` (`fop_id_fk`, `sp_id_fk`, `cpm_id_fk`) VALUES ('2', '3', '3');


--
-- Dumping data for table `spl_hpft_patient_master_tbl`
--

INSERT INTO `spl_hpft_patient_master_tbl` (`uuid`, `cpm_id_fk`, `patient_reg_no`, `fname`, `lname`, `mob_no`, `age`, `blood_grp`, `gender`, `updated_by`) 
VALUES 
('PM001', '3', 'P12B12213', 'Amol', 'Patil', '9812xxxxxx', '22', 'AB+', '1', 1),
('PM002', '3', 'P12B12214', 'Sagar', 'Patil', '9982xxxxxx', '24', 'O+', '1', 1),
('PM003', '3', 'P12B12215', 'Shubham', 'Lunia', '9832xxxxxx', '34', 'A+', '1', 1),
('PM004', '3', 'P12B12216', 'Mayuri', 'Jain', '9212xxxxxx', '27', 'A+', '2', 1),
('PM005', '3', 'P12B12217', 'Sanjay', 'Sawant', '9644xxxxxx', '33', 'A+', '1', 1),
('PM006', '3', 'P12B12218', 'Pooja', 'Lokare', '9522xxxxxx', '25', 'AB-', '2', 1),
('PM007', '3', 'P12B12219', 'Mandar', 'Bhong', '9012xxxxxx', '38', 'O-', '1', 1),
('PM008', '3', 'P12B12220', 'Praveen', 'Pandey', '9442xxxxxx', '29', 'B+', '1', 1),
('PM009', '3', 'P12B12221', 'Shashank', 'Atre', '9642xxxxxx', '21', 'O+', '1', 1),
('PM010', '3', 'P12B12222', 'Tejal', 'Deshmukh', '9412xxxxxx', '25', 'AB-', '2', 1),
('PM011', '3', 'P12B12223', 'Shahuraj', 'Patil', '9572xxxxxx', '21', 'O+', '1', 1),
('PM012', '3', 'P12B12224', 'Abhijeet', 'Kalbhor', '9042xxxxxx', '24', 'O+', '1', 1);


--
-- Dumping data for table `spl_hpft_patient_admission_tbl`
--

INSERT INTO `spl_hpft_patient_admission_tbl` (`uuid`, `cpm_id_fk`, `patient_id_fk`, `patient_reg_no`, `bed_no`, `status`, `sp_id_fk`, `dr_incharge`, `admitted_on`, `updated_by`) 
VALUES 
('PA001', '3', 1, 'P12B12213', '3A/312', '1', '3', '8', '2018-12-04 14:37:53', 2),
('PA002', '3', 2, 'P12B12214', '3B/323', '1', '3', '8', '2018-12-04 12:47:53', 2),
('PA003', '3', 3, 'P12B12215', '2A/643', '1', '3', '8', '2018-12-04 09:17:53', 2),
('PA004', '3', 4, 'P12B12216', '4A/415', '2', '3', '8', '2018-12-04 11:00:53', 2),
('PA005', '3', 5, 'P12B12217', '5A/616', '2', '3', '8', '2018-12-04 01:11:53', 2),
('PA006', '3', 6, 'P12B12218', '6A/317', '1', '3', '8', '2018-12-04 14:32:53', 2),
('PA007', '3', 7, 'P12B12219', '7A/312', '3', '7', '8', '2018-12-04 16:44:53', 2);


--
-- Dumping data for table `spl_hpft_patient_personal_details_tbl`
--

INSERT INTO `spl_hpft_patient_personal_details_tbl` (`cpm_id_fk`, `patient_id`, `admission_id_fk`, `uuid`, `age`, `person_accompanying`, `updated_by`) 
VALUES 
('3', '1', '1', 'PPD001', '22', '{"name": "Ashish", "gender": 1, "contact": "9843xxxxxx", "pesrsonage": "23", "personaddress": "warje", "alternatecontact": "9923xxxxxx", "relationshipwithpatient": "cousin"}', '2'),
('3', '2', '2', 'PPD002', '22', '{"name": "Sanket", "gender": 1, "contact": "9712xxxxxx", "pesrsonage": "23", "personaddress": "warje", "alternatecontact": "8123xxxxxx", "relationshipwithpatient": "cousin"}', '2'),
('3', '3', '3', 'PPD003', '22', '{"name": "Ashish", "gender": 1, "contact": "9932xxxxxx", "pesrsonage": "23", "personaddress": "warje", "alternatecontact": "8993xxxxxx", "relationshipwithpatient": "cousin"}', '2'),
('3', '4', '4', 'PPD004', '22', '{"name": "Rohit", "gender": 2, "contact": "8812xxxxxx", "pesrsonage": "23", "personaddress": "warje", "alternatecontact": "9901xxxxxx", "relationshipwithpatient": "cousin"}', '2'),
('3', '5', '5', 'PPD005', '22', '{"name": "Priya", "gender": 2, "contact": "9453xxxxxx", "pesrsonage": "23", "personaddress": "warje", "alternatecontact": "9456xxxxxx", "relationshipwithpatient": "cousin"}', '2'),
('3', '6', '6', 'PPD006', '22', '{"name": "Sonal", "gender": 2, "contact": "8663xxxxxx", "pesrsonage": "23", "personaddress": "warje", "alternatecontact": "8723xxxxxx", "relationshipwithpatient": "cousin"}', '2'),
('3', '7', '7', 'PPD007', '22', '{"name": "Kaushik", "gender": 1, "contact": "9813xxxxxx", "pesrsonage": "23", "personaddress": "warje", "alternatecontact": "8884xxxxxx", "relationshipwithpatient": "cousin"}', '2');


--
-- Dumping data for table `spl_hpft_patient_medical_details_tbl`
--

 INSERT INTO `spl_hpft_patient_medical_details_tbl` (`uuid`, `cpm_id_fk`, `patient_id`, `admission_id_fk`, `present_complaints`, `reason_for_admission`, `history_present_illness`, `past_history`, `treatment_before_admission`, `investigation_before_admission`, `family_history`, `allergies`, `personal_history`, `updated_by`) 
 VALUES 
  ('PMD001', '3', '1', '1', 
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "throat infection"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "throat infection"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "history of present illness"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "past history"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "treatment"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "investigation"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "family history"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "allergies"}]}',
 	'{"data":{"weight":{"weight":40,"weight_tendency":"Increasing"},"alcohol":{"applicable":true,"quantity":"30 ml","remarks":"drinks sometimes"},"smoking":{"applicable":false,"quantity":null,"remarks":null},"others":null}}',
  	'2'),
 ('PMD002', '3', '1', '1', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD003', '3', '1', '1', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD004', '3', '1', '1', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD005', '3', '1', '1', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD006', '3', '1', '1', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD007', '3', '1', '1', null, null, null, null, null, null, null, null, null, '2');





