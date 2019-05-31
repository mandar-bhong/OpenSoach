import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { RegistrationComponent } from './registration/registration.component';
import { UserMgmtRoutingModule } from './user-mgmt-routing.module';
import { AppCommonModule } from '../../app-common.module';

@NgModule({
  imports: [
    CommonModule,
    UserMgmtRoutingModule,
    AppCommonModule
  ],
  declarations: [
  RegistrationComponent
  ]
})
export class UserMgmtModule { }
