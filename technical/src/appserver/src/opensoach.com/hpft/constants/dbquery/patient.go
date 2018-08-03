package dbquery

const QUERY_GET_PATIENT_LIST = `select id,patient_details,medical_details,patient_file_template,status,sp_id_fk from spl_hpft_patient_master_tbl where cpm_id_fk=?`
