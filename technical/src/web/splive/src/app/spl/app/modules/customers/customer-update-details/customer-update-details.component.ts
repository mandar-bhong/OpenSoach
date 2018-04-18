import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';
import { CustomerAddDetailsRequest, CustomerDetailsResponse } from '../../../../../shared/models/api/customer-models';
import { CustomerDetailsModel } from '../../../../../shared/models/ui/customer-models';
import { CustomerService } from '../../../services/customer.service';
@Component({
  selector: 'app-customer-update-details',
  templateUrl: './customer-update-details.component.html',
  styleUrls: ['./customer-update-details.component.css']
})
export class CustomerUpdateDetailsComponent implements OnInit {
  myForm: FormGroup;
  dataModel = new CustomerDetailsModel();
  constructor(private customerService: CustomerService) { }

  ngOnInit() {
    this.createControls();

    // TODO get customer details api call
    const customerDetailsResponse = new CustomerDetailsResponse();
    this.dataModel.copyFrom(customerDetailsResponse);

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
}
