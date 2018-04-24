import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';
import { CustomerAddDetailsRequest, CustomerDetailsResponse } from '../../../../../shared/models/api/customer-models';
import { CustomerDetailsModel } from '../../../../../shared/models/ui/customer-models';
import { CustomerService } from '../../../services/customer.service';
import { ActivatedRoute } from '@angular/router';
import { OnDestroy } from '@angular/core/src/metadata/lifecycle_hooks';
import { Subscription } from 'rxjs/Subscription';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
@Component({
  selector: 'app-customer-update-details',
  templateUrl: './customer-update-details.component.html',
  styleUrls: ['./customer-update-details.component.css']
})
export class CustomerUpdateDetailsComponent implements OnInit, OnDestroy {
  myForm: FormGroup;
  dataModel = new CustomerDetailsModel();
  routeSubscription: Subscription;
  formMode = 0; // 0:view, 1:add, 2:update

  constructor(private customerService: CustomerService,
    private route: ActivatedRoute,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.createControls();

    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.custid = Number(params['id']);
      this.getCustomerDetails();
    });
  }
  createControls(): void {
    this.myForm = new FormGroup({
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

    this.setFormMode(0);
  }

  getCustomerDetails() {
    this.customerService.getCustomerDetails({ recid: this.dataModel.custid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          console.log('payloadResponse.data', payloadResponse.data);
          this.dataModel.copyFrom(payloadResponse.data);
          console.log('this.dataModel', this.dataModel);
        } else {
          this.appNotificationService.info(this.translatePipe.transform('InfoMessageDetailsNotAvailable'));
          this.setFormMode(1);
        }
      }
    });
  }

  save() {
    const customerAddDetailsRequest = new CustomerAddDetailsRequest();
    this.dataModel.copyTo(customerAddDetailsRequest);
    this.customerService.updateCustomerDetails(customerAddDetailsRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        // TODO: navigate to list
      }
    });
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }

  editForm() {
    this.setFormMode(2);
  }

  setFormMode(mode: number) {
    this.formMode = mode;
    switch (this.formMode) {
      case 0:
        this.myForm.disable();
        break;
      case 1:
      case 2:
        this.myForm.enable();
        break;
    }
  }
}
