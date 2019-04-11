import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserListComponent } from '../../../../prod-shared/modules/users/user-list/user-list.component';
import { UserAddComponent } from '../../../../prod-shared/modules/users/user-add/user-add.component';
import { UserDetailsComponent } from '../../../../prod-shared/modules/users/user-details/user-details.component';
import { ChangePasswordComponent } from '../../../../prod-shared/modules/users/change-password/change-password.component';
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
    path: 'change-password',
    component: ChangePasswordComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }
