import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DefaultDashboardComponent } from './default-dashboard/default-dashboard.component';

const routes: Routes = [
  {
    path: '',
    component: DefaultDashboardComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }
