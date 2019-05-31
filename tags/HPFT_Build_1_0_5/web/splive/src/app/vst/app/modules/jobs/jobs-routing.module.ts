import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { JobListComponent } from './job-list/job-list.component';
import { JobDetailsComponent } from './job-details/job-details.component';

const routes: Routes = [
  {
    path: '',
    component: JobListComponent
  },
  {
    path: 'details',
    component: JobDetailsComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class JobsRoutingModule { }
