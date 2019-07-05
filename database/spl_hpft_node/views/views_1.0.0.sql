--
-- view_get_monitor_patient
--

create view view_get_monitor_patient as 

select * from (

select 1 as monitored,padmsn.id as admission_id,upmm.id as upmmid,upmm.patient_id_fk as upmm_patient_id_fk,upmm.sp_id_fk as upmm_sp_id_fk,patient.fname,patient.lname,padmsn.patient_reg_no,padmsn.bed_no,padmsn.cpm_id_fk,upmm.usr_id_fk,padmsn.sp_id_fk,sp.sp_name,padmsn.patient_id_fk from spl_hpft_patient_admission_tbl padmsn
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk 
inner join  spl_hpft_user_patient_monitor_mapping upmm on padmsn.sp_id_fk = upmm.sp_id_fk and padmsn.patient_id_fk = upmm.patient_id_fk
inner join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk

union

select 1 as monitored, padmsn.id as admission_id,upmm.id as upmmid,upmm.patient_id_fk as upmm_patient_id_fk,upmm.sp_id_fk as upmm_sp_id_fk, patient.fname,patient.lname, padmsn.patient_reg_no,padmsn.bed_no, padmsn.cpm_id_fk,upmm.usr_id_fk,padmsn.sp_id_fk,sp.sp_name,padmsn.patient_id_fk from  spl_hpft_patient_admission_tbl padmsn  
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk 
inner join  spl_hpft_user_patient_monitor_mapping upmm on padmsn.sp_id_fk = upmm.sp_id_fk and  upmm.patient_id_fk is null
inner join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk

union

select 1 as monitored, padmsn.id as admission_id,upmm.id as upmmid,upmm.patient_id_fk as upmm_patient_id_fk,upmm.sp_id_fk as upmm_sp_id_fk, patient.fname,patient.lname, padmsn.patient_reg_no,padmsn.bed_no, padmsn.cpm_id_fk,upmm.usr_id_fk,padmsn.sp_id_fk,sp.sp_name,padmsn.patient_id_fk from  spl_hpft_patient_admission_tbl padmsn  
inner join spl_hpft_patient_master_tbl patient on patient.id = padmsn.patient_id_fk 
inner join  spl_hpft_user_patient_monitor_mapping upmm on upmm.sp_id_fk is null and  upmm.patient_id_fk is null
inner join spl_node_sp_tbl sp on sp.sp_id_fk = padmsn.sp_id_fk

) t;


--
-- view_monitored_patients
--

create view view_monitored_patients as

select * from (
 
select padmsn.uuid,padmsn.id as admission_id,padmsn.patient_id_fk,padmsn.sp_id_fk,upmm.usr_id_fk   from spl_hpft_patient_admission_tbl padmsn 
inner join  spl_hpft_user_patient_monitor_mapping upmm on padmsn.sp_id_fk = upmm.sp_id_fk and padmsn.patient_id_fk = upmm.patient_id_fk

union

select padmsn.uuid,padmsn.id as admission_id,padmsn.patient_id_fk,padmsn.sp_id_fk,upmm.usr_id_fk from spl_hpft_patient_admission_tbl padmsn 
inner join  spl_hpft_user_patient_monitor_mapping upmm on padmsn.sp_id_fk = upmm.sp_id_fk and  upmm.patient_id_fk is null

union

select padmsn.uuid,padmsn.id as admission_id,padmsn.patient_id_fk,padmsn.sp_id_fk,upmm.usr_id_fk from spl_hpft_patient_admission_tbl padmsn 
inner join  spl_hpft_user_patient_monitor_mapping upmm on upmm.sp_id_fk is null and  upmm.patient_id_fk is null

) t ;

