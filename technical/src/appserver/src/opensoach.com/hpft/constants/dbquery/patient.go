package dbquery

const QUERY_PATIENT_MASTER_TABLE_SELECT_BY_ID = `select * from spl_hpft_patient_master_tbl where id = ?`

const QUERY_GET_SPL_PATIENT_TOTAL_FILTERED_COUNT = `select count(*) as count
from spl_hpft_patient_master_tbl patient
left join spl_hpft_patient_admission_tbl padmsn on patient.id = padmsn.patient_id_fk
$WhereCondition$`

const QUERY_SPL_PATIENT_SELECT_BY_FILTER = `
select patient.id as patient_id_fk, padmsn.id,padmsn.cpm_id_fk,fname,lname,mob_no,bed_no,status,sp_id_fk,dr_incharge,admitted_on,discharged_on 
from spl_hpft_patient_master_tbl patient
left join spl_hpft_patient_admission_tbl padmsn on patient.id = padmsn.patient_id_fk
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_PATIENT_CONF_LIST = `select * from spl_hpft_patient_conf_tbl where cpm_id_fk=?`

const QUERY_GET_SPL_MASTER_PATIENT_TOTAL_FILTERED_COUNT = `select count(*) as count
from spl_hpft_patient_master_tbl
$WhereCondition$`

const QUERY_SPL_MASTER_PATIENT_SELECT_BY_FILTER = `select id,patient_reg_no,fname,lname,mob_no,age,blood_grp,gender,created_on,updated_on from spl_hpft_patient_master_tbl
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_PATIENT_ADMISSION_TABLE_SELECT_BY_ID = `select * from spl_hpft_patient_admission_tbl where id = ?`
