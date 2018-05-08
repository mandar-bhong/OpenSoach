package dbquery

const QUERY_EP_PROC_GET_DB_INST = `select dbi.connection_string from spl_master_database_instance_tbl as dbi
inner join spl_master_cust_prod_mapping_tbl as cpm on dbi.id = cpm.dbi_id_fk
where cpm.id = ?`

const QUERY_EP_PROC_GET_DEVICE_SP = `select sp.sp_id_fk, sp.sp_name,spc.spc_name from spl_node_dev_sp_mapping devsp
inner join spl_node_sp_tbl sp  on sp.sp_id_fk = devsp.sp_id_fk
inner join spl_node_sp_category_tbl spc on spc.id = sp.spc_id_fk
where devsp.cpm_id_fk = ? and devsp.dev_id_fk = ?`

const QUERY_EP_PROC_GET_SP_OPERATOR = `select fopcode from spl_node_field_operator_tbl as fop
inner join spl_node_fop_sp_tbl as fopsp on fop.id = fopsp.fop_id_fk
where fopsp.cpm_id_fk = ? and fopsp.sp_id_fk = ?`
