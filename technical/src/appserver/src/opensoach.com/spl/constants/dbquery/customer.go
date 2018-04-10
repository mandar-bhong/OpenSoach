package dbquery

const QUERY_GET_CUSTOMER_INFO_BY_ID = `Select id,corp_id_fk,cust_name,cust_state,cust_state_since,created_on,updated_on 
										FROM spl_master_customer_tbl 
										WHERE id = ?`
