import { Component, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';

import { ServiceConfigurationRequest } from '../../../../../../prod-shared/models/api/service-configuration-models';
import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { ServiceConfigurationUpdateRequest } from '../../../../models/api/chart-conf-models';
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
    private chartConfigureService: ChartConfigureService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

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
    if (this.dataModel.servconfid) {
      // tod update
      const request = new ServiceConfigurationUpdateRequest();
      this.dataModel.copyToUpdate(request);
      this.chartConfigureService.updateConfiguration(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.chartConfigureService.setDataModel(this.dataModel);
          this.dynamicContextService.onAction(true);
        }
      });
    } else {
      const serviceConfigurationRequest = new ServiceConfigurationRequest();
      this.dataModel.copyTo(serviceConfigurationRequest);
      this.chartConfigureService.addChartData(serviceConfigurationRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.servconfid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.chartConfigureService.setDataModel(this.dataModel);
          this.dynamicContextService.onAction(true);
        }
      });
    }
  }

  previousClick() {
    this.dynamicContextService.onAction(false);
    this.editableForm.controls['categoryControl'].disable();
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

