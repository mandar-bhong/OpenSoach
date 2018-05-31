import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ServicepointListComponent } from '../../../../prod-shared/modules/servicepoints/servicepoint-list/servicepoint-list.component';
import {
  ServicepointDetailsComponent
} from '../../../../prod-shared/modules/servicepoints/servicepoint-details/servicepoint-details.component';

const routes: Routes = [
  {
    path: '',
    component: ServicepointListComponent
  },
  {
    path: 'details',
    component: ServicepointDetailsComponent
  },

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ServicepointsRoutingModule { }
