package dbquery

const QUERY_SPL_NODE_DASHBOARD_DEVICE_SUMMARY = `select connection_state,count(*) as count from
	spl_node_dev_status_tbl as devstate
	inner join spl_node_dev_tbl as dev on dev.dev_id_fk = devstate.dev_id_fk
	where dev.cpm_id_fk = ?
	group by connection_state`

const QUERY_SPL_NODE_DASHBOARD_LOCATION_SUMMARY = `select sp_state,count(*) as count from spl_node_sp_tbl where cpm_id_fk= ? group by sp_state`

const QUERY_SPL_NODE_DASHBOARD_FEEDBACK = `select feedback , count(*) as count from spl_node_feedback_tbl $WhereCondition$ group by feedback`

const QUERY_SPL_NODE_DASHBOARD_TASK = `select status,count(status) as count from spl_node_service_in_txn_tbl serv_in_txn
inner join spl_node_service_instance_tbl serv_conf_in on serv_in_txn.serv_in_id_fk = serv_conf_in.id
$WhereCondition$ group by status`

const QUERY_SPL_NODE_DASHBOARD_COMPLAINT_SUMMARY = `select  complaint_state,count(*) as count from spl_hkt_sp_complaint_tbl $WhereCondition$ group by complaint_state`

const QUERY_GET_FEEDBACK_BY_DATE = `select  id,sp_id_fk,feedback,feedback_comment,raised_on from spl_node_feedback_tbl $WhereCondition$`

const QUERY_SPL_NODE_DASHBOARD_IN_USE_LOCATION_COUNT = `select count( distinct sp.sp_id_fk) as count  from spl_node_sp_tbl sp
inner join spl_node_dev_sp_mapping devsp on devsp.sp_id_fk = sp.sp_id_fk 
inner join spl_node_service_instance_tbl serv_conf_in on serv_conf_in.sp_id_fk = sp.sp_id_fk 
where sp.cpm_id_fk= ?`
