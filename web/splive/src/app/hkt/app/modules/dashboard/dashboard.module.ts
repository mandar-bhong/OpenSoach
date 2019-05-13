import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { NgxChartsModule } from '@swimlane/ngx-charts';

import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { ComplaintTrendComponent } from './complaint-trend/complaint-trend.component';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { DefaultDashboardComponent } from './default-dashboard/default-dashboard.component';
import { FeedbackSummaryComponent } from './feedback-summary/feedback-summary.component';
import { FeedbackTrendComponent } from './feedback-trend/feedback-trend.component';
import { TaskSummaryComponent } from './task-summary/task-summary.component';
import { TaskTrendComponent } from './task-trend/task-trend.component';
import { ComplaintSummaryComponent } from './complaint-summary/complaint-summary.component';

@NgModule({
  imports: [
    CommonModule,
    DashboardRoutingModule,
    MaterialModules,
    AppCommonModule,
    FormsModule,
    NgxChartsModule
  ],
  declarations: [
    DefaultDashboardComponent,
    FeedbackSummaryComponent,
    TaskSummaryComponent,
    TaskTrendComponent,
    FeedbackTrendComponent,
    ComplaintTrendComponent,
    ComplaintSummaryComponent
  ]
})
export class DashboardModule { }
