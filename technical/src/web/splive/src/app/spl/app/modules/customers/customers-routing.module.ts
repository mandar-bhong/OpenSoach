import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { CustomerAddComponent } from './customer-add/customer-add.component';
import { CustomerAssociateProductComponent } from './customer-associate-product/customer-associate-product.component';
import { CustomerListComponent } from './customer-list/customer-list.component';
import { CustomerUpdateDetailsComponent } from './customer-update-details/customer-update-details.component';

const routes: Routes = [
  {
    path: '',
    component: CustomerListComponent
  },
  {
    path: 'add',
    component: CustomerAddComponent
  },
  {
    path: 'update',
    component: CustomerUpdateDetailsComponent
  },
  {
    path: 'products',
    component: CustomerAssociateProductComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CustomersRoutingModule { }
