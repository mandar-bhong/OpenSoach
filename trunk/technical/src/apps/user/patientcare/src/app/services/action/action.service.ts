import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import  { ActionDBModel, ActionTxnDBModel } from "~/app/models/ui/action-models";

@Injectable()
export class ActionService {

    constructor(private database: DatabaseService) {
        
    }

    public getActionList() :any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("actionList").then(
                (val)=> {
                    // console.log("action data",val);                  
                    resolve(val);          
                },
                (error)=>{
                    reject(error);
                }
            );      

        });

    }

    public getActionTxnList() :any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("actionTxnList").then(
                (val)=> {
                    console.log("action data service",val);                  
                    resolve(val);          
                },
                (error)=>{
                    reject(error);
                }
            );      

        });

    }
    

    public insertActionItem(data: ActionDBModel) {

        return new Promise((resolve, reject) => {

        const listData = new Array<any>();

        listData.push(data.uuid);
        listData.push(data.admission_uuid);
        listData.push(data.conf_type_code);
        listData.push(data.schedule_uuid);
        listData.push(data.exec_time);

        this.database.update("actionInsert",listData).then(
        (val)=> {
            // console.log("chart data",val);                  
            resolve(val);          
        },
        (error)=>{
            reject(error);
        }
        ); 

        });
    }

    public insertActionTxnItem(data: ActionTxnDBModel) {

        return new Promise((resolve, reject) => {

        const listData = new Array<any>();

        listData.push(data.uuid);
        listData.push(data.schedule_uuid);
        listData.push(data.txn_data);
        listData.push(data.txn_date);
        listData.push(data.txn_state);
        listData.push(data.conf_type_code);
        listData.push(data.runtime_config_data);
        listData.push(data.status);

        this.database.update("actionTxnInsert",listData).then(
        (val)=> {
            // console.log("chart data",val);                  
            resolve(val);          
        },
        (error)=>{
            reject(error);
        }
        ); 

        });
    }

}