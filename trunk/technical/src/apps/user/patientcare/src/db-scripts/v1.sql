CREATE TABLE IF NOT EXISTS sync_tbl (store_name TEXT, sync_order INTEGER, last_synced DATETIME, sync_type INTEGER, sync_to_server_pending INTEGER,sync_to_server_pending_time DATETIME,sync_from_server_pending INTEGER,sync_from_server_pending_time DATETIME);
CREATE TABLE IF NOT EXISTS patient_master_tbl (uuid TEXT, patient_reg_no Text, fname TEXT, lname TEXT, mob_no TEXT, age TEXT, blood_grp TEXT, gender INTEGER, updated_on DATETIME, sync_pending INTEGER, client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS patient_admission_tbl (uuid TEXT, patient_uuid TEXT, patient_reg_no TEXT, bed_no TEXT, status INTEGER, attended DATETIME, sp_uuid INTEGER, dr_incharge INTEGER, admitted_on DATETIME, discharged_on DATETIME, updated_on DATETIME, sync_pending INTEGER, client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS patient_personal_details_tbl (patient_uuid TEXT, admission_uuid INTEGER, uuid TEXT, age TEXT, weight TEXT, other_details TEXT, updated_on DATETIME, sync_pending INTEGER , client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS patient_medical_details_tbl (patient_uuid TEXT, admission_uuid INTEGER, uuid TEXT, reason_for_admission TEXT, patient_medical_hist TEXT, treatment_recieved_before TEXT, family_hist TEXT, menstrual_hist TEXT, allergies TEXT, personal_hist TEXT, general_physical_exam TEXT, systematic_exam TEXT, updated_on DATETIME, sync_pending INTEGER, client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS schedule_tbl (admission_uuid TEXT, uuid TEXT, conf_type_code TEXT, conf TEXT,end_date TEXT, updated_on DATETIME, sync_pending INTEGER, client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS conf_tbl (uuid TEXT,conf_type_code TEXT, conf TEXT, updated_on DATETIME, sync_pending INTEGER, client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS action_tbl (uuid TEXT,admission_uuid TEXT, conf_type_code TEXT, schedule_uuid TEXT, exec_time DATETIME, sync_pending INTEGER, client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS action_txn_tbl (uuid TEXT,admission_uuid TEXT,schedule_uuid TEXT,txn_data TEXT,txn_date DATETIME, txn_state INTEGER, conf_type_code TEXT, runtime_config_data TEXT, updated_on DATETIME, sync_pending INTEGER, status INTEGER, client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS service_point_tbl (uuid TEXT, sp_name TEXT, short_desc TEXT, sp_state INTEGER, sp_state_since DATETIME, updated_on DATETIME, sync_pending INTEGER , client_updated_at DATETIME);
CREATE TABLE IF NOT EXISTS device_access_tbl (userid TEXT, user_fname TEXT, user_lname TEXT, email TEXT, pin TEXT);
CREATE TABLE IF NOT EXISTS doctors_orders_tbl (uuid TEXT, admission_uuid TEXT, doctor_id INTEGER, doctors_orders TEXT, document_uuid TEXT, updated_by INTEGER, updated_on DATETIME, sync_pending INTEGER , client_updated_at DATETIME);
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('service_point_tbl', 1,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('conf_tbl', 2,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_master_tbl', 3,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type,sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('schedule_tbl', 4,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_admission_tbl', 5,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_personal_details_tbl', 6,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_medical_details_tbl', 7,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('action_txn_tbl', 8,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('doctors_orders_tbl', 9,3,0,'',0,'');
INSERT INTO conf_tbl (uuid, conf_type_code , conf , updated_on , sync_pending) 
VALUES ('C0001', 'Monitor' ,'{"tasks":[{"name":"Temperature"},{"name":"Blood Pressure"},{"name":"Pulse Rate"},{"name":"Respiration Rate"}]}' , '2018-12-04 14:37:53' , 0);