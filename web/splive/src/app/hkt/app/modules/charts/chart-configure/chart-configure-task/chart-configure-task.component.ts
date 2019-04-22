import { log } from 'util';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { DragulaService } from 'ng2-dragula/components/dragula.provider';

import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { TaskTemplateRequest } from '../../../../models/api/chart-conf-models';
import { ChartConfigurationModel, ChartTaskModel } from '../../../../models/ui/chart-conf-models';
import { ChartConfigureService } from '../../../../services/chart-configure.service';

@Component({
  selector: 'app-chart-configure-task',
  templateUrl: './chart-configure-task.component.html',
  styleUrls: ['./chart-configure-task.component.css']
})
export class ChartConfigureTaskComponent implements OnInit {

  taskLibrary: Array<ChartTaskModel> = [];
  chartTasks: Array<ChartTaskModel> = [];
  addTaskForm: FormGroup;
  dataModel: ChartConfigurationModel;
  isTaskAdd = false;
  addNewTaskName: string;
  addtask: Map<string, string>;
  libtasklist: Map<string, string>;
  constructor(private dynamicContextService: DynamicContextService,
    private chartConfigureService: ChartConfigureService,
    private dragulaService: DragulaService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.createControls();
    this.dataModel = this.chartConfigureService.getDataModel();
    this.dataModel.variableconf.taskconf.tasks.forEach(item => {
      this.chartTasks.push(item);
    });
    console.log(' this.dataModel.variableconf.taskconf.tasks');
    console.log(this.dataModel.variableconf.taskconf.tasks);
    this.addtask = new Map(this.dataModel.variableconf.taskconf.tasks.map(x => [x.taskname, x.taskname] as [string, string]));
    console.log('map');
    console.log(this.addtask);
    this.getTaskLibrary();
  }

  createControls(): void {
    this.addTaskForm = new FormGroup({
      taskNameControl: new FormControl('', [Validators.required])
    });
  }

  getTaskLibrary() {
    // TODO:call api to get task library
    this.chartConfigureService.getTaskDataList({ recid: this.dataModel.spcid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        console.log('TaskTemplateResponse', payloadResponse.data);
        payloadResponse.data.forEach(item => {
          if (!this.chartTasks.find(a => a.taskname === item.taskname)) {
            const chartTaskConfModel = new ChartTaskModel();
            chartTaskConfModel.copyFrom(item);
            this.taskLibrary.push(chartTaskConfModel);
          }
        });
        console.log(' this.chartTaskss');
        console.log(this.taskLibrary);
        this.libtasklist = new Map(this.taskLibrary.map(x => [x.taskname, x.taskname] as [string, string]));
        console.log('map2');
        console.log(this.libtasklist);
      }
    });
  }

  nextClick() {
    if (this.chartTasks.length > 0) {
      this.dataModel.variableconf.taskconf.tasks.length = 0;
      this.chartTasks.forEach(item => {
        this.dataModel.variableconf.taskconf.tasks.push(item);
      });
      this.chartConfigureService.setDataModel(this.dataModel);
      this.dynamicContextService.onAction(true);
    } else {
      this.appNotificationService.info(this.translatePipe.transform('TASK_NOT_AVAILABLE'));
    }
  }

  previousClick() {
    this.dynamicContextService.onAction(false);
  }

  addTask() {
    if (this.addTaskForm.invalid) { return; }
    const addtaskobj = this.addtask.get(this.addNewTaskName);
    console.log('testobj', addtaskobj);
    const listtaskobj = this.libtasklist.get(this.addNewTaskName);
    console.log('listtaskobj', listtaskobj);
    console.log('this.addNewTaskName', this.addNewTaskName);
    if (listtaskobj) {
      this.appNotificationService.info(this.translatePipe.transform('TASK_NAME_ALL_READY_EXISTS_TASk_LIB'));
    } else {
      if (addtaskobj) {
        this.appNotificationService.info(this.translatePipe.transform('TASK_NAME_ALL_READY_EXISTS'));
      } else {
        const task = new ChartTaskModel();
        task.taskname = this.addNewTaskName;
        this.chartTasks.push(task);
        const taskTemplateRequest = new TaskTemplateRequest();
        taskTemplateRequest.taskname = this.addNewTaskName;
        taskTemplateRequest.spcid = this.dataModel.spcid;
        taskTemplateRequest.shortdesc = this.dataModel.shortdesc;
        this.chartConfigureService.addTask(taskTemplateRequest).subscribe(payloadResponse => {
          if (payloadResponse && payloadResponse.issuccess) {
          }
        });
        this.addNewTaskName = null;
        this.isTaskAdd = false;
      }
    }
  }

  showAddTask() {
    this.isTaskAdd = true;
  }

  cancelAddTask() {
    this.addNewTaskName = null;
    this.isTaskAdd = false;
  }
}
