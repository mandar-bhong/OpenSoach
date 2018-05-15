import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ChartListComponent } from './chart-list/chart-list.component';
import { ChartConfigureComponent } from './chart-configure/chart-configure.component';

const routes: Routes = [
  {
    path: '',
    component: ChartConfigureComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ChartsRoutingModule { }
