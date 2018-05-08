import { Component, OnInit, OnDestroy } from '@angular/core';
import { FormControl, Validators, FormBuilder, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../shared/views/edit-record-base';
import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { DeviceDetailsModel } from '../../../../spl/app/models/ui/device-models';
import { DeviceAddDetailsRequest } from '../../../../spl/app/models/api/device-models';
import { DeviceService } from '../../../../spl/app/services/device.service';
@Component({
  selector: 'app-device-update',
  templateUrl: './device-update.component.html',
  styleUrls: ['./device-update.component.css']
})
export class DeviceUpdateComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new DeviceDetailsModel();
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
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      console.log('devid recieved', params['id']);
      this.dataModel.devid = Number(params['id']);
      this.callbackUrl = params['callbackurl'];
      this.getDeviceDetails();
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      devnameControl: new FormControl('', [Validators.required]),
      makeControl: new FormControl('', [Validators.required]),
      technologyControl: new FormControl('', [Validators.required]),
      techversionControl: new FormControl('', [Validators.required])
    });
  }
  getDeviceDetails() {
    this.deviceService.getDeviceDetails({ recid: this.dataModel.devid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        } else {
          this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    const deviceAddDetailsRequest = new DeviceAddDetailsRequest();
    this.dataModel.copyTo(deviceAddDetailsRequest);
    this.deviceService.updateDeviceDetails(deviceAddDetailsRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_USERS_DETAILS_SAVED'));
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
      } else {}
    });
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  editRecord(id: number) {
    this.router.navigate(['devices', 'update'], { queryParams: { id: id, callbackurl: 'devices' }, skipLocationChange: true });
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
