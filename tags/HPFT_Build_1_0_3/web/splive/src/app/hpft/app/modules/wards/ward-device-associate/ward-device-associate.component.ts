import { Component, Inject, OnInit, ChangeDetectorRef } from '@angular/core';
import { MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef } from '@angular/material';
import { DeviceListItemResponse } from '../../../../../prod-shared/models/api/device-models';
import { ProdDeviceService } from '../../../../../prod-shared/services/device/prod-device.service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { ProdServicepointService } from '../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';

@Component({
  selector: 'app-ward-device-associate',
  templateUrl: './ward-device-associate.component.html',
  styleUrls: ['./ward-device-associate.component.css']
})
export class WardDeviceAssociateComponent implements OnInit {
  devices: DeviceListItemResponse[] = [];
  selecteddevice: DeviceListItemResponse;
  spid: number;
  constructor(private bottomSheetRef: MatBottomSheetRef<WardDeviceAssociateComponent>,
    private deviceService: ProdDeviceService, private appNotificationService: AppNotificationService,
    private prodServicepointService: ProdServicepointService,
    private translatePipe: TranslatePipe,
    private changeDetectorRef: ChangeDetectorRef,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: any) {
    this.spid = Number(data);
    console.log('in associate');
  }

  ngOnInit() {
    this.getDeviceList();
  }
  getDeviceList() {
    this.deviceService.getDeviceList().subscribe(payloadResponse => {
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
            this.bottomSheetRef.dismiss({ devid: this.selecteddevice.devid, devname: this.selecteddevice.devname });
          }
          else{
            this.appNotificationService.error(this.translatePipe.transform('ALREADY_DEVICE_ADDED'));
          }
        });
    }
  }

  close() {
    this.bottomSheetRef.dismiss();
  }
}
