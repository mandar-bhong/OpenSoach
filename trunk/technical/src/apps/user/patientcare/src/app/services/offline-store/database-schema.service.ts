import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";

@Injectable()
export class DatabaseSchemaService {

    dbConnection: any;
    dbVersion: any;
    dbInError: boolean;
    schema = [
        "CREATE TABLE IF NOT EXISTS sync_tbl (store_name TEXT, sync_order INTEGER, last_synced DATETIME, sync_type INTEGER, sync_to_server_pending INTEGER,sync_to_server_pending_time DATETIME,sync_from_server_pending INTEGER,sync_from_server_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS patient_master_tbl (uuid TEXT, patient_reg_no Text, fname TEXT, lname TEXT, mob_no TEXT, age TEXT, blood_grp TEXT, gender INTEGER, updated_on DATETIME, sync_pending INTEGER, sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS patient_admission_tbl (uuid TEXT, patient_uuid TEXT, patient_reg_no TEXT, bed_no TEXT, status INTEGER, attended DATETIME, sp_uuid INTEGER, dr_incharge INTEGER, admitted_on DATETIME, discharged_on DATETIME, updated_on DATETIME, sync_pending INTEGER, sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS patient_personal_details_tbl (patient_uuid TEXT, admission_uuid INTEGER, uuid TEXT, age TEXT, weight TEXT, other_details TEXT, updated_on DATETIME, sync_pending INTEGER , sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS patient_medical_details_tbl (patient_uuid TEXT, admission_uuid INTEGER, uuid TEXT, reason_for_admission TEXT, patient_medical_hist TEXT, treatment_recieved_before TEXT, family_hist TEXT, menstrual_hist TEXT, allergies TEXT, personal_hist TEXT, general_physical_exam TEXT, systematic_exam TEXT, updated_on DATETIME, sync_pending INTEGER, sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS schedule_tbl (admission_uuid TEXT, uuid TEXT, conf_type_code TEXT, conf TEXT, updated_on DATETIME, sync_pending INTEGER, sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS conf_tbl (uuid TEXT,conf_type_code TEXT, conf TEXT, updated_on DATETIME, sync_pending INTEGER, sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS action_tbl (uuid TEXT,admission_uuid TEXT, conf_type_code TEXT, schedule_uuid TEXT, exec_time INTEGER, sync_pending INTEGER, sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS action_txn_tbl (uuid TEXT,admission_uuid TEXT,schedule_uuid TEXT,txn_data TEXT,txn_date DATETIME, txn_state INTEGER, conf_type_code TEXT, runtime_config_data TEXT, updated_on DATETIME, sync_pending INTEGER, status INTEGER, sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS service_point_tbl (uuid TEXT, sp_name TEXT, short_desc TEXT, sp_state INTEGER, sp_state_since DATETIME, updated_on DATETIME, sync_pending INTEGER , sync_pending_time DATETIME)",
        "CREATE TABLE IF NOT EXISTS user_account_tbl (id TEXT, user_fname TEXT, user_lname TEXT, email TEXT, pin TEXT)"
    ]

    seedData = [

        // sync tbl
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('service_point_tbl', 1,1,0,'',0,'')",  //from
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('conf_tbl', 2,1,0,'',0,'')",           // from
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_master_tbl', 3,3,0,'',0,'')",
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type,sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('schedule_tbl', 4,3,0,'',0,'')",
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_admission_tbl', 5,3,0,'',0,'')",
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_personal_details_tbl', 6,3,0,'',0,'')",
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('patient_medical_details_tbl', 7,3,0,'',0,'')",
        "INSERT INTO sync_tbl (store_name, sync_order, sync_type, sync_to_server_pending,sync_to_server_pending_time,sync_from_server_pending,sync_from_server_pending_time) VALUES ('action_txn_tbl', 8,3,0,'',0,'')",


        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM001','P12B12213', 'Amol', 'Patil', '9812xxxxxx', '22', 'AB+', '1', '2018-12-03 12:22:57' , 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM002','P12B12214', 'Sagar', 'Patil', '9982xxxxxx', '24', 'O+', '1', '2018-12-03 12:22:57' , 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM003','P12B12215', 'Shubham', 'Lunia', '9832xxxxxx', '34', 'A+', '1', '2018-12-03 12:22:57' , 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM004','P12B12216', 'Mayuri', 'Jain', '9212xxxxxx', '27', 'A+', '2', '2018-12-03 12:22:57' , 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM005','P12B12217', 'Sanjay', 'Sawant', '9644xxxxxx', '33', 'A+', '1', '2018-12-03 12:22:57' , 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM006','P12B12218', 'Pooja', 'Lokare', '9522xxxxxx', '25', 'AB-', '2', '2018-12-03 12:22:57' , 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM007','P12B12219', 'Mandar', 'Bhong', '9012xxxxxx', '38', 'O-', '1', '2018-12-03 12:22:57', 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM008','P12B12220', 'Praveen', 'Pandey', '9442xxxxxx', '29', 'B+', '1', '2018-12-03 12:22:57', 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM009','P12B12221', 'Shashank', 'Atre', '9642xxxxxx', '21', 'O+', '1', '2018-12-03 12:22:57', 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM010','P12B12222', 'Tejal', 'Deshmukh', '9412xxxxxx', '25', 'AB-', '2', '2018-12-03 12:22:57', 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM011','P12B12234', 'Shahuraj', 'Patil', '9572xxxxxx', '21', 'O+', '1', '2018-12-03 12:22:57', 0 )",
        "INSERT INTO patient_master_tbl ( uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending) VALUES ('PM012','P12B12224', 'Abhijeet', 'Kalbhor', '9042xxxxxx', '24', 'O+', '1', '2018-12-03 12:22:57', 0 )",

        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA001', 'PM001', 'P12B12213', '3A/312', '1', '2018-12-04 14:37:53', 'SP001', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA002', 'PM002', 'P12B12214', '3B/323', '1', '2018-12-04 12:47:53', 'SP001', '1', '2018-12-05 14:37:53', '', '2018-12-05 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA003', 'PM003', 'P12B12213', '2A/643', '1', '2018-12-04 09:17:53', 'SP001', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA004', 'PM004', 'P12B12213', '4A/415', '2', '2018-12-04 11:00:53', 'SP001', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA005', 'PM005', 'P12B12213', '5A/616', '3', '2018-12-04 01:11:53', 'SP001', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA006', 'PM006', 'P12B12213', '6A/317', '1', '2018-12-04 14:32:53', 'SP001', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA007', 'PM007', 'P12B12213', '7A/312', '2', '2018-12-04 16:44:53', 'SP002', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA008', 'PM008', 'P12B12213', '3A/319', '3', '2018-12-04 11:12:53', 'SP002', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA009', 'PM009', 'P12B12213', '8A/314', '2', '2018-12-04 04:54:53', 'SP002', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA010', 'PM010', 'P12B12213', '4A/309', '1', '2018-12-04 15:55:53', 'SP002', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA011', 'PM011', 'P12B12213', '2B/231', '4', '2018-12-04 21:35:53', 'SP003', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",
        "INSERT INTO patient_admission_tbl (uuid , patient_uuid , patient_reg_no , bed_no , status , attended, sp_uuid , dr_incharge , admitted_on , discharged_on , updated_on , sync_pending) VALUES ('PA012', 'PM012', 'P12B12213', '2B/232', '1', '2018-12-04 19:33:53', 'SP003', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53' , 0)",


        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC001','Medicine' ,'{"name":"Sinarest","desc":"3 times a day after meal"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC002','Medicine' ,'{"name":"Aspirin","desc":"Incase of high body temperature"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC003','Medicine' ,'{"name":"Zofran","desc":"Incase of continuos vomitting and nausea"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC004','Intake' ,'{"name":"Saline","desc":"200ml"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC005','Output' ,'{"name":"Output","desc":"200ml"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC006','Monitor' ,'{"name":"Temperature","desc":"Monitor every 2 hours"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC007','Monitor' ,'{"name":"Blood pressure","desc":"Monitor every 3 hours"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC008','Monitor' ,'{"name":"Blood pressure","desc":"Monitor every 3 hours"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC009','Monitor' ,'{"name":"Pulse Rate","desc":"Monitor every 15 mins"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC010','Monitor' ,'{"name":"Respiration Rate","desc":"Monitor every 30 mins"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO schedule_tbl (admission_uuid, uuid , conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('PA001', 'PC011','Medicine' ,'{"name":"Acetaminophen","desc":"3 times a day after meal"}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO conf_tbl (uuid, conf_type_code , conf , updated_on , sync_pending) 
            VALUES ('C0001', 'Monitor' ,'{"tasks":[{"name":"Temperature"},{"name":"Blood Pressure"},{"name":"Pulse Rate"},{"name":"Respiration Rate"}]}' , '2018-12-04 14:37:53' , 0)`,

        `INSERT INTO service_point_tbl (uuid,sp_name,short_desc,sp_state,sp_state_since,updated_on, sync_pending ) values ('SP001',"General Ward 1","",1,'2018-12-04 14:37:53','2018-12-04 14:37:53', 0)`,
        `INSERT INTO service_point_tbl (uuid,sp_name,short_desc,sp_state,sp_state_since,updated_on, sync_pending ) values ('SP002', "General Ward 2","",1,'2018-12-04 14:37:53','2018-12-04 14:37:53', 0)`,
        `INSERT INTO service_point_tbl (uuid,sp_name,short_desc,sp_state,sp_state_since,updated_on, sync_pending ) values ('SP003', "General Ward 3","",1,'2018-12-04 14:37:53','2018-12-04 14:37:53', 0)`,

        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A001', 'PA001' ,'Medicine' ,'PC001',  1548218525 ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A002', 'PA002' ,'Medicine' ,'PC001', 1548218525  ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A003', 'PA003' ,'Medicine' ,'PC001',   1548240000  ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A004', 'PA004' ,'Medicine' ,'PC001', 1548240000 ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A005', 'PA005' ,'Intake' ,'PC001',  1548243000  ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A006', 'PA006' ,'Monitor' ,'PC001', 1548243000 ,0)`,

        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A007', 'PA007' ,'Output' ,'PC001',   1548243000  ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A008', 'PA008' ,'Medicine' ,'PC001', 1548243000 ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A009', 'PA009' ,'Intake' ,'PC001',   1548241200  ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A010', 'PA010' ,'Monitor' ,'PC001', 1548241200 ,0)`,
        `INSERT INTO action_tbl (uuid, admission_uuid , conf_type_code, schedule_uuid ,exec_time, sync_pending) 
            VALUES ('A011', 'PA011' ,'Output' ,'PC001',  1548218525  ,0)`,



        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT001', 'PA001' ,'PC006','{"comment":"test1","value":"98"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT002', 'PA001' ,'PC006','{"comment":"test2","value":"98.3 "}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT003', 'PA001' ,'PC006','{"comment":"test3","value":"99"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT004', 'PA001' ,'PC006','{"comment":"test4","value":"101"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT005', 'PA001' ,'PC006','{"comment":"test5","value":"97"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT006', 'PA001' ,'PC006','{"comment":"test6","value":"100"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,


        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT006', 'PA001' ,'PC008','{"comment":"test1","value":"100"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT007', 'PA001' ,'PC008','{"comment":"test2","value":"150"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT008', 'PA001' ,'PC008','{"comment":"test3","value":"130"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT009', 'PA001' ,'PC008','{"comment":"test4","value":"87"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT010', 'PA001' ,'PC008','{"comment":"test5","value":"110"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT011', 'PA001' ,'PC008','{"comment":"test6","value":"140"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,


        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT012', 'PA001' ,'PC010','{"comment":"test1","value":"110"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT013', 'PA001' ,'PC010','{"comment":"test2","value":"150"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT014', 'PA001' ,'PC010','{"comment":"test3","value":"90"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT015', 'PA001' ,'PC010','{"comment":"test4","value":"77"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT016', 'PA001' ,'PC010','{"comment":"test5","value":"110"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT017', 'PA001' ,'PC010','{"comment":"test6","value":"140"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,


        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT018', 'PA001' ,'PC009','{"comment":"test1","value":"150"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT019', 'PA001' ,'PC009','{"comment":"test2","value":"100"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT020', 'PA001' ,'PC009','{"comment":"test3","value":"130"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT021', 'PA001' ,'PC009','{"comment":"test4","value":"89"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT022', 'PA001' ,'PC009','{"comment":"test5","value":"110"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,
        `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
            VALUES ('AT023', 'PA001' ,'PC009','{"comment":"test6","value":"140"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,

        // `INSERT INTO action_txn_tbl (uuid, admission_uuid ,schedule_uuid, txn_data,txn_date,txn_state, conf_type_code, runtime_config_data, updated_on , sync_pending, status) 
        //     VALUES ('A002', 'PA001' ,'PC007','{"comment":"test","value":"70"}' ,'2018-12-04 09:17:53','1', 'Monitor' ,'','',1, 0)`,


        `INSERT INTO user_account_tbl (id, user_fname ,user_lname, email, pin ) 
            VALUES ('U001', 'Amol' ,'Patil', 'amol.patil@gmail.com', '1111')`,
        `INSERT INTO user_account_tbl (id, user_fname ,user_lname, email, pin) 
            VALUES ('U002', 'Sarjerao' ,'Patil', 'sarjerao@gmail.com', '2222')`,
        `INSERT INTO user_account_tbl (id, user_fname ,user_lname, email, pin) 
            VALUES ('U003', 'Sanjay' ,'Sawant', 'sanjay.sawant@gmail.com', '3333')`,
        `INSERT INTO user_account_tbl (id, user_fname ,user_lname, email, pin) 
            VALUES ('U004', 'Sumeet' ,'Karnde', 'sumeet.karnde@gmail.com', '4444')`,
        `INSERT INTO user_account_tbl (id, user_fname ,user_lname, email, pin) 
            VALUES ('U005', 'Chandan' ,'Pal', 'chandan.pal@gmail.com', '5555')`,

    ]


    constructor(private database: DatabaseService) {

    }

    public setOfflineDB() {

        this.database.deleteDatabaseInDebugMode();
        // var promise1 = new Promise(function(resolve, reject) {
        //     setTimeout(function() {
        //       // check if everything
        //     }, 3000);
        //   });

        this.database.getdbConnection()
            .then(db => {
                this.dbConnection = db;
                this.dbConnection.version().then(
                    (version) => {
                        this.dbVersion = version[0];
                        console.log('db version', this.dbVersion);
                        // TODO for debugging
                        //this.dbVersion = 0;
                        if (this.dbVersion === "0") {
                            this.createSchema();
                            this.createSeedData();

                            this.dbVersion = 1;
                            this.dbConnection.version(this.dbVersion);
                            console.log('set db version', this.dbVersion);
                        }
                    },
                    (error) => {
                        console.log('get version error', error);
                        this.dbInError = true;
                    }
                );
            },
                (error) => {
                    console.log('db connection error');
                    this.dbInError = true;
                });

    }

    createSchema() {
        this.schema.forEach(query => {
            this.dbConnection.execSQL(query).then(() => {
                // console.log("TABLE CREATED", query);
            });
        }, (error) => {
            console.error("CREATE TABLE ERROR", error);
            this.dbInError = true;
        });
    }

    createSchemaPromise(): Promise<any> {
        return new Promise<any>((resolve, reject) => {
            this.schema.forEach(query => {
                this.dbConnection.execSQL(query).then(() => {
                    // console.log('Table created', query);
                });
            }, (error) => {
                console.log("CREATE TABLE ERROR", error);
                this.dbInError = true;
            });
        });
    }

    createSeedData() {
        this.seedData.forEach(query => {
            this.dbConnection.execSQL(query).then(() => {
                //  console.log("SEED DATA CREATED", query);
            });
        }, (error) => {
            console.error("CREATE TABLE ERROR", error);
            this.dbInError = true;
        });
    }


}