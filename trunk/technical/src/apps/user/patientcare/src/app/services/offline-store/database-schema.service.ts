import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";

@Injectable()
export class DatabaseSchemaService {

    dbConnection: any;
    dbVersion: any;
    dbInError: boolean;
    schema = [
        "CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY AUTOINCREMENT, item_name TEXT, user_id TEXT)",
        "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id TEXT UNIQUE, password TEXT)",
        "CREATE TABLE IF NOT EXISTS sync_tbl (id INTEGER PRIMARY KEY, store_name TEXT)",
        "CREATE TABLE IF NOT EXISTS patient_master_tbl (id INTEGER PRIMARY KEY AUTOINCREMENT, uuid TEXT, patient_reg_no Text, fname TEXT, lname TEXT, mob_no TEXT, age TEXT, blood_grp TEXT, gender INTEGER, created_on DATETIME ,updated_on DATETIME )",
        "CREATE TABLE IF NOT EXISTS patient_admission_tbl (id INTEGER PRIMARY KEY AUTOINCREMENT, uuid TEXT, patient_id INTEGER, patient_reg_no TEXT, bed_no TEXT, status INTEGER, attended DATETIME, sp_id INTEGER, dr_incharge INTEGER, admitted_on DATETIME, discharged_on DATETIME, created_on DATETIME, updated_on DATETIME)",
        "CREATE TABLE IF NOT EXISTS patient_personal_details_tbl (id INTEGER PRIMARY KEY AUTOINCREMENT, patient_id INTEGER, admission_id INTEGER, uuid TEXT, age TEXT, weight TEXT, other_details TEXT, created_on DATETIME, updated_on DATETIME )",
        "CREATE TABLE IF NOT EXISTS patient_medical_details_tbl (id INTEGER PRIMARY KEY AUTOINCREMENT,patient_id INTEGER, admission_id INTEGER, uuid TEXT, reason_for_admission TEXT, patient_medical_hist TEXT, treatment_recieved_before TEXT, family_hist TEXT, menstrual_hist TEXT, allergies TEXT, perdsonal_hist TEXT, general_physical_exam TEXT, systematic_exam TEXT, created_on DATETIME, updated_on DATETIME)",    
        "CREATE TABLE IF NOT EXISTS patient_chart_conf_tbl (id INTEGER PRIMARY KEY AUTOINCREMENT,admission_id INTEGER,  uuid TEXT, conf_type_code TEXT, conf TEXT, created_on TEXT DEFAULT CURRENT_TIMESTAMP, updated_on TEXT DEFAULT CURRENT_TIMESTAMP)",
        "CREATE TABLE IF NOT EXISTS patient_conf_tbl (id INTEGER PRIMARY KEY AUTOINCREMENT, uuid TEXT,conf_type_code TEXT, conf TEXT, created_on TEXT DEFAULT CURRENT_TIMESTAMP, updated_on TEXT DEFAULT CURRENT_TIMESTAMP)",
        "CREATE TABLE IF NOT EXISTS action_tbl (id INTEGER PRIMARY KEY AUTOINCREMENT, uuid TEXT,admission_id INTEGER,chart_conf_id INTEGER, exec_time TEXT)",
        "CREATE TABLE IF NOT EXISTS service_point_tbl (sp_id INTEGER, spc_id INTEGER, sp_name TEXT, short_desc TEXT, sp_state INTEGER, sp_state_since DATETIME, created_on DATETIME,updated_on DATETIME )"
    ]

    seedData = [
        "INSERT INTO items (item_name, user_id) VALUES ('Apple','Sanjay')",
        "INSERT INTO sync_tbl (id, store_name) VALUES (1, 'patient_master')",
        "INSERT INTO sync_tbl (id, store_name) VALUES (2, 'patient_admission')",
        "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on,updated_on) VALUES (1, 'PM001','P12B12213', 'Amol', 'Patil', '9812xxxxxx', '22', 'AB+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",    
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (2, 'PM002','P12B12214', 'Sagar', 'Patil', '9982xxxxxx', '24', 'O+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (3, 'PM003','P12B12215', 'Shubham', 'Lunia', '9832xxxxxx', '34', 'A+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (4, 'PM004','P12B12216', 'Mayuri', 'Jain', '9212xxxxxx', '27', 'A+', '2', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (5, 'PM005','P12B12217', 'Sanjay', 'Sawant', '9644xxxxxx', '33', 'A+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
        "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (6, 'PM006','P12B12218', 'Pooja', 'Lokare', '9522xxxxxx', '25', 'AB-', '2', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (7, 'PM007','P12B12219', 'Mandar', 'Bhong', '9012xxxxxx', '38', 'O-', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (8, 'PM008','P12B12220', 'Praveen', 'Pandey', '9442xxxxxx', '29', 'B+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (9, 'PM009','P12B12221', 'Shashank', 'Atre', '9642xxxxxx', '21', 'O+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (10, 'PM010','P12B12222', 'Tejal', 'Deshmukh', '9412xxxxxx', '25', 'AB-', '2', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
	    "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (11, 'PM011','P12B1223', 'Shahuraj', 'Patil', '9572xxxxxx', '21', 'O+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",	
        "INSERT INTO patient_master_tbl (id, uuid, patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, created_on, updated_on) VALUES (12, 'PM012','P12B12224', 'Abhijeet', 'Kalbhor', '9042xxxxxx', '24', 'O+', '1', '2018-12-03 12:22:57', '2018-12-03 12:22:57' )",    
        "INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (1, 'PA001', 1, 'P12B12213', '3A/312', '1', '2018-12-04 14:37:53', '1', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",        
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (2, 'PA002', 2, 'P12B12214', '3B/323', '1', '2018-12-04 12:47:53', '1', '1', '2018-12-05 14:37:53', '', '2018-12-05 14:37:53', '2018-12-05 14:37:53')",		
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (3, 'PA001', 3, 'P12B12213', '2A/643', '1', '2018-12-04 09:17:53', '1', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",        
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (4, 'PA001', 4, 'P12B12213', '4A/415', '2', '2018-12-04 11:00:53', '1', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",        
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (5, 'PA001', 5, 'P12B12213', '5A/616', '3', '2018-12-04 01:11:53', '1', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",	
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (6, 'PA001', 6, 'P12B12213', '6A/317', '1', '2018-12-04 14:32:53', '1', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",		
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (7, 'PA001', 7, 'P12B12213', '7A/312', '2', '2018-12-04 16:44:53', '2', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",      
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (8, 'PA001', 8, 'P12B12213', '3A/319', '3', '2018-12-04 11:12:53', '2', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",       
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (9, 'PA001', 9, 'P12B12213', '8A/314', '2', '2018-12-04 04:54:53', '2', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",        
    	"INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (10, 'PA001', 10, 'P12B12213', '4A/309', '1', '2018-12-04 15:55:53', '2', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",       
	    "INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (11, 'PA001', 11, 'P12B12213', '2B/231', '4', '2018-12-04 21:35:53', '3', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",        
	    "INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (12, 'PA001', 12, 'P12B12213', '2B/232', '1', '2018-12-04 19:33:53', '3', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",        
        "INSERT INTO patient_admission_tbl (id , uuid , patient_id , patient_reg_no , bed_no , status , attended, sp_id , dr_incharge , admitted_on , discharged_on , created_on , updated_on ) VALUES (13, 'PA001', 1, 'P12B12213', '4A/416', '2', '2018-12-04 13:32:53', '3', '1', '2018-12-04 14:37:53', '', '2018-12-04 14:37:53', '2018-12-04 14:37:53')",

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (1, 1,'Medicine' ,'{"name":"Sinarest","desc":"3 times a day after meal"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (2, 1,'Medicine' ,'{"name":"Aspirin","desc":"Incase of high body temperature"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (3, 1,'Medicine' ,'{"name":"Zofran","desc":"Incase of continuos vomitting and nausea"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (4, 1,'Intake' ,'{"name":"Saline","desc":"200ml"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (5, 1,'Output' ,'{"name":"Output","desc":"200ml"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (6, 1,'Monitor' ,'{"name":"Temperature","desc":"Monitor every 2 hours"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (7, 1,'Monitor' ,'{"name":"Blood pressure","desc":"Monitor every 3 hours"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (8, 1,'Monitor' ,'{"name":"Blood pressure","desc":"Monitor every 3 hours"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (9, 1,'Monitor' ,'{"name":"Pulse Rate","desc":"Monitor every 15 mins"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (10, 1,'Monitor' ,'{"name":"Respiration Rate","desc":"Monitor every 30 mins"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_chart_conf_tbl (id , admission_id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (11, 1,'Medicine' ,'{"name":"Acetaminophen","desc":"3 times a day after meal"}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO patient_conf_tbl (id , conf_type_code , conf , created_on , updated_on ) 
            VALUES (1, 'Monitor' ,'{"tasks":[{"name":"Temperature"},{"name":"Blood Pressure"},{"name":"Pulse Rate"},{"name":"Respiration Rate"}]}' , '2018-12-04 14:37:53', '2018-12-04 14:37:53')`,

        `INSERT INTO service_point_tbl (sp_id,spc_id,sp_name,short_desc,sp_state,sp_state_since,created_on,updated_on ) values (1,1,"General Ward 1","",1,'2018-12-04 14:37:53','2018-12-04 14:37:53','2018-12-04 14:37:53')`,
        `INSERT INTO service_point_tbl (sp_id,spc_id,sp_name,short_desc,sp_state,sp_state_since,created_on,updated_on ) values (2,1,"General Ward 2","",1,'2018-12-04 14:37:53','2018-12-04 14:37:53','2018-12-04 14:37:53')`,
        `INSERT INTO service_point_tbl (sp_id,spc_id,sp_name,short_desc,sp_state,sp_state_since,created_on,updated_on ) values (3,1,"General Ward 3","",1,'2018-12-04 14:37:53','2018-12-04 14:37:53','2018-12-04 14:37:53')`,        
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
                console.log("TABLE CREATED", query);
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
                    console.log('Table created', query);
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
                console.log("SEED DATA CREATED", query);
            });
        }, (error) => {
            console.error("CREATE TABLE ERROR", error);
            this.dbInError = true;
        });
    }


}