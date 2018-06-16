import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';

import { SpServiceConfService } from '../../../../../prod-shared/services/spservice/sp-service-conf.service';
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
  count = 0;
  // completedcount = 0;
  steps: StepData[] = [
    { name: 'Basic', color: '', icon: 'fa fa-cog', componentselector: 'app-chart-configure-basic' },
    { name: 'Time', color: '', icon: 'fa fa-clock-o', componentselector: 'app-chart-configure-time' },
    { name: 'Task', color: '', icon: 'fa fa-tasks', componentselector: 'app-chart-configure-task' },
    { name: 'Preview', color: '', icon: 'fa fa-table', componentselector: 'app-chart-configure-preview' }
  ];
  step: StepData;
  dataModel: ChartConfigurationModel;

  routeSubscription: Subscription;
  dynamicComponentActionSubcription: Subscription;
  constructor(private stepperService: StepperService,
    private dynamicContextService: DynamicContextService,
    private route: ActivatedRoute,
    private chartConfigureService: ChartConfigureService,
    private spServiceConfService: SpServiceConfService) {
  }

  ngOnInit() {
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel = this.chartConfigureService.createDataModel();

      if (params['id']) {
        this.dataModel.servconfid = Number(params['id']);
      }

      if (params['spid']) {
        this.dataModel.spid = Number(params['spid']);
      }

      if (params['mode']) {
        this.dataModel.mode = Number(params['mode']);
      } else {
        this.dataModel.mode = 0; // default add mode;
      }

      if (this.dataModel.servconfid) {
        this.getConfiguration();
      } else {
        this.chartConfigureService.setDataModel(this.dataModel);
        this.step = this.steps[0];
      }
    });

    this.dynamicComponentActionSubcription = this.dynamicContextService.triggerAction.subscribe(params => {
      this.changeSteps(params);
    });
  }

  getConfiguration() {
    this.spServiceConfService.getServiceConf({ recid: this.dataModel.servconfid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.copyFrom(payloadResponse.data);
        this.chartConfigureService.setDataModel(this.dataModel);
        this.step = this.steps[0];
      }
    });
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
