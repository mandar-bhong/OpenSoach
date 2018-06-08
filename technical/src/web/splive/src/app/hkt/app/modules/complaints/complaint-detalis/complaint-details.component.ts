import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { COMPLAINT_STATE } from '../../../../../shared/app-common-constants';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppDataStoreService } from '../../../../../shared/services/app-data-store/app-data-store-service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { ComplaintUpdateRequest } from '../../../models/api/complaint-models';
import { ComplaintDetailsModel } from '../../../models/ui/complaint-models';
import { ProdComplaintService } from '../../../services/complaint/prod-complaint.service';

@Component({
  selector: 'app-complaint-details',
  templateUrl: './complaint-details.component.html',
  styleUrls: ['./complaint-details.component.css']
})
export class ComplaintDetailsComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new ComplaintDetailsModel();
  routeSubscription: Subscription;
  complStates: EnumDataSourceItem<number>[];
  seveStates: EnumDataSourceItem<number>[];
  showCat = true;
  show = false;
  formModeEnum;
  constructor(private prodComplaintService: ProdComplaintService,
    private route: ActivatedRoute,
    private appDataStoreService: AppDataStoreService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe
  ) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.complStates = this.prodComplaintService.getComplaintStates();
    this.seveStates = this.prodComplaintService.getSeveritiesStates();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.complaintid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getComplaintDetails();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      complaintstateControl: new FormControl('', [Validators.required]),
      remarksControl: new FormControl(''),
      severityControl: new FormControl('', [Validators.required])
    });
  }
  getComplaintDetails() {
    this.prodComplaintService.getComplaintDetails({ recid: this.dataModel.complaintid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          if (this.dataModel.complaintstate === COMPLAINT_STATE.INPROGRESS) {
            this.complStates = this.complStates.filter(a => a.value !== COMPLAINT_STATE.OPEN);
          } else if (this.dataModel.complaintstate === COMPLAINT_STATE.CLOSE) {
            this.showCat = false;
            this.show = true;
          }
        } else {
          this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
        }
      }
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    const complaintUpdateRequest = new ComplaintUpdateRequest();
    this.dataModel.copyToUpdateRequest(complaintUpdateRequest);
    this.prodComplaintService.updateComplaintDetails(complaintUpdateRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_USERS_DETAILS_SAVED'));
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
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
