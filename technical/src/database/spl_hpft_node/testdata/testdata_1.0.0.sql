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
INSERT INTO `spl_node_sp_category_tbl` (`id`, `cpm_id_fk` ,`spc_name`) VALUES (4,3, 'Emergency Ward');


--
-- Dumping data for table `spl_node_sp_tbl`
--

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `uuid`,`cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('3', 'SP001','3', '1', 'General Ward', '1', UTC_TIMESTAMP);
INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `uuid`,`cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('7', 'SP002','3', '1', 'Emergency', '1', UTC_TIMESTAMP);
INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `uuid`,`cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('8', 'SP003','3', '1', 'Maternity', '1', UTC_TIMESTAMP);


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
('PM002', '3', 'P12B12214', 'Sumeet', 'Karande', '9982xxxxxx', '24', 'O+', '1', 1),
('PM003', '3', 'P12B12215', 'Sarjerao', 'Ghadage', '9832xxxxxx', '34', 'A+', '1', 1),
('PM004', '3', 'P12B12216', 'Tejal', 'Deshmukh', '9212xxxxxx', '27', 'A+', '2', 1),
('PM005', '3', 'P12B12217', 'Sanjay', 'Sawant', '9644xxxxxx', '33', 'A+', '1', 1),
('PM006', '3', 'P12B12218', 'Mandar', 'Bhong', '9522xxxxxx', '25', 'AB-', '1', 1),
('PM007', '3', 'P12B12219', 'Chandan', 'Pal', '9012xxxxxx', '38', 'O-', '1', 1),
('PM008', '3', 'P12B12220', 'Praveen', 'Pandey', '9442xxxxxx', '29', 'B+', '1', 1),
('PM009', '3', 'P12B12221', 'Shashank', 'Atre', '9642xxxxxx', '21', 'O+', '1', 1),
('PM010', '3', 'P12B12222', 'Mayuri', 'Jain', '9412xxxxxx', '25', 'AB-', '2', 1),
('PM011', '3', 'P12B12223', 'Shahuraj', 'Patil', '9572xxxxxx', '21', 'O+', '1', 1),
('PM012', '3', 'P12B12224', 'Abhijeet', 'Kalbhor', '9042xxxxxx', '24', 'O+', '1', 1);


--
-- Dumping data for table `spl_hpft_patient_admission_tbl`
--

INSERT INTO `spl_hpft_patient_admission_tbl` (`uuid`, `cpm_id_fk`, `patient_id_fk`, `patient_reg_no`, `bed_no`, `status`, `sp_id_fk`, `dr_incharge`, `admitted_on`, `updated_by`) 
VALUES 
('PA001', '3', 1, 'P12B12213', '3A/312', '1', '7', '8', '2018-12-04 14:37:53', 2),
('PA002', '3', 2, 'P12B12214', '3B/323', '1', '7', '8', '2018-12-04 12:47:53', 2),
('PA003', '3', 3, 'P12B12215', '2A/643', '1', '3', '8', '2018-12-04 09:17:53', 2),
('PA004', '3', 4, 'P12B12216', '4A/415', '1', '3', '8', '2018-12-04 11:00:53', 2),
('PA005', '3', 5, 'P12B12217', '5A/616', '1', '3', '8', '2018-12-04 01:11:53', 2),
('PA006', '3', 6, 'P12B12218', '6A/317', '1', '3', '8', '2018-12-04 14:32:53', 2),
('PA007', '3', 7, 'P12B12219', '7A/312', '1', '3', '8', '2018-12-04 16:44:53', 2);


--
-- Dumping data for table `spl_hpft_patient_personal_details_tbl`
--

INSERT INTO `spl_hpft_patient_personal_details_tbl` (`cpm_id_fk`, `patient_id`, `admission_id_fk`, `uuid`, `age`, `person_accompanying`, `updated_by`) 
VALUES 
('3', '1', '1', 'PPD001', '22', '{"data": [{"name": "Ashish", "gender": 1, "contact": "9843xxxxxx", "age": "23", "address": "warje", "alternatecontact": "9923xxxxxx", "relationshipwithpatient": "cousin"}],"version":1}', '2'),
('3', '2', '2', 'PPD002', '24', '{"data": [{"name": "Sanket", "gender": 1, "contact": "9712xxxxxx", "age": "26", "address": "warje", "alternatecontact": "8123xxxxxx", "relationshipwithpatient": "brother"}],"version":1}', '2'),
('3', '3', '3', 'PPD003', '34', '{"data": [{"name": "Ashish", "gender": 1, "contact": "9932xxxxxx", "age": "32", "address": "warje", "alternatecontact": "8993xxxxxx", "relationshipwithpatient": "friend"}],"version":1}', '2'),
('3', '4', '4', 'PPD004', '27', '{"data": [{"name": "Rohit", "gender": 2, "contact": "8812xxxxxx", "age": "53", "address": "warje", "alternatecontact": "9901xxxxxx", "relationshipwithpatient": "father"}],"version":1}', '2'),
('3', '5', '5', 'PPD005', '33', '{"data": [{"name": "Priya", "gender": 2, "contact": "9453xxxxxx", "age": "28", "address": "warje", "alternatecontact": "9456xxxxxx", "relationshipwithpatient": "wife"}],"version":1}', '2'),
('3', '6', '6', 'PPD006', '25', '{"data": [{"name": "Sonal", "gender": 2, "contact": "8663xxxxxx", "age": "30", "address": "warje", "alternatecontact": "8723xxxxxx", "relationshipwithpatient": "sister"}],"version":1}', '2'),
('3', '7', '7', 'PPD007', '38', '{"data": [{"name": "Kaushik", "gender": 1, "contact": "9813xxxxxx", "age": "23", "address": "warje", "alternatecontact": "8884xxxxxx", "relationshipwithpatient": "cousin"}],"version":1}', '2');


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
 ('PMD002', '3', '2', '2', '{"data": [{"date": "2019-03-22T13:11:05.383Z", "text": "This is test complaint"}, {"date": "2019-03-22T13:11:11.541Z", "text": "This is test complaint 2"}]}', '{"data": [{"date": "2019-03-22T13:12:10.448Z", "text": "Fever and back pain"}, {"date": "2019-03-22T13:12:31.052Z", "text": "Headache"}]}', '{"data": [{"date": "2019-03-22T13:12:58.582Z", "text": "Headache from last few months"}]}', '{"data": [{"date": "2019-03-22T13:13:31.222Z", "text": "2 years back headache started"}]}', '{"data": [{"date": "2019-03-22T13:13:53.912Z", "text": "Not done any treatment before addmission"}]}', '{"data": [{"date": "2019-03-22T13:14:03.581Z", "text": "No investigation"}]}', '{"data": [{"date": "2019-03-22T13:14:18.882Z", "text": "Grand father was suffering from headache"}]}', '{"data": [{"date": "2019-03-22T13:14:35.110Z", "text": "Having dust allergy"}]}', '{"data": {"other": "", "weight": {"weight": "52", "weighttendency": "Increasing"}, "alcohol": {"aplicable": false, "alcoholcomment": "No"}, "smoking": {"aplicable": false, "smokingcomment": "No"}}, "version": 1}', '2'),
 ('PMD003', '3', '3', '3', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD004', '3', '4', '4', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD005', '3', '5', '5', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD006', '3', '6', '6', null, null, null, null, null, null, null, null, null, '2'),
 ('PMD007', '3', '7', '7', null, null, null, null, null, null, null, null, null, '2');
 
 
--
-- Dumping data for table `spl_hpft_patient_conf_tbl`
--

INSERT INTO `spl_hpft_patient_conf_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `conf_type_code`, `conf`, `start_date`, `end_date`, `status`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, '1a1e4f86-6764-4a05-868e-a47648ec04e6', 3, 1, 'Monitor', '{"desc": " 4 times a day after every 3 hours for 4 days.", "name": "Temperature", "remark": null, "duration": "4", "interval": 180, "frequency": 0, "startTime": 620, "numberofTimes": "4"}', '2019-04-05 04:50:53', '2019-04-08 13:50:00', 0, '2019-04-05 04:51:24.893', '2019-04-05 04:52:06.463', '2019-04-05 04:52:06.463', 2),
	(2, 'b1790afa-5c72-45ee-8dd5-b5a85c2734db', 3, 1, 'Medicine', '{"desc": "Every morning & night after meal for 4 days", "name": "Crocin", "remark": null, "duration": "4", "foodInst": 1, "frequency": 0, "medicinetype": "Tablet", "mornFreqInfo": {"freqMorn": true, "mornFreqQuantity": "1"}, "aftrnFreqInfo": {"freqAftrn": false, "aftrnFreqQuantity": 0}, "nightFreqInfo": {"freqNight": true, "nightFreqQuantity": "1"}}', '2019-04-05 04:52:16', '2019-04-09 04:00:00', 0, '2019-04-05 04:52:39.714', '2019-04-05 04:52:51.603', '2019-04-05 04:52:51.603', 2),
	(3, 'd965560c-1a72-4206-a023-ff0dcf7941d4', 3, 1, 'Intake', '{"desc": " 3 times a day after every 4 hours for 3 days.", "name": "Saline", "remark": null, "duration": "3", "interval": 240, "quantity": "1", "frequency": 0, "startTime": 624, "intakeType": "Oral", "numberofTimes": "3", "specificTimes": []}', '2019-04-05 04:54:27', '2019-04-07 12:54:00', 0, '2019-04-05 04:55:08.062', '2019-04-05 04:55:17.345', '2019-04-05 04:55:17.345', 2),
	(4, 'e2b8d1a2-1c1c-4f9a-9f78-05926dce6faa', 3, 1, 'Output', '{"desc": "Check Urine for 2 days", "name": "Urine", "remark": null, "duration": "2"}', '2019-04-05 07:15:42', '2019-04-07 07:16:25', 0, '2019-04-05 07:15:46.947', '2019-04-05 07:16:30.683', '2019-04-05 07:16:30.683', 2);
	
	
--
-- Dumping data for table `spl_hpft_action_tbl`
--

INSERT INTO `spl_hpft_action_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `patient_conf_id_fk`, `conf_type_code`, `scheduled_time`, `is_deleted`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, 'a86a0fef-0222-456e-8918-9d8310a341bf', 3, 1, 1, 'Monitor', '2019-04-05 04:50:00', 0, '2019-04-05 04:52:03.897', '2019-04-05 04:52:07.915', '2019-04-05 04:52:07.915', 2),
	(2, '1a0230c9-d6d5-466d-9df1-73ca0639f319', 3, 1, 1, 'Monitor', '2019-04-05 07:50:00', 0, '2019-04-05 04:52:03.900', '2019-04-05 04:52:08.060', '2019-04-05 04:52:08.060', 2),
	(3, 'f6c87370-5e79-4f17-aba7-eac266ad584e', 3, 1, 1, 'Monitor', '2019-04-05 10:50:00', 0, '2019-04-05 04:52:03.901', '2019-04-05 04:52:08.128', '2019-04-05 04:52:08.128', 2),
	(4, 'ba31b8ba-628b-47f2-bae5-656fdda05d70', 3, 1, 1, 'Monitor', '2019-04-05 13:50:00', 0, '2019-04-05 04:52:03.902', '2019-04-05 04:52:08.263', '2019-04-05 04:52:08.263', 2),
	(5, 'b69301f5-efd4-4682-9327-c93a886787cf', 3, 1, 1, 'Monitor', '2019-04-06 04:50:00', 0, '2019-04-05 04:52:03.903', '2019-04-05 04:52:08.311', '2019-04-05 04:52:08.311', 2),
	(6, 'a15df74a-4c3f-4d07-93d5-99587f762acb', 3, 1, 1, 'Monitor', '2019-04-06 07:50:00', 0, '2019-04-05 04:52:03.904', '2019-04-05 04:52:08.340', '2019-04-05 04:52:08.340', 2),
	(7, '8a2ce59c-cad7-4efa-9983-aed7e720a54c', 3, 1, 1, 'Monitor', '2019-04-06 10:50:00', 0, '2019-04-05 04:52:03.905', '2019-04-05 04:52:08.377', '2019-04-05 04:52:08.377', 2),
	(8, 'b8e20710-d3ba-43f3-85af-4fdc84856448', 3, 1, 1, 'Monitor', '2019-04-06 13:50:00', 0, '2019-04-05 04:52:03.906', '2019-04-05 04:52:08.409', '2019-04-05 04:52:08.409', 2),
	(9, '607072d1-9293-4759-b19b-85b15e607b7f', 3, 1, 1, 'Monitor', '2019-04-07 04:50:00', 0, '2019-04-05 04:52:03.907', '2019-04-05 04:52:08.456', '2019-04-05 04:52:08.456', 2),
	(10, '4fcc9d9a-3188-4a25-a0da-6abe4948b962', 3, 1, 1, 'Monitor', '2019-04-07 07:50:00', 0, '2019-04-05 04:52:03.908', '2019-04-05 04:52:08.488', '2019-04-05 04:52:08.488', 2),
	(11, 'b2ac01ed-8397-4025-8fc1-7dec27ec4994', 3, 1, 1, 'Monitor', '2019-04-07 10:50:00', 0, '2019-04-05 04:52:03.909', '2019-04-05 04:52:08.523', '2019-04-05 04:52:08.523', 2),
	(12, '7b189743-acba-4d79-b4e6-2f7f5c32c40b', 3, 1, 1, 'Monitor', '2019-04-07 13:50:00', 0, '2019-04-05 04:52:03.910', '2019-04-05 04:52:08.554', '2019-04-05 04:52:08.554', 2),
	(13, '6708cfb8-5863-48ea-b43e-27377cc8b175', 3, 1, 1, 'Monitor', '2019-04-08 04:50:00', 0, '2019-04-05 04:52:03.911', '2019-04-05 04:52:08.589', '2019-04-05 04:52:08.589', 2),
	(14, '5f12ca37-fd9d-42fa-b883-a16a2fb62bcd', 3, 1, 1, 'Monitor', '2019-04-08 07:50:00', 0, '2019-04-05 04:52:03.912', '2019-04-05 04:52:08.636', '2019-04-05 04:52:08.636', 2),
	(15, '6c1165fa-74c1-4490-973a-bdcee0bcf872', 3, 1, 1, 'Monitor', '2019-04-08 10:50:00', 0, '2019-04-05 04:52:03.913', '2019-04-05 04:52:08.676', '2019-04-05 04:52:08.676', 2),
	(16, '6355fbeb-c2ee-4f50-8f26-67ea4373580b', 3, 1, 1, 'Monitor', '2019-04-08 13:50:00', 0, '2019-04-05 04:52:03.914', '2019-04-05 04:52:08.707', '2019-04-05 04:52:08.707', 2),
	(17, 'eef668e1-466a-4b3a-92fe-2df36c1c8ac4', 3, 1, 2, 'Medicine', '2019-04-05 15:00:00', 0, '2019-04-05 04:52:49.376', '2019-04-05 04:52:51.847', '2019-04-05 04:52:51.847', 2),
	(18, '0f171cc1-69df-4f07-95f4-33d28c731bd4', 3, 1, 2, 'Medicine', '2019-04-06 04:00:00', 0, '2019-04-05 04:52:49.378', '2019-04-05 04:52:51.903', '2019-04-05 04:52:51.903', 2),
	(19, 'ebcb42af-5fcd-4de2-8de0-869b0e621eed', 3, 1, 2, 'Medicine', '2019-04-06 15:00:00', 0, '2019-04-05 04:52:49.379', '2019-04-05 04:52:51.948', '2019-04-05 04:52:51.948', 2),
	(20, '0283c9ce-c13d-4232-9a5a-f16529247814', 3, 1, 2, 'Medicine', '2019-04-07 04:00:00', 0, '2019-04-05 04:52:49.380', '2019-04-05 04:52:51.995', '2019-04-05 04:52:51.995', 2),
	(21, '81535076-7ed6-43d6-87e1-a2ab9b9ed1a3', 3, 1, 2, 'Medicine', '2019-04-07 15:00:00', 0, '2019-04-05 04:52:49.381', '2019-04-05 04:52:52.115', '2019-04-05 04:52:52.115', 2),
	(22, '26362d5f-7d54-4312-b07e-7b31b520dd35', 3, 1, 2, 'Medicine', '2019-04-08 04:00:00', 0, '2019-04-05 04:52:49.382', '2019-04-05 04:52:52.168', '2019-04-05 04:52:52.168', 2),
	(23, '7b986282-6283-410c-82b9-36f55c142a0f', 3, 1, 2, 'Medicine', '2019-04-08 15:00:00', 0, '2019-04-05 04:52:49.384', '2019-04-05 04:52:52.228', '2019-04-05 04:52:52.228', 2),
	(24, 'f0d9a4b5-04bb-4ed6-91a6-e23d31b23d18', 3, 1, 2, 'Medicine', '2019-04-09 04:00:00', 0, '2019-04-05 04:52:49.385', '2019-04-05 04:52:52.269', '2019-04-05 04:52:52.269', 2),
	(25, '0cc747d3-693f-4b11-ac78-c1b54c8140de', 3, 1, 3, 'Intake', '2019-04-05 04:54:00', 0, '2019-04-05 04:55:15.069', '2019-04-05 04:55:17.600', '2019-04-05 04:55:17.600', 2),
	(26, 'a77ee3e4-778e-468a-a9cc-61833d9caf55', 3, 1, 3, 'Intake', '2019-04-05 08:54:00', 0, '2019-04-05 04:55:15.070', '2019-04-05 04:55:17.629', '2019-04-05 04:55:17.629', 2),
	(27, '66a8b279-868a-400b-96c0-163cfaef3349', 3, 1, 3, 'Intake', '2019-04-05 12:54:00', 0, '2019-04-05 04:55:15.071', '2019-04-05 04:55:17.684', '2019-04-05 04:55:17.684', 2),
	(28, 'ae7d7b31-34ad-4607-a2af-5b75ed773975', 3, 1, 3, 'Intake', '2019-04-06 04:54:00', 0, '2019-04-05 04:55:15.072', '2019-04-05 04:55:17.715', '2019-04-05 04:55:17.715', 2),
	(29, '454ccee8-5f68-48d8-a6a3-bc47ec6ab0dc', 3, 1, 3, 'Intake', '2019-04-06 08:54:00', 0, '2019-04-05 04:55:15.074', '2019-04-05 04:55:17.761', '2019-04-05 04:55:17.761', 2),
	(30, '20e4450e-d5ea-480c-8e37-a98b65f87698', 3, 1, 3, 'Intake', '2019-04-06 12:54:00', 0, '2019-04-05 04:55:15.075', '2019-04-05 04:55:17.883', '2019-04-05 04:55:17.883', 2),
	(31, 'b7de1c2f-e520-4ace-aa4e-bee8504ed862', 3, 1, 3, 'Intake', '2019-04-07 04:54:00', 0, '2019-04-05 04:55:15.076', '2019-04-05 04:55:17.915', '2019-04-05 04:55:17.915', 2),
	(32, 'e6bb7fc0-c0d3-4882-92fb-4608989feaf4', 3, 1, 3, 'Intake', '2019-04-07 08:54:00', 0, '2019-04-05 04:55:15.077', '2019-04-05 04:55:17.950', '2019-04-05 04:55:17.950', 2),
	(33, '5438ba0b-b6e9-4b9a-84b7-93eccda86513', 3, 1, 3, 'Intake', '2019-04-07 12:54:00', 0, '2019-04-05 04:55:15.078', '2019-04-05 04:55:17.992', '2019-04-05 04:55:17.992', 2);
	
	
--
-- Dumping data for table `spl_hpft_action_txn_tbl`
--

INSERT INTO `spl_hpft_action_txn_tbl` (`id`, `uuid`, `cpm_id_fk`, `patient_conf_id_fk`, `admission_id_fk`, `txn_data`, `runtime_config_data`, `scheduled_time`, `txn_state`, `conf_type_code`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, '8972e3d6-1a06-4b9e-8732-87fc218fa0bd', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, '2019-04-05 04:54:00', 1, 'Intake', '2019-04-05 04:57:10.496', '2019-04-05 04:57:16.885', '2019-04-05 04:57:16.885', 2),
	(2, 'a23b740a-b84d-42c4-a791-3281ef296b33', 3, 1, 1, '{"value": "98", "comment": "normal temperature"}', NULL, '2019-04-05 04:50:00', 1, 'Monitor', '2019-04-05 04:58:04.767', '2019-04-05 04:58:11.166', '2019-04-05 04:58:11.166', 2);

	
--
-- Dumping data for table `spl_hpft_document_tbl`
--

INSERT INTO `spl_hpft_document_tbl` (`id`, `cpm_id_fk`, `uuid`, `name`, `doctype`, `store_name`, `location`, `location_type`, `persisted`, `updated_by`, `client_updated_at`, `created_on`, `updated_on`) VALUES
	(1, 3, '7baefe06-597a-4d0a-934f-a3fcce54494e', 'NSIMG_20190109_14959.jpg', 'image/jpeg', 'doctors_orders_tbl', '/resources/documents/3/7baefe06-597a-4d0a-934f-a3fcce54494e', 1, 1, 2, '2019-04-05 05:01:22.881', '2019-04-05 05:01:31.388', '2019-04-05 05:01:31.388'),
	(2, 3, 'DB3D7B0E50AC47EBB0AF8A680340B58B45', 'patient.pdf', 'application/pdf', NULL, '/resources/documents/3/DB3D7B0E50AC47EBB0AF8A680340B58B45', 1, 1, 0, NULL, '2019-04-05 05:03:08.455', '2019-04-05 05:03:08.455'),
	(3, 3, 'DBB78F762C1144505002441318C93BCF5E', '2018-07-31 11_56_08-Window.png', 'image/png', NULL, '/resources/documents/3/DBB78F762C1144505002441318C93BCF5E', 1, 1, 0, NULL, '2019-04-05 05:03:54.926', '2019-04-05 05:03:54.926'),
	(4, 3, 'DB595C45752507DEDE67F12A93744704FA', '2018-07-31 11_56_08-Window.png', 'image/png', NULL, '/resources/documents/3/DB595C45752507DEDE67F12A93744704FA', 1, 1, 0, NULL, '2019-04-05 05:04:20.460', '2019-04-05 05:04:20.460');
	
--
-- Dumping data for table `spl_hpft_doctors_orders_tbl`
--

INSERT INTO `spl_hpft_doctors_orders_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `doctor_id_fk`, `doctors_orders`, `comment`, `ack_by`, `ack_time`, `status`, `order_created_time`, `order_type`, `document_id_fk`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, '8d9e75dd-6387-4c77-9f01-9c52619d6acf', 3, 1, 9, 'Aspirin', 'Incase of headache', NULL, NULL, 0, '2019-04-05 05:00:08', 'Prescription', NULL, '2019-04-05 05:00:08.303', '2019-04-05 05:00:14.568', '2019-04-05 05:00:14.568', 2),
	(2, '0cf2430c-7b05-49a8-b96b-86fed8fc104f', 3, 1, 9, 'Diet', 'Follow this diet', NULL, NULL, 0, '2019-04-05 05:01:23', 'General', 1, '2019-04-05 05:01:22.876', '2019-04-05 05:01:32.634', '2019-04-05 05:01:32.634', 2);
	

--
-- Dumping data for table `spl_hpft_pathology_record_tbl`
--

INSERT INTO `spl_hpft_pathology_record_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `test_performed`, `test_performed_time`, `test_result`, `comments`, `updated_by`, `client_updated_at`, `created_on`, `updated_on`) VALUES
	(1, 'DBCBFE021839FC638DF4727F77FE700BFB', 3, 1, 'test1', '2019-04-04 18:30:00', 'result1', 'comment1', 2, NULL, '2019-04-05 05:03:10.035', '2019-04-05 05:03:10.035'),
	(2, 'DB2B143D6B5E644F399D548D86611C7856', 3, 1, 'test2', '2019-04-04 18:30:00', 'result2', 'comment2', 2, NULL, '2019-04-05 05:03:56.411', '2019-04-05 05:03:56.411');
	
	
--
-- Dumping data for table `spl_hpft_pathology_record_doc_tbl`
--

INSERT INTO `spl_hpft_pathology_record_doc_tbl` (`pathology_id_fk`, `document_id_fk`, `created_on`, `updated_on`) VALUES
	(1, 2, '2019-04-05 05:03:10.057', '2019-04-05 05:03:10.057'),
	(2, 3, '2019-04-05 05:03:56.423', '2019-04-05 05:03:56.423');
	
--
-- Dumping data for table `spl_hpft_treatment_tbl`
--

INSERT INTO `spl_hpft_treatment_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `treatment_done`, `treatment_performed_time`, `details`, `post_observation`, `updated_by`, `client_updated_at`, `created_on`, `updated_on`) VALUES
	(1, 'DB856147B493D5F5476C19CD77419E634D', 3, 1, 'treatment1', '2019-04-04 18:30:00', 'details1', 'observation1', 2, NULL, '2019-04-05 05:04:21.229', '2019-04-05 05:04:21.229');
	
	
--
-- Dumping data for table `spl_hpft_treatment_doc_tbl`
--

INSERT INTO `spl_hpft_treatment_doc_tbl` (`treatment_id_fk`, `document_id_fk`, `created_on`, `updated_on`) VALUES
	(1, 4, '2019-04-05 05:04:21.231', '2019-04-05 05:04:21.231');
 
 
 
 
 