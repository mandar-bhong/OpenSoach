import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";

@Injectable()
export class DatabaseSchemaService {

    dbConnection: any;
    dbVersion: any;
    dbInError: boolean;
    schema = [
        "CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY AUTOINCREMENT, item_name TEXT, user_id TEXT)",
        "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id TEXT UNIQUE, password TEXT)"
    ]

    seedData = [
        "INSERT INTO items (item_name, user_id) VALUES ('Apple','Sanjay')"
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