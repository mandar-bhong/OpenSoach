import { Component, Inject, OnInit, OnDestroy } from '@angular/core';
import { MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef } from '@angular/material';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../shared/views/edit-record-base';

import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { ProdDeviceService } from '../../../services/device/prod-device.service';
import { ProdServicepointService } from '../../../services/servicepoint/prod-servicepoint.service';
import {
  ServicepointListResponse,
  ServicepointDetailsUpdateRequest, ServicepointDetailsResponse, SPCategoriesShortDataResponse
} from '../../../models/api/servicepoint-models';
import { ServicePointDetailsModel } from '../../../models/ui/servicepoint-models';
import { Subscription } from 'rxjs';
import { EnumDataSourceItem } from '../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { ServicepointConfigureListResponse } from '../../../models/api/service-configuration-models';

@Component({
  selector: 'app-servicepoint-update',
  templateUrl: './servicepoint-update.component.html',
  styleUrls: ['./servicepoint-update.component.css']
})
export class ServicepointUpdateComponent extends EditRecordBase implements OnInit {
  dataModel = new ServicePointDetailsModel();
  routeSubscription: Subscription;
  spStates: EnumDataSourceItem<number>[];
  categories: SPCategoriesShortDataResponse[] = [];

  constructor(
    private bottomSheetRef: MatBottomSheetRef<ServicepointUpdateComponent>,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private prodServicepointService: ProdServicepointService,
    private router: Router,
    private route: ActivatedRoute,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: any) {
    super();
    this.dataModel.spid = Number(data);
    console.log('in associate');
  }

  ngOnInit() {
    this.createControls();
    this.getCategoriesList();
    this.getServicepointDetails();
    this.spStates = this.prodServicepointService.getServicepointStates();
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      spnameControl: new FormControl('', [Validators.required]),
      spStateControl: new FormControl(''),
      categoryControl: new FormControl(''),
    });
  }
  getServicepointDetails() {
    console.log('test');
    this.prodServicepointService.getServicepointDetails({ recid: this.dataModel.spid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
        }
      }
    });
  }
  getCategoriesList() {
    this.prodServicepointService.getCategoriesShortDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.categories = payloadResponse.data;
      }
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    const servicepointDetailsUpdateRequest = new ServicepointDetailsUpdateRequest();
    this.dataModel.copyTo(servicepointDetailsUpdateRequest);
    this.prodServicepointService.updateServicepointDetails(servicepointDetailsUpdateRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        this.bottomSheetRef.dismiss({
          spid: this.dataModel.spid,
          spname: this.dataModel.spname,
          spstate: this.dataModel.spstate,
          spcid: this.dataModel.spcid,
          spcname: this.categories.find(c => c.spcid === this.dataModel.spcid).spcname
        });
      }
    });
  }
  closeForm() {
    this.bottomSheetRef.dismiss();
  }
}
