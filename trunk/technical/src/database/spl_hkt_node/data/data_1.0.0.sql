INSERT INTO `spl_node_report_template_tbl` (`id`, `report_code`, `report_desc`, `report_header`, `report_query_params`, `report_query`) VALUES
	(1, 'TASK_LIST_SP', 'This will produce task list', '{"en": ["Task Name", "On Time", "Delayed", "Total"]}', '1', 'select taskname, count(ontime) as ontime, count(delay) as \'delayed\', count(total) as total from (
	select
	txn_data ->> \'$.taskname\' as taskname,
	  case status
			when 1 then 1
	  end as ontime,
	  case status
			when 2 then 1
	  end as delay,
	  status as total
	from spl_node_service_in_txn_tbl serv_in_txn
	inner join spl_node_service_instance_tbl serv_conf_in on serv_conf_in.id = serv_in_txn.serv_in_id_fk
	where serv_in_txn.cpm_id_fk = $WhereCpmIdValue$ and sp_id_fk = ? and txn_date between ? and ?
	) as tbl
	group by taskname'),
	(2, 'TASK_SUMMARY_SP', 'This will produce task list', '{"en": ["Loction Name", "Location Category", "Completed Tasks", "Delayed Tasks"]}', '1','select 
sp.sp_name,
spc.spc_name,
count(serv_in_txn.`status`) as completed,
count(case serv_in_txn.`status` 
when 2 then 1 end) as \'delayed\' 
from spl_node_sp_tbl sp
inner join spl_node_sp_category_tbl spc on sp.spc_id_fk = spc.id
inner join spl_node_service_instance_tbl serv_conf_in on serv_conf_in.sp_id_fk = sp.sp_id_fk
inner join spl_node_service_in_txn_tbl serv_in_txn on serv_conf_in.id = serv_in_txn.serv_in_id_fk 
where sp.cpm_id_fk = $WhereCpmIdValue$ and sp.sp_id_fk = ? and txn_date between ? and ?'),
	(3, 'TASK_LIST_ALL', 'This will produce task list', '{"en": ["Location Name", "Task Name", "On Time", "Delayed", "Total"]}', '1','select locationname, taskname, count(ontime) as ontime, count(delay) as \'delayed\', count(total) as total from (
	select
	sp.sp_name as locationname,
	txn_data ->> \'$.taskname\' as taskname,
	  case status
			when 1 then 1
	  end as ontime,
	  case status
			when 2 then 1
	  end as delay,
	  status as total
	from spl_node_service_in_txn_tbl serv_in_txn
	inner join spl_node_service_instance_tbl serv_conf_in on serv_conf_in.id = serv_in_txn.serv_in_id_fk
	inner join spl_node_sp_tbl sp on serv_conf_in.sp_id_fk = sp.sp_id_fk
		where serv_in_txn.cpm_id_fk = $WhereCpmIdValue$ and txn_date between ? and ?
	) as tbl
	group by taskname'),
	(4, 'TASK_SUMMARY_ALL', 'This will produce task list', '{"en": ["Completed Tasks", "Delayed Tasks"]}', '1','select 
count(serv_in_txn.`status`) as completed,
count(case serv_in_txn.`status` 
when 2 then 1 end) as \'delayed\' 
from spl_node_service_in_txn_tbl serv_in_txn
where serv_in_txn.cpm_id_fk = $WhereCpmIdValue$ and txn_date between ? and ?');