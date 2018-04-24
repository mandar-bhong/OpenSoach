import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';
import { CustomerAddDetailsRequest, CustomerDetailsResponse } from '../../../../../shared/models/api/customer-models';
import { CustomerDetailsModel } from '../../../../../shared/models/ui/customer-models';
import { CustomerService } from '../../../services/customer.service';
import { ActivatedRoute } from '@angular/router';
import { OnDestroy } from '@angular/core/src/metadata/lifecycle_hooks';
import { Subscription } from 'rxjs/Subscription';
@Component({
  selector: 'app-customer-update-details',
  templateUrl: './customer-update-details.component.html',
  styleUrls: ['./customer-update-details.component.css']
})
export class CustomerUpdateDetailsComponent implements OnInit, OnDestroy {
  myForm: FormGroup;
  dataModel = new CustomerDetailsModel();
  routeSubscription: Subscription;
  constructor(private customerService: CustomerService,
    private route: ActivatedRoute) { }

  ngOnInit() {
    this.createControls();

    // TODO get customer details api call
    const customerDetailsResponse = new CustomerDetailsResponse();
    this.dataModel.copyFrom(customerDetailsResponse);

    this.routeSubscription = this.route.params.subscribe(params => {

    });
  }
  createControls(): void {
    this.myForm = new FormGroup({
      customerName: new FormControl('', [Validators.required]),
      emailControl: new FormControl('', [Validators.required]),
      mobileNumber: new FormControl('', [Validators.required]),
      customerAddress: new FormControl('', [Validators.required]),
      customerCity: new FormControl('', [Validators.required]),
      customerState: new FormControl('', [Validators.required]),
      pinCode: new FormControl('', [Validators.required])
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
}
