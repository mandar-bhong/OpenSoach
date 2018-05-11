package dbquery

const QUERY_SPL_MASTER_CUSTOMER_TABLE_INSERT = "INSERT INTO spl_master_customer_tbl (corp_id_fk,cust_name,cust_state,cust_state_since) values ( :corp_id_fk,:cust_name,:cust_state,:cust_state_since)"

const QUERY_SPL_MASTER_CUST_DETAILS_TABLE_SELECT_All = "SELECT cust_id_fk,poc1_name,poc1_email_id,poc1_mobile_no,poc2_name,poc2_email_id,poc2_mobile_no,address,address_state,address_city,address_pincode FROM spl_master_cust_details_tbl"

const QUERY_SPL_MASTER_CUST_DETAILS_TABLE_SELECT_BY_ID = "SELECT cust_id_fk,poc1_name,poc1_email_id,poc1_mobile_no,poc2_name,poc2_email_id,poc2_mobile_no,address,address_state,address_city,address_pincode,created_on,updated_on FROM spl_master_cust_details_tbl WHERE cust_id_fk =?"
const QUERY_SPL_CU_MASTER_CUST_DETAILS_TABLE_SELECT_BY_ID = "SELECT cust_id_fk from spl_master_cust_details_tbl where cust_id_fk =?"
const QUERY_SPL_MASTER_CUST_DETAILS_TABLE_INSERT = "INSERT INTO spl_master_cust_details_tbl (cust_id_fk,poc1_name,poc1_email_id,poc1_mobile_no,poc2_name,poc2_email_id,poc2_mobile_no,address,address_state,address_city,address_pincode) values (:cust_id_fk,:poc1_name,:poc1_email_id,:poc1_mobile_no,:poc2_name,:poc2_email_id,:poc2_mobile_no,:address,:address_state,:address_city,:address_pincode)"

const QUERY_SPL_MASTER_CUST_DETAILS_TABLE_UPDATE = "UPDATE spl_master_cust_details_tbl SET poc1_name = :poc1_name, poc1_email_id = :poc1_email_id, poc1_mobile_no = :poc1_mobile_no, poc2_name = :poc2_name, poc2_email_id = :poc2_email_id, poc2_mobile_no = :poc2_mobile_no, address = :address, address_state = :address_state, address_city = :address_city, address_pincode = :address_pincode WHERE cust_id_fk = :cust_id_fk"

const QUERY_GET_CUSTOMER_TABLE_INFO_BY_ID = `Select id,corp_id_fk,cust_name,cust_state,cust_state_since,created_on,updated_on 
											FROM spl_master_customer_tbl 
											WHERE id = ?`

const QUERY_GET_CORP_TABLE_INFO_BY_CUSTOMER_ID = `SELECT corp.id,corp_name,corp_mobile_no,corp_email_id,corp_landline_no,corp.created_on,corp.updated_on from spl_master_corp_tbl corp
											INNER JOIN spl_master_customer_tbl cust  ON corp.id = cust.corp_id_fk
											WHERE cust.id = ?`

const QUERY_SPL_MASTER_CUSTOMER_TABLE_SELECT_All = "SELECT id,corp_id_fk,cust_name,cust_state,cust_state_since,created_on,updated_on FROM spl_master_customer_tbl"

const QUERY_SPL_MASTER_CUSTOMER_TABLE_SELECT_BY_FILTER = `SELECT cust.id as id,cust.corp_id_fk as corp_id_fk ,cust.cust_name as cust_name ,cust.cust_state as cust_state,corp.corp_name as corp_name ,custdetails.poc1_name as poc1_name,custdetails.poc1_email_id as poc1_email_id,custdetails.poc1_mobile_no as poc1_mobile_no ,cust.created_on as created_on,cust.updated_on as updated_on FROM spl_master_customer_tbl cust INNER JOIN spl_master_corp_tbl corp ON corp.id = cust.corp_id_fk 
LEFT JOIN spl_master_cust_prod_mapping_tbl cpm ON cust.id = cpm.id
LEFT JOIN spl_master_product_tbl prod ON prod.id = cpm.prod_id_fk
LEFT JOIN spl_master_cust_details_tbl custdetails ON custdetails.cust_id_fk = cust.id $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_SPL_MASTER_CUSTOMER_TABLE_TOTAL_FILTERED_COUNT = `SELECT count(*) as count FROM spl_master_customer_tbl cust INNER JOIN spl_master_corp_tbl corp ON corp.id = cust.corp_id_fk 
LEFT JOIN spl_master_cust_prod_mapping_tbl cpm ON cust.id = cpm.id
LEFT JOIN spl_master_product_tbl prod ON prod.id = cpm.prod_id_fk
LEFT JOIN spl_master_cust_details_tbl custdetails ON custdetails.cust_id_fk = cust.id $WhereCondition$`

const QUERY_GET_PRODUCT_ASSOCIATION_BY_CUST_ID = `Select cpm.id as id,cpm.cpm_state as cpm_state,cpm.prod_id_fk as prod_id_fk,prod.prod_code as prod_code,cpm.dbi_id_fk as dbi_id_fk,dbi.dbi_name as dbi_name From spl_master_cust_prod_mapping_tbl cpm
Inner Join spl_master_product_tbl prod on cpm.prod_id_fk = prod.id
Inner Join spl_master_database_instance_tbl dbi on cpm.dbi_id_fk= dbi.id
Where cust_id_fk = ?`

const QUERY_SPL_MASTER_CUST_TABLE_SELECT_SHORT_DATA_LIST = `select id,cust_name from spl_master_customer_tbl`
