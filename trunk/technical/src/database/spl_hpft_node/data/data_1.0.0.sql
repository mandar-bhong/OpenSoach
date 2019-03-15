--
-- Dumping data for table `spl_node_sync_config_tbl`
--


INSERT INTO `spl_node_sync_config_tbl` (`store_name`, `updated_on`, `has_qry`, `select_count_qry`, `select_qry`, `insert_qry`, `update_qry`,`data_source`, `query_data`) VALUES
	('service_point_tbl','2018-01-01 00:00:00','select count(*) as count from spl_node_sp_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_node_sp_tbl where cpm_id_fk = :cpmid and updated_on > :updatedon','select * from spl_node_sp_tbl where cpm_id_fk = :cpmid and updated_on > :updatedon','insert_qry','update_qry',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('conf_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_conf_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_conf_tbl where cpm_id_fk = :cpmid and updated_on > :updatedon','select * from spl_hpft_conf_tbl where cpm_id_fk = :cpmid and updated_on > :updatedon','insert_qry','update_qry',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('patient_master_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_master_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_master_tbl where cpm_id_fk = :cpmid and updated_on > :updatedon','select * from spl_hpft_patient_master_tbl where cpm_id_fk = :cpmid and updated_on > :updatedon','insert into spl_hpft_patient_master_tbl 
(uuid,cpm_id_fk,patient_reg_no,fname,lname,mob_no,age,blood_grp,gender,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,:patient_reg_no,:fname,:lname,:mob_no,:age,:blood_grp,:gender, STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_master_tbl 
set cpm_id_fk = :cpm_id_fk, patient_reg_no = :patient_reg_no, fname = :fname, lname = :lname, mob_no = :mob_no, age = :age, blood_grp = :blood_grp, gender = :gender,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by WHERE uuid = :uuid',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('schedule_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_conf_tbl where uuid = ?','select count(*) as count, max(pconf.updated_on) as max_updated_on from spl_hpft_patient_conf_tbl pconf left join spl_hpft_patient_admission_tbl  padmsn on padmsn.id = pconf.admission_id_fk where pconf.cpm_id_fk = :cpmid and pconf.updated_on > :updatedon','select pconf.uuid, padmsn.uuid as admission_uuid,pconf.conf_type_code,pconf.conf,pconf.end_date,pconf.status,pconf.updated_by,pconf.updated_on from spl_hpft_patient_conf_tbl pconf
			left join spl_hpft_patient_admission_tbl  padmsn on padmsn.id = pconf.admission_id_fk where pconf.cpm_id_fk = :cpmid and pconf.updated_on > :updatedon','insert into spl_hpft_patient_conf_tbl 
(uuid,cpm_id_fk,admission_id_fk,conf_type_code,conf,end_date,status,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1) ,:conf_type_code,:conf,STR_TO_DATE(:end_date ,"%Y-%m-%dT%T.%xZ"),:status,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_conf_tbl set cpm_id_fk = :cpm_id_fk, admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), conf_type_code = :conf_type_code, conf = :conf, end_date= STR_TO_DATE(:end_date ,"%Y-%m-%dT%T.%xZ"),status=:status,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by WHERE uuid = :uuid',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('patient_admission_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_admission_tbl where uuid = ?','select count(*) as count, max(padmsn.updated_on) as max_updated_on from spl_hpft_patient_admission_tbl padmsn 
left join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk
left join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk where padmsn.cpm_id_fk = :cpmid and padmsn.updated_on > :updatedon','select padmsn.uuid,patient.uuid as patient_uuid,padmsn.patient_reg_no,padmsn.bed_no,padmsn.status,sp.uuid as sp_uuid,padmsn.dr_incharge,padmsn.admitted_on,padmsn.discharged_on,padmsn.updated_by,padmsn.updated_on from spl_hpft_patient_admission_tbl padmsn 
left join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk
left join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk where padmsn.cpm_id_fk = :cpmid and padmsn.updated_on > :updatedon','insert into spl_hpft_patient_admission_tbl 
(uuid,cpm_id_fk,patient_id_fk,patient_reg_no,bed_no,status,sp_id_fk,dr_incharge,admitted_on,discharged_on,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,(select id as patient_id_fk from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),:patient_reg_no,:bed_no,:status,(select sp_id_fk from spl_node_sp_tbl where uuid = :sp_uuid limit 1),:dr_incharge,:admitted_on,:discharged_on,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_admission_tbl 
set uuid = :uuid, cpm_id_fk = :cpm_id_fk, patient_id_fk = (select id as patient_id_fk from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), patient_reg_no = :patient_reg_no, bed_no = :bed_no, status = :status, sp_id_fk = (select sp_id_fk from spl_node_sp_tbl where uuid = :sp_uuid limit 1), dr_incharge = :dr_incharge, admitted_on = :admitted_on, discharged_on = :discharged_on,client_updated_at = STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by WHERE uuid = :uuid',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('patient_personal_details_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_personal_details_tbl where uuid = ?','select count(*) as count, max(pdetails.updated_on) as max_updated_on from spl_hpft_patient_personal_details_tbl pdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = pdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where pdetails.cpm_id_fk = :cpmid and pdetails.updated_on > :updatedon','select pdetails.uuid,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid, pdetails.age,pdetails.other_details,pdetails.person_accompanying,pdetails.updated_by,pdetails.updated_on from spl_hpft_patient_personal_details_tbl pdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = pdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where pdetails.cpm_id_fk = :cpmid and pdetails.updated_on > :updatedon','insert into spl_hpft_patient_personal_details_tbl 
(cpm_id_fk,patient_id,admission_id_fk,uuid,age,other_details,person_accompanying,client_updated_at,updated_by) 
values (:cpm_id_fk,(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:uuid,:age,:other_details,:person_accompanying,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_personal_details_tbl 
set cpm_id_fk = :cpm_id_fk, patient_id =(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), age = :age, other_details = :other_details,person_accompanying = :person_accompanying, client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by WHERE uuid = :uuid',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('patient_medical_details_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_medical_details_tbl where uuid = ?','select count(*) as count, max(mdetails.updated_on) as max_updated_on from spl_hpft_patient_medical_details_tbl mdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = mdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where mdetails.cpm_id_fk = :cpmid and mdetails.updated_on > :updatedon','select mdetails.uuid, patient.uuid as patient_uuid,padmsn.uuid as admission_uuid,mdetails.present_complaints,mdetails.reason_for_admission,mdetails.history_present_illness,mdetails.past_history,mdetails.treatment_before_admission,mdetails.investigation_before_admission,mdetails.family_history,mdetails.allergies,mdetails.personal_history,mdetails.updated_by,mdetails.updated_on from spl_hpft_patient_medical_details_tbl mdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = mdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where mdetails.cpm_id_fk = :cpmid and mdetails.updated_on > :updatedon','insert into spl_hpft_patient_medical_details_tbl 
(uuid,cpm_id_fk,patient_id,admission_id_fk,present_complaints,reason_for_admission,history_present_illness,past_history,treatment_before_admission,investigation_before_admission,family_history,allergies,personal_history,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk,(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:present_complaints,:reason_for_admission,:history_present_illness,:past_history,:treatment_before_admission,:investigation_before_admission,:family_history,:allergies,:personal_history,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_patient_medical_details_tbl 
set  cpm_id_fk = :cpm_id_fk, patient_id = (select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), present_complaints= :present_complaints,reason_for_admission = :reason_for_admission,history_present_illness = :history_present_illness, past_history = :past_history, treatment_before_admission = :treatment_before_admission,investigation_before_admission:investigation_before_admission, family_history = :family_history, allergies = :allergies, personal_history = :personal_history,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"), updated_by = :updated_by 
WHERE uuid = :uuid',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('action_txn_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_action_txn_tbl where uuid = ?','select count(*) as count, max(atxn.updated_on) as max_updated_on from spl_hpft_action_txn_tbl atxn
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = atxn.admission_id_fk
left join spl_hpft_patient_conf_tbl pconf on pconf.id = atxn.patient_conf_id_fk where atxn.cpm_id_fk = :cpmid and atxn.updated_on > :updatedon','select atxn.uuid,padmsn.uuid as admission_uuid,pconf.uuid as schedule_uuid,atxn.txn_data,atxn.scheduled_time,atxn.txn_state,atxn.conf_type_code,atxn.runtime_config_data,atxn.updated_by,atxn.updated_on from spl_hpft_action_txn_tbl atxn
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = atxn.admission_id_fk
left join spl_hpft_patient_conf_tbl pconf on pconf.id = atxn.patient_conf_id_fk where atxn.cpm_id_fk = :cpmid and atxn.updated_on > :updatedon','insert into spl_hpft_action_txn_tbl 
(uuid,cpm_id_fk,patient_conf_id_fk,admission_id_fk,txn_data,runtime_config_data,scheduled_time,txn_state,conf_type_code,client_updated_at,updated_by) 
values 
(:uuid,:cpm_id_fk, (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:txn_data,:runtime_config_data,STR_TO_DATE(:scheduled_time ,"%Y-%m-%dT%T.%xZ"),:txn_state,:conf_type_code,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_action_txn_tbl set cpm_id_fk = :cpm_id_fk, patient_conf_id_fk = (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1) ,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), txn_data = :txn_data, runtime_config_data = :runtime_config_data, scheduled_time = STR_TO_DATE(:scheduled_time ,"%Y-%m-%dT%T.%xZ"), txn_state = :txn_state, conf_type_code = :conf_type_code,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by WHERE uuid = :uuid',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),

	('doctors_orders_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_doctors_orders_tbl where uuid = ?','select count(*) as count, max(doc_ordrs.updated_on) as max_updated_on from spl_hpft_doctors_orders_tbl doc_ordrs
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = doc_ordrs.admission_id_fk
left join spl_hpft_document_tbl doc on doc_ordrs.document_id_fk = doc.id where doc_ordrs.cpm_id_fk = :cpmid and doc_ordrs.updated_on > :updatedon','select doc_ordrs.uuid,padmsn.uuid as admission_uuid,doc_ordrs.doctor_id_fk as doctor_id,doc_ordrs.doctors_orders,doc_ordrs.comment,doc_ordrs.ack_by, doc_ordrs.ack_time, doc_ordrs.status, doc_ordrs.order_created_time, doc_ordrs.order_type,doc.uuid as document_uuid,doc.name as document_name,doc.doctype,doc_ordrs.updated_on,doc_ordrs.updated_by from spl_hpft_doctors_orders_tbl doc_ordrs
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = doc_ordrs.admission_id_fk
left join spl_hpft_document_tbl doc on doc_ordrs.document_id_fk = doc.id where doc_ordrs.cpm_id_fk = :cpmid and doc_ordrs.updated_on > :updatedon','insert into spl_hpft_doctors_orders_tbl 
(uuid,cpm_id_fk,admission_id_fk,doctor_id_fk,doctors_orders,comment,ack_by,ack_time,status,order_created_time,order_type,document_id_fk,client_updated_at,updated_by)
values
(:uuid,:cpm_id_fk, (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),(select id as doctor_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),:doctors_orders,:comment,:ack_by,STR_TO_DATE(:ack_time ,"%Y-%m-%dT%T.%xZ"),:status,STR_TO_DATE(:order_created_time ,"%Y-%m-%dT%T.%xZ"),:order_type,(select id as document_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),:updated_by)','update spl_hpft_doctors_orders_tbl set cpm_id_fk = :cpm_id_fk,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),doctor_id_fk = :doctor_id_fk,doctors_orders = :doctors_orders,comment = :comment,ack_by = :ack_by,ack_time = STR_TO_DATE(:ack_time ,"%Y-%m-%dT%T.%xZ"),status = :status,order_created_time = STR_TO_DATE(:order_created_time ,"%Y-%m-%dT%T.%xZ"),order_type = :order_type,document_id_fk = (select id as document_id_fk from spl_hpft_document_tbl where uuid = :document_uuid limit 1),client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"),updated_by = :updated_by where uuid = :uuid',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),

	('document_tbl', '2018-01-01 00:00:00', 'select count(*) as count from spl_hpft_document_tbl where uuid = ?', 'select_count_qry', 'select_qry', 'insert into spl_hpft_document_tbl (uuid,cpm_id_fk,name,doctype,store_name,persisted,updated_by,client_updated_at) values (:uuid,:cpm_id_fk,:doc_name,:doc_type,:datastore,0,:updated_by,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))', 'update_qry',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('treatment_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_treatment_tbl where uuid = ?','select count(*) as count, max(trtmnt.updated_on) as max_updated_on from spl_hpft_treatment_tbl trtmnt
inner join spl_hpft_patient_admission_tbl padmsn on padmsn.id = trtmnt.admission_id_fk
where trtmnt.cpm_id_fk = :cpmid and trtmnt.updated_on > :updatedon','select trtmnt.uuid,padmsn.uuid  as admission_uuid,treatment_done,treatment_performed_time,details,post_observation,trtmnt.updated_by,trtmnt.updated_on 
from spl_hpft_treatment_tbl trtmnt
inner join spl_hpft_patient_admission_tbl padmsn on padmsn.id = trtmnt.admission_id_fk
where trtmnt.cpm_id_fk =:cpmid and trtmnt.updated_on > :updatedon','insert_qry','update_qry',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('treatment_doc_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_treatment_doc_tbl where uuid = ?','select count(*) as count, max(trtmnt.updated_on) as max_updated_on from spl_hpft_treatment_tbl trtmnt
left join spl_hpft_treatment_doc_tbl tdoc on tdoc.treatment_id_fk = trtmnt.id
left join spl_hpft_document_tbl doc on doc.id = tdoc.document_id_fk where trtmnt.updated_on > :updatedon','select trtmnt.uuid as treatment_uuid,doc.uuid as document_uuid from spl_hpft_treatment_tbl trtmnt
left join spl_hpft_treatment_doc_tbl tdoc on tdoc.treatment_id_fk = trtmnt.id
left join spl_hpft_document_tbl doc on doc.id = tdoc.document_id_fk where trtmnt.updated_on > :updatedon','insert_qry','update_qry',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('pathology_record_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_pathology_record_tbl where uuid = ?','select count(*) as count, max(prec.updated_on) as max_updated_on from spl_hpft_pathology_record_tbl prec
inner join spl_hpft_patient_admission_tbl padmsn on padmsn.id = prec.admission_id_fk 
where prec.cpm_id_fk = :cpmid and prec.updated_on > :updatedon','select prec.uuid, padmsn.uuid as admission_uuid,test_performed,test_performed_time,test_result,comments,prec.updated_by,prec.updated_on 
from spl_hpft_pathology_record_tbl prec
inner join spl_hpft_patient_admission_tbl padmsn on padmsn.id = prec.admission_id_fk 
where prec.cpm_id_fk = :cpmid and prec.updated_on > :updatedon','insert_qry','update_qry',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('pathology_record_doc_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_pathology_record_doc_tbl where uuid = ?','select count(*) as count, max(prec.updated_on) as max_updated_on from spl_hpft_pathology_record_tbl prec
left join spl_hpft_pathology_record_doc_tbl precdoc on precdoc.pathology_id_fk = prec.id
left join spl_hpft_document_tbl doc on doc.id = precdoc.document_id_fk where prec.updated_on > :updatedon','select prec.uuid as pathology_record_uuid,doc.uuid as document_uuid from spl_hpft_pathology_record_tbl prec
left join spl_hpft_pathology_record_doc_tbl precdoc on precdoc.pathology_id_fk = prec.id
left join spl_hpft_document_tbl doc on doc.id = precdoc.document_id_fk where prec.updated_on > :updatedon','insert_qry','update_qry',2,'{"select": {"filters": ["cpm", "updatedon"]}}'),

('mst_user_tbl','2018-01-01 00:00:00','has_qry','SELECT count(usr.id) as count
	FROM spl_master_user_tbl usr
	INNER  JOIN spl_master_usr_cpm_tbl ucpm ON usr.id = ucpm.user_id_fk
	INNER  JOIN spl_master_user_role_tbl urole ON urole.id = ucpm.urole_id_fk
	WHERE urole.prod_id_fk = 2 AND ucpm.cpm_id_fk = :cpmid and usr.updated_on > :updatedon','SELECT usr.id as usr_id, usr.usr_name, urole.urole_code, urole.urole_name ,fname,lname,
	ucpm.cpm_id_fk
	FROM spl_master_user_tbl usr
	INNER  JOIN spl_master_usr_cpm_tbl ucpm ON usr.id = ucpm.user_id_fk
	INNER  JOIN spl_master_user_role_tbl urole ON urole.id = ucpm.urole_id_fk
	INNER JOIN spl_master_usr_details_tbl usrd on usrd.usr_id_fk = usr.id
	WHERE urole.prod_id_fk = 2 AND  ucpm.cpm_id_fk = :cpmid and usr.updated_on > :updatedon','insert_qry','update_qry',1,'{"select": {"filters": ["cpm", "updatedon"]}}'),
	
	('action_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_action_tbl where uuid = ?','select count(*) as count, max(actn.updated_on) as max_updated_on from spl_hpft_action_tbl actn
inner join spl_hpft_patient_admission_tbl padmsn on padmsn.id = actn.admission_id_fk
inner join spl_hpft_patient_conf_tbl pconf on pconf.id = actn.patient_conf_id_fk where actn.cpm_id_fk = :cpmid and actn.updated_on > :updatedon','select actn.uuid,padmsn.uuid as admission_uuid,pconf.uuid as schedule_uuid,actn.conf_type_code,actn.scheduled_time,actn.is_deleted,actn.updated_on,actn.updated_by 
from spl_hpft_action_tbl actn
inner join spl_hpft_patient_admission_tbl padmsn on padmsn.id = actn.admission_id_fk
inner join spl_hpft_patient_conf_tbl pconf on pconf.id = actn.patient_conf_id_fk where actn.cpm_id_fk =:cpmid and actn.updated_on > :updatedon','insert into spl_hpft_action_tbl 
(uuid,cpm_id_fk,admission_id_fk,patient_conf_id_fk,conf_type_code,scheduled_time,is_deleted,client_updated_at,updated_by)
values
(:uuid,:cpm_id_fk,(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),(select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1),:conf_type_code,STR_TO_DATE(:scheduled_time ,"%Y-%m-%dT%T.%xZ"),:is_deleted,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ")),:updated_by)','update spl_hpft_action_tbl set cpm_id_fk = :cpm_id_fk,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),patient_conf_id_fk = (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1),conf_type_code = :conf_type_code,scheduled_time = :scheduled_time,is_deleted = :is_deleted,client_updated_at = STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ")),updated_by = :updated_by',2,'{"select": {"filters": ["cpm", "updatedon"]}}');
	

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




