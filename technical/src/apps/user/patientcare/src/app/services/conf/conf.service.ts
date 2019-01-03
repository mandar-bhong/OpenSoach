import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";

@Injectable()
export class ConfService {

    constructor(private database: DatabaseService) {
        
    }

    public getMonitorConf() :any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("monitorConfList").then(
                (val)=> {                  
                    resolve(val);          
                },
                (error)=>{
                    reject(error);
                }
            );      

        });

    }

}