package dbquery

const QUERY_SPL_NODE_DASHBOARD_DEVICE_SUMMARY = `select connection_state,count(*) as count from
	spl_node_dev_status_tbl as devstate
	inner join spl_node_dev_tbl as dev on dev.dev_id_fk = devstate.dev_id_fk
	where dev.cpm_id_fk = ?
	group by connection_state`

const QUERY_SPL_NODE_DASHBOARD_LOCATION_SUMMARY = `select sp_state,count(*) as count from spl_node_sp_tbl where cpm_id_fk= ? group by sp_state`

const QUERY_SPL_NODE_DASHBOARD_FEEDBACK = `select feedback , count(*) as count from spl_node_feedback_tbl $WhereCondition$ group by feedback`
