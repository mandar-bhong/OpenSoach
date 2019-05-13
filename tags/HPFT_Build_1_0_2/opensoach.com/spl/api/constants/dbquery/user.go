package dbquery

const QUERY_SPL_MASTER_USER_TABLE_INSERT = "INSERT INTO spl_master_user_tbl (usr_name,usr_password,usr_category,urole_id_fk,usr_state,usr_state_since) values ( :usr_name,:usr_password,:usr_category,:urole_id_fk,:usr_state,:usr_state_since)"

const QUERY_SPL_MASTER_USR_DETAILS_TABLE_SELECT_BY_ID = "SELECT fname,lname,mobile_no,alternate_contact_no,created_on,updated_on FROM spl_master_usr_details_tbl WHERE usr_id_fk = ?"

const QUERY_SPL_MASTER_USR_DETAILS_TABLE_INSERT = "INSERT INTO spl_master_usr_details_tbl (usr_id_fk,fname,lname,mobile_no,alternate_contact_no) values (:usr_id_fk,:fname,:lname,:mobile_no,:alternate_contact_no)"

const QUERY_SPL_MASTER_USR_DETAILS_TABLE_UPDATE = "UPDATE spl_master_usr_details_tbl SET fname = :fname, lname = :lname, mobile_no = :mobile_no, alternate_contact_no = :alternate_contact_no, created_on = :created_on, updated_on = :updated_on WHERE usr_id_fk = :usr_id_fk"

const QUERY_SPL_MASTER_USER_TABLE_UPDATE_STATE = "UPDATE spl_master_user_tbl SET usr_state = :usr_state, usr_state_since = :usr_state_since WHERE id = :id"

const QUERY_SPL_MASTER_USER_TABLE_SELECT_BY_ID_PASSWORD = "SELECT usr_name,usr_password,usr_category,urole_id_fk,usr_state,usr_state_since FROM spl_master_user_tbl WHERE id = ? AND usr_password = ?"

const QUERY_SPL_MASTER_USER_TABLE_CHANGE_PASSWORD = "UPDATE spl_master_user_tbl SET usr_password = :usr_password WHERE id = :id"

const QUERY_CU_GET_SPL_MASTER_USER_TABLE_TOTAL_FILTERED_COUNT = `SELECT count(*) as count FROM spl_master_usr_cpm_tbl ucpm INNER JOIN spl_master_user_tbl as usr on ucpm.user_id_fk = usr.id Left Join spl_master_usr_details_tbl usrd on usrd.usr_id_fk = usr.id Left Join spl_master_user_role_tbl urole on urole.id = ucpm.urole_id_fk $WhereCondition$`

const QUERY_CU_SPL_MASTER_USER_TABLE_SELECT_BY_FILTER = `SELECT usr.id as id,usr.usr_name,usrd.fname as fname,usrd.lname as lname,usrd.mobile_no as mobile_no,usr.usr_category,ucpm.urole_id_fk,urole.urole_name as urole_name,usr.usr_state,usr.usr_state_since,usr.created_on,usr.updated_on 
FROM spl_master_usr_cpm_tbl ucpm 
INNER JOIN spl_master_user_tbl as usr on ucpm.user_id_fk = usr.id 
Left Join spl_master_usr_details_tbl usrd on usrd.usr_id_fk = usr.id 
Left Join spl_master_user_role_tbl urole on urole.id = ucpm.urole_id_fk $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_OSU_GET_SPL_MASTER_USER_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From spl_master_user_tbl usr Left Join spl_master_usr_details_tbl usrd On usrd.usr_id_fk = usr.id Left Join spl_master_user_role_tbl urole On urole.id = usr.urole_id_fk $WhereCondition$`

const QUERY_OSU_SPL_MASTER_USER_TABLE_SELECT_BY_FILTER = `Select usr.id as id,usr.usr_name as usr_name ,usrd.fname as fname,usrd.lname as lname,usrd.mobile_no as mobile_no,usr.usr_category as usr_category,usr.urole_id_fk as urole_id_fk,urole.urole_name as urole_name,usr.usr_state as usr_state,usr.usr_state_since as usr_state_since,usr.created_on as created_on,usr.updated_on as updated_on 
From spl_master_user_tbl usr 
Left Join spl_master_usr_details_tbl usrd On usrd.usr_id_fk = usr.id 
Left Join spl_master_user_role_tbl urole On urole.id = usr.urole_id_fk $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_USERID_BY_USERNAME = `Select id From spl_master_user_tbl where usr_name = ? and usr_category = 2`

const QUERY_GET_USER_TABLE_INFO_BY_ID = `Select id,usr_name,usr_category,urole_id_fk,usr_state,usr_state_since,created_on,updated_on  From spl_master_user_tbl where id =?`

const QUERY_GET_CU_USER_TABLE_INFO_BY_ID = `select usr.id as id,usr_name,usr_category,usr_state,usr_state_since,ucpm.urole_id_fk,urole.urole_name,usr.created_on,usr.updated_on 
from spl_master_user_tbl usr 
inner join spl_master_usr_cpm_tbl ucpm on usr.id = ucpm.user_id_fk
inner join spl_master_user_role_tbl urole on ucpm.urole_id_fk = urole.id
where usr.id = ? and ucpm.cpm_id_fk= ?`

const QUERY_GET_SPL_MASTER_USER_DETAILS_TABLE_SELECT_BY_ID = `Select usr_id_fk,fname,lname,gender,mobile_no,alternate_contact_no,created_on,updated_on From spl_master_usr_details_tbl where usr_id_fk = ?`

const QUERY_GET_UROLE_LIST_OSU = `Select urole.id as id,urole.urole_code as urole_code,urole.urole_name as urole_name,prod.prod_code as prod_code From spl_master_user_role_tbl urole
Left Join spl_master_product_tbl prod on prod.id=urole.prod_id_fk`

const QUERY_GET_UROLE_LIST = `Select urole.id as id ,urole.urole_code as urole_code,urole.urole_name as urole_name From spl_master_user_role_tbl urole
Inner Join spl_master_product_tbl prod On prod.id = urole.prod_id_fk
Where prod_code = ?`

const QUERY_GET_PRODUCT_ASSOCIATION_BY_USER_ID = `Select ucpm.id as id,cust.cust_name,prod.prod_code,urole.urole_code as urole_code,ucpm.ucpm_state as ucpm_state,ucpm.ucpm_state_since as ucpm_state_since From spl_master_usr_cpm_tbl ucpm
Left Join spl_master_cust_prod_mapping_tbl cpm on cpm.id = ucpm.cpm_id_fk
Left Join spl_master_customer_tbl cust on cust.id = cpm.cust_id_fk
Left Join spl_master_product_tbl prod on prod.id = cpm.prod_id_fk
Left Join spl_master_user_role_tbl urole on urole.id =  ucpm.urole_id_fk
Where ucpm.user_id_fk = ?`

const QUERY_GET_USER_ID_BY_ACTIVATION_CODE = `select id,usr_id_fk from spl_master_usr_activation_tbl where code = ?`

const QUERY_DELETE_USER_ACTIVATION_TABLE_ROW = `Delete From spl_master_usr_activation_tbl Where code = :code`

const QUERY_DELETE_USER_OTP_TABLE_ROW = `Delete From spl_master_usr_otp_tbl Where otp = :otp`

const QUERY_GET_USR_OTP_TBL_BY_OTP = `Select * From spl_master_usr_otp_tbl where otp = ? and usr_name = ?`
