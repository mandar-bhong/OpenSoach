import { Component, Inject, OnInit, ChangeDetectorRef } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef } from '@angular/material';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EditRecordBase } from '../../../../shared/views/edit-record-base';
import { DeviceDetailsUpdateRequest, DeviceDataListResponse } from '../../../models/api/device-models';
import { DeviceDetailsModel } from '../../../models/ui/device-models';
import { ProdDeviceService } from '../../../services/device/prod-device.service';


@Component({
  selector: 'app-device-details-view',
  templateUrl: './device-details-view.component.html',
  styleUrls: ['./device-details-view.component.css']
})
export class DeviceDetailsViewComponent extends EditRecordBase implements OnInit {
  dataModel = new DeviceDetailsModel();
  routeSubscription: Subscription;
  selecteddevice: DeviceDataListResponse;
  constructor(private bottomSheetRef: MatBottomSheetRef<DeviceDetailsViewComponent>,
    private prodDeviceService: ProdDeviceService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private changeDetectorRef: ChangeDetectorRef,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: any
  ) {
    super();
    this.dataModel.devid = Number(data);
  }

  ngOnInit() {
    this.createControls();
    this.getServicepointDetails();
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      devnameControl: new FormControl('', [Validators.required]),
    });
  }
  getServicepointDetails() {
    console.log('test');
    this.prodDeviceService.getDeviceDetails({ recid: this.dataModel.devid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          // TODO: This is work around for the bug in angular material: #11351 Mat-sheet can not update mat-field from promise.
          this.changeDetectorRef.markForCheck();
        }
      }
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    const deviceDetailsUpdateRequest = new DeviceDetailsUpdateRequest();
    this.dataModel.copyTo(deviceDetailsUpdateRequest);
    this.prodDeviceService.updateDeviceDetails(deviceDetailsUpdateRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        this.bottomSheetRef.dismiss({
          devid: this.dataModel.devid, devname: this.dataModel.devname,
        });
      }
    });
  }
  closeForm() {
    this.bottomSheetRef.dismiss();
  }
}
