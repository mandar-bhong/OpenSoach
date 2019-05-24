package dbquery

const QUERY_SPL_MASTER_PRODUCT_TBL_SELECT_ALL = `SELECT id,prod_code,created_on,updated_on FROM spl_master_product_tbl`

const QUERY_SPL_MASTER_DATABASE_INSTANCE_SELECT_ALL = `SELECT id,dbi_name,connection_string,prod_id_fk,created_on,updated_on FROM spl_master_database_instance_tbl`
