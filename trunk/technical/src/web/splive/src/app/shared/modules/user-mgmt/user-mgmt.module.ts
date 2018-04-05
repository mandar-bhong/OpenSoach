import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserMgmtRoutingModule } from './user-mgmt-routing.module';
import { RegistrationComponent } from './registration/registration.component';
@NgModule({
  imports: [
    CommonModule,
    UserMgmtRoutingModule
  ],
  declarations: [
  RegistrationComponent
  ]
})
export class UserMgmtModule { }
