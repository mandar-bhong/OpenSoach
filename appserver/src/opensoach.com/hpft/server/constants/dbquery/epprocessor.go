package dbquery

const QUERY_EP_PROC_GET_DB_INST = `select dbi.connection_string from spl_master_database_instance_tbl as dbi
inner join spl_master_cust_prod_mapping_tbl as cpm on dbi.id = cpm.dbi_id_fk
where cpm.id = ?`

const QUERY_EP_PROC_GET_DEVICE_SP = `select sp.sp_id_fk, sp.sp_name,spc.spc_name from spl_node_dev_sp_mapping devsp
inner join spl_node_sp_tbl sp  on sp.sp_id_fk = devsp.sp_id_fk
inner join spl_node_sp_category_tbl spc on spc.id = sp.spc_id_fk
where devsp.cpm_id_fk = ? and devsp.dev_id_fk = ?`

const QUERY_EP_PROC_GET_SP_OPERATOR = `select fopcode from spl_node_field_operator_tbl as fop
where fop.fop_area =1
union
select fopcode from spl_node_field_operator_tbl as fop
inner join spl_node_fop_sp_tbl as fopsp on fop.id = fopsp.fop_id_fk
where fopsp.cpm_id_fk = ? and fopsp.sp_id_fk = ?`

const QUERY_EP_PROC_GET_SP_SERV_CONF = `select serv_conf_in.id as id,serv_conf_in.serv_conf_id_fk as serv_conf_id_fk,serv_conf.conf_type_code as conf_type_code,serv_conf.serv_conf_name as serv_conf_name,serv_conf.serv_conf as serv_conf From spl_node_service_instance_tbl serv_conf_in
inner join spl_node_service_conf_tbl serv_conf on serv_conf.id = serv_conf_in.serv_conf_id_fk
where serv_conf_in.cpm_id_fk = ? and serv_conf_in.sp_id_fk = ?`

const QUERY_EP_PROC_GET_SP_PATIENT_CONF = `select serv_conf_in.id as id,serv_conf_in.serv_conf_id_fk as serv_conf_id_fk,serv_conf.conf_type_code as conf_type_code,serv_conf.serv_conf_name as serv_conf_name,serv_conf.serv_conf as serv_conf,patient_details,medical_details
From spl_node_service_instance_tbl serv_conf_in
inner join spl_node_service_conf_tbl serv_conf on serv_conf.id = serv_conf_in.serv_conf_id_fk
inner join spl_hpft_patient_master_tbl patient on patient.serv_in_id_fk = serv_conf_in.id
where serv_conf_in.cpm_id_fk = ? and serv_conf_in.sp_id_fk = ? and patient.status = 1`
