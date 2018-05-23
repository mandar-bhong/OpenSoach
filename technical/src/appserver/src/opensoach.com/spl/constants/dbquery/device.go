package dbquery

const QUERY_SPL_MASTER_DEVICE_DETAILS_TABLE_SELECT_BY_ID = `SELECT * FROM spl_master_dev_details_tbl WHERE dev_id_fk =?`

const QUERY_GET_SPL_MASTER_DEVICE_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From spl_master_device_tbl dev
Left Join spl_master_customer_tbl  cust on cust.id = dev.cust_id_fk
Left Join spl_master_dev_details_tbl devd on devd.dev_id_fk = dev.id
Left Join spl_master_dev_status_tbl devstate on devstate.dev_id_fk = dev.id
Left Join spl_master_cpm_dev_mapping_tbl cpmd on cpmd.dev_id_fk = dev.id
Left Join spl_master_cust_prod_mapping_tbl cpm on cpm.id = cpmd.cpm_id_fk
$WhereCondition$`

const QUERY_SPL_MASTER_DEVICE_TABLE_SELECT_BY_FILTER = `Select dev.id as id,dev.cust_id_fk as cust_id_fk,cust.cust_name as cust_name,dev.serialno as serialno,dev.dev_state as dev_state,dev.dev_state_since as dev_state_since,dev.created_on as created_on,dev.updated_on as updated_on,devstate.connection_state as connection_state, devstate.connection_state_since as connection_state_since,devstate.sync_state as sync_state,devstate.sync_state_since as sync_state_since,devstate.battery_level as battery_level,devstate.battery_level_since 
From spl_master_device_tbl dev
Left Join spl_master_customer_tbl  cust on cust.id = dev.cust_id_fk
Left Join spl_master_dev_details_tbl devd on devd.dev_id_fk = dev.id
Left Join spl_master_dev_status_tbl devstate on devstate.dev_id_fk = dev.id
Left Join spl_master_cpm_dev_mapping_tbl cpmd on cpmd.dev_id_fk = dev.id
Left Join spl_master_cust_prod_mapping_tbl cpm on cpm.id = cpmd.cpm_id_fk
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_DEV_ID_BY_CPM_ID = `Select dev_id_fk From spl_master_cpm_dev_mapping_tbl where cpm_id_fk = ? and dev_id_fk = ?`

const QUERY_GET_MASTER_DEVICE_TABLE_BY_ID = `Select id,cust_id_fk,serialno,dev_state,dev_state_since,created_on,updated_on From spl_master_device_tbl Where id = ?`

const QUERY_GET_PRODUCT_ASSOCIATION_BY_DEVICE_ID = `select cust.cust_name,prod.prod_code From spl_master_cpm_dev_mapping_tbl cpdm
inner join spl_master_cust_prod_mapping_tbl cpm on cpm.id=cpdm.cpm_id_fk
inner join spl_master_customer_tbl cust on cust.id = cpm.cust_id_fk
inner join spl_master_product_tbl prod on prod.id = cpm.prod_id_fk
where cpdm.dev_id_fk=?`

const QUERY_GET_CUST_ID_BY_CPM_ID = `Select cust_id_fk From spl_master_cust_prod_mapping_tbl where id = ?`

const QUERY_SPL_MASTER_DEVICE_TABLE_SELECT_SHORT_DATA_LIST = `select dev.id,dev.serialno from spl_master_device_tbl`
