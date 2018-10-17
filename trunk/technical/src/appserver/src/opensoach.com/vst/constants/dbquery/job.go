package dbquery

const QUERY_GET_JOB_LIST_BY_FILTERED_COUNT = `select count(*) as count
from spl_node_service_in_txn_tbl serv_in_txn
inner join spl_vst_token token on token.id = serv_in_txn.txn_data->'$.tokenid'
inner join spl_vst_vehicle_master_tbl vehicle on vehicle.id = token.vhl_id_fk
where serv_in_txn.id in (
	select max(id) from spl_node_service_in_txn_tbl group by txn_data->'$.tokenid'
) $WhereCondition$ `

const QUERY_JOB_LIST_SELECT_BY_FILTER = `select vehicle.id as vehicleid,vehicle.vehicle_no,token.id as tokenid,token.token,token.state,token.generated_on,token.generated_on as intime, 
case when status = 6 then txn_date
			 else null 
			 end as outtime
from spl_node_service_in_txn_tbl serv_in_txn
inner join spl_vst_token token on token.id = serv_in_txn.txn_data->'$.tokenid'
inner join spl_vst_vehicle_master_tbl vehicle on vehicle.id = token.vhl_id_fk
where serv_in_txn.id in (
	select max(id) from spl_node_service_in_txn_tbl group by txn_data->'$.tokenid'
) $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_JOB_DETAILS_BY_ID = "select * from spl_node_service_in_txn_tbl where txn_data->'$.tokenid'= ?"

const QUERY_GET_JOB_VHL_INFO_BY_VHL_ID = `select 
vehicle.vehicle_no,
vehicle.details->>'$.ownerdetails.mobno' as mobileno,
concat(vehicle.details->>'$.ownerdetails.firstname'," ",vehicle.details->>'$.ownerdetails.lastname') as name,
txn_data->>'$.vehicledetails.km' as kms,
txn_data->>'$.vehicledetails.petrol' as petrol
from spl_vst_vehicle_master_tbl vehicle
inner join spl_vst_token token on token.vhl_id_fk = vehicle.id
inner join spl_node_service_in_txn_tbl serv_in_txn on serv_in_txn.txn_data->'$.tokenid'=token.id
where token.id = ? and serv_in_txn.status = 3`
