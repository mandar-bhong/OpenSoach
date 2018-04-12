package dbquery

const QUERY_GET_CUSTOMER_TABLE_INFO_BY_ID = `Select id,corp_id_fk,cust_name,cust_state,cust_state_since,created_on,updated_on 
											FROM spl_master_customer_tbl 
											WHERE id = ?`

const QUERY_GET_CUSTOMER_DETAILS_TABLE_INFO_BY_ID = `SELECT cust_id_fk,poc1_name,poc1_email_id,poc1_mobile_no,poc2_name,poc2_email_id,poc2_mobile_no,address,address_state,address_city,address_pincode 
											FROM spl_master_cust_details_tbl
											WHERE cust_id_fk = ?`

const QUERY_GET_CORP_TABLE_INFO_BY_CUSTOMER_ID = `SELECT corp.id,corp_name,corp_mobile_no,corp_email_id,corp_landline_no,corp.created_on,corp.updated_on from spl_master_corp_tbl corp
											INNER JOIN spl_master_customer_tbl cust  ON corp.id = cust.corp_id_fk
											WHERE cust.id = ?`
