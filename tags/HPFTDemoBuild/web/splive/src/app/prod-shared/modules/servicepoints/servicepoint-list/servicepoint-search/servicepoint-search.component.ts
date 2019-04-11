import { Component, OnInit } from '@angular/core';

import { SPCategoriesShortDataResponse } from '../../../../models/api/servicepoint-models';
import { ServicepointFilterModel } from '../../../../models/ui/servicepoint-models';
import { ProdServicepointService } from '../../../../services/servicepoint/prod-servicepoint.service';

@Component({
  selector: 'app-servicepoint-search',
  templateUrl: './servicepoint-search.component.html',
  styleUrls: ['./servicepoint-search.component.css']
})
export class ServicepointSearchComponent implements OnInit {
  dataModel = new ServicepointFilterModel();
  isExpanded = false;
  categories: SPCategoriesShortDataResponse[] = [];
  constructor(public prodServicepointService: ProdServicepointService,
    private prodServicePointService: ProdServicepointService) { }

  ngOnInit() {
    // this.spStates = this.prodServicepointService.getServicepointStates();
    this.getCategoriesList();
  }
  getCategoriesList() {
    this.prodServicePointService.getCategoriesShortDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.categories = payloadResponse.data;
      }
    });
  }
  search() {
    this.isExpanded = false;
    const servicepointFilterModel = new ServicepointFilterModel();
    this.dataModel.copyTo(servicepointFilterModel);
    this.prodServicepointService.dataListSubjectTrigger(servicepointFilterModel);
  }
  panelOpened() {
    this.isExpanded = true;
  }
}
