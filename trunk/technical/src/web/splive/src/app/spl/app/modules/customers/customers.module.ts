import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CustomersRoutingModule } from './customers-routing.module';
import { CustomerAddComponent } from './customer-add/customer-add.component';
import { CustomerListComponent } from './customer-list/customer-list.component';
import { CustomerUpdateDetailsComponent } from './customer-update-details/customer-update-details.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { CustomerSearchComponent } from './customer-list/customer-search/customer-search.component';
import { CustomerListViewComponent } from './customer-list/list-view/customer-list-view.component';

@NgModule({
  imports: [
    CommonModule,
    CustomersRoutingModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialModules
  ],
  declarations: [CustomerAddComponent, CustomerListComponent, CustomerUpdateDetailsComponent, CustomerSearchComponent, CustomerListViewComponent]
})
export class CustomersModule { }
