import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppDeviceService } from '../../../../../shared/services/device/app-device.service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { DeviceAddRequest, DeviceMasterUpdateResponse } from '../../../models/api/device-models';
import { DeviceAddModel } from '../../../models/ui/device-models';
import { DeviceService } from '../../../services/device.service';

@Component({
  selector: 'app-device-add',
  templateUrl: './device-add.component.html',
  styleUrls: ['./device-add.component.css']
})
export class DeviceAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new DeviceAddModel();
  deviceStates: EnumDataSourceItem<number>[];
  routeSubscription: Subscription;
  constructor(private deviceService: DeviceService,
    private deviceSharedService: AppDeviceService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
    this.iconCss = 'fa fa-tablet';
    this.pageTitle = 'Device Details';
  }

  ngOnInit() {
    this.createControls();
    this.deviceStates = this.deviceSharedService.getDeviceStates();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.devid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getDeviceMaster();
        this.removeControl();
      } else {
        this.subTitle = 'Add Details of Device';
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      serialnoControl: new FormControl('', [Validators.required]),
      devStateControl: new FormControl('', [Validators.required])
    });

  }
  getDeviceMaster() {
    this.deviceService.getDeviceMaster({ recid: this.dataModel.devid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.subTitle = this.dataModel.serialno;
        }
      }
    });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const deviceAddRequest = new DeviceAddRequest();
      this.dataModel.copyTo(deviceAddRequest);
      this.deviceService.addDevice(deviceAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.custid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.serialno;
          this.removeControl();
        }
      });
      this.inProgress = false;
    } else  {
      const request = new DeviceMasterUpdateResponse();
      this.dataModel.copyToUpdateRequest(request);
      this.deviceService.updateDeviceMaster(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.serialno;
        }
      });
      this.inProgress = false;
    }
  }

  getdeviceState(value: number) {
    if (this.deviceStates && value) {
      return this.deviceStates.find(a => a.value === value).text;
    }
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  editRecord(id: number) {
    this.router.navigate(['devices', 'add'], { queryParams: { id: id, callbackurl: 'devices' }, skipLocationChange: true });
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
  removeControl() {
    this.editableForm.removeControl('serialnoControl');
  }
}
