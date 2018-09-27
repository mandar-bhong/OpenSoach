import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ServicepointListComponent } from './servicepoint-list/servicepoint-list.component';
import {
  ServicepointServiceAssociateComponent,
} from './servicepoint-service-associate/servicepoint-service-associate.component';

const routes: Routes = [
  {
    path: '',
    component: ServicepointListComponent
  },
  {
    path: 'service-associate',
    component: ServicepointServiceAssociateComponent
  },

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ServicepointsRoutingModule { }
