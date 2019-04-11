import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { ActionDBModel, ActionTxnDBModel } from "~/app/models/ui/action-models";
import { ActionDataStoreModel } from "~/app/models/db/action-datastore";
import { MonitorType } from "~/app/app-constants";

@Injectable()
export class MonitorService {

    constructor(private database: DatabaseService) {

    }

    public getTempActionTxn(admission_uuid): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();        
            paramList.push(admission_uuid); 
            this.database.selectByID("monitorTxnList",paramList).then(
                (val) => {
                    const filterdata  = val.filter(a => JSON.parse(a.conf).name ==='Temperature');
                   resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
    public getBloodPreActionTxn(admission_uuid:string): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();        
            paramList.push(admission_uuid);           
            this.database.selectByID("monitorTxnList",paramList).then(                
                (val) => {                   
                     const filterdata  = val.filter(data => JSON.parse(data.conf).name ===MonitorType.Blood_Pressure);
                     resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getRespirationActionTxn(admission_uuid): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();        
            paramList.push(admission_uuid); 
            this.database.selectByID("monitorTxnList",paramList).then(
                (val) => {   
                    console.log('getRespirationActionTxn data received',val);             
                    const filterdata  = val.filter(a => JSON.parse(a.conf).name ==='Respiration Rate');
                   console.log('filterdata',filterdata);
                    resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getPulseActionTxn(admission_uuid): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();        
            paramList.push(admission_uuid);
            this.database.selectByID("monitorTxnList",paramList).then(
                (val) => {
                    const filterdata  = val.filter(a => JSON.parse(a.conf).name ==='Pulse Rate');
                     resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    public getUserAccountList(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("userList").then(
                (val) => {                
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
}