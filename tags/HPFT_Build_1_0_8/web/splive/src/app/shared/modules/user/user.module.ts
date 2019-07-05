import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { ChangePasswordComponent } from './change-password/change-password.component';
import { UserRoutingModule } from './user-routing.module';
import { AppCommonModule } from '../../app-common.module';

@NgModule({
  imports: [
    CommonModule,
    UserRoutingModule,
    AppCommonModule
  ],
  declarations: [
    ChangePasswordComponent
  ]
})
export class UserModule { }
