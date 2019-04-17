import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ForgotPasswordComponent } from './forgot-password/forgot-password.component';
import { LoginComponent } from './login/login.component';
import { LoginLayoutComponent } from './login-layout/login-layout.component';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { AuthGuard } from '../../auth-guard';
import { ChangePasswordSuccessComponent } from './change-password-success/change-password-success.component';

const routes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'loginlayout',
         component: LoginComponent
      },
      {
        path: 'forgotpassword',
         component: ForgotPasswordComponent
      },
      {
        path: 'login',
         component: LoginLayoutComponent
      },
      {
        path:'change-password/:code',component: ChangePasswordComponent
      },
      {
        path:'change-password-success',component: ChangePasswordSuccessComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthRoutingModule { }
