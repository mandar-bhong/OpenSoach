import { Injectable } from '@angular/core';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { ImageAsset } from 'tns-core-modules/image-asset/image-asset';
import { Subject } from 'rxjs/internal/Subject';
import { action } from 'tns-core-modules/ui/dialogs/dialogs';
import { IDeviceAuthResult } from '../modules/idevice-auth-result';


import {
    CFAlertDialog,
    DialogOptions,
    CFAlertActionAlignment,
    CFAlertActionStyle,
    CFAlertStyle
} from "nativescript-cfalert-dialog";
import { RouterExtensions } from 'nativescript-angular/router';
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
    backalert: boolean;
    getAuthUserId: any;

    private cfalertDialog: CFAlertDialog;
    constructor(private routerExtensions: RouterExtensions) {
        //  this.patientListViewModel = new PatientListViewModel();
        console.log('service initiated');
        this.cfalertDialog = new CFAlertDialog();
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
    getPatientID() {
        return this.patientListViewModel.dbmodel.patient_uuid;
    }
    // fucntion for create actions
    createActions(actions: boolean) {
        this.createActionsSubject.next(actions);
    }
    // end of fucntion

    showNotification(): void {
        let onSelection = response => {
            // this.routerExtensions.back();
        };
        let onSelection1 = response => {
        };
        const options: DialogOptions = {
            dialogStyle: CFAlertStyle.NOTIFICATION,
            title: "Unsaved Changes!",
            message: "Do you wish to Disactive changes?",
            backgroundBlur: true,
            onDismiss: () => console.log("showAlert dismissed"),
            buttons: [
                {
                    text: "Yes",
                    buttonStyle: CFAlertActionStyle.NEGATIVE,
                    buttonAlignment: CFAlertActionAlignment.JUSTIFIED,
                    onClick: onSelection
                },
                {
                    text: "No, Thanks.",
                    buttonStyle: CFAlertActionStyle.POSITIVE,
                    buttonAlignment: CFAlertActionAlignment.JUSTIFIED,
                    onClick: onSelection1
                }]
        };
        this.cfalertDialog.show(options);
    }
}