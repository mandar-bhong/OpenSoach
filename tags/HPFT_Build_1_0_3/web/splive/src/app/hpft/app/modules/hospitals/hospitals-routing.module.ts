import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HospitalListComponent } from './hospital-list/hospital-list.component';
import { PathologyReportComponent } from '../patients/pathology-report/pathology-report.component';

const routes: Routes = [
  {
    path: '',
    component: HospitalListComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class HospitalsRoutingModule { }
