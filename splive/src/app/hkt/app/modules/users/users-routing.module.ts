import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserListComponent } from '../../../../prod-shared/modules/users/user-list/user-list.component';
import { UserAddComponent } from '../../../../prod-shared/modules/users/user-add/user-add.component';
import { UserDetailsComponent } from '../../../../prod-shared/modules/users/user-details/user-details.component';
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
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }
