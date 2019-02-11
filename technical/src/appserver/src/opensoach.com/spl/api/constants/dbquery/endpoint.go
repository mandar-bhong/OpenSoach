package dbquery

const QUERY_GET_DEVICE_INFO_BY_DEVICE_SERIAL_NO = `Select id,dev_state from spl_master_device_tbl where serialno = ?`
const QUERY_GET_DEVICE_AUTH_INFO = `Select cpm.id as id ,sreg.server_address as server_address From spl_master_cust_prod_mapping_tbl cpm
 Inner join spl_master_cpm_dev_mapping_tbl cpmd on cpmd.cpm_id_fk = cpm.id
 Inner join spl_master_server_register sreg on sreg.prod_id_fk = cpm.prod_id_fk
 Inner join spl_master_product_tbl prod on prod.id = cpm.prod_id_fk
 where dev_id_fk = ? and prod.prod_code = ?`

const QUERY_GET_DEVICE_USER_INFO = `select usr.id,usr_name,fname,lname 
from spl_master_user_tbl usr
left join spl_master_usr_details_tbl usrd on usr.id = usrd.usr_id_fk 
inner join spl_master_usr_cpm_tbl ucpm on usr.id = ucpm.user_id_fk where usr.usr_name = ? and usr.usr_password = ? and ucpm.cpm_id_fk = ?`
