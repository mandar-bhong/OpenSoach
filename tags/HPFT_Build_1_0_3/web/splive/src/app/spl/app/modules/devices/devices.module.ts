import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DevicesRoutingModule } from './devices-routing.module';
import { DeviceUpdateComponent } from './device-update/device-update.component';
import { DeviceListComponent } from './device-list/device-list.component';
import { DeviceAddComponent } from './device-add/device-add.component';
import { DeviceSearchComponent } from './device-list/device-search/device-search.component';
import { DeviceListViewComponent } from './device-list/device-list-view/device-list-view.component';
import { DeviceAssociateProductComponent} from './device-associate-product/device-associate-product.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
@NgModule({
  imports: [
    CommonModule,
    DevicesRoutingModule,
    AppCommonModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialModules
  ],
  declarations: [
    DeviceUpdateComponent,
    DeviceListComponent,
    DeviceAddComponent,
    DeviceSearchComponent,
    DeviceListViewComponent,
    DeviceAssociateProductComponent
  ]
})
export class DevicesModule { }
