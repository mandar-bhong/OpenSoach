import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { RegistrationComponent } from './registration/registration.component';
import { UserMgmtRoutingModule } from './user-mgmt-routing.module';

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
