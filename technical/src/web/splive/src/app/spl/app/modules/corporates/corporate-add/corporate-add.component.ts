import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { CorporateAddRequest, CorporateUpdateRequest } from '../../../models/api/corporate-models';
import { CorporateDetailsModel } from '../../../models/ui/corporate-models';
import { CorporateService } from '../../../services/corporate.service';

@Component({
  selector: 'app-corporate-add',
  templateUrl: './corporate-add.component.html',
  styleUrls: ['./corporate-add.component.css']
})
export class CorporateAddComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new CorporateDetailsModel();
  routeSubscription: Subscription;
  constructor(private corporateService: CorporateService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }
  ngOnInit() {
    this.createControls();
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      corpnameControl: new FormControl('', [Validators.required]),
      corpmobilenoControl: new FormControl('', [Validators.required]),
      corpemailidControl: new FormControl('', [Validators.required]),
      corplandlinenoControl: new FormControl('', [Validators.required])
    });
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.corpid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getCorporateDetails();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const request = new CorporateAddRequest();
      this.dataModel.copyToAddRequest(request);
      this.corporateService.addCorporate(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.corpid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        }
      });
    } else {
      const request = new CorporateUpdateRequest();
      this.dataModel.copyToUpdateRequest(request);
      this.corporateService.updateCorporateDetails(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.setFormMode(FORM_MODE.VIEW);
        }
      });
    }
  }

  getCorporateDetails() {
    this.corporateService.getCorporateDetails({ recid: this.dataModel.corpid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
        }
      }
    });
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }

}
