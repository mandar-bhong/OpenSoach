import { Injectable } from '@angular/core';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
@Injectable()

export class PassDataService {
   public patientListViewModel: PatientListViewModel;
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
}