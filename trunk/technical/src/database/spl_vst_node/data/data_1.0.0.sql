INSERT INTO `spl_node_report_template_tbl` (`id`, `report_code`, `report_desc`, `report_header`, `report_query_params`, `report_query`) VALUES
	(1, 'CONSOLIDATED_REPORT', 'This will produce consolidated report', '{"en": ["Vehicle No", "Contact No", "Service Time", "In Time", "Out Time", "Wait Time", "Job Creation Time", "Execution Time", "Delivery Time", "Tentative Price", "Billed Amount"]}', '1', 'select 
vehicle_no,
vehicle.details->>\'$.ownerdetails.mobileno\' as contactno,
convert(token.generated_on,char(50)) as servicetime,
convert(token.generated_on,char(50)) as intime,
case when status = 6 then convert(txn_date,char(50))
			 else null 
			 end as outtime,
(select 
case when (time_to_sec(timediff(b.txn_date,a.txn_date))/60) > 60 
	then
		concat(hour(TIMEDIFF(b.txn_date,a.txn_date)),\'h\',minute(TIMEDIFF(b.txn_date,a.txn_date)),\'m\' )
	else 	
		concat(minute(TIMEDIFF(b.txn_date,a.txn_date)),\'m\' )
	end  as wt 
from spl_node_service_in_txn_tbl as a,spl_node_service_in_txn_tbl as b
where a.`status` = 1 and b.`status` = 2 and a.txn_data->\'$.tokenid\'= token.id and b.txn_data->\'$.tokenid\'= token.id) as waittime,
(select 
case when (time_to_sec(timediff(d.txn_date,c.txn_date))/60) > 60 
	then
		concat(hour(TIMEDIFF(d.txn_date,c.txn_date)),\'h\',minute(TIMEDIFF(d.txn_date,c.txn_date)),\'m\' )
	else 	
		concat(minute(TIMEDIFF(d.txn_date,c.txn_date)),\'m\' )
	end  as jobcreationtime 
from spl_node_service_in_txn_tbl as c,spl_node_service_in_txn_tbl as d
where c.`status` = 2 and d.`status` = 3 and c.txn_data->\'$.tokenid\'= token.id and d.txn_data->\'$.tokenid\'= token.id) as jobcreationtime,
(select 
case when (time_to_sec(timediff(f.txn_date,e.txn_date))/60) > 60 
	then
		concat(hour(TIMEDIFF(f.txn_date,e.txn_date)),\'h\',minute(TIMEDIFF(f.txn_date,e.txn_date)),\'m\' )
	else 	
		concat(minute(TIMEDIFF(f.txn_date,e.txn_date)),\'m\' )
	end  as jobexetime 
from spl_node_service_in_txn_tbl as e,spl_node_service_in_txn_tbl as f
where e.`status` = 2 and f.`status` = 4 and e.txn_data->\'$.tokenid\'= token.id and f.txn_data->\'$.tokenid\'= token.id limit 1) as jobexetime,
(select 
case when (time_to_sec(timediff(h.txn_date,g.txn_date))/60) > 60 
	then
		concat(hour(TIMEDIFF(h.txn_date,g.txn_date)),\'h\',minute(TIMEDIFF(h.txn_date,g.txn_date)),\'m\' )
	else 	
		concat(minute(TIMEDIFF(h.txn_date,g.txn_date)),\'m\' )
	end  as deliverytime 
from spl_node_service_in_txn_tbl as g,spl_node_service_in_txn_tbl as h
where g.`status` = 5 and h.`status` = 6 and g.txn_data->\'$.tokenid\'= token.id and h.txn_data->\'$.tokenid\'= token.id) as deliverytime,
(select txn_data->\'$.tentcost\' from spl_node_service_in_txn_tbl where txn_data->\'$.tokenid\'= token.id and status = 2) as tentativeprice,
txn_data->\'$.billedamount\' as billedamount
from spl_vst_token token
inner join spl_vst_vehicle_master_tbl vehicle on vehicle.id = token.vhl_id_fk
inner join spl_node_service_in_txn_tbl serv_in_txn on token.id = serv_in_txn.txn_data->\'$.tokenid\'
where serv_in_txn.id in (
	select max(id) from spl_node_service_in_txn_tbl group by txn_data->\'$.tokenid\'
) and serv_in_txn.cpm_id_fk = $WhereCpmIdValue$ and txn_date between ? and ?');