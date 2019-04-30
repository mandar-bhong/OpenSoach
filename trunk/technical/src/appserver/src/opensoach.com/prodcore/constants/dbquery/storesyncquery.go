package dbquery

const QUERY_SELECT_SYNC_CONFIG_ON = `select * from spl_node_sync_config_tbl where store_name = ?`

const QUERY_SELECT_SYNC_STORE_QUERY = `select * from (
	select * from ($SyncQuery$) as t
	where admission_id in 
	(
		select admission_id from view_monitored_patients where usr_id_fk = :usrid
		group by admission_id,sp_id_fk,patient_id_fk,usr_id_fk
	) 
	)t`
