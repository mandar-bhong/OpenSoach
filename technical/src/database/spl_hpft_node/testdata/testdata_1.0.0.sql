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

INSERT INTO `spl_node_dev_status_tbl` (`dev_id_fk`,`connection_state`,`connection_state_since`,`sync_state`,`sync_state_since`,`battery_level`,`battery_level_since`) VALUES 
('1', '0', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP),
('2', '0', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP),
('3', '0', UTC_TIMESTAMP, '1', UTC_TIMESTAMP, '0', UTC_TIMESTAMP);

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

INSERT INTO `spl_hpft_patient_admission_tbl` (`uuid`, `cpm_id_fk`, `patient_id_fk`, `patient_reg_no`, `bed_no`, `status`, `sp_id_fk`, `dr_incharge`, `admitted_on`, `updated_by`, `created_on`, `updated_on`) 
VALUES 
('PA001', '3', 1, 'P12B12213', '3A/312', '1', '7', '8', '2018-12-04 14:37:53', 2, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'01:50:53')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'01:50:53'))),
('PA002', '3', 2, 'P12B12214', '3B/323', '1', '7', '8', '2018-12-04 12:47:53', 2,UTC_TIMESTAMP,UTC_TIMESTAMP),
('PA003', '3', 3, 'P12B12215', '2A/643', '1', '3', '8', '2018-12-04 09:17:53', 2,UTC_TIMESTAMP,UTC_TIMESTAMP),
('PA004', '3', 4, 'P12B12216', '4A/415', '1', '3', '8', '2018-12-04 11:00:53', 2,UTC_TIMESTAMP,UTC_TIMESTAMP),
('PA005', '3', 5, 'P12B12217', '5A/616', '1', '3', '8', '2018-12-04 01:11:53', 2,UTC_TIMESTAMP,UTC_TIMESTAMP),
('PA006', '3', 6, 'P12B12218', '6A/317', '1', '3', '8', '2018-12-04 14:32:53', 2,UTC_TIMESTAMP,UTC_TIMESTAMP),
('PA007', '3', 7, 'P12B12219', '7A/312', '1', '3', '8', '2018-12-04 16:44:53', 2,UTC_TIMESTAMP,UTC_TIMESTAMP);


--
-- Dumping data for table `spl_hpft_patient_personal_details_tbl`
--

INSERT INTO `spl_hpft_patient_personal_details_tbl` (`cpm_id_fk`, `patient_id`, `admission_id_fk`, `uuid`, `age`, `person_accompanying`, `updated_by`, `created_on`, `updated_on`) 
VALUES 
('3', '1', '1', 'PPD001', '22', '{"data": [{"name": "Ashish", "gender": 1, "contact": "9843xxxxxx", "age": "23", "address": "warje", "alternatecontact": "9923xxxxxx", "relationshipwithpatient": "cousin"}],"version":1}', '2',timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'01:50:53')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'01:50:53'))),
('3', '2', '2', 'PPD002', '24', '{"data": [{"name": "Sanket", "gender": 1, "contact": "9712xxxxxx", "age": "26", "address": "warje", "alternatecontact": "8123xxxxxx", "relationshipwithpatient": "brother"}],"version":1}', '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
('3', '3', '3', 'PPD003', '34', '{"data": [{"name": "Ashish", "gender": 1, "contact": "9932xxxxxx", "age": "32", "address": "warje", "alternatecontact": "8993xxxxxx", "relationshipwithpatient": "friend"}],"version":1}', '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
('3', '4', '4', 'PPD004', '27', '{"data": [{"name": "Rohit", "gender": 2, "contact": "8812xxxxxx", "age": "53", "address": "warje", "alternatecontact": "9901xxxxxx", "relationshipwithpatient": "father"}],"version":1}', '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
('3', '5', '5', 'PPD005', '33', '{"data": [{"name": "Priya", "gender": 2, "contact": "9453xxxxxx", "age": "28", "address": "warje", "alternatecontact": "9456xxxxxx", "relationshipwithpatient": "wife"}],"version":1}', '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
('3', '6', '6', 'PPD006', '25', '{"data": [{"name": "Sonal", "gender": 2, "contact": "8663xxxxxx", "age": "30", "address": "warje", "alternatecontact": "8723xxxxxx", "relationshipwithpatient": "sister"}],"version":1}', '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
('3', '7', '7', 'PPD007', '38', '{"data": [{"name": "Kaushik", "gender": 1, "contact": "9813xxxxxx", "age": "23", "address": "warje", "alternatecontact": "8884xxxxxx", "relationshipwithpatient": "cousin"}],"version":1}', '2',UTC_TIMESTAMP,UTC_TIMESTAMP);


--
-- Dumping data for table `spl_hpft_patient_medical_details_tbl`
--

 INSERT INTO `spl_hpft_patient_medical_details_tbl` (`uuid`, `cpm_id_fk`, `patient_id`, `admission_id_fk`, `present_complaints`, `reason_for_admission`, `history_present_illness`, `past_history`, `treatment_before_admission`, `investigation_before_admission`, `family_history`, `allergies`, `personal_history`, `updated_by`, `created_on`, `updated_on`) 
 VALUES 
  ('PMD001', '3', '1', '1', 
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "Having severe throat infection"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "Reason for admission due to throat infection"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "Throat infection since last two months"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "No past history"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "Undergoing Alopathic treatment since one month"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "No investigation before admission"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "No family history"}]}',
  '{"data": [{"date": "2019-03-07T08:52:36.474Z", "text": "No allergies"}]}',
 	'{"data":{"weight":{"weight":40,"weight_tendency":"Increasing"},"alcohol":{"applicable":true,"quantity":"30 ml","remarks":"drinks sometimes"},"smoking":{"applicable":false,"quantity":null,"remarks":null},"others":null}}',
  	'2',timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'01:50:53')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'01:50:53'))),
 ('PMD002', '3', '2', '2', '{"data": [{"date": "2019-03-22T13:11:05.383Z", "text": "This is test complaint"}, {"date": "2019-03-22T13:11:11.541Z", "text": "This is test complaint 2"}]}', '{"data": [{"date": "2019-03-22T13:12:10.448Z", "text": "Fever and back pain"}, {"date": "2019-03-22T13:12:31.052Z", "text": "Headache"}]}', '{"data": [{"date": "2019-03-22T13:12:58.582Z", "text": "Headache from last few months"}]}', '{"data": [{"date": "2019-03-22T13:13:31.222Z", "text": "2 years back headache started"}]}', '{"data": [{"date": "2019-03-22T13:13:53.912Z", "text": "Not done any treatment before addmission"}]}', '{"data": [{"date": "2019-03-22T13:14:03.581Z", "text": "No investigation"}]}', '{"data": [{"date": "2019-03-22T13:14:18.882Z", "text": "Grand father was suffering from headache"}]}', '{"data": [{"date": "2019-03-22T13:14:35.110Z", "text": "Having dust allergy"}]}', '{"data": {"other": "", "weight": {"weight": "52", "weighttendency": "Increasing"}, "alcohol": {"aplicable": false, "alcoholcomment": "No"}, "smoking": {"aplicable": false, "smokingcomment": "No"}}, "version": 1}', '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
 ('PMD003', '3', '3', '3', null, null, null, null, null, null, null, null, null, '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
 ('PMD004', '3', '4', '4', null, null, null, null, null, null, null, null, null, '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
 ('PMD005', '3', '5', '5', null, null, null, null, null, null, null, null, null, '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
 ('PMD006', '3', '6', '6', null, null, null, null, null, null, null, null, null, '2',UTC_TIMESTAMP,UTC_TIMESTAMP),
 ('PMD007', '3', '7', '7', null, null, null, null, null, null, null, null, null, '2',UTC_TIMESTAMP,UTC_TIMESTAMP);
 
 
--
-- Dumping data for table `spl_hpft_patient_conf_tbl`
--

INSERT INTO `spl_hpft_patient_conf_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `conf_type_code`, `conf`, `start_date`, `end_date`, `status`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, '1a1e4f86-6764-4a05-868e-a47648ec04e6', 3, 1, 'Monitor', '{"desc": " 4 times a day after every 3 hours for 4 days.", "name": "Temperature", "remark": null, "duration": "4", "interval": 180, "frequency": 0, "startTime": 620, "numberofTimes": "4"}', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:53')), timestamp(ADDTIME(( curdate() + INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:51:24')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), 9),
	(2, 'b1790afa-5c72-45ee-8dd5-b5a85c2734db', 3, 1, 'Medicine', '{"desc": "Every morning & night after meal for 4 days", "name": "Crocin", "remark": null, "duration": "4", "foodInst": 1, "frequency": 0, "medicinetype": "Tablet", "mornFreqInfo": {"freqMorn": true, "mornFreqQuantity": "1"}, "aftrnFreqInfo": {"freqAftrn": false, "aftrnFreqQuantity": 0}, "nightFreqInfo": {"freqNight": true, "nightFreqQuantity": "1"}}', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:16')), timestamp(ADDTIME(( curdate() + INTERVAL 0 DAY),'04:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:39')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), 9),
	(3, 'd965560c-1a72-4206-a023-ff0dcf7941d4', 3, 1, 'Intake', '{"desc": " 3 times a day after every 4 hours for 3 days.", "name": "Saline", "remark": null, "duration": "3", "interval": 240, "quantity": "1", "frequency": 0, "startTime": 624, "intakeType": "Oral", "numberofTimes": "3", "specificTimes": []}', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:54:27')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'12:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	(4, 'e2b8d1a2-1c1c-4f9a-9f78-05926dce6faa', 3, 1, 'Output', '{"desc": "Check Urine for 2 days", "name": "Urine", "remark": null, "duration": "2"}', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:15:42')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:16:25')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:15:46')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:16:30')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:16:30')), 9),
	(5, '1a1e4f86-6764-4a05-868e-a47648ec04e5', 3, 1, 'Monitor', '{"desc": " 4 times a day after every 3 hours for 4 days.", "name": "Blood Pressure", "remark": null, "duration": "4", "interval": 180, "frequency": 0, "startTime": 620, "numberofTimes": "4"}', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:53')), timestamp(ADDTIME(( curdate() + INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:51:24')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), 9),
	(6, '1a1e4f86-6764-4a05-868e-a47648ec04e7', 3, 1, 'Monitor', '{"desc": " 4 times a day after every 3 hours for 4 days.", "name": "Pulse Rate", "remark": null, "duration": "4", "interval": 180, "frequency": 0, "startTime": 620, "numberofTimes": "4"}', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:53')), timestamp(ADDTIME(( curdate() + INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:51:24')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), 9),
	(7, '1a1e4f86-6764-4a05-868e-a47648ec04e8', 3, 1, 'Monitor', '{"desc": " 4 times a day after every 3 hours for 4 days.", "name": "Respiration Rate", "remark": null, "duration": "4", "interval": 180, "frequency": 0, "startTime": 620, "numberofTimes": "4"}', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:53')), timestamp(ADDTIME(( curdate() + INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:51:24')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:06')), 9);
	
	
--
-- Dumping data for table `spl_hpft_action_tbl`
--

INSERT INTO `spl_hpft_action_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `patient_conf_id_fk`, `conf_type_code`, `scheduled_time`, `is_deleted`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, 'a86a0fef-0222-456e-8918-9d8310a341bf', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), 2),	
	(2, '1a0230c9-d6d5-466d-9df1-73ca0639f319', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(3, 'f6c87370-5e79-4f17-aba7-eac266ad584e', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),	
	(4, 'ba31b8ba-628b-47f2-bae5-656fdda05d70', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	(5, 'b69301f5-efd4-4682-9327-c93a886787cf', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(6, 'a15df74a-4c3f-4d07-93d5-99587f762acb', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(7, '8a2ce59c-cad7-4efa-9983-aed7e720a54c', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(8, 'b8e20710-d3ba-43f3-85af-4fdc84855548', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	(9, '607072d1-9293-4759-b19b-85b15e607b7f', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(10, '4fcc9d9a-3188-4a25-a0da-6abe4948b962', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(11, 'b2ac01ed-8397-4025-8fc1-7dec27ec4994', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(12, '7b189743-acba-4d79-b4e6-2f7f5c32c40b', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	(13, '6708cfb8-5863-48ea-b43e-27377cc8b175', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(14, '5f12ca37-fd9d-42fa-b883-a16a2fb62bcd', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(15, '6c1165fa-74c1-4490-973a-bdcee0bcf872', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(16, '6355fbeb-c2ee-4f50-8f26-67ea4373580b', 3, 1, 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	(17, 'eef668e1-466a-4b3a-92fe-2df36c1c8ac4', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'15:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), 2),
	
	(18, '0f171cc1-69df-4f07-95f4-33d28c731bd4', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), 2),
	(19, 'ebcb42af-5fcd-4de2-8de0-869b0e621eed', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'15:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), 2),
	
	(20, '0283c9ce-c13d-4232-9a5a-f16529247814', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:51')), 2),
	(21, '81535076-7ed6-43d6-87e1-a2ab9b9ed1a3', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'15:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), 2),
	
	(22, '26362d5f-7d54-4312-b07e-7b31b520dd35', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'04:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), 2),
	(23, '7b986282-6283-410c-82b9-36f55c142a0f', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'15:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), 2),
	
	(24, 'f0d9a4b5-04bb-4ed6-91a6-e23d31b23d18', 3, 1, 2, 'Medicine', timestamp(ADDTIME(( curdate() + INTERVAL 1 DAY),'04:00:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:49')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:52')), 2),
	
	(25, '0cc747d3-693f-4b11-ac78-c1b54c8140de', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	(26, 'a77ee3e4-778e-468a-a9cc-61833d9caf55', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'08:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	(27, '66a8b279-868a-400b-96c0-163cfaef3349', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'12:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	
	(28, 'ae7d7b31-34ad-4607-a2af-5b75ed773975', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	(29, '454ccee8-5f68-48d8-a6a3-bc47ec6ab0dc', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'08:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	(30, '20e4450e-d5ea-480c-8e37-a98b65f87698', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'12:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	
	(31, 'b7de1c2f-e520-4ace-aa4e-bee8504ed862', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	(32, 'e6bb7fc0-c0d3-4882-92fb-4608989feaf4', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'08:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),
	(33, '5438ba0b-b6e9-4b9a-84b7-93eccda86513', 3, 1, 3, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'12:54:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:15')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:55:17')), 2),

	
	(34, 'a86a0fef-0222-456e-8918-9d8310a341b1', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), 2),	
	(35, '1a0230c9-d6d5-466d-9df1-73ca0639f312', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(36, 'f6c87370-5e79-4f17-aba7-eac266ad5843', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),	
	(37, 'ba31b8ba-628b-47f2-bae5-656fdda05d74', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),

	(38, 'b69301f5-efd4-4682-9327-c93a886787c5', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(39, 'a15df74a-4c3f-4d07-93d5-99587f762ac6', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(40, '8a2ce59c-cad7-4efa-9983-aed7e720a547', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(41, 'b8e20710-d3ba-43f3-85af-4fdc84856448', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	(42, '607072d1-9293-4759-b19b-85b15e607b10', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(43, '4fcc9d9a-3188-4a25-a0da-6abe4948b911', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(44, 'b2ac01ed-8397-4025-8fc1-7dec27ec4912', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(45, '7b189743-acba-4d79-b4e6-2f7f5c32c413', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),

	(46, '6708cfb8-5863-48ea-b43e-27377cc8b114', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(47, '5f12ca37-fd9d-42fa-b883-a16a2fb62b15', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(48, '6c1165fa-74c1-4490-973a-bdcee0bcf816', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(49, '6355fbeb-c2ee-4f50-8f26-67ea43735817', 3, 1, 5, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	
	(50, 'a86a0fef-0222-456e-8918-9d8310a34118', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), 2),	
	(51, '1a0230c9-d6d5-466d-9df1-73ca0639f322', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(52, 'f6c87370-5e79-4f17-aba7-eac266ad5823', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),	
	(53, 'ba31b8ba-628b-47f2-bae5-656fdda05d24', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),

	(54, 'b69301f5-efd4-4682-9327-c93a88678725', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(55, 'a15df74a-4c3f-4d07-93d5-99587f762a26', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(56, '8a2ce59c-cad7-4efa-9983-aed7e720a527', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(57, 'b8e20710-d3ba-43f3-85af-4fdc84856428', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	(58, '607072d1-9293-4759-b19b-85b15e607b29', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(59, '4fcc9d9a-3188-4a25-a0da-6abe4948b930', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(60, 'b2ac01ed-8397-4025-8fc1-7dec27ec4931', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(61, '7b189743-acba-4d79-b4e6-2f7f5c32c432', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),

	(62, '6708cfb8-5863-48ea-b43e-27377cc8b133', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(63, '5f12ca37-fd9d-42fa-b883-a16a2fb62b34', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(64, '6c1165fa-74c1-4490-973a-bdcee0bcf835', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(65, '6355fbeb-c2ee-4f50-8f26-67ea43735836', 3, 1, 6, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	
	(66, 'a86a0fef-0222-456e-8918-9d8310a3437', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:07')), 2),	
	(67, '1a0230c9-d6d5-466d-9df1-73ca0639f338', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(68, 'f6c87370-5e79-4f17-aba7-eac266ad5839', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),	
	(69, 'ba31b8ba-628b-47f2-bae5-656fdda05d40', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),

	(70, 'b69301f5-efd4-4682-9327-c93a88678741', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(71, 'a15df74a-4c3f-4d07-93d5-99587f762a42', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(72, '8a2ce59c-cad7-4efa-9983-aed7e720a543', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(73, 'b8e20710-d3ba-43f3-85af-4fdc84856444', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	
	(74, '607072d1-9293-4759-b19b-85b15e607b45', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(75, '4fcc9d9a-3188-4a25-a0da-6abe4948b946', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(76, 'b2ac01ed-8397-4025-8fc1-7dec27ec4947', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(77, '7b189743-acba-4d79-b4e6-2f7f5c32c448', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),

	(78, '6708cfb8-5863-48ea-b43e-27377cc8b149', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'04:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(79, '5f12ca37-fd9d-42fa-b883-a16a2fb62b50', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'07:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(80, '6c1165fa-74c1-4490-973a-bdcee0bcf851', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'10:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2),
	(81, '6355fbeb-c2ee-4f50-8f26-67ea43735852', 3, 1, 7, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 0 DAY),'13:50:00')), 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:03')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:52:08')), 2);
	
	
	
	
	
--
-- Dumping data for table `spl_hpft_action_txn_tbl`
--

INSERT INTO `spl_hpft_action_txn_tbl` (`id`, `uuid`, `cpm_id_fk`, `patient_conf_id_fk`, `admission_id_fk`, `txn_data`, `runtime_config_data`, `scheduled_time`, `txn_state`, `conf_type_code`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, 'a23b740a-b84d-42c4-a791-3281ef296b34', 3, 1, 1, '{"value": "99", "comment": "high temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), 11),
	(2, 'a23b740a-b84d-42c4-a791-3281ef296b35', 3, 1, 1, '{"value": "97", "comment": "normal temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), 11),
	(3, 'a23b740a-b84d-42c4-a791-3281ef296b36', 3, 1, 1, '{"value": "97.5", "comment": "normal temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 11),
	(4, 'a23b740a-b84d-42c4-a791-3281ef296b37', 3, 1, 1, '{"value": "96", "comment": "low temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), 11),
	
	(5, 'a23b740a-b84d-42c4-a791-3281ef296b38', 3, 1, 1, '{"value": "96", "comment": "normal temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), 11),
	(6, 'a23b740a-b84d-42c4-a791-3281ef296b39', 3, 1, 1, '{"value": "99", "comment": "high temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), 11),
	(7, 'a23b740a-b84d-42c4-a791-3281ef296b40', 3, 1, 1, '{"value": "98", "comment": "normal temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 11),
	(8, 'a23b740a-b84d-42c4-a791-3281ef296b41', 3, 1, 1, '{"value": "97", "comment": "normal temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), 11),
	
	(9, 'a23b740a-b84d-42c4-a791-3281ef296b42', 3, 1, 1, '{"value": "97", "comment": "normal temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), 11),
	(10, 'a23b740a-b84d-42c4-a791-3281ef296b343', 3, 1, 1, '{"value": "97.5", "comment": "normal temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), 11),
	(11, 'a23b740a-b84d-42c4-a791-3281ef296b44', 3, 1, 1, '{"value": "96", "comment": "low temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 11),
	(12, 'a23b740a-b84d-42c4-a791-3281ef296b45', 3, 1, 1, '{"value": "96.5", "comment": "low temperature"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), 11),
	
	(13, 'a23b740a-b84d-42c4-a791-3281ef296b50', 3, 2, 1, '{"value": null, "comment": "medicine given"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'15:00:00')), 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'15:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'15:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'15:00:00')), 11),
	
	(14, 'a23b740a-b84d-42c4-a791-3281ef296b51', 3, 2, 1, '{"value": null, "comment": "medicine given"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:00:00')), 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:00:00')), 11),
	(15, 'a23b740a-b84d-42c4-a791-3281ef296b52', 3, 2, 1, '{"value": null, "comment": "medicine given"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'15:00:00')), 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'15:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'15:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'15:00:00')), 11),
	
	(16, 'a23b740a-b84d-42c4-a791-3281ef296b53', 3, 2, 1, '{"value": null, "comment": "medicine given"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:00:00')), 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:00:00')), 11),
	(17, 'a23b740a-b84d-42c4-a791-3281ef296b54', 3, 2, 1, '{"value": null, "comment": "medicine given"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'15:00:00')), 2, 'Medicine', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'15:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'15:00:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'15:00:00')), 11),
	
	(18, '8972e3d6-1a06-4b9e-8732-87fc218fa157', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:57:10')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:57:16')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:57:16')), 11),
	(19, '8972e3d6-1a06-4b9e-8732-87fc218fa257', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'08:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'08:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'08:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'08:54:00')), 11),
	(20, '8972e3d6-1a06-4b9e-8732-87fc218fa357', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'12:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'12:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'12:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'12:54:00')), 11),
	
	(21, '8972e3d6-1a06-4b9e-8732-87fc218fa457', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:57:10')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:57:16')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:57:16')), 11),
	(22, '8972e3d6-1a06-4b9e-8732-87fc218fa058', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'08:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'08:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'08:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'08:54:00')), 11),
	(23, '8972e3d6-1a06-4b9e-8732-87fc218fa059', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'12:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'12:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'12:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'12:54:00')), 11),
	
	(24, '8972e3d6-1a06-4b9e-8732-87fc218fa060', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:57:10')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:57:16')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:57:16')), 11),
	(25, '8972e3d6-1a06-4b9e-8732-87fc218fa061', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'08:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'08:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'08:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'08:54:00')), 11),
	(26, '8972e3d6-1a06-4b9e-8732-87fc218fa062', 3, 3, 1, '{"comment": "saline given 100ml"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'12:54:00')), 1, 'Intake', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'12:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'12:54:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'12:54:00')), 11),


	(27, 'a23b740a-b84d-42c4-aaa1-3281ef296b01', 3, 5, 1, '{"value": {"systolic":"82","diastolic":"125"}, "comment": " normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), 11),
	(28, 'a23b740a-b84d-42c4-aa11-3281ef296b32', 3, 5, 1, '{"value": {"systolic":"85","diastolic":"123"}, "comment": "normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), 11),
	(29, 'a23b740a-b84d-42c4-a791-3281ef296b33', 3, 5, 1, '{"value": {"systolic":"83","diastolic":"121"}, "comment": "normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 11),
	(30, 'a23b740a-b84d-42c4-a791-3281ef296csw', 3, 5, 1, '{"value": {"systolic":"87","diastolic":"135"}, "comment": "high blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), 11),
                                                                      
	(31, 'a23b740a-b84d-42c4-a791-3281ef296111', 3, 5, 1, '{"value": {"systolic":"89","diastolic":"137"}, "comment": "high blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), 11),
	(32, 'a23b740a-b84d-42c4-a791-3281ef296112', 3, 5, 1, '{"value": {"systolic":"82","diastolic":"122"}, "comment": "normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), 11),
	(33, 'a23b740a-b84d-42c4-a791-3281ef296b11', 3, 5, 1, '{"value": {"systolic":"86","diastolic":"130"}, "comment": "high blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 11),
	(34, 'a23b740a-b84d-42c4-a791-3281ef296b12', 3, 5, 1, '{"value": {"systolic":"87","diastolic":"132"}, "comment": "high blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), 11),
                                                                     
	(35, 'a23b740a-b84d-42c4-a791-3281ef296b13', 3, 5, 1, '{"value": {"systolic":"85","diastolic":"124"}, "comment": "normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), 11),
	(36, 'a23b740a-b84d-42c4-a791-3281ef296b314', 3, 5, 1,'{"value": {"systolic":"82","diastolic":"123"}, "comment": "normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), 11),
	(37, 'a23b740a-b84d-42c4-a791-3281ef296b15', 3, 5, 1, '{"value": {"systolic":"81","diastolic":"120"}, "comment": "normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 11),
	(38, 'a23b740a-b84d-42c4-a791-3281ef296b16', 3, 5, 1, '{"value": {"systolic":"81","diastolic":"120"}, "comment": "normal blood pressure observed"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), 11),
	                                                                                                 
	
	(39, 'a23b740a-b84d-42c4-a791-3281ef29617', 3, 6, 1, '{"value": "80", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), 11),
	(40, 'a23b740a-b84d-42c4-a791-3281ef296b18', 3, 6, 1, '{"value": "82", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), 11),
	(41, 'a23b740a-b84d-42c4-a791-3281ef296b19', 3, 6, 1, '{"value": "84", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 11),
	(42, 'a23b740a-b84d-42c4-a791-3281ef296b20', 3, 6, 1, '{"value": "86", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), 11),
	
	(43, 'a23b740a-b84d-42c4-a791-3281ef296b21', 3, 6, 1, '{"value": "84", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), 11),
	(44, 'a23b740a-b84d-42c4-a791-3281ef296b22', 3, 6, 1, '{"value": "83", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), 11),
	(45, 'a23b740a-b84d-42c4-a791-3281ef296b23', 3, 6, 1, '{"value": "80", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 11),
	(46, 'a23b740a-b84d-42c4-a791-3281ef296b24', 3, 6, 1, '{"value": "78", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), 11),

	(47, 'a23b740a-b84d-42c4-a791-3281ef296b25', 3, 6, 1, '{"value": "82", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), 11),
	(48, 'a23b740a-b84d-42c4-a791-3281ef296b326', 3, 6, 1, '{"value": "70", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), 11),
	(49, 'a23b740a-b84d-42c4-a791-3281ef296b27', 3, 6, 1, '{"value": "65", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 11),
	(50, 'a23b740a-b84d-42c4-a791-3281ef296b28', 3, 6, 1, '{"value": "62", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), 11),
	
	(51, 'a23b740a-b84d-42c4-a791-3281ef296b29', 3, 7, 1, '{"value": "20", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'04:58:11')), 11),
	(52, 'a23b740a-b84d-42c4-a791-3281ef296119', 3, 7, 1, '{"value": "19", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'07:48:00')), 11),
	(53, 'a23b740a-b84d-42c4-a791-3281ef296b113', 3, 7, 1, '{"value": "17", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'10:50:00')), 11),
	(54, 'a23b740a-b84d-42c4-a791-3281ef296b123', 3, 7, 1, '{"value": "16", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'13:55:00')), 11),
	
	(55, 'a23b740a-b84d-42c4-a791-3281ef296133', 3, 7, 1, '{"value": "15", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'04:58:11')), 11),
	(56, 'a23b740a-b84d-42c4-a791-3281ef296134', 3, 7, 1, '{"value": "14", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'07:48:00')), 11),
	(57, 'a23b740a-b84d-42c4-a791-3281ef296135', 3, 7, 1, '{"value": "12", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'10:50:00')), 11),
	(58, 'a23b740a-b84d-42c4-a791-3281ef296136', 3, 7, 1, '{"value": "15", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 2 DAY),'13:55:00')), 11),

	(59, 'a23b740a-b84d-42c4-a791-3281ef296137', 3, 7, 1, '{"value": "16", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:04')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'04:58:11')), 11),
	(60, 'a23b740a-b84d-42c4-a791-3281ef296b138', 3, 7, 1, '{"value": "14", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'07:48:00')), 11),
	(61, 'a23b740a-b84d-42c4-a791-3281ef296b139', 3, 7, 1, '{"value": "15", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'10:50:00')), 11),
	(62, 'a23b740a-b84d-42c4-a791-3281ef296140', 3, 7, 1, '{"value": "16", "comment": "normal"}', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:50:00')), 1, 'Monitor', timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), timestamp(ADDTIME(( curdate() - INTERVAL 1 DAY),'13:55:00')), 11);
	
	

	
--
-- Dumping data for table `spl_hpft_document_tbl`
--

INSERT INTO `spl_hpft_document_tbl` (`id`, `cpm_id_fk`, `uuid`, `name`, `doctype`, `store_name`, `location`, `location_type`, `persisted`, `updated_by`, `client_updated_at`, `created_on`, `updated_on`) VALUES
	(1, 3, '7baefe06-597a-4d0a-934f-a3fcce54494e', 'DSC122312.jpg', 'image/jpeg', 'doctors_orders_tbl', '/resources/documents/3/7baefe06-597a-4d0a-934f-a3fcce54494e', 1, 1, 2, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:01:22')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:01:31')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:01:31'))),
	(2, 3, 'DB3D7B0E50AC47EBB0AF8A680340B58B45', 'DSC122313.pdf', 'application/pdf', NULL, '/resources/documents/3/DB3D7B0E50AC47EBB0AF8A680340B58B45', 1, 1, 0, NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:08'))),
	(3, 3, 'DBB78F762C1144505002441318C93BCF5E', 'DSC122314.png', 'image/png', NULL, '/resources/documents/3/DBB78F762C1144505002441318C93BCF5E', 1, 1, 0, NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:54')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:54'))),
	(4, 3, 'DB595C45752507DEDE67F12A93744704FA', 'DSC122315.png', 'image/png', NULL, '/resources/documents/3/DB595C45752507DEDE67F12A93744704FA', 1, 1, 0, NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:04:20')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:04:20')));
	
--
-- Dumping data for table `spl_hpft_doctors_orders_tbl`
--

INSERT INTO `spl_hpft_doctors_orders_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `doctor_id_fk`, `doctors_orders`, `comment`, `ack_by`, `ack_time`, `status`, `order_created_time`, `order_type`, `document_id_fk`, `client_updated_at`, `created_on`, `updated_on`, `updated_by`) VALUES
	(1, '8d9e75dd-6387-4c77-9f01-9c52619d6acf', 3, 1, 9, 'Aspirin', 'Incase of headache', NULL, NULL, 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:00:08')), 'Prescription', NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:00:08')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:00:14')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:00:14')), 2),
	(2, '0cf2430c-7b05-49a8-b96b-86fed8fc104f', 3, 1, 9, 'Diet', 'Follow this diet', NULL, NULL, 0, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:01:23')), 'General', 1, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:01:22')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:01:32')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:01:32')), 2);
	

--
-- Dumping data for table `spl_hpft_pathology_record_tbl`
--

INSERT INTO `spl_hpft_pathology_record_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `test_performed`, `test_performed_time`, `test_result`, `comments`, `updated_by`, `client_updated_at`, `created_on`, `updated_on`) VALUES
	(1, 'DBCBFE021839FC638DF4727F77FE700BFB', 3, 1, 'Full Blood Examination', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'18:30:00')), 'normal', 'normal', 2, NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:10')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:10'))),
	(2, 'DB2B143D6B5E644F399D548D86611C7856', 3, 1, 'TSH (Thyroid Stimulating Hormone) Quantification', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'18:30:00')), 'normal', 'normal', 2, NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:56')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:56')));
	
	
--
-- Dumping data for table `spl_hpft_pathology_record_doc_tbl`
--

INSERT INTO `spl_hpft_pathology_record_doc_tbl` (`pathology_id_fk`, `document_id_fk`, `created_on`, `updated_on`) VALUES
	(1, 2, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:10')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:10'))),
	(2, 3, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:56')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:03:56')));
	
--
-- Dumping data for table `spl_hpft_treatment_tbl`
--

INSERT INTO `spl_hpft_treatment_tbl` (`id`, `uuid`, `cpm_id_fk`, `admission_id_fk`, `treatment_done`, `treatment_performed_time`, `details`, `post_observation`, `updated_by`, `client_updated_at`, `created_on`, `updated_on`) VALUES
	(1, 'DB856147B493D5F5476C19CD77419E634D', 3, 1, 'Allopathy Treatment', timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'18:30:00')), 'allopathy medicine course', 'infection under control', 2, NULL, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:04:21')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:04:21')));
	
	
--
-- Dumping data for table `spl_hpft_treatment_doc_tbl`
--

INSERT INTO `spl_hpft_treatment_doc_tbl` (`treatment_id_fk`, `document_id_fk`, `created_on`, `updated_on`) VALUES
	(1, 4, timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:04:21')), timestamp(ADDTIME(( curdate() - INTERVAL 3 DAY),'05:04:21')));
 
 
 --
-- Dumping data for table `spl_hpft_user_patient_monitor_mapping`
--

INSERT INTO `spl_hpft_user_patient_monitor_mapping` (`id`, `uuid`, `cpm_id_fk`, `usr_id_fk`, `sp_id_fk`, `patient_id_fk`, `created_on`, `updated_on`, `updated_by`) 
VALUES 
('1', 'UP001', '3', '2', 3, null, '2019-04-26 16:41:06', '2019-04-26 16:41:07', '2'),
('2', 'UP002', '3', '2', 7, null, '2019-04-26 16:41:06', '2019-04-26 16:41:07', '2'),
('3', 'UP003', '3', '8', null, null, '2019-04-26 16:41:06', '2019-04-26 16:41:07', '2'),
('4', 'UP004', '3', '9', 7, 1, '2019-04-26 16:41:06', '2019-04-26 16:41:07', '2');
 
 