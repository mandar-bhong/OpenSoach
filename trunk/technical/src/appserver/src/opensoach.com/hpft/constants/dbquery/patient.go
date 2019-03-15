package dbquery

const QUERY_PATIENT_MASTER_TABLE_SELECT_BY_ID = `select * from spl_hpft_patient_master_tbl where id = ?`

const QUERY_GET_SPL_PATIENT_TOTAL_FILTERED_COUNT = `select count(*) as count
from spl_hpft_patient_master_tbl patient
left join spl_hpft_patient_admission_tbl padmsn on patient.id = padmsn.patient_id_fk
$WhereCondition$`

const QUERY_SPL_PATIENT_SELECT_BY_FILTER = `select patient.id as patient_id_fk, padmsn.id, padmsn.patient_reg_no,padmsn.cpm_id_fk,fname,lname,mob_no,ppd.person_accompanying->>'$.data[0].contact' as emergency_contact_no,bed_no,status,sp_id_fk,dr_incharge,admitted_on,discharged_on 
from spl_hpft_patient_master_tbl patient
left join spl_hpft_patient_admission_tbl padmsn on patient.id = padmsn.patient_id_fk
left join spl_hpft_patient_personal_details_tbl ppd on ppd.admission_id_fk = padmsn.id
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_PATIENT_CONF_LIST = `select * from spl_hpft_patient_conf_tbl where cpm_id_fk=?`

const QUERY_GET_SPL_MASTER_PATIENT_TOTAL_FILTERED_COUNT = `select count(*) as count
from spl_hpft_patient_master_tbl
$WhereCondition$`

const QUERY_SPL_MASTER_PATIENT_SELECT_BY_FILTER = `select id,patient_reg_no,fname,lname,mob_no,age,blood_grp,gender,created_on,updated_on from spl_hpft_patient_master_tbl
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_PATIENT_ADMISSION_TABLE_SELECT_BY_ID = `select * from spl_hpft_patient_admission_tbl where id = ?`

const QUERY_PATIENT_PERSONAL_DETAILS_TABLE_SELECT_BY_ID = `select * from spl_hpft_patient_personal_details_tbl where id = ?`

const QUERY_PATIENT_MEDICAL_DETAILS_TABLE_SELECT_BY_ID = `select * from spl_hpft_patient_medical_details_tbl where id = ?`

const QUERY_PATIENT_ADMISSION_TABLE_STATUS_SELECT_BY_ID = `select padmsn.status,patient.id from spl_hpft_patient_master_tbl patient
left join spl_hpft_patient_admission_tbl padmsn on padmsn.patient_id_fk = patient.id
where patient.id = ? order by admitted_on desc limit 1`

const QUERY_GET_PATIENT_PERSONAL_DETAILS_BY_ADMISSION_ID = `select * from spl_hpft_patient_personal_details_tbl where admission_id_fk = ?`

const QUERY_GET_PATIENT_MEDICAL_DETAILS_BY_ADMISSION_ID = `select * from spl_hpft_patient_medical_details_tbl where admission_id_fk = ?`

const QUERY_GET_SPL_PATIENT_ACTION_TXN_TOTAL_FILTERED_COUNT = `select count(*) as count
from spl_hpft_action_txn_tbl actn_txn
inner join spl_hpft_patient_conf_tbl pconf on pconf.id = actn_txn.patient_conf_id_fk
$WhereCondition$`

const QUERY_SPL_PATIENT_ACTION_TXN_SELECT_BY_FILTER = `select actn_txn.admission_id_fk,actn_txn.patient_conf_id_fk,txn_date,txn_data,txn_state,actn_txn.conf_type_code,actn_txn.updated_by,pconf.conf->>'$.name' as action_name from spl_hpft_action_txn_tbl as actn_txn
inner join spl_hpft_patient_conf_tbl pconf on pconf.id = actn_txn.patient_conf_id_fk
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_SELECT_PATIENT_USER_INFO = `select fname,lname from spl_master_usr_details_tbl where usr_id_fk = ?`

const QUERY_SELECT_PATIENT_CONF_BY_ID = `select * from spl_hpft_patient_conf_tbl where id = ?`
const QUERY_SELECT_PATIENT_DOCTORS_ORDERS_BY_ID = `select * from spl_hpft_doctors_orders_tbl where id = ?`
const QUERY_SELECT_PATIENT_PATHOLGY_RECORDS_BY_ID = `select * from spl_hpft_pathology_record_tbl where id = ?`
const QUERY_SELECT_PATIENT_TREATMENT_BY_ID = `select * from spl_hpft_treatment_tbl where id = ?`
const QUERY_SELECT_PATIENT_DOCTORS_ORDERS_BY_ADMISSION_ID = `select * from spl_hpft_doctors_orders_tbl where admission_id_fk = ?`

const QUERY_GET_SPL_PATIENT_DOCTORS_ORDERS_TOTAL_FILTERED_COUNT = `select count(*) as count from spl_hpft_doctors_orders_tbl dorders
left join spl_hpft_document_tbl doc on doc.id = dorders.document_id_fk $WhereCondition$`

const QUERY_SPL_PATIENT_DOCTORS_ORDERS_SELECT_BY_FILTER = `select dorders.id,dorders.uuid,admission_id_fk,doctor_id_fk,doctors_orders,ack_by,ack_time,comment,order_created_time,order_type,document_id_fk,doc.uuid as document_uuid,doc.name 
from spl_hpft_doctors_orders_tbl dorders 
left join spl_hpft_document_tbl doc on doc.id = dorders.document_id_fk
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_SPL_PATIENT_TREATMENT_TOTAL_FILTERED_COUNT = `SELECT count(*) as count FROM spl_hpft_treatment_tbl trtmnt $WhereCondition$`

const QUERY_SPL_PATIENT_TREATMENT_SELECT_BY_FILTER = `select id,admission_id_fk,treatment_done,treatment_performed_time,details,post_observation,created_on 
from spl_hpft_treatment_tbl 
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_SPL_PATIENT_PATHOLOGY_RECORD_TOTAL_FILTERED_COUNT = `SELECT count(*) as count FROM spl_hpft_pathology_record_tbl $WhereCondition$`

const QUERY_SPL_PATIENT_PATHOLOGY_RECORD_SELECT_BY_FILTER = `select id,admission_id_fk,test_performed,test_performed_time,test_result,comments,created_on 
from spl_hpft_pathology_record_tbl 
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_SELECT_TREATMENT_DOCUMENTS_BY_TREATMENT_ID = `select doc.uuid, doc.name from spl_hpft_treatment_tbl trtmnt 
inner join spl_hpft_treatment_doc_tbl tdoc on tdoc.treatment_id_fk = trtmnt.id
inner join spl_hpft_document_tbl doc on doc.id = tdoc.document_id_fk
where trtmnt.id = ?`

const QUERY_SELECT_PATHOLOGY_REOCORD_DOCUMENTS_BY_ID = `select doc.uuid, doc.name from spl_hpft_pathology_record_tbl prec 
inner join spl_hpft_pathology_record_doc_tbl precdoc on precdoc.pathology_id_fk = prec.id
inner join spl_hpft_document_tbl doc on doc.id = precdoc.document_id_fk
where prec.id = ?`

const QUERY_GET_SPL_PATIENT_CONF_TOTAL_FILTERED_COUNT = `select count(*) as count from spl_hpft_patient_conf_tbl $WhereCondition$`

const QUERY_SPL_PATIENT_CONF_SELECT_BY_FILTER = `select id,admission_id_fk,conf_type_code,conf,end_date,status 
from spl_hpft_patient_conf_tbl 
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_DOCUMENT_BY_DOCUMENT_UUID = "select * from spl_hpft_document_tbl where uuid = ?"

const QUERY_SELECT_PATIENT_TREATMENT_BY_ADMISSION_ID = `select * from spl_hpft_treatment_tbl where admission_id_fk = ?`
const QUERY_SELECT_PATIENT_PATHOLOGY_RECORDS_BY_ADMISSION_ID = `select * from spl_hpft_pathology_record_tbl where admission_id_fk = ?`
