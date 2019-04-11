import { Component, OnInit } from '@angular/core';

import { EnumDataSourceItem } from '../../../../../../shared/models/ui/enum-datasource-item';
import { CorporateShortDataResponse } from '../../../../models/api/corporate-models';
import { CustomerFilterRequest } from '../../../../models/api/customer-models';
import { CustomerFilterModel } from '../../../../models/ui/customer-models';
import { CorporateService } from '../../../../services/corporate.service';
import { CustomerService } from '../../../../services/customer.service';

@Component({
  selector: 'app-customer-search',
  templateUrl: './customer-search.component.html',
  styleUrls: ['./customer-search.component.css']
})
export class CustomerSearchComponent implements OnInit {

  dataModel = new CustomerFilterModel();
  isExpanded = false;
  corporates: CorporateShortDataResponse[] = [];
  customerStates: EnumDataSourceItem<number>[];
  constructor(private customerService: CustomerService,
    private corporateService: CorporateService) { }

  ngOnInit() {
    this.getCorporateList();
    this.customerStates = this.customerService.getCustomerStates();
  }

  search() {
    this.isExpanded = false;
    const customerFilterRequest = new CustomerFilterRequest();
    this.dataModel.copyTo(customerFilterRequest);
    this.customerService.dataListSubjectTrigger(customerFilterRequest);
  }

  getCorporateList() {
    this.corporateService.getCorporateShortDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.corporates = payloadResponse.data;
      }
    });
  }
  panelOpened() {
    this.isExpanded = true;
  }

}
