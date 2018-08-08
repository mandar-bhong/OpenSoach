package dbquery

const QUERY_GET_PATIENT_LIST = `select id,patient_details,medical_details,patient_file_template,status,discharged_on,patient.sp_id_fk,sp.sp_name
from spl_hpft_patient_master_tbl patient
inner join spl_node_sp_tbl sp on sp.sp_id_fk = patient.sp_id_fk
where patient.cpm_id_fk=?`

const QUERY_PATIENT_MASTER_TABLE_SELECT_BY_ID = `select * from spl_hpft_patient_master_tbl where id = ?`
