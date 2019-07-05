import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { DashboardRoutingModule } from './dashboard-routing.module';
import { DefaultDashboardComponent } from './default-dashboard/default-dashboard.component';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';

@NgModule({
  imports: [
    CommonModule,
    DashboardRoutingModule,
    MaterialModules
  ],
  declarations: [DefaultDashboardComponent]
})
export class DashboardModule { }
