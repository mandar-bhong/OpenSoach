import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { DeviceAddRequest } from '../../../models/api/device-models';
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
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.deviceStates = this.deviceService.getDeviceStates();

  }
  createControls(): void {
    this.editableForm = new FormGroup({
      serialnoControl: new FormControl('', [Validators.required]),
      devStateControl: new FormControl('', [Validators.required])
    });
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.recordState = EDITABLE_RECORD_STATE.ADD;
      this.setFormMode(FORM_MODE.EDITABLE);
      this.callbackUrl = params['callbackurl'];
    });
  }

  save() {
    const deviceAddRequest = new DeviceAddRequest();
    this.dataModel.copyTo(deviceAddRequest);
    this.deviceService.addDevice(deviceAddRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.appNotificationService.success(this.translatePipe.transform('SUCCESS_ADD_DEVICE_SAVED'));
          this.recordState = EDITABLE_RECORD_STATE.ADD;
          this.setFormMode(FORM_MODE.VIEW);
        } else {
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
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
}
