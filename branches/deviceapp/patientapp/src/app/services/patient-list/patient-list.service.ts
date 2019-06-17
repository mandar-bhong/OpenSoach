import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { PatientListViewModel } from "~/app/models/ui/patient-view-models";
import { Subscription, Subject } from "rxjs";
import { WorkerService } from "../worker.service";
import { JSONBaseDataModel } from "~/app/models/ui/json-base-data-model";
import { PersonAccompanyModel } from "~/app/models/ui/person-accompany-model";
import * as appSettings from "tns-core-modules/application-settings";
@Injectable({providedIn:'root'})
export class PatientListService {
    patientlistviewmodel: PatientListViewModel;
    val1: any;
    patientListChangedSubject = new Subject<PatientListViewModel[]>();

    constructor(private database: DatabaseService,
        private workerService: WorkerService) {
        this.workerService.patientMasterDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getPatientListDataById(dataStoreModel.uuid, 'patientlistbymasteruuid');
        });

        this.workerService.patientAdmissionDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getPatientListDataById(dataStoreModel.uuid, 'patientlistbyadmissionuuid');
        });

        this.workerService.patientPersonalDetailsDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getPatientListDataById(dataStoreModel.admission_uuid, 'patientlistbyadmissionuuid');
        });

        this.workerService.scheduleDataReceivedSubject.subscribe((dataStoreModel) => {
            this.getPatientListDataById(dataStoreModel.admission_uuid, 'patientlistbyadmissionuuid');
        });
    }


    public getData(): Promise<PatientListViewModel[]> {
        return new Promise((resolve, reject) => {
            const key = 'patientlist'
            const paramList = new Array<any>();
            let patientlist: PatientListViewModel[] = [];
            paramList.push(appSettings.getNumber("CPM_ID"));
            this.database.selectByID(key, paramList).then(
                (val) => {
                    const list: PatientListViewModel[] = [];
                    val.forEach(item => {
                        const patientListItem = new PatientListViewModel();
                        patientListItem.dbmodel = item;
                        this.fillPersonAccompanyingDetails(patientListItem.dbmodel.person_accompanying, patientListItem);
                        list.push(patientListItem);
                    });
                    resolve(list);
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
        paramList.push(appSettings.getNumber("CPM_ID"));
        this.database.selectByID(key, paramList).then(
            (val) => {
                val.forEach(item => {
                    const patientListItem = new PatientListViewModel();
                    patientListItem.dbmodel = item;
                    this.fillPersonAccompanyingDetails(patientListItem.dbmodel.person_accompanying, patientListItem);
                    patientlist.push(patientListItem);
                });

                if (val.length == 0) {
                    const patientListItem = new PatientListViewModel();
                    patientListItem.deleteuuid = uuid;
                    patientlist.push(patientListItem);
                }

                this.patientListChangedSubject.next(patientlist);

            },
            (error) => {

            }
        );
    }
    public getAdmissionDetailsByUUID(patient_uuid: string): any {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(patient_uuid);

            this.database.selectByID("patient_admission_details", paramList).then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getPatientDetailsByUUID(admission_uuid: string): any {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(admission_uuid);

            this.database.selectByID("patient_personal_details", paramList).then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    public getPersonAccompanyingByUUID(admission_uuid: string): any {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(admission_uuid);

            this.database.selectByID("patient_person_accompanying_details", paramList).then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getMedicalDetailsByUUID(admission_uuid: string): any {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(admission_uuid);

            this.database.selectByID("patient_medical_details", paramList).then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    fillPersonAccompanyingDetails(personAccompanyJSON: string, viewModel: PatientListViewModel) {
        const personAccompanyingDetails = new JSONBaseDataModel<PersonAccompanyModel[]>();
        Object.assign(personAccompanyingDetails, JSON.parse(personAccompanyJSON));
        if (personAccompanyingDetails && personAccompanyingDetails.data
            && personAccompanyingDetails.data.length > 0) {
            viewModel.personAccompanyContact = personAccompanyingDetails.data[0].contact;
            if (personAccompanyingDetails.data[0].alternatecontact) {
                viewModel.personAccompanyContact += ', ' + personAccompanyingDetails.data[0].alternatecontact;
            }
        }
    }
}