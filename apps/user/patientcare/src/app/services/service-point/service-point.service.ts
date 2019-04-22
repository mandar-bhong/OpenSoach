import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";

@Injectable()
export class ServicePointService {

    constructor(private database: DatabaseService) {
        
    }

    public getServicePointList() :any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("servicePointList").then(
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