import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';
import { CustomerAddRequest } from '../../../models/api/customer-models';
import { CustomerAddModel } from '../../../models/ui/customer-models';
import { CustomerService } from '../../../services/customer.service';
@Component({
  selector: 'app-customer-add',
  templateUrl: './customer-add.component.html',
  styleUrls: ['./customer-add.component.css']
})
export class CustomerAddComponent implements OnInit {
  myForm: FormGroup;
  dataModel = new CustomerAddModel();

  constructor(private customerService: CustomerService) { }

  ngOnInit() {
    this.createControls();

  }
  createControls(): void {
    this.myForm = new FormGroup({
      emailControl: new FormControl('', [Validators.required]),
      userState: new FormControl('', [Validators.required]),
      corprateId: new FormControl('', [Validators.required])
    });
  }
  save() {
    // if (this.myForm.valid) {
      console.log('form is valid');
      const customerAddRequest = new CustomerAddRequest();
      this.dataModel.copyTo(customerAddRequest);
      this.customerService.addCustomer(customerAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          // TODO: navigate to list
        }
      });

    // } else {
      console.log('invalid ');
      // console.log(this.myForm.errors);
      // console.log(this.customerAddRequest);
      // const controls = this.myForm.controls;
      // for (const name in controls){
      //   console.log(name,controls[name].status);
      // }
    }


}
