import { Component, OnInit } from '@angular/core';

import { CustomerFilterRequest } from '../../../../models/api/customer-models';
import { CustomerFilterModel } from '../../../../models/ui/customer-models';
import { CustomerService } from '../../../../services/customer.service';

@Component({
  selector: 'app-customer-search',
  templateUrl: './customer-search.component.html',
  styleUrls: ['./customer-search.component.css']
})
export class CustomerSearchComponent implements OnInit {

  dataModel = new CustomerFilterModel();
  isExpanded = false;
  constructor(private customerService: CustomerService) { }

  ngOnInit() {
  }

  search() {
    this.isExpanded = false;
    const customerFilterRequest = new CustomerFilterRequest();
    this.dataModel.copyTo(customerFilterRequest);
    this.customerService.dataListSubjectTrigger(customerFilterRequest);
  }

  panelOpened() {
    this.isExpanded = true;
  }

}
