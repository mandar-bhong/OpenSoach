import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AmazingTimePickerModule, AmazingTimePickerService } from 'amazing-time-picker';
import { DragulaModule } from 'ng2-dragula';

import { AppCommonModule } from '../../../../shared/app-common.module';
import {
  DynamicComponentLoaderModule,
} from '../../../../shared/modules/dynamic-component-loader/dynamic-component-loader.module';
import { DynamicContextService } from '../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { StepperModule } from '../../../../shared/modules/stepper/stepper.module';
import { ChartConfigureBasicComponent } from './chart-configure/chart-configure-basic/chart-configure-basic.component';
import { ChartConfigurePreviewComponent } from './chart-configure/chart-configure-preview/chart-configure-preview.component';
import { ChartConfigureTaskComponent } from './chart-configure/chart-configure-task/chart-configure-task.component';
import { ChartConfigureTimeComponent } from './chart-configure/chart-configure-time/chart-configure-time.component';
import { ChartConfigureComponent } from './chart-configure/chart-configure.component';
import { ChartDataComponent } from './chart-data/chart-data.component';
import { ChartListComponent } from './chart-list/chart-list.component';
import { ChartsRoutingModule } from './charts-routing.module';
import { ChartSearchComponent } from './chart-list/chart-search/chart-search.component';
import { ChartViewComponent } from './chart-list/chart-view/chart-view.component';
import { Ng2CarouselamosModule } from 'ng2-carouselamos';
@NgModule({
  imports: [
    CommonModule,
    ChartsRoutingModule,
    MaterialModules,
    FormsModule,
    ReactiveFormsModule,
    AppCommonModule,
    DynamicComponentLoaderModule,
    StepperModule,
    AmazingTimePickerModule,
    DragulaModule,
    Ng2CarouselamosModule
  ],
  declarations: [
    ChartListComponent,
    ChartConfigureComponent,
    ChartConfigureTaskComponent,
    ChartConfigureBasicComponent,
    ChartConfigureTimeComponent,
    ChartConfigurePreviewComponent,
    ChartSearchComponent,
    ChartViewComponent,
    ChartConfigurePreviewComponent,
    ChartConfigurePreviewComponent,
    ChartDataComponent
  ],
  entryComponents: [
    ChartConfigureTaskComponent,
    ChartConfigureBasicComponent,
    ChartConfigureTimeComponent,
    ChartConfigurePreviewComponent,
  ],
  providers: [
    AmazingTimePickerService
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
