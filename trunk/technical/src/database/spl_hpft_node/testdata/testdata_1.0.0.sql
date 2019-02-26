--
-- Dumping data for table `spl_node_cpm_tbl`
--

INSERT INTO `spl_node_cpm_tbl` (`cpm_id_fk`) VALUES ('3');


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
-- Dumping data for table `spl_node_sync_config_tbl`
--


INSERT INTO `spl_node_sync_config_tbl` (`store_name`, `updated_on`, `has_qry`, `select_count_qry`, `select_qry`, `insert_qry`, `update_qry`) VALUES
	('service_point_tbl','2018-01-01 00:00:00','select count(*) as count from spl_node_sp_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_node_sp_tbl where updated_on > ?','select * from spl_node_sp_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('conf_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_conf_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_conf_tbl where updated_on > ?','select * from spl_hpft_conf_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('patient_master_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_master_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_master_tbl where updated_on > ?','select * from spl_hpft_patient_master_tbl where updated_on > ?','insert: insert into spl_hpft_patient_master_tbl 
(uuid,cpm_id_fk,patient_reg_no,fname,lname,mob_no,age,blood_grp,gender,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,:patient_reg_no,:fname,:lname,:mob_no,:age,:blood_grp,:gender, STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update: update spl_hpft_patient_master_tbl 
set cpm_id_fk = :cpm_id_fk, patient_reg_no = :patient_reg_no, fname = :fname, lname = :lname, mob_no = :mob_no, age = :age, blood_grp = :blood_grp, gender = :gender,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by WHERE uuid = :uuid'),
	
	('schedule_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_conf_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_conf_tbl where updated_on > ?','select pconf.uuid, padmsn.uuid as admission_uuid,pconf.conf_type_code,pconf.conf,pconf.end_date,pconf.updated_by,pconf.updated_on from spl_hpft_patient_conf_tbl pconf
			left join spl_hpft_patient_admission_tbl  padmsn on padmsn.id = pconf.admission_id_fk where pconf.updated_on > ?','insert into spl_hpft_patient_conf_tbl 
(uuid,cpm_id_fk,admission_id_fk,conf_type_code,conf,end_date,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1) ,:conf_type_code,:conf,STR_TO_DATE(:end_date ,"%Y-%m-%dT%T.%xZ"),STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_conf_tbl set cpm_id_fk = :cpm_id_fk, admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), conf_type_code = :conf_type_code, conf = :conf, end_date= STR_TO_DATE(:end_date ,"%Y-%m-%dT%T.%xZ"),client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by WHERE uuid = :uuid'),
	
	('patient_admission_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_admission_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_admission_tbl where updated_on > ?','select padmsn.uuid,patient.uuid as patient_uuid,padmsn.patient_reg_no,padmsn.bed_no,padmsn.status,sp.uuid as sp_uuid,padmsn.dr_incharge,padmsn.admitted_on,padmsn.discharged_on,padmsn.updated_by,padmsn.updated_on from spl_hpft_patient_admission_tbl padmsn 
left join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk
left join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk where padmsn.updated_on > ?','insert into spl_hpft_patient_admission_tbl 
(uuid,cpm_id_fk,patient_id_fk,patient_reg_no,bed_no,status,sp_id_fk,dr_incharge,admitted_on,discharged_on,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,(select id as patient_id_fk from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),:patient_reg_no,:bed_no,:status,(select sp_id_fk from spl_node_sp_tbl where uuid = :sp_uuid limit 1),:dr_incharge,:admitted_on,:discharged_on,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_admission_tbl 
set uuid = :uuid, cpm_id_fk = :cpm_id_fk, patient_id_fk = (select id as patient_id_fk from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), patient_reg_no = :patient_reg_no, bed_no = :bed_no, status = :status, sp_id_fk = (select sp_id_fk from spl_node_sp_tbl where uuid = :sp_uuid limit 1), dr_incharge = :dr_incharge, admitted_on = :admitted_on, discharged_on = :discharged_on,client_updated_at = STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by WHERE uuid = :uuid'),
	
	('patient_personal_details_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_personal_details_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_personal_details_tbl where updated_on > ?','select pdetails.uuid,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid, pdetails.age,pdetails.other_details,pdetails.person_accompanying,pdetails.updated_by,pdetails.updated_on from spl_hpft_patient_personal_details_tbl pdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = pdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where pdetails.updated_on > ?','insert into spl_hpft_patient_personal_details_tbl 
(cpm_id_fk,patient_id,admission_id_fk,uuid,age,other_details,person_accompanying,client_updated_at,updated_by) 
values (:cpm_id_fk,(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:uuid,:age,:other_details,:person_accompanying,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_personal_details_tbl 
set cpm_id_fk = :cpm_id_fk, patient_id =(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), age = :age, other_details = :other_details,person_accompanying = :person_accompanying, client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by WHERE uuid = :uuid'),
	
	('patient_medical_details_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_medical_details_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_medical_details_tbl where updated_on > ?','select mdetails.uuid, patient.uuid as patient_uuid,padmsn.uuid as admission_uuid,mdetails.present_complaints,mdetails.reason_for_admission,mdetails.history_present_illness,mdetails.past_history,mdetails.treatment_before_admission,mdetails.investigation_before_admission,mdetails.family_history,mdetails.allergies,mdetails.personal_history,mdetails.updated_by,mdetails.updated_on from spl_hpft_patient_medical_details_tbl mdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = mdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where mdetails.updated_on > ?','insert into spl_hpft_patient_medical_details_tbl 
(uuid,cpm_id_fk,patient_id,admission_id_fk,present_complaints,reason_for_admission,history_present_illness,past_history,treatment_before_admission,investigation_before_admission,family_history,allergies,personal_history,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:present_complaints,:reason_for_admission,:history_present_illness,:past_history,:treatment_before_admission,:investigation_before_admission,:family_history,:allergies,:personal_history,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_medical_details_tbl 
set  cpm_id_fk = :cpm_id_fk, patient_id = (select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), present_complaints= :present_complaints,reason_for_admission = :reason_for_admission,history_present_illness = :history_present_illness, past_history = :past_history, treatment_before_admission = :treatment_before_admission,investigation_before_admission:investigation_before_admission, family_history = :family_history, allergies = :allergies, personal_history = :personal_history,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by 
WHERE uuid = :uuid'),
	
	('action_txn_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_action_txn_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_action_txn_tbl where updated_on > ?','select atxn.uuid,padmsn.uuid as admission_uuid,pconf.uuid as schedule_uuid,atxn.txn_data,atxn.txn_date,atxn.txn_state,atxn.conf_type_code,atxn.runtime_config_data,atxn.updated_by,atxn.updated_on from spl_hpft_action_txn_tbl atxn
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = atxn.admission_id_fk
left join spl_hpft_patient_conf_tbl pconf on pconf.id = atxn.patient_conf_id_fk where atxn.updated_on > ?','insert into spl_hpft_action_txn_tbl 
(uuid,cpm_id_fk,patient_conf_id_fk,:admission_id_fk,txn_data,runtime_config_data,txn_date,txn_state,conf_type_code,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk, (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:txn_data,:runtime_config_data,:txn_date,:txn_state,:conf_type_code,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_action_txn_tbl set cpm_id_fk = :cpm_id_fk, patient_conf_id_fk = (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1) ,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), txn_data = :txn_data, runtime_config_data = :runtime_config_data, txn_date = :txn_date, txn_state = :txn_state, conf_type_code = :conf_type_code,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by WHERE uuid = :uuid'),

	('doctors_orders_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_doctors_orders_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_doctors_orders_tbl where updated_on > ?','select doc_ordrs.uuid,padmsn.uuid as admission_uuid,doc_ordrs.doctor_id_fk as doctor_id,doc_ordrs.doctors_orders,doc.uuid as document_uuid,doc.name as document_name,doc.doctype,doc_ordrs.updated_on,doc_ordrs.updated_by from spl_hpft_doctors_orders_tbl doc_ordrs
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = doc_ordrs.admission_id_fk
left join spl_hpft_document_tbl doc on doc_ordrs.document_id_fk = doc.id where doc_ordrs.updated_on > ?','insert into spl_hpft_doctors_orders_tbl 
(uuid,cpm_id_fk,admission_id_fk,doctor_id_fk,doctors_orders,document_id_fk,client_updated_at,updated_by)
values
(:uuid,:cpm_id_fk, (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),(select id as doctor_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),:doctors_orders,(select id as document_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_doctors_orders_tbl set cpm_id_fk = :cpm_id_fk,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),doctor_id_fk = :doctor_id_fk,doctors_orders = :doctors_orders,document_id_fk = (select id as document_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by where uuid = :uuid'),

	('document_tbl', '2018-01-01 00:00:00', 'select count(*) as count from spl_hpft_document_tbl where uuid = ?', 'select_count_qry', 'select_qry', 'insert into spl_hpft_document_tbl (uuid,cpm_id_fk,name,doctype,store_name,persisted,updated_by,client_updated_at) values (:uuid,:cpm_id_fk,:doc_name,:doc_type,:datastore,0,:updated_by,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))', 'update_qry'),
	
	('treatment_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_treatment_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_treatment_tbl where updated_on > ?','select * from spl_hpft_treatment_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('treatment_doc_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_treatment_doc_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_treatment_doc_tbl where updated_on > ?','select * from spl_hpft_treatment_doc_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('pathology_record_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_pathology_record_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_pathology_record_tbl where updated_on > ?','select * from spl_hpft_pathology_record_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('pathology_record_doc_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_pathology_record_doc_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_pathology_record_doc_tbl where updated_on > ?','select * from spl_hpft_pathology_record_doc_tbl where updated_on > ?','insert_qry','update_qry');


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
('PA001', '3', 1, 'P12B12213', '3A/312', '1', '3', '8', '2018-12-04 14:37:53', 1),
('PA002', '3', 2, 'P12B12214', '3B/323', '1', '3', '8', '2018-12-04 12:47:53', 1),
('PA003', '3', 3, 'P12B12215', '2A/643', '1', '3', '8', '2018-12-04 09:17:53', 1),
('PA004', '3', 4, 'P12B12216', '4A/415', '2', '3', '8', '2018-12-04 11:00:53', 1),
('PA005', '3', 5, 'P12B12217', '5A/616', '3', '3', '8', '2018-12-04 01:11:53', 1),
('PA006', '3', 6, 'P12B12218', '6A/317', '1', '3', '8', '2018-12-04 14:32:53', 1),
('PA007', '3', 7, 'P12B12219', '7A/312', '2', '7', '8', '2018-12-04 16:44:53', 1),
('PA008', '3', 8, 'P12B12210', '3A/319', '3', '7', '8', '2018-12-04 11:12:53', 1),
('PA009', '3', 9, 'P12B12221', '8A/314', '2', '7', '8', '2018-12-04 04:54:53', 1),
('PA010', '3', 10, 'P12B12222', '4A/309', '1', '7', '8', '2018-12-04 15:55:53', 1),
('PA011', '3', 11, 'P12B12223', '2B/231', '4', '8', '8', '2018-12-04 21:35:53', 1),
('PA012', '3', 12, 'P12B12224', '2B/232', '1', '8', '8', '2018-12-04 19:33:53', 1);


--
-- Dumping data for table `spl_hpft_conf_tbl`
--

INSERT INTO `spl_hpft_conf_tbl` (`uuid`, `cpm_id_fk`, `conf_type_code`, `conf`, `updated_by`) 
VALUES 
('C0001', '3', 'Monitor', '{"tasks":[{"name":"Temperature"},{"name":"Blood Pressure"},{"name":"Pulse Rate"},{"name":"Respiration Rate"}]}',1);





