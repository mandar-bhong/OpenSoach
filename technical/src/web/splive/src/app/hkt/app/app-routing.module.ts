import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {AppContainerComponent} from '../../shared/applayout/app-container/app-container.component';
import {ContentComponent} from '../../shared/applayout/content/content.component';
import {FooterComponent} from '../../shared/applayout/footer/footer.component';
import {SideBarComponent} from '../../shared/applayout/side-bar/side-bar.component';
import {TopHeaderComponent} from '../../shared/applayout/top-header/top-header.component';
const routes : Routes = [
  {
    path: '',
    children: [
      {
        path: '',
        component:AppContainerComponent,
        children: [
          {
            path: 'devices',
            loadChildren: '../../shared/devices/devices.module#DevicesModule'
          },
          {
            path: 'charts',
            loadChildren: './charts/charts.module#ChartsModule'
          },          
        ]
      },
      {
        path: 'auth',
        loadChildren: '../../shared/auth/auth.module#AuthModule'
      },

    ]
  }
  
  // {
  //   path: '',
  //   redirectTo: '',
  //   pathMatch: 'full',
  // }
];

// const routes: Routes = [
//   {
//     path: 'devices',
//     loadChildren: '../../../app/shared/devices/devices.module#DevicesModule'
//   },
//   {
//     path: 'charts',
//     loadChildren: 'app/hkt/app/charts/charts.module#ChartsModule'
//   },
//   {
//     path: '',
//     redirectTo: '',
//     pathMatch: 'full',
//   }
// ];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
