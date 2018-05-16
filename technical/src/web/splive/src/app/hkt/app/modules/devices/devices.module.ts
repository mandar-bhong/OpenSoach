import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import {
  DeviceDetailsViewComponent,
} from '../../../../shared/views/devices/device-details-view/device-details-view.component';
import {
  DeviceListViewComponent,
} from '../../../../shared/views/devices/device-list/device-list-view/device-list-view.component';
import { DeviceListComponent } from '../../../../shared/views/devices/device-list/device-list.component';
import { DeviceSearchComponent } from '../../../../shared/views/devices/device-list/device-search/device-search.component';
import { DevicesRoutingModule } from './devices-routing.module';

@NgModule({
  imports: [
    CommonModule,
    DevicesRoutingModule,
    AppCommonModule,
    MaterialModules,
    FormsModule,
    ReactiveFormsModule,
  ],
  declarations: [
    DeviceDetailsViewComponent,
    DeviceListComponent,
    DeviceListViewComponent,
    DeviceSearchComponent
  ]
})
export class DevicesModule { }
