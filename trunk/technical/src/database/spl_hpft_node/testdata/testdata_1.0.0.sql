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

INSERT INTO `spl_node_sp_tbl` (`sp_id_fk`, `uuid`,`cpm_id_fk`, `spc_id_fk`, `sp_name`, `sp_state`, `sp_state_since`) VALUES ('3', 'SP001','3', '1', 'General Ward 3', '1', UTC_TIMESTAMP);


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
	
	('patient_master_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_master_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_master_tbl where updated_on > ?','select * from spl_hpft_patient_master_tbl where updated_on > ?','insert into spl_hpft_patient_master_tbl 
(uuid,cpm_id_fk,patient_reg_no,fname,lname,mob_no,age,blood_grp,gender,client_updated_at) 
values 
(:uuid,:cpm_id_fk,:patient_reg_no,:fname,:lname,:mob_no,:age,:blood_grp,:gender, STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))','update spl_hpft_patient_master_tbl 
set cpm_id_fk = :cpm_id_fk, patient_reg_no = :patient_reg_no, fname = :fname, lname = :lname, mob_no = :mob_no, age = :age, blood_grp = :blood_grp, gender = :gender,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ") WHERE uuid = :uuid'),
	
	('schedule_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_conf_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_conf_tbl where updated_on > ?','select pconf.uuid, padmsn.uuid as admission_uuid,pconf.conf_type_code,pconf.conf,pconf.end_date,pconf.updated_on from spl_hpft_patient_conf_tbl pconf
			left join spl_hpft_patient_admission_tbl  padmsn on padmsn.id = pconf.admission_id_fk where pconf.updated_on > ?','insert into spl_hpft_patient_conf_tbl 
(uuid,cpm_id_fk,admission_id_fk,conf_type_code,conf,end_date,client_updated_at) 
values 
(:uuid,:cpm_id_fk,(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1) ,:conf_type_code,:conf,STR_TO_DATE(:end_date ,"%Y-%m-%dT%T.%xZ"),STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))','update spl_hpft_patient_conf_tbl set cpm_id_fk = :cpm_id_fk, admission_id_fk =(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), conf_type_code = :conf_type_code, conf = :conf, end_date= STR_TO_DATE(:end_date ,"%Y-%m-%dT%T.%xZ"),client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ") WHERE uuid = :uuid'),
	
	('patient_admission_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_admission_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_admission_tbl where updated_on > ?','select padmsn.uuid,patient.uuid as patient_uuid,padmsn.patient_reg_no,padmsn.bed_no,padmsn.status,sp.uuid as sp_uuid,padmsn.dr_incharge,padmsn.admitted_on,padmsn.discharged_on,padmsn.updated_on from spl_hpft_patient_admission_tbl padmsn 
left join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk
left join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk where padmsn.updated_on > ?','insert into spl_hpft_patient_admission_tbl 
(uuid,cpm_id_fk,patient_id_fk,patient_reg_no,bed_no,status,sp_id_fk,dr_incharge,admitted_on,discharged_on,client_updated_at) 
values 
(:uuid,:cpm_id_fk,(select id as patient_id_fk from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),:patient_reg_no,:bed_no,:status,(select sp_id_fk from spl_node_sp_tbl where uuid = :sp_uuid limit 1),:dr_incharge,STR_TO_DATE(:admitted_on ,"%Y-%m-%dT%T.%xZ"),STR_TO_DATE(:discharged_on ,"%Y-%m-%dT%T.%xZ"),STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))','update spl_hpft_patient_admission_tbl 
set uuid = :uuid, cpm_id_fk = :cpm_id_fk, patient_id_fk = (select id as patient_id_fk from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), patient_reg_no = :patient_reg_no, bed_no = :bed_no, status = :status, sp_id_fk = (select sp_id_fk from spl_node_sp_tbl where uuid = :sp_uuid limit 1), dr_incharge = :dr_incharge, admitted_on = STR_TO_DATE(:admitted_on ,"%Y-%m-%dT%T.%xZ"), discharged_on = STR_TO_DATE(:discharged_on ,"%Y-%m-%dT%T.%xZ"),client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ") WHERE uuid = :uuid'),
	
	('patient_personal_details_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_personal_details_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_personal_details_tbl where updated_on > ?','select pdetails.uuid,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid, pdetails.age,pdetails.weight,pdetails.other_details,pdetails.updated_on from spl_hpft_patient_personal_details_tbl pdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = pdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where pdetails.updated_on > ?','insert into spl_hpft_patient_personal_details_tbl 
(cpm_id_fk,patient_id,admission_id_fk,uuid,age,weight,other_details,client_updated_at) 
values (:cpm_id_fk,(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:uuid,:age,:weight,:other_details,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))','update spl_hpft_patient_personal_details_tbl 
set cpm_id_fk = :cpm_id_fk, patient_id =(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), age = :age, weight = :weight, other_details = :other_details, client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ") WHERE uuid = :uuid'),
	
	('patient_medical_details_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_patient_medical_details_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_patient_medical_details_tbl where updated_on > ?','select mdetails.uuid, patient.uuid as patient_uuid,padmsn.uuid as admission_uuid,mdetails.reason_for_admission,mdetails.patient_medical_hist,mdetails.treatment_recieved_before,mdetails.family_hist,mdetails.menstrual_hist,mdetails.allergies,mdetails.personal_history,mdetails.general_physical_exam,mdetails.systematic_exam,mdetails.updated_on from spl_hpft_patient_medical_details_tbl mdetails
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = mdetails.admission_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk where mdetails.updated_on > ?','insert into spl_hpft_patient_medical_details_tbl 
(uuid,cpm_id_fk,patient_id,admission_id_fk,reason_for_admission,patient_medical_hist,treatment_recieved_before,family_hist,menstrual_hist,allergies,personal_history,general_physical_exam,systematic_exam,client_updated_at) 
values 
(:uuid,:cpm_id_fk,(select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:reason_for_admission,:patient_medical_hist,:treatment_recieved_before,:family_hist,:menstrual_hist,:allergies,:personal_history,:general_physical_exam,:systematic_exam,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))','update spl_hpft_patient_medical_details_tbl 
set  cpm_id_fk = :cpm_id_fk, patient_id = (select id as patient_id from spl_hpft_patient_master_tbl where uuid = :patient_uuid limit 1 ), admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), reason_for_admission = :reason_for_admission, patient_medical_hist = :patient_medical_hist, treatment_recieved_before = :treatment_recieved_before, family_hist = :family_hist, menstrual_hist = :menstrual_hist, allergies = :allergies, personal_history = :personal_history, general_physical_exam = :general_physical_exam, systematic_exam = :systematic_exam,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ" 
WHERE uuid = :uuid'),
	
	('action_txn_tbl','2018-01-01 00:00:00','select count(*) as count from spl_hpft_action_txn_tbl where uuid = ?','select count(*) as count, max(updated_on) as max_updated_on from spl_hpft_action_txn_tbl where updated_on > ?','select atxn.uuid,padmsn.uuid as admission_uuid,pconf.uuid as schedule_uuid,atxn.txn_data,atxn.txn_date,atxn.txn_state,atxn.conf_type_code,atxn.runtime_config_data,atxn.updated_on from spl_hpft_action_txn_tbl atxn
left join spl_hpft_patient_admission_tbl padmsn on padmsn.id = atxn.admission_id_fk
left join spl_hpft_patient_conf_tbl pconf on pconf.id = atxn.patient_conf_id_fk where atxn.updated_on > ?','insert into spl_hpft_action_txn_tbl 
(uuid,cpm_id_fk,patient_conf_id_fk,:admission_id_fk,txn_data,runtime_config_data,txn_date,txn_state,conf_type_code,client_updated_at) 
values 
(:uuid,:cpm_id_fk, (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1),(select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1),:txn_data,:runtime_config_data,:txn_date,:txn_state,:conf_type_code,STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ"))','update spl_hpft_action_txn_tbl set uuid = :uuid, cpm_id_fk = :cpm_id_fk, patient_conf_id_fk = (select id as patient_conf_id_fk from spl_hpft_patient_conf_tbl where uuid = :schedule_uuid limit 1) ,admission_id_fk = (select id as admission_id_fk from spl_hpft_patient_admission_tbl where uuid = :admission_uuid limit 1), txn_data = :txn_data, runtime_config_data = :runtime_config_data, txn_date = :txn_date, txn_state = :txn_state, conf_type_code = :conf_type_code,client_updated_at=STR_TO_DATE(:client_updated_at ,"%Y-%m-%dT%T.%xZ") WHERE uuid = :uuid');
	
	



