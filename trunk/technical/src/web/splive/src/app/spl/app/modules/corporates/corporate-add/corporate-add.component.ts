import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';

import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { CorporateService } from '../../../services/corporate.service';
import { CorporateShortDataResponse, CorporateUpdateRequest, CorpDetailsResponse } from '../../../models/api/corporate-models';
import { CorporateAddRequest } from '../../../models/api/corporate-models';
import { CorporateAddModel, CorporateDetailsModel } from '../../../models/ui/corporate-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';

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
          this.appNotificationService.success(this.translatePipe.transform('SUCCESS_USERS_DETAILS_SAVED'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        }
      });
    } else {
      const request = new CorporateUpdateRequest();
      this.dataModel.copyToUpdateRequest(request);
      this.corporateService.updateCorporateDetails(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success(this.translatePipe.transform('SUCCESS_USERS_DETAILS_SAVED'));
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
