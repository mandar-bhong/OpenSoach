import { Component, OnInit } from '@angular/core';
import { OnDestroy } from '@angular/core/src/metadata/lifecycle_hooks';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/edit-record-base';
import { CustomerAddDetailsRequest } from '../../../../../shared/models/api/customer-models';
import { CustomerDetailsModel } from '../../../../../shared/models/ui/customer-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
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
  }

  ngOnInit() {
    this.createControls();

    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.custid = Number(params['id']);
      this.callbackUrl = params['callbackurl'];
      this.getCustomerDetails();
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      poc1nameControl: new FormControl('', [Validators.required]),
      poc1emailidControl: new FormControl('', [Validators.required]),
      poc1mobilenoControl: new FormControl('', [Validators.required]),
      poc2nameControl: new FormControl(''),
      poc2emailidControl: new FormControl(''),
      poc2mobilenoControl: new FormControl(''),
      addressControl: new FormControl('', [Validators.required]),
      addressstateControl: new FormControl('', [Validators.required]),
      addresscityControl: new FormControl('', [Validators.required]),
      addresspincodeControl: new FormControl('', [Validators.required])
    });

  }

  getCustomerDetails() {
    this.customerService.getCustomerDetails({ recid: this.dataModel.custid }).subscribe(payloadResponse => {
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
    const customerAddDetailsRequest = new CustomerAddDetailsRequest();
    this.dataModel.copyTo(customerAddDetailsRequest);
    this.customerService.updateCustomerDetails(customerAddDetailsRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_CUSTOMER_DETAILS_SAVED'));
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
