import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { PatientListViewModel } from "~/app/models/ui/patient-view-models";
import { Subscription, Subject } from "rxjs";
import { WorkerService } from "../worker.service";

@Injectable()
export class PatientListService {
    // patientlist: PatientListViewModel[];
    patientlistviewmodel: PatientListViewModel;
    val1: any;
    public patientMasterSubscription: Subscription;
    public patientAdmissionSubscription: Subscription;
    patientListChangedSubject = new Subject<PatientListViewModel[]>();

    constructor(private database: DatabaseService,
        private workerService: WorkerService) {
        this.patientMasterSubscription = this.workerService.patientMasterDataReceivedSubject.subscribe((uuid) => {
                this.getPatientListDataById(uuid, 'patientlistbymasteruuid');
        });

        this.patientAdmissionSubscription = this.workerService.patientAdmissionDataReceivedSubject.subscribe((uuid) => {
            console.log('subscriber invoked in patient list', uuid);
                this.getPatientListDataById(uuid, 'patientlistbyadmissionuuid');
        });
    }


    public getData(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("patientlist").then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }

    public getPatientListDataById(uuid, key): any {
        const paramList = new Array<any>();
        let patientlist: PatientListViewModel[] = [];

        paramList.push(uuid);
        this.database.selectByID(key, paramList).then(
            (val) => {
                console.log("patient list by master item", val);
                val.forEach(item => {
                    // console.log("val master", val);
                    const patientMasterListItem = new PatientListViewModel();
                    patientMasterListItem.dbmodel = item;
                    patientlist.push(patientMasterListItem);
                    console.log(' this.patientmaster patient list service', patientlist);
                });
                this.patientListChangedSubject.next(patientlist);

            },
            (error) => {

            }
        );


    }
}