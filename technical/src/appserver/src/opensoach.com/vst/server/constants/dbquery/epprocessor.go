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

const QUERY_EP_PROC_GET_VEHICLE_ID_BY_NAME = `select id from spl_vst_vehicle_master_tbl where vehicle_no = ?`
const QUERY_EP_PROC_GET_LAST_VEHICLE_RECORD = `select * from spl_vst_token order by id DESC limit 1`
const QUERY_EP_PROC_GET_TOKEN_MAPPING_DETAILS_BY_ID = `select mapping_details from spl_vst_token where id = ?`

const QUERY_EP_PROC_GET_VHL_TOKEN_LIST = `select token.id,token,vhl_id_fk,vehicle_no,state,generated_on from spl_vst_token token
inner join spl_vst_vehicle_master_tbl vehicle on token.vhl_id_fk = vehicle.id`

const QUERY_EP_PROC_GET_VHL_TOKEN_BY_TOKEN_ID = `select token.id,token,vhl_id_fk,vehicle_no,state,generated_on from spl_vst_token token
inner join spl_vst_vehicle_master_tbl vehicle on token.vhl_id_fk = vehicle.id where token.id = ?`
