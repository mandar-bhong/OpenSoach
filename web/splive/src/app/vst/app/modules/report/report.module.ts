import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { ReportContainerComponent } from './report-container/report-container.component';
import { ReportRoutingModule } from './report-routing.module';

@NgModule({
  imports: [
    CommonModule,
    ReportRoutingModule,
    MaterialModules,
    FormsModule,
    ReactiveFormsModule,
    AppCommonModule,
  ],
  declarations: [ReportContainerComponent]
})
export class ReportModule { }
