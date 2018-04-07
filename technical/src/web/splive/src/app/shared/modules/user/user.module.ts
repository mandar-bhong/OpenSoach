import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { ChangePasswordComponent } from './change-password/change-password.component';
import { UserRoutingModule } from './user-routing.module';

@NgModule({
  imports: [
    CommonModule,
    UserRoutingModule
  ],
  declarations: [
    ChangePasswordComponent
  ]
})
export class UserModule { }
