import { Component, Inject, OnInit, ChangeDetectorRef } from '@angular/core';
import { MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef } from '@angular/material';

import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { DeviceListItemResponse } from '../../../models/api/device-models';
import { ProdDeviceService } from '../../../services/device/prod-device.service';
import { ProdServicepointService } from '../../../services/servicepoint/prod-servicepoint.service';

@Component({
  selector: 'app-servicepoint-device-associate',
  templateUrl: './servicepoint-device-associate.component.html',
  styleUrls: ['./servicepoint-device-associate.component.css']
})
export class ServicepointDeviceAssociateComponent implements OnInit {

  devices: DeviceListItemResponse[] = [];
  selecteddevice: DeviceListItemResponse;
  spid: number;
  constructor(private bottomSheetRef: MatBottomSheetRef<ServicepointDeviceAssociateComponent>,
    private deviceService: ProdDeviceService, private appNotificationService: AppNotificationService,
    private prodServicepointService: ProdServicepointService,
    private changeDetectorRef: ChangeDetectorRef,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: any) {
    this.spid = Number(data);
    console.log('in associate');
  }

  ngOnInit() {
    this.getDeviceList();
  }

  getDeviceList() {
    this.deviceService.getDevicesNotAssociatedWithSP().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.devices = payloadResponse.data;
        // TODO: This is work around for the bug in angular material: #11351 Mat-sheet can not update mat-field from promise.
        this.changeDetectorRef.markForCheck();
      }
    });
  }

  save() {
    console.log('in save');
    if (this.selecteddevice != null) {
      this.prodServicepointService.associateDeviceServicePoint(
        { devid: this.selecteddevice.devid, spid: this.spid }).subscribe(payloadResponse => {
          if (payloadResponse && payloadResponse.issuccess) {
            this.appNotificationService.success();
            console.log('in dismiss', this.selecteddevice.devid);
            this.bottomSheetRef.dismiss({ devid: this.selecteddevice.devid, devname: this.selecteddevice.devname });
          }
        });
    }
  }

  close() {
    this.bottomSheetRef.dismiss();
  }

}
