package dbquery

const QUERY_GET_SER_CONFIG_BY_SER_CONF_INS_ID = `select serv_conf_in.id as serv_conf_in_id, 
devsp.dev_id_fk, devsp.sp_id_fk,serv_conf_in.serv_conf_id_fk,serv_conf.conf_type_code,
serv_conf.serv_conf_name, serv_conf.serv_conf 
from spl_node_dev_sp_mapping devsp 
inner join spl_node_service_instance_tbl serv_conf_in 
on devsp.sp_id_fk = serv_conf_in.sp_id_fk 
and devsp.cpm_id_fk = serv_conf_in.cpm_id_fk
inner join spl_node_service_conf_tbl serv_conf
on serv_conf.id = serv_conf_in.serv_conf_id_fk
where serv_conf_in.id=?`

const QUERY_GET_SER_CONFIG_BY_CPM_DEV_SP = `select serv_conf_in.id as serv_conf_in_id, 
devsp.dev_id_fk, devsp.sp_id_fk, 
serv_conf_in.serv_conf_id_fk,
serv_conf.conf_type_code,serv_conf.serv_conf_name,
serv_conf.serv_conf from spl_node_dev_sp_mapping devsp 
inner join spl_node_service_instance_tbl serv_conf_in 
on devsp.sp_id_fk = serv_conf_in.sp_id_fk 
and devsp.cpm_id_fk = serv_conf_in.cpm_id_fk
inner join spl_node_service_conf_tbl serv_conf
on serv_conf.id = serv_conf_in.serv_conf_id_fk
where devsp.cpm_id_fk=? and devsp.dev_id_fk =? and devsp.sp_id_fk=?`

const QUERY_GET_SER_CONFIG_BY_SER_CONF_ID = `select serv_conf_in.id as serv_conf_in_id, 
devsp.dev_id_fk, devsp.sp_id_fk,serv_conf_in.serv_conf_id_fk,serv_conf.conf_type_code,
serv_conf.serv_conf_name, serv_conf.serv_conf 
from spl_node_dev_sp_mapping devsp 
inner join spl_node_service_instance_tbl serv_conf_in 
on devsp.sp_id_fk = serv_conf_in.sp_id_fk 
and devsp.cpm_id_fk = serv_conf_in.cpm_id_fk
inner join spl_node_service_conf_tbl serv_conf
on serv_conf.id = serv_conf_in.serv_conf_id_fk
where serv_conf.id=?`
