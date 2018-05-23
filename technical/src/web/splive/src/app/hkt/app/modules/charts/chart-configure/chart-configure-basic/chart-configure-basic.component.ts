import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';
import { SERVICE_CONF_TYPE } from '../../../../app-constants';
import { CategoriesShortDataResponse } from '../../../../models/api/chart-conf-models';
import { ChartConfigurationModel } from '../../../../models/ui/chart-conf-models';
import { ChartConfigureService } from '../../../../services/chart-configure.service';

@Component({
  selector: 'app-chart-configure-basic',
  templateUrl: './chart-configure-basic.component.html',
  styleUrls: ['./chart-configure-basic.component.css']
})
export class ChartConfigureBasicComponent implements OnInit {

  editableForm: FormGroup;
  dataModel: ChartConfigurationModel;
  categories: CategoriesShortDataResponse[] = [];
  constructor(private dynamicContextService: DynamicContextService,
    private chartConfigureService: ChartConfigureService) { }

  ngOnInit() {
    this.createControls();
    this.getCategoriesList();
    this.dataModel = this.chartConfigureService.getDataModel();
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      configNameControl: new FormControl('', [Validators.required]),
      categoryControl: new FormControl('', [Validators.required])
    });
  }
  getCategoriesList() {
    this.chartConfigureService.getCategoriesShortDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.categories = payloadResponse.data;
      }
    });
  }

  nextClick() {
    if (this.editableForm.invalid) { return; }
    this.dataModel.conftypecode = SERVICE_CONF_TYPE.SERVICE_DAILY_CHART;

    this.chartConfigureService.setDataModel(this.dataModel);
    this.dynamicContextService.onAction(true);
  }
}
