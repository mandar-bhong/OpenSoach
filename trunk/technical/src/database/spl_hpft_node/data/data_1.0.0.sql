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
(uuid,cpm_id_fk,patient_conf_id_fk,admission_id_fk,txn_data,runtime_config_data,txn_date,txn_state,conf_type_code,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk, (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:txn_data,:runtime_config_data,STR_TO_DATE(:txn_date ,"%Y-%m-%dT%T.%xZ"),:txn_state,:conf_type_code,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_action_txn_tbl set cpm_id_fk = :cpm_id_fk, patient_conf_id_fk = (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1) ,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), txn_data = :txn_data, runtime_config_data = :runtime_config_data, txn_date = STR_TO_DATE(:txn_date ,"%Y-%m-%dT%T.%xZ"), txn_state = :txn_state, conf_type_code = :conf_type_code,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by WHERE uuid = :uuid'),

	('doctors_orders_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_doctors_orders_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_doctors_orders_tbl where updated_on > ?','select doc_ordrs.uuid,padmsn.uuid as admission_uuid,doc_ordrs.doctor_id_fk as doctor_id,doc_ordrs.doctors_orders,doc.uuid as document_uuid,doc.name as document_name,doc.doctype,doc_ordrs.updated_on,doc_ordrs.updated_by from spl_hpft_doctors_orders_tbl doc_ordrs
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = doc_ordrs.admission_id_fk
left join spl_hpft_document_tbl doc on doc_ordrs.document_id_fk = doc.id where doc_ordrs.updated_on > ?','insert into spl_hpft_doctors_orders_tbl 
(uuid,cpm_id_fk,admission_id_fk,doctor_id_fk,doctors_orders,document_id_fk,client_updated_at,updated_by)
values
(:uuid,:cpm_id_fk, (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),(select id as doctor_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),:doctors_orders,(select id as document_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_doctors_orders_tbl set cpm_id_fk = :cpm_id_fk,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),doctor_id_fk = :doctor_id_fk,doctors_orders = :doctors_orders,document_id_fk = (select id as document_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by where uuid = :uuid'),

	('document_tbl', '2018-01-01 00:00:00', 'select count(*) as count from spl_hpft_document_tbl where uuid = ?', 'select_count_qry', 'select_qry', 'insert into spl_hpft_document_tbl (uuid,cpm_id_fk,name,doctype,store_name,persisted,updated_by,client_updated_at) values (:uuid,:cpm_id_fk,:doc_name,:doc_type,:datastore,0,:updated_by,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))', 'update_qry'),
	
	('treatment_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_treatment_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_treatment_tbl where updated_on > ?','select uuid,(select uuid from spl_hpft_patient_admission_tbl where id = admission_id_fk limit 1) as admission_uuid,treatment_done,details,post_observation,updated_by,updated_on from spl_hpft_treatment_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('treatment_doc_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_treatment_doc_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_treatment_doc_tbl where updated_on > ?','select (select uuid as treatment_uuid from spl_hpft_treatment_tbl where id = treatment_id_fk limit 1) as treatment_uuid,(select uuid as document_uuid from spl_hpft_document_tbl where id = document_id_fk limit 1) as document_uuid,created_on,updated_on from spl_hpft_treatment_doc_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('pathology_record_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_pathology_record_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_pathology_record_tbl where updated_on > ?','select uuid,(select uuid from spl_hpft_patient_admission_tbl where id = admission_id_fk limit 1) as admission_uuid,test_performed,test_result,comments,updated_by,updated_on from spl_hpft_pathology_record_tbl where updated_on > ?','insert_qry','update_qry'),
	
	('pathology_record_doc_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_pathology_record_doc_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_pathology_record_doc_tbl where updated_on > ?','select (select uuid as pathology_record_uuid from spl_hpft_pathology_record_tbl where id = pathology_id_fk limit 1) as pathology_record_uuid,(select uuid as document_uuid from spl_hpft_document_tbl where id = document_id_fk limit 1) as document_uuid,created_on,updated_on from spl_hpft_treatment_doc_tbl where updated_on > ?','insert_qry','update_qry');
	

--
-- Dumping data for table `spl_node_cpm_tbl`
--

INSERT INTO `spl_node_cpm_tbl` (`cpm_id_fk`) VALUES ('3');
	
	
--
-- Dumping data for table `spl_hpft_conf_tbl`
--

INSERT INTO `spl_hpft_conf_tbl` (`uuid`, `cpm_id_fk`, `conf_type_code`, `conf`, `updated_by`) 
VALUES 
('C0001', '3', 'Monitor', '{"tasks":[{"name":"Temperature"},{"name":"Blood Pressure"},{"name":"Pulse Rate"},{"name":"Respiration Rate"}]}',1);



INSERT INTO `spl_hpft_conf_tbl` (`uuid`, `cpm_id_fk`, `conf_type_code`, `conf`, `updated_by`) 
VALUES 
('22072dfe99c04ba6', '3', 'VITAL_PARAMETERS', '[
{"name":"TEMP","displayname":"Temperature",  
"fields":[{"name":"Temperature","type":"Number", "unitdisplay":"°F", "range":{"low":"96","high":"98"}}]},
{"name":"BLOOD_PRESSURE","displayname":"Blood Pressure", 
"fields":[{"name":"Systolic","type":"Number", "unitdisplay":"mmHg", "range":{"low":"120","high":"140"}}, 
{"name":"Diastolic","type":"Number", "unitdisplay":"mmHg", "range":{"low":"70","high":"90"}}]},
{"name":"PLUSE_RATE", "displayname":"Pulse Rate", "fields":[{"name":"Pulse Rate", "unitdisplay":"bpm", "range":{"low":"60","high":"100"}}]},
{"name":"RESP_RATE", "displayname":"Respiration Rate", "fields":[{"name":"Respiration Rate", "unitdisplay":"bpm", "range":{"low":"60","high":"100"}}]}]',1);

INSERT INTO `spl_hpft_conf_tbl` (`uuid`, `cpm_id_fk`, `conf_type_code`, `conf`, `updated_by`) 
VALUES 
('e9db1c9f0ac44b97', '3', 'MEDICINE_TYPE', '["Tablet", "Capsules", "Syrup", "Injection", "Ointment", "Eye Drop", "Drop","Puff"]',1);

INSERT INTO `spl_hpft_conf_tbl` (`uuid`, `cpm_id_fk`, `conf_type_code`, `conf`, `updated_by`) 
VALUES 
('cc47eb5cbd4d4a1a', '3', 'INPUT_TYPE', '["Oral", "LV"]',1);

INSERT INTO `spl_hpft_conf_tbl` (`uuid`, `cpm_id_fk`, `conf_type_code`, `conf`, `updated_by`) 
VALUES 
('f9d1d51e56e64455', '3', 'OUTPUT_TYPE', '["Urine", "Gastric", "Stool"]',1);




