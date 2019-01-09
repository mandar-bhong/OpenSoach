import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import  { ActionDBModel } from "~/app/models/ui/action-models";

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
    

    public insertActionItem(data: ActionDBModel) {

        return new Promise((resolve, reject) => {

        const listData = new Array<any>();

        listData.push(data.admission_id);
        listData.push(data.chart_conf_id);
        listData.push(data.exec_time);

        this.database.insert("actionInsert",listData).then(
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