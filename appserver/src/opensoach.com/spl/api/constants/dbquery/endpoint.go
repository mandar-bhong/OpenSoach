package dbquery

const QUERY_GET_DEVICE_INFO_BY_DEVICE_SERIAL_NO = `Select id,dev_state from spl_master_device_tbl where serialno = ?`
const QUERY_GET_DEVICE_AUTH_INFO = `Select cpm.id as id ,sreg.server_address as server_address,dbi.connection_string From spl_master_cust_prod_mapping_tbl cpm
Inner join spl_master_cpm_dev_mapping_tbl cpmd on cpmd.cpm_id_fk = cpm.id
Inner join spl_master_server_register sreg on sreg.prod_id_fk = cpm.prod_id_fk
Inner join spl_master_product_tbl prod on prod.id = cpm.prod_id_fk
Inner join spl_master_database_instance_tbl dbi on dbi.prod_id_fk = prod.id
where dev_id_fk = ? and prod.prod_code = ?`

const QUERY_GET_DEVICE_USER_INFO = `select cpm.id as cpm_id, cpm.cust_id_fk,usrcpm.urole_id_fk,role.urole_code,dbinst.connection_string,sreg.server_address from 
spl_master_cust_prod_mapping_tbl cpm
INNER JOIN spl_master_product_tbl prod  ON cpm.prod_id_fk = prod.id
INNER JOIN spl_master_usr_cpm_tbl usrcpm ON  usrcpm.cpm_id_fk = cpm.id
INNER JOIN spl_master_user_role_tbl role ON usrcpm.urole_id_fk = role.id
INNER JOIN spl_master_database_instance_tbl dbinst ON cpm.dbi_id_fk = dbinst.id
inner join spl_master_server_register sreg on sreg.prod_id_fk = prod.id
where prod.prod_code = ? AND usrcpm.user_id_fk = ?`

const QUERY_GET_DEVCIE_USER_LIST_DATA = `select usr.id as usr_id,usr.usr_name,urole.urole_name,usrd.fname,usrd.lname from spl_master_user_tbl usr
left join spl_master_usr_details_tbl usrd on usr.id = usrd.usr_id_fk
left join spl_master_usr_cpm_tbl ucpm on usr.id = ucpm.user_id_fk 
left join spl_master_user_role_tbl urole on urole.id = ucpm.urole_id_fk
where ucpm.cpm_id_fk = ?`

const QUERY_GET_DEVICE_SHARED_USER_INFO = `select usr.id,usr_name,fname,lname 
from spl_master_user_tbl usr
left join spl_master_usr_details_tbl usrd on usr.id = usrd.usr_id_fk 
inner join spl_master_usr_cpm_tbl ucpm on usr.id = ucpm.user_id_fk where usr.usr_name = ? and usr.usr_password = ? and ucpm.cpm_id_fk = ?`

const QUERY_GET_DEVICE_USER_CPM_LIST = `select ucpm.cpm_id_fk,prod.prod_code from spl_master_user_tbl usr 
left join spl_master_usr_cpm_tbl ucpm on ucpm.user_id_fk = usr.id
left join spl_master_cust_prod_mapping_tbl cpm on ucpm.cpm_id_fk = cpm.id
left join spl_master_product_tbl prod on cpm.prod_id_fk = prod.id
where usr.usr_name = ?`
