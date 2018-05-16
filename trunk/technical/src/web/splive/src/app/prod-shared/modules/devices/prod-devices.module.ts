import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppCommonModule } from '../../../shared/app-common.module';
import { MaterialModules } from '../../../shared/modules/material/material-modules';
import { ProdCommonModule } from '../../prod-common.module';
import { DeviceDetailsViewComponent } from './device-details-view/device-details-view.component';
import { DeviceListViewComponent } from './device-list/device-list-view/device-list-view.component';
import { DeviceListComponent } from './device-list/device-list.component';
import { DeviceSearchComponent } from './device-list/device-search/device-search.component';

@NgModule({
  imports: [
    CommonModule,
    AppCommonModule,
    MaterialModules,
    FormsModule,
    ReactiveFormsModule,
    ProdCommonModule
  ],
  declarations: [
    DeviceDetailsViewComponent,
    DeviceListComponent,
    DeviceListViewComponent,
    DeviceSearchComponent
  ]
})
export class ProdDevicesModule { }
