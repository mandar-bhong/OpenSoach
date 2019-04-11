import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';

import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { CustomerAddComponent } from './customer-add/customer-add.component';
import { CustomerAssociateProductComponent } from './customer-associate-product/customer-associate-product.component';
import { CustomerListComponent } from './customer-list/customer-list.component';
import { CustomerSearchComponent } from './customer-list/customer-search/customer-search.component';
import { CustomerListViewComponent } from './customer-list/list-view/customer-list-view.component';
import { CustomerUpdateDetailsComponent } from './customer-update-details/customer-update-details.component';
import { CustomersRoutingModule } from './customers-routing.module';
import { CustomerServicepointAssociateComponent } from './customer-servicepoint-associate/customer-servicepoint-associate.component';

@NgModule({
  imports: [
    CommonModule,
    CustomersRoutingModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialModules,
    AppCommonModule
  ],
  declarations: [
    CustomerAddComponent,
    CustomerListComponent,
    CustomerUpdateDetailsComponent,
    CustomerSearchComponent,
    CustomerListViewComponent,
    CustomerAssociateProductComponent,
    CustomerServicepointAssociateComponent
  ]
})
export class CustomersModule { }
