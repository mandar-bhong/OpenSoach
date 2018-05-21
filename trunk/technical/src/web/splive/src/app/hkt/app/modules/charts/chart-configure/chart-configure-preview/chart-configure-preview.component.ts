import { Component, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { ChartConfigurationModel } from '../../../../models/ui/chart-conf-models';
import { ChartConfigureService } from '../../../../services/chart-configure.service';

@Component({
  selector: 'app-chart-configure-preview',
  templateUrl: './chart-configure-preview.component.html',
  styleUrls: ['./chart-configure-preview.component.css']
})
export class ChartConfigurePreviewComponent implements OnInit {

  editableForm: FormGroup;
  dataModel: ChartConfigurationModel;
  slots: string[] = [];
  constructor(private dynamicContextService: DynamicContextService,
    private chartConfigureService: ChartConfigureService) { }

  ngOnInit() {
    this.createControls();
    this.dataModel = this.chartConfigureService.getDataModel();
    this.createChartSlots();
  }

  createControls(): void {
    this.editableForm = new FormGroup({});
  }

  save() {
    if (this.editableForm.invalid) { return; }

    // TODO: Call API to save configuration.

    this.chartConfigureService.setDataModel(this.dataModel);
    this.dynamicContextService.onAction(true);
  }

  previousClick() {
    this.dynamicContextService.onAction(false);
  }

  createChartSlots() {
    const d = new Date();
    const startTime = new Date(d.getFullYear(), d.getMonth(), d.getDate(), 0, this.dataModel.variableconf.timeconf.starttime, 0, 0);
    const endTime = new Date(d.getFullYear(), d.getMonth(), d.getDate(), 0, this.dataModel.variableconf.timeconf.endtime, 0, 0);
    const options = { hour12: true, hour: '2-digit', minute: '2-digit' };

    const i = startTime;
    while (i < endTime) {
      const startSlotTime = i.toLocaleTimeString('en-us', options);
      i.setMinutes(i.getMinutes() + this.dataModel.variableconf.timeconf.interval);
      const endSlotTime = i.toLocaleTimeString('en-us', options);
      this.slots.push(startSlotTime);
    }
  }
}

