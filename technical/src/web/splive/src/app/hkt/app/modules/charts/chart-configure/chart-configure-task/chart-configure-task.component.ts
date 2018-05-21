import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { DragulaService } from 'ng2-dragula/components/dragula.provider';

import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { TaskTemplateModel } from '../../../../models/api/chart-conf-models';
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
  constructor(private dynamicContextService: DynamicContextService,
    private chartConfigureService: ChartConfigureService,
    private dragulaService: DragulaService) { }

  ngOnInit() {
    this.createControls();
    this.dataModel = this.chartConfigureService.getDataModel();
    this.dataModel.variableconf.taskconf.tasks.forEach(item => {
      this.chartTasks.push(item);
    });

    this.getTaskLibrary();
  }

  createControls(): void {
    this.addTaskForm = new FormGroup({
      taskNameControl: new FormControl('', [Validators.required])
    });
  }

  getTaskLibrary() {
    // TODO:call api to get task library

    const tasks: TaskTemplateModel[] = [];

    for (let i = 1; i <= 10; i++) {
      const a = new TaskTemplateModel();
      a.taskname = 'task ' + i;
      tasks.push(a);
    }

    tasks.forEach(item => {
      if (!this.chartTasks.find(a => a.taskname === item.taskname)) {
        const chartTaskConfModel = new ChartTaskModel();
        chartTaskConfModel.copyFrom(item);
        this.taskLibrary.push(chartTaskConfModel);
      }
    });
  }

  nextClick() {
    this.dataModel.variableconf.taskconf.tasks.length = 0;
    this.chartTasks.forEach(item => {
      this.dataModel.variableconf.taskconf.tasks.push(item);
    });

    this.chartConfigureService.setDataModel(this.dataModel);
    this.dynamicContextService.onAction(true);
  }

  previousClick() {
    this.dynamicContextService.onAction(false);
  }

  addTask() {
    if (this.addTaskForm.invalid) { return; }
    const task = new ChartTaskModel();
    task.taskname = this.addNewTaskName;
    this.chartTasks.push(task);
    this.addNewTaskName = null;
    this.isTaskAdd = false;

    // TODO: call API to save task in library
  }

  showAddTask() {
    this.isTaskAdd = true;
  }

  cancelAddTask() {
    this.addNewTaskName = null;
    this.isTaskAdd = false;
  }
}
