import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {
  ServicepointListComponent,
} from '../../../../prod-shared/modules/servicepoints/servicepoint-list/servicepoint-list.component';
import {
  ServicepointServiceAssociateComponent,
} from '../../../../prod-shared/modules/servicepoints/servicepoint-service-associate/servicepoint-service-associate.component';
// import { WordServiceAssociateComponent } from '../wards/word-service-associate/word-service-associate.component';
import { WordListComponent } from '../wards/word-list/word-list.component';

const routes: Routes = [
  {
    path: '',
    component: WordListComponent
  }
  // ,
  // {
  //   path: 'service-associate',
  //   component: WordServiceAssociateComponent
  // },

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ServicepointsRoutingModule { }
