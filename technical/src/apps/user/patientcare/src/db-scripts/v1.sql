CREATE TABLE IF NOT EXISTS sync_tbl (store_name TEXT, sync_order INTEGER, last_synced TEXT, sync_type INTEGER, sync_to_server_pending INTEGER,sync_to_server_pending_time TEXT,sync_from_server_pending INTEGER,sync_from_server_pending_time TEXT);
CREATE TABLE IF NOT EXISTS patient_master_tbl (uuid TEXT, patient_reg_no Text, fname TEXT, lname TEXT, mob_no TEXT, age TEXT, blood_grp TEXT, gender INTEGER, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS patient_admission_tbl (uuid TEXT, patient_uuid TEXT, patient_reg_no TEXT, bed_no TEXT, status INTEGER, sp_uuid INTEGER, dr_incharge INTEGER, admitted_on DATETIME, discharged_on DATETIME, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS patient_personal_details_tbl (patient_uuid TEXT, admission_uuid INTEGER, uuid TEXT, age TEXT, other_details TEXT,  person_accompanying TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER , client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS patient_medical_details_tbl (patient_uuid TEXT, admission_uuid INTEGER, uuid TEXT, present_complaints TEXT, reason_for_admission TEXT, history_present_illness TEXT, past_history TEXT, treatment_before_admission TEXT, investigation_before_admission TEXT, family_history TEXT, allergies TEXT, personal_history TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS schedule_tbl (admission_uuid TEXT, uuid TEXT, conf_type_code TEXT, conf TEXT, start_date TEXT,end_date TEXT, status INTEGER, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS conf_tbl (uuid TEXT,conf_type_code TEXT, conf TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS action_tbl (uuid TEXT,admission_uuid TEXT, conf_type_code TEXT, schedule_uuid TEXT, scheduled_time DATETIME, is_deleted INTEGER, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS action_txn_tbl (uuid TEXT,admission_uuid TEXT,schedule_uuid TEXT,txn_data TEXT,scheduled_time DATETIME, txn_state INTEGER, conf_type_code TEXT, runtime_config_data TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, status INTEGER, client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS service_point_tbl (uuid TEXT, sp_name TEXT, short_desc TEXT, sp_state INTEGER, sp_state_since TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER , client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS device_access_tbl (userid INTEGER, user_fname TEXT, user_lname TEXT, email TEXT, pin TEXT);
CREATE TABLE IF NOT EXISTS doctors_orders_tbl (uuid TEXT, admission_uuid TEXT, doctor_id INTEGER, doctors_orders TEXT, comment TEXT, ack_by INTEGER, ack_time TEXT, status INTEGER, order_created_time TEXT, order_type TEXT, document_uuid TEXT,document_name Text, doctype Text, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER , client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS usr_tbl (usr_id INTEGER, usr_name TEXT, urole_name Text, fname TEXT, lname TEXT ,updated_on TEXT, sync_pending INTEGER , client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS document_tbl (uuid TEXT, doc_path TEXT, doc_name TEXT, doc_type TEXT, datastore TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER, client_updated_at TEXT );
CREATE TABLE IF NOT EXISTS treatment_tbl (uuid TEXT, admission_uuid TEXT, treatment_done TEXT,treatment_performed_time TEXT, details TEXT, post_observation TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER , client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS treatment_doc_tbl (treatment_uuid TEXT, document_uuid TEXT);
CREATE TABLE IF NOT EXISTS pathology_record_tbl (uuid TEXT, admission_uuid TEXT, test_performed TEXT, test_performed_time TEXT, test_result TEXT, comments TEXT, updated_by INTEGER, updated_on TEXT, sync_pending INTEGER , client_updated_at TEXT);
CREATE TABLE IF NOT EXISTS pathology_record_doc_tbl (pathology_record_uuid TEXT,document_uuid TEXT);
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('mst_user_tbl', 100,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('document_tbl', 101,2,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('service_point_tbl', 200,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('conf_tbl', 201,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_master_tbl', 300,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_admission_tbl', 301,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_personal_details_tbl', 302,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_medical_details_tbl', 303,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('schedule_tbl', 304,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('action_tbl', 305,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('action_txn_tbl', 306,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('doctors_orders_tbl', 307,3,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('treatment_tbl', 308,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('treatment_doc_tbl', 309,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('pathology_record_tbl', 310,1,0,'',0,'');
INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('pathology_record_doc_tbl', 311,1,0,'',0,'');