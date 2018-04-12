import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthGuard } from '../../shared/auth-guard';
import { AppContainerComponent } from '../../shared/layouts/app-layout/app-container/app-container.component';
import { AuthLayoutComponent } from '../../shared/layouts/auth-layout/auth-layout.component';

const routes: Routes = [
  {
    path: '',
    children: [
      {
        path: '',
        canActivate: [AuthGuard],
        component: AppContainerComponent,
        children: [
          {
            path: 'devices',
            loadChildren: '../../shared/modules/devices/devices.module#DevicesModule'
          },
          {
            path: 'charts',
            loadChildren: './modules/charts/charts.module#ChartsModule'
          },
        ]
      },
    ]
  },
  {
    path: 'auth',
    component: AuthLayoutComponent,
    loadChildren: '../../shared/modules/auth/auth.module#AuthModule'
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
