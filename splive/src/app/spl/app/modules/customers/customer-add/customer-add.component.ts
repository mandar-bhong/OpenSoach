import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { CustomerMasterResponse, CustomerMasterUpdateRequest } from '../../../../../shared/models/api/customer-models';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { CorporateShortDataResponse } from '../../../models/api/corporate-models';
import { CustomerAddRequest } from '../../../models/api/customer-models';
import { CustomerAddModel } from '../../../models/ui/customer-models';
import { CorporateService } from '../../../services/corporate.service';
import { CustomerService } from '../../../services/customer.service';

@Component({
  selector: 'app-customer-add',
  templateUrl: './customer-add.component.html',
  styleUrls: ['./customer-add.component.css']
})
export class CustomerAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new CustomerAddModel();
  customerStates: EnumDataSourceItem<number>[];
  corporates: CorporateShortDataResponse[] = [];
  routeSubscription: Subscription;
  currentRecord: CustomerMasterResponse;
  constructor(private customerService: CustomerService,
    private corporateService: CorporateService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.customerStates = this.customerService.getCustomerStates();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.custid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getCustomerEdit();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
    this.getCorporateList();

  }
  createControls(): void {
    this.editableForm = new FormGroup({
      emailControl: new FormControl('', [Validators.required]),
      userState: new FormControl('', [Validators.required]),
      corprateId: new FormControl('', [Validators.required])
    });
  }

  getCorporateList() {
    this.corporateService.getCorporateShortDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.corporates = payloadResponse.data;
      }
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const customerAddRequest = new CustomerAddRequest();
      this.dataModel.copyTo(customerAddRequest);
      this.customerService.addCustomer(customerAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.custid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        }
      });
    } else {
      const request = new CustomerMasterUpdateRequest();
      this.dataModel.copyToUpdateRequest(request);
      this.customerService.updateCustomerMaster(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.setFormMode(FORM_MODE.VIEW);
        }
      });
    }
  }

  getCustomerEdit() {
    this.customerService.getCustomerMaster({ recid: this.dataModel.custid }).subscribe(payloadResponse => {
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

  edit() {
    this.editForm();
    this.editableForm.controls['corprateId'].disable();
  }
}
