import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ChartConfigureComponent } from './chart-configure/chart-configure.component';
import { ChartDataComponent } from './chart-data/chart-data.component';
import { ChartListComponent } from './chart-list/chart-list.component';

const routes: Routes = [
  {
    path: '',
    component: ChartDataComponent
  },
  {
    path: 'configure',
    component: ChartConfigureComponent
  },
  {
    path: 'templatelist',
    component: ChartListComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ChartsRoutingModule { }
