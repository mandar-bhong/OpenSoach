import { Component, OnInit } from '@angular/core';
import { StepData } from '../../../../../shared/modules/stepper/step-data';
import { StepperService } from '../../../../../shared/modules/stepper/stepper.service';
import { DynamicContextService } from '../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { ChartConfigurationModel, VariableServiceConfig, TimeConfigurationModel } from '../../../models/ui/chart-conf-models';
import { ServiceConfigurationModel } from '../../../../../shared/models/ui/service-configuration-models';
import { ServiceConfigurationRequest } from '../../../../../shared/models/api/service-configuration-models';
import { Subscription } from 'rxjs/Subscription';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-chart-configure',
  templateUrl: './chart-configure.component.html',
  styleUrls: ['./chart-configure.component.css']
})
export class ChartConfigureComponent implements OnInit {

  step = 'app-chart-configure-basic';
  mode = 0; // 0:add, 1:update
  steps = ['app-chart-configure-basic', 'app-chart-configure-time', 'app-chart-configure-task', 'app-chart-configure-preview'];
  count = 0;
  completedcount = 0;
  dataModel: ChartConfigurationModel;
  stepperList: StepData[] = [
    { name: 'Basic', color: '', icon: 'fa fa-certificate', componentselector: 'app-banking-outlet-actions' },
    { name: 'Time', color: '', icon: 'fa fa-table', componentselector: 'app-banking-channel-details' },
    { name: 'Task', color: '', icon: 'fa fa-hourglass-start', componentselector: 'app-test-branch-hour' },
    { name: 'Preview', color: '', icon: 'fa fa-map-marker', componentselector: 'app-banking-channel-location' }
  ];
  routeSubscription: Subscription;
  constructor(private stepperService: StepperService,
    private dynamicContextService: DynamicContextService,
    private route: ActivatedRoute) {
    //this.dynamicContextService
  }
  ngOnInit() {

    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.servconfid = Number(params['id']);
        this.getConfiguration();
      } else {
        this.dataModel = new ChartConfigurationModel();
      }
    });
  }

  getConfiguration() {
    // call API to get existing configuration
  }

  nextClick() {
    this.dynamicContextService.save();
    this.count++;
    this.step = this.steps[this.count];
    console.log('step', this.step);
    this.stepperService.stepperCount(this.count);
  }

  previousClick() {
    this.count--;
    this.step = this.steps[this.count];
    this.stepperService.stepperCount(this.count);
  }
}
