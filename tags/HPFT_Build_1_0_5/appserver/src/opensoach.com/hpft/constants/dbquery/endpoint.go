package dbquery

const QUERY_DELETE_USER_PATIENT_ASSOCIATION = `delete from spl_hpft_user_patient_monitor_mapping where usr_id_fk = :usr_id_fk and sp_id_fk = :sp_id_fk and patient_id_fk = :patient_id_fk and cpm_id_fk = :cpm_id_fk`

const QUERY_DELETE_USER_PATIENT_ASSOCIATION_BY_SP = `delete from spl_hpft_user_patient_monitor_mapping where usr_id_fk = :usr_id_fk and sp_id_fk = :sp_id_fk and cpm_id_fk = :cpm_id_fk`

const QUERY_DELETE_USER_PATIENT_ASSOCIATION_BY_USER_ID = `delete from spl_hpft_user_patient_monitor_mapping where usr_id_fk = :usr_id_fk and cpm_id_fk = :cpm_id_fk`

const QUERY_SELECT_USER_PATIENT_ASSOCIATION_BY_USER_ID = `select * from spl_hpft_user_patient_monitor_mapping where usr_id_fk = ?`

const QUERY_SELECT_USER_PATIENT_ASSOCIATION_BY_USER_ID_SP_ID = `select * from spl_hpft_user_patient_monitor_mapping where usr_id_fk = ? and sp_id_fk = ?`

const QUERY_DELETE_USER_PATIENT_ASSOCIATION_BY_PATIENT_ID = `delete from spl_hpft_user_patient_monitor_mapping where patient_id_fk = :patient_id_fk and cpm_id_fk = :cpm_id_fk`
