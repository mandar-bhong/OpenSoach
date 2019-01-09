import { Injectable } from '@angular/core';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { ImageAsset } from 'tns-core-modules/image-asset/image-asset';
@Injectable()

export class PassDataService {
    public patientListViewModel: PatientListViewModel;
    // for stroing image
    selectedPatient: PatientListViewModel;
    pickedImage: ImageAsset;
    uploadedImage: ImageAsset[] = [];
    patientName: string;
    constructor() {
        //  this.patientListViewModel = new PatientListViewModel();
        console.log('service initiated');
    }
    setPatientData(data) {
        this.patientListViewModel = new PatientListViewModel();
        this.patientListViewModel = data;
        console.log('patientdata set');
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
}