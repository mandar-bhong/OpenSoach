package dbquery

const QUERY_SELECT_DOCTOR_USERS = `select id,usr_name,fname,lname,urole_code,urole_name from (
	select usr.id,usr.usr_name,usrd.fname,usrd.lname,urole.urole_code,urole.urole_name from spl_master_user_tbl usr
inner join spl_master_usr_cpm_tbl ucpm on usr.id = ucpm.user_id_fk
inner join spl_master_usr_details_tbl usrd on usr.id = usrd.usr_id_fk
inner join spl_master_user_role_tbl urole on urole.id = ucpm.urole_id_fk
where urole.urole_code in ('EX_DOC','IN_DOC') and ucpm.cpm_id_fk = 3
) as tbl
group by id`
