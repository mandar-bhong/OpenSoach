import { Injectable } from '@angular/core';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { ImageAsset } from 'tns-core-modules/image-asset/image-asset';
import { Subject } from 'rxjs/internal/Subject';
import { action } from 'tns-core-modules/ui/dialogs/dialogs';
import { IDeviceAuthResult } from '../modules/idevice-auth-result';
@Injectable()

export class PassDataService {
    public patientListViewModel: PatientListViewModel;
    // for stroing image
    selectedPatient: PatientListViewModel;
    pickedImage: ImageAsset;
    uploadedImage: ImageAsset[] = [];
    patientName: string;
    createActionsSubject: Subject<boolean> = new Subject<boolean>();
    authResultReuested: IDeviceAuthResult;    
    backalert = [];
    constructor() {
        //  this.patientListViewModel = new PatientListViewModel();
        console.log('service initiated');
    }
    setPatientData(data) {
        this.patientListViewModel = new PatientListViewModel();
        this.patientListViewModel = data;
        // console.log('patientdata set');
        console.log(this.patientListViewModel);
    }
    getpatientData() {
        return this.patientListViewModel;
    }
    getHeaderName() {
        this.selectedPatient = this.patientListViewModel;
        this.patientName = this.selectedPatient.dbmodel.bed_no + ', ' + this.selectedPatient.dbmodel.fname + ' ' + this.selectedPatient.dbmodel.lname;
        return this.patientName;
    }
    getAdmissionID() {
        return this.patientListViewModel.dbmodel.admission_uuid;
    }
    // fucntion for create actions
    createActions(actions: boolean) {
        this.createActionsSubject.next(actions);
    }
    // end of fucntion
}