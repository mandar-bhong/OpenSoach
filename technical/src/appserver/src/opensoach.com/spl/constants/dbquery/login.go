package dbquery

const QUERY_MUST_CHECK_USER_LOGIN = "SELECT id,usr_name,usr_password,usr_category,usr_state,urole_id_fk FROM spl_master_user_tbl WHERE usr_name = ? AND usr_password = ?"

const QUERY_GET_USER_AUTH_INFO = `select cpm.id as cpm_id, cpm.cust_id_fk,role.urole_code,di.connection_string from 
spl_master_cust_prod_mapping_tbl cpm
INNER JOIN spl_master_product_tbl prod  ON cpm.prod_id_fk = prod.id
INNER JOIN spl_master_usr_cpm_tbl usrcpm ON  usrcpm.cpm_id_fk = cpm.id
INNER JOIN spl_master_user_role_tbl role ON usrcpm.urole_id_fk = role.id
INNER JOIN spl_master_database_instance_tbl di ON cpm.dbi_id_fk = di.id
where prod.prod_code = ? AND usrcpm.user_id_fk = ?`

const QUERY_GET_USER_LOGIN_INFO = `SELECT fname,lname FROM spl_master_usr_details_tbl where usr_id_fk = ?`

const QUERY_GET_CUSTOMER_LOGIN_INFO = `SELECT corp_name,cust_name FROM spl_master_customer_tbl cust
INNER JOIN spl_master_corp_tbl corp ON corp.id = cust.corp_id_fk
WHERE cust.id = ?`
