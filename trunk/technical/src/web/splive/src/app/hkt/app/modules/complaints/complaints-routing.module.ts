import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ComplaintListComponent } from '../complaints/complaint-list/complaint-list.component';
import { ComplaintDetailsComponent } from '../complaints/complaint-detalis/complaint-details.component';

const routes: Routes = [
  {
    path: '',
    component: ComplaintListComponent
  },
  {
    path: 'detail',
    component: ComplaintDetailsComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ComplaintsRoutingModule { }
