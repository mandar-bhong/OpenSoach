import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserListComponent } from './user-list/user-list.component';
import { UserAddComponent } from './user-add/user-add.component';
import { UserDetailsComponent } from './user-details/user-details.component';
import { UserAssociateProductComponent } from './user-associate-product/user-associate-product.component';

const routes: Routes = [
  {
    path: '',
    component: UserListComponent
  },
  {
    path: 'add-user',
    component: UserAddComponent
  },
  {
    path: 'user-detail',
    component: UserDetailsComponent
  },
  {
    path: 'products',
    component: UserAssociateProductComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }
