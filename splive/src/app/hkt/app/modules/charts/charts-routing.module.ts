import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ChartConfigureComponent } from './chart-configure/chart-configure.component';
import { ChartDataComponent } from './chart-data/chart-data.component';
import { ChartListComponent } from './chart-list/chart-list.component';

const routes: Routes = [
  {
    path: '',
    component: ChartListComponent
  },
  {
    path: 'configure',
    component: ChartConfigureComponent
  },
  {
    path: 'data',
    component: ChartDataComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ChartsRoutingModule { }
