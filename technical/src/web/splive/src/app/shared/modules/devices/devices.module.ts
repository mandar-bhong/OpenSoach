import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { DeviceListComponent } from './device-list/device-list.component';
import { DevicesRoutingModule } from './devices-routing.module';
import { AppCommonModule } from '../../app-common.module';

@NgModule({
  imports: [
    CommonModule,
    DevicesRoutingModule,
    AppCommonModule
  ],
  declarations: [DeviceListComponent]
})
export class DevicesModule { }
