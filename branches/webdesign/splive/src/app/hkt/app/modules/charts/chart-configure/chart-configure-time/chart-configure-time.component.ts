import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AmazingTimePickerService } from 'amazing-time-picker';

import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { ChartConfigurationModel } from '../../../../models/ui/chart-conf-models';
import { ChartConfigureService } from '../../../../services/chart-configure.service';

@Component({
  selector: 'app-chart-configure-time',
  templateUrl: './chart-configure-time.component.html',
  styleUrls: ['./chart-configure-time.component.css']
})
export class ChartConfigureTimeComponent implements OnInit {

  editableForm: FormGroup;
  dataModel: ChartConfigurationModel;
  chartConfModel: ChartConfigurationModel;
  selectedStartTime: string;
  selectedEndTime: string;
  constructor(private dynamicContextService: DynamicContextService,
    private chartConfigureService: ChartConfigureService,
    private amazingtimepicker: AmazingTimePickerService) { }

  ngOnInit() {
    this.createControls();
    this.dataModel = this.chartConfigureService.getDataModel();
    console.log('get datamodel', this.dataModel);
    this.selectedStartTime = this.minutesToTimeString(this.dataModel.variableconf.timeconf.starttime);
    this.selectedEndTime = this.minutesToTimeString(this.dataModel.variableconf.timeconf.endtime);
    console.log('selectedEndTime', this.selectedEndTime);
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      startTimeControl: new FormControl('', [Validators.required]),
      endTimeControl: new FormControl('', [Validators.required]),
      intervalControl: new FormControl('', [Validators.required])
    });
  }

  nextClick() {
    // TODO: uncomment
    if (this.editableForm.invalid) { return; }
    console.log('this.selectedStartTime', this.selectedStartTime);
    console.log('this.selectedEndTime', this.selectedEndTime);
    this.dataModel.variableconf.timeconf.starttime = this.timeStringToMinutes(this.selectedStartTime);
    this.dataModel.variableconf.timeconf.endtime = this.timeStringToMinutes(this.selectedEndTime);

    this.chartConfigureService.setDataModel(this.dataModel);
    this.dynamicContextService.onAction(true);
  }

  previousClick() {
    this.dynamicContextService.onAction(false);
  }

  openStartTime() {
    const amazingTimePicker = this.amazingtimepicker.open({
      time: this.selectedStartTime,
      theme: 'material-orange',
    });
    amazingTimePicker.afterClose().subscribe(time => {
      this.selectedStartTime = time;
    });
  }

  openEndTime() {
    const amazingTimePicker = this.amazingtimepicker.open({
      time: this.selectedEndTime,
      theme: 'material-orange',
    });
    amazingTimePicker.afterClose().subscribe(time => {
      this.selectedEndTime = time;
    });
  }

  timeStringToMinutes(time: string) {
    const hoursMinutes = time.split(':');
    const hours = parseInt(hoursMinutes[0], 10);
    const minutes = hoursMinutes[1] ? parseInt(hoursMinutes[1], 10) : 0;
    return hours * 60 + minutes;
  }

  minutesToTimeString(time: number) {
    if (time) {
      const hours = Math.floor(time / 60);
      const hourstr = hours < 10 ? '0' + hours : hours;
      const minutes = time % 60;
      const minutestr = minutes < 10 ? '0' + minutes : minutes;
      return hourstr + ':' + minutestr;
    }
  }

}
