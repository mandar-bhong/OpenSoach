
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DeviceListComponent } from './device-list/device-list.component';
import { DeviceAddComponent} from './device-add/device-add.component';
import { DeviceUpdateComponent} from './device-update/device-update.component';
import { DeviceAssociateProductComponent} from './device-associate-product/device-associate-product.component';
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
  },
  {
    path: 'products',
    component: DeviceAssociateProductComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DevicesRoutingModule { }
