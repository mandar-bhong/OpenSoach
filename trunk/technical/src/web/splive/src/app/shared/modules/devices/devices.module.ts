import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { DeviceListComponent } from './device-list/device-list.component';
import { DevicesRoutingModule } from './devices-routing.module';

@NgModule({
  imports: [
    CommonModule,
    DevicesRoutingModule
  ],
  declarations: [DeviceListComponent]
})
export class DevicesModule { }
