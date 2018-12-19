import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { PatientListViewModel } from "~/app/models/ui/patient-view-models";

@Injectable()
export class PatientListService {
    patientlist:PatientListViewModel[];
    patientlistviewmodel:PatientListViewModel;
    val1:any;

    constructor(private database: DatabaseService) {
        
    }

    public getData() :any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("patientlist").then(
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