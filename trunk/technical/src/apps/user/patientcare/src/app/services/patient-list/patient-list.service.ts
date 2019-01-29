import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { PatientListViewModel } from "~/app/models/ui/patient-view-models";
import { Subscription, Subject } from "rxjs";
import { WorkerService } from "../worker.service";

@Injectable()
export class PatientListService {
    patientlist: PatientListViewModel[];
    patientlistviewmodel: PatientListViewModel;
    val1: any;
    public patientMasterSubscription: Subscription;
    public patientAdmissionSubscription: Subscription;
    patientListChangedSubject = new Subject<PatientListViewModel[]>();

    patientmaster = [];
    patientadmission = [];
    constructor(private database: DatabaseService,
        private workerService: WorkerService) {
        this.patientMasterSubscription = this.workerService.patientMasterDataReceivedSubject.subscribe((uuid) => {
            this.getPatientListItemByMasterid(uuid);
        });

        this.patientAdmissionSubscription = this.workerService.patientAdmissionDataReceivedSubject.subscribe((uuid) => {
            console.log('subscriber invoked in patient list', uuid);
            this.getPatientListItemByAdmissionid(uuid);
        });
    }

    public getPatientListItemByMasterid(uuid: string) {
        setTimeout(() => {
            return this.getPatientListDataMasterById(uuid);
        }, 4000)

    }

    public getPatientListItemByAdmissionid(uuid: string) {
        setTimeout(() => {
            return this.getPatientListDataAdmissionById(uuid);
        }, 4000)
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

    public getPatientListDataMasterById(uuid: string): any {

        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(uuid);
            this.database.selectByID("patientlistbymasteruuid", paramList).then(
                (val) => {
                    console.log("patient list by master item", val);

                    val.forEach(item => {
                        // console.log("val master", val);
                        const patientMasterListItem = new PatientListViewModel();
                        patientMasterListItem.dbmodel = item;
                        this.patientmaster.push(patientMasterListItem);
                        console.log(' this.patientmaster patient list service', this.patientmaster);
                    });
                    this.patientListChangedSubject.next(this.patientmaster);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
    public getPatientListDataAdmissionById(uuid: string): any {


        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();

            paramList.push(uuid);
            this.database.selectByID("patientlistbyuuid", paramList).then(
                (val) => {
                    val.forEach(item => {
                        // console.log("val Admission", val);
                        const patientAdmissionListItem = new PatientListViewModel();
                        patientAdmissionListItem.dbmodel = item;
                        // console.log(' patient admission patient list service', patientAdmissionListItem.dbmodel);
                        this.patientadmission.push(patientAdmissionListItem);
                        console.log('all data', this.patientadmission);
                    });

                    this.patientListChangedSubject.next(this.patientadmission);
                    resolve(val);
                },
                (error) => {
                    console.log('error', error);
                    reject(error);
                }
            );

        });

    }

}