import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AppCustomPreloader } from '../../shared/app-custom-preloader';
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
        redirectTo: '/dashboard', pathMatch: 'full'
      },
      {
        path: 'dashboard',
        loadChildren: './modules/dashboard/dashboard.module#DashboardModule',
        data: { preload: true }
      },
      {
        path: 'devices',
        loadChildren: './modules/devices/devices.module#DevicesModule',
        data: { preload: false }
      },
      {
        path: 'servicepoints',
        children: [{
          path: '',
          loadChildren: './modules/servicepoints/servicepoints.module#ServicepointsModule',
          data: { preload: true }
        },
        {
          path: 'charts',
          loadChildren: './modules/charts/charts.module#ChartsModule',
          data: { preload: false }
        }]
      },
      {
        path: 'users',
        loadChildren: './modules/users/users.module#UsersModule',
        data: { preload: false }
      },
      {
        path: 'foperators',
        loadChildren: './modules/operators/operators.module#OperatorsModule',
        data: { preload: false }
      },
      {
        path: 'complaints',
        loadChildren: './modules/complaints/complaints.module#ComplaintsModule',
        data: { preload: false }
      },
    ]
  },
  {
    path: '',
    component: AuthLayoutComponent,
    children: [
      {
        path: 'auth',
        loadChildren: '../../shared/modules/auth/auth.module#AuthModule',
        data: { preload: true }
      }
    ]
  },
  {
    path: '**',
    redirectTo: '/dashboard', pathMatch: 'full'
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { preloadingStrategy: AppCustomPreloader, initialNavigation: false })],
  exports: [RouterModule],
  providers: [AppCustomPreloader]
})
export class AppRoutingModule { }
