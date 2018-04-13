package dbquery

const QUERY_SPL_MASTER_USER_TABLE_INSERT = "INSERT INTO spl_master_user_tbl (usr_name,usr_password,usr_category,urole_id_fk,usr_state,usr_state_since) values ( :usr_name,:usr_password,:usr_category,:urole_id_fk,:usr_state,:usr_state_since)"

const QUERY_SPL_MASTER_USR_DETAILS_TABLE_SELECT_BY_ID = "SELECT fname,lname,mobile_no,alternate_contact_no,created_on,updated_on FROM spl_master_usr_details_tbl WHERE usr_id_fk = ?"

const QUERY_SPL_MASTER_USR_DETAILS_TABLE_INSERT = "INSERT INTO spl_master_usr_details_tbl (usr_id_fk,fname,lname,mobile_no,alternate_contact_no) values (:usr_id_fk,:fname,:lname,:mobile_no,:alternate_contact_no)"

const QUERY_SPL_MASTER_USR_DETAILS_TABLE_UPDATE = "UPDATE spl_master_usr_details_tbl SET fname = :fname, lname = :lname, mobile_no = :mobile_no, alternate_contact_no = :alternate_contact_no, created_on = :created_on, updated_on = :updated_on WHERE usr_id_fk = :usr_id_fk"

const QUERY_SPL_MASTER_USER_TABLE_UPDATE_STATE = "UPDATE spl_master_user_tbl SET usr_state = :usr_state WHERE id = :id"

const QUERY_SPL_MASTER_USER_TABLE_SELECT_BY_ID_PASSWORD = "SELECT usr_name,usr_password,usr_category,urole_id_fk,usr_state,usr_state_since,created_on,updated_on FROM spl_master_user_tbl WHERE id = ? AND usr_password = ?"

const QUERY_SPL_MASTER_USER_TABLE_CHANGE_PASSWORD = "UPDATE spl_master_user_tbl SET usr_password = :usr_password WHERE id = :id"
