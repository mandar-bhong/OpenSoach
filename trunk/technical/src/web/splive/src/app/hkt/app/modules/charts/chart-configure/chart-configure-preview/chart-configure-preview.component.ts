import { Component, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';

import {
  ServiceConfigurationRequest,
  ServiceConfigurationUpdateRequest,
} from '../../../../../../prod-shared/models/api/service-configuration-models';
import { SpServiceConfService } from '../../../../../../prod-shared/services/spservice/sp-service-conf.service';
import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { ChartConfigurationModel } from '../../../../models/ui/chart-conf-models';
import { ChartConfigureService } from '../../../../services/chart-configure.service';
import { ServicepointAssociateRequest } from '../../../../../../prod-shared/models/api/servicepoint-models';
import { ProdServicepointService } from '../../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';

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
    private spServiceConfService: SpServiceConfService,
    private prodServicepointService: ProdServicepointService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe ) { }

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
      this.spServiceConfService.updateServiceConf(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.chartConfigureService.setDataModel(this.dataModel);
          this.dynamicContextService.onAction(true);
          this.associate();
        }
      });
    } else {
      const serviceConfigurationRequest = new ServiceConfigurationRequest();
      this.dataModel.copyTo(serviceConfigurationRequest);
      this.spServiceConfService.addServiceConf(serviceConfigurationRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.servconfid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.chartConfigureService.setDataModel(this.dataModel);
          this.dynamicContextService.onAction(true);
        }
      });
    }
  }
  associate() {
    const request = new ServicepointAssociateRequest();
    this.dataModel.copyToAssociateRequest(request);
    this.prodServicepointService.associateConfigure(request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
      }
    });
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

