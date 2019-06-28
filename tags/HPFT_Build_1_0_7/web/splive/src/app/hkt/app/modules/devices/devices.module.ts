import { NgModule } from '@angular/core';

import { ProdDevicesModule } from '../../../../prod-shared/modules/devices/prod-devices.module';
import { DevicesRoutingModule } from './devices-routing.module';

@NgModule({
  imports: [
    DevicesRoutingModule,
    ProdDevicesModule
  ],
  declarations: [
  ]
})
export class DevicesModule { }
