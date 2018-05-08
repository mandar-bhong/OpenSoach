
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DeviceListComponent } from './device-list/device-list.component';
import { DeviceAddComponent} from './device-add/device-add.component';
import { DeviceUpdateComponent} from './device-update/device-update.component';
const routes: Routes = [
  {
    path: '',
    component: DeviceListComponent
  },
  {
    path: 'add',
    component: DeviceAddComponent
  },
  {
    path: 'update',
    component: DeviceUpdateComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DevicesRoutingModule { }
