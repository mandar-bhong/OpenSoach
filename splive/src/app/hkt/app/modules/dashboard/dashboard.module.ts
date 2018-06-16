import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { DefaultDashboardComponent } from './default-dashboard/default-dashboard.component';
import { FeedbackSummaryComponent } from './feedback-summary/feedback-summary.component';
import { TaskSummaryComponent } from './task-summary/task-summary.component';

@NgModule({
  imports: [
    CommonModule,
    DashboardRoutingModule,
    MaterialModules,
    AppCommonModule,
    FormsModule
  ],
  declarations: [
    DefaultDashboardComponent,
    FeedbackSummaryComponent,
    TaskSummaryComponent
  ]
})
export class DashboardModule { }
