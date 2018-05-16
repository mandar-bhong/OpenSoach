import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {
  DeviceDetailsViewComponent,
} from '../../../../prod-shared/modules/devices/device-details-view/device-details-view.component';
import { DeviceListComponent } from '../../../../prod-shared/modules/devices/device-list/device-list.component';

const routes: Routes = [
  {
    path: '',
    component: DeviceListComponent
  },
  {
    path: 'details',
    component: DeviceDetailsViewComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DevicesRoutingModule { }
