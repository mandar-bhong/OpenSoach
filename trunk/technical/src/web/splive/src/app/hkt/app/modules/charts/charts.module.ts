import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { ChartConfigureComponent } from './chart-configure/chart-configure.component';
import { ChartListComponent } from './chart-list/chart-list.component';
import { ChartsRoutingModule } from './charts-routing.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppCommonModule } from '../../../../shared/app-common.module';
import { DynamicContextService } from '../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { DynamicComponentLoaderModule } from '../../../../shared/modules/dynamic-component-loader/dynamic-component-loader.module';
import { ChartConfigureTaskComponent } from './chart-configure/chart-configure-task/chart-configure-task.component';
import { ChartConfigureBasicComponent } from './chart-configure/chart-configure-basic/chart-configure-basic.component';
import { ChartConfigureTimeComponent } from './chart-configure/chart-configure-time/chart-configure-time.component';
import { ChartConfigurePreviewComponent } from './chart-configure/chart-configure-preview/chart-configure-preview.component';
import { StepperModule } from '../../../../shared/modules/stepper/stepper.module';

@NgModule({
  imports: [
    CommonModule,
    ChartsRoutingModule,
    MaterialModules,
    FormsModule,
    ReactiveFormsModule,
    AppCommonModule,
    DynamicComponentLoaderModule,
    StepperModule
  ],
  declarations: [
    ChartListComponent,
    ChartConfigureComponent,
    ChartConfigureTaskComponent,
    ChartConfigureBasicComponent,
    ChartConfigureTimeComponent,
    ChartConfigurePreviewComponent],
  entryComponents: [
    ChartConfigureTaskComponent,
    ChartConfigureBasicComponent,
    ChartConfigureTimeComponent,
    ChartConfigurePreviewComponent,
  ]
})
export class ChartsModule {
  constructor(private dynamicContextService: DynamicContextService) {
    dynamicContextService.addDynamicComponentMaping('app-chart-configure-basic', ChartConfigureBasicComponent);
    dynamicContextService.addDynamicComponentMaping('app-chart-configure-task', ChartConfigureTaskComponent);
    dynamicContextService.addDynamicComponentMaping('app-chart-configure-time', ChartConfigureTimeComponent);
    dynamicContextService.addDynamicComponentMaping('app-chart-configure-preview', ChartConfigurePreviewComponent);
  }
}
