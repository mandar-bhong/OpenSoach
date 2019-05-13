package dbquery

const QUERY_SELECT_SYNC_CONFIG_ON = `select * from spl_node_sync_config_tbl where store_name = ?`

const QUERY_SELECT_SYNC_STORE_USER_DEVICE_QUERY = `select * from (
	select * from ($SyncQuery$) as t
	where admission_id in 
	(
		select admission_id from view_monitored_patients where usr_id_fk = :usrid
		group by admission_id,sp_id_fk,patient_id_fk,usr_id_fk
	) 
	)t`

const QUERY_SELECT_SYNC_STORE_USER_DEVICE_COUNT_QUERY = `select count(*) as count,max(updated_on) as max_updated_on from (
		select * from ($SyncQuery$) as t
		where admission_id in 
		(
			select admission_id from view_monitored_patients where usr_id_fk = :usrid
			group by admission_id,sp_id_fk,patient_id_fk,usr_id_fk
		) 
		)t`

const QUERY_SELECT_SYNC_STORE_COUNT_QUERY = `select count(*) as count,max(updated_on) as max_updated_on from (
			$SyncQuery$
		)t`

const QUERY_SELECT_SYNC_STORE_PATIENT_ADMISSION_STORE = `select if(padmsn.updated_on > upmm.updated_on,padmsn.updated_on,upmm.updated_on) as updated_on, padmsn.id as admission_id,padmsn.uuid,patient.uuid as patient_uuid,padmsn.patient_reg_no,padmsn.bed_no,padmsn.status,sp.uuid as sp_uuid,padmsn.dr_incharge,padmsn.admitted_on,padmsn.discharged_on,padmsn.updated_by, padmsn.client_updated_at from spl_hpft_patient_admission_tbl padmsn 
left join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk
left join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk 
left join spl_hpft_user_patient_monitor_mapping upmm on upmm.sp_id_fk = padmsn.sp_id_fk and upmm.patient_id_fk = padmsn.patient_id_fk
left join spl_hpft_user_patient_monitor_mapping upmm1 on upmm1.sp_id_fk = padmsn.sp_id_fk and upmm1.patient_id_fk is null
left join spl_hpft_user_patient_monitor_mapping upmm2 on upmm2.sp_id_fk is null and upmm2.patient_id_fk is null
where padmsn.cpm_id_fk = :cpmid and (padmsn.updated_on >= :updatedon or upmm.updated_on >= :updatedon ) and padmsn.status = 1
group by admission_id`

const QUERY_SELECT_SYNC_STORE_PATIENT_MONITOR_MAPPING_USER_DEVICE_QUERY = `select * from (
	select * from ($SyncQuery$) as t
	where admission_id not in 
	(
		select admission_id from view_monitored_patients where usr_id_fk = :usrid
		group by admission_id,sp_id_fk,patient_id_fk,usr_id_fk
	) 
	)t`

const QUERY_SELECT_SYNC_STORE_PATIENT_MONITOR_MAPPING_USER_DEVICE_COUNT_QUERY = `select count(*) as count,max(updated_on) as max_updated_on from (
		select * from ($SyncQuery$) as t
		where admission_id not in 
		(
			select admission_id from view_monitored_patients where usr_id_fk = :usrid
			group by admission_id,sp_id_fk,patient_id_fk,usr_id_fk
		) 
		)t`
