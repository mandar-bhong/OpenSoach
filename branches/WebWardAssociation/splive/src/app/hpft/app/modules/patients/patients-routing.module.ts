import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { PatientAddComponent } from './patient-add/patient-add.component';
import { PatientChartComponent } from './patient-chart/patient-chart.component';
import { PatientCheckSearchComponent } from './patient-check-search/patient-check-search.component';
import { PatientDetailsComponent } from './patient-details/patient-details.component';
import { PatientListComponent } from './patient-list/patient-list.component';
import { PatientsPersonalDetailComponent } from './patients-personal-detail/patients-personal-detail.component';


const routes: Routes = [
  {
    path: '',
    component: PatientListComponent
  },
  {
    path: 'add',
    component: PatientAddComponent
  },
  {
    path: 'patient_chart',
    component: PatientChartComponent
  },
  {
    path: 'patient_admission',
    component: PatientDetailsComponent
  },
  {
    path: 'patient_search',
    component: PatientCheckSearchComponent
  },
  {
    path: 'patient_detail',
    component: PatientsPersonalDetailComponent
  },

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PatientsRoutingModule { }
