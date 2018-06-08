import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppCommonModule } from '../../../shared/app-common.module';
import { MaterialModules } from '../../../shared/modules/material/material-modules';
import { ProdCommonModule } from '../../prod-common.module';
import {
  ServicepointDeviceAssociateComponent,
} from './servicepoint-device-associate/servicepoint-device-associate.component';
import { ServicepointListViewComponent } from './servicepoint-list/servicepoint-list-view/servicepoint-list-view.component';
import { ServicepointListComponent } from './servicepoint-list/servicepoint-list.component';
import { ServicepointSearchComponent } from './servicepoint-list/servicepoint-search/servicepoint-search.component';
import {
  ServicepointServiceAssociateComponent,
} from './servicepoint-service-associate/servicepoint-service-associate.component';
import { ServicepointUpdateComponent } from './servicepoint-update/servicepoint-update.component';

@NgModule({
  imports: [
    CommonModule,
    AppCommonModule,
    MaterialModules,
    ProdCommonModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  declarations: [
    ServicepointServiceAssociateComponent,
    ServicepointListComponent,
    ServicepointListViewComponent,
    ServicepointSearchComponent,
    ServicepointDeviceAssociateComponent,
    ServicepointUpdateComponent,
  ],
  entryComponents: [
    ServicepointDeviceAssociateComponent,
    ServicepointUpdateComponent
  ]
})
export class ProdServicepointsModule { }
