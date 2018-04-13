import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthGuard } from '../../shared/auth-guard';
import { AppContainerComponent } from '../../shared/layouts/app-layout/app-container/app-container.component';
import { AuthLayoutComponent } from '../../shared/layouts/auth-layout/auth-layout.component';

const routes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    component: AppContainerComponent,
    children: [
      {
        path: '',
        loadChildren: './modules/dashboard/dashboard.module#DashboardModule',
      },
      {
        path: 'dashboard',
        loadChildren: './modules/dashboard/dashboard.module#DashboardModule',
      },
      {
        path: 'devices',
        loadChildren: '../../shared/modules/devices/devices.module#DevicesModule'
      },
      {
        path: 'customers',
        loadChildren: './modules/customers/customers.module#CustomersModule',
      },
    ]
  },
  {
    path: '',
    component: AuthLayoutComponent,
    children: [
      {
        path: 'auth',
        loadChildren: '../../shared/modules/auth/auth.module#AuthModule'
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
