import { Component, OnInit } from '@angular/core';
import { OnDestroy } from '@angular/core/src/metadata/lifecycle_hooks';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { CustomerAddDetailsRequest } from '../../../../../shared/models/api/customer-models';
import { CustomerDetailsModel } from '../../../../../shared/models/ui/customer-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { CustomerService } from '../../../services/customer.service';

@Component({
  selector: 'app-customer-update-details',
  templateUrl: './customer-update-details.component.html',
  styleUrls: ['./customer-update-details.component.css']
})
export class CustomerUpdateDetailsComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new CustomerDetailsModel();
  routeSubscription: Subscription;
  constructor(private customerService: CustomerService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Customer Details';
  }

  ngOnInit() {
    this.createControls();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.custid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getCustomerDetails();
        this.subTitle = 'Add Details of Customer';
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      poc1nameControl: new FormControl('', [Validators.required]),
      poc1emailidControl: new FormControl('', [Validators.required, Validators.email]),
      poc1mobilenoControl: new FormControl('', [Validators.required, Validators.pattern(/^\d+$/)]),
      poc2nameControl: new FormControl(''),
      poc2emailidControl: new FormControl('', Validators.email),
      poc2mobilenoControl: new FormControl('', [Validators.pattern(/^\d+$/)]),
      addressControl: new FormControl('', [Validators.required]),
      addressstateControl: new FormControl('', [Validators.required]),
      addresscityControl: new FormControl('', [Validators.required]),
      addresspincodeControl: new FormControl('', [Validators.required, Validators.pattern(/^\d+$/)])
    });

  }

  getCustomerDetails() {
    this.customerService.getCustomerDetails({ recid: this.dataModel.custid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.poc1name;
        } else {
          this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.ADD;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    const customerAddDetailsRequest = new CustomerAddDetailsRequest();
    this.dataModel.copyTo(customerAddDetailsRequest);
    this.customerService.updateCustomerDetails(customerAddDetailsRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.subTitle = this.dataModel.poc1name;
      }
    });
    this.inProgress = false;
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
