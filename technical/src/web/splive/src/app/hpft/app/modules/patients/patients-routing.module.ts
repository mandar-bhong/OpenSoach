import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { PatientListComponent } from './patient-list/patient-list.component';
import { PatientAddComponent } from './patient-add/patient-add.component';

const routes: Routes = [
  {
    path: '',
    component: PatientListComponent
  },
  {
    path: 'add',
    component: PatientAddComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PatientsRoutingModule { }
