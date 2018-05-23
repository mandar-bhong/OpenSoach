import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { DynamicContextService } from '../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { StepData } from '../../../../../shared/modules/stepper/step-data';
import { StepperService, StepperState } from '../../../../../shared/modules/stepper/stepper.service';
import { ChartConfigurationModel } from '../../../models/ui/chart-conf-models';
import { ChartConfigureService } from '../../../services/chart-configure.service';

@Component({
  selector: 'app-chart-configure',
  templateUrl: './chart-configure.component.html',
  styleUrls: ['./chart-configure.component.css']
})
export class ChartConfigureComponent implements OnInit, OnDestroy {
  mode = 0; // 0:add, 1:update
  count = 0;
  completedcount = 0;
  steps: StepData[] = [
    { name: 'Basic', color: '', icon: 'fa fa-cog', componentselector: 'app-chart-configure-basic' },
    { name: 'Time', color: '', icon: 'fa fa-clock-o', componentselector: 'app-chart-configure-time' },
    { name: 'Task', color: '', icon: 'fa fa-tasks', componentselector: 'app-chart-configure-task' },
    { name: 'Preview', color: '', icon: 'fa fa-table', componentselector: 'app-chart-configure-preview' }
  ];
  step: StepData;

  routeSubscription: Subscription;
  dynamicComponentActionSubcription: Subscription;
  constructor(private stepperService: StepperService,
    private dynamicContextService: DynamicContextService,
    private route: ActivatedRoute,
    private chartConfigureService: ChartConfigureService) {
  }

  ngOnInit() {
    this.step = this.steps[0];
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.getConfiguration(Number(params['id']));
      } else {
        this.chartConfigureService.createDataModel();
      }
    });

    this.dynamicComponentActionSubcription = this.dynamicContextService.triggerAction.subscribe(params => {
      this.changeSteps(params);
    });
  }

  getConfiguration(configid: number) {
    // TODO: call API to get existing configuration
    // this.chartConfigureService.getConfigList({ recid: configid }).subscribe(payloadResponse => {
    //   if (payloadResponse && payloadResponse.issuccess) {

    //   }
    // });
    // this.chartConfigureService.dataModel = new ChartConfigurationModel();
  }

  changeSteps(direction: any) {
    const currentState = new StepperState();
    if (direction === true) {
      if (this.count < this.steps.length - 1) {
        this.count++;
      } else {
        currentState.completed = true;
      }
    } else {
      this.count--;
    }

    currentState.stepindex = this.count;
    this.stepperService.stepperCount(currentState);
    this.step = this.steps[this.count];
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }

    if (this.dynamicComponentActionSubcription) {
      this.dynamicComponentActionSubcription.unsubscribe();
    }
  }
}
