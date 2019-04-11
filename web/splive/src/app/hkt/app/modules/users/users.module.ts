import { NgModule } from '@angular/core';

import { ProdUsersModule } from '../../../../prod-shared/modules/users/prod-users.module';
import { UsersRoutingModule } from './users-routing.module';

@NgModule({
  imports: [
    UsersRoutingModule,
    ProdUsersModule
  ],
  declarations: []
})
export class UsersModule { }
