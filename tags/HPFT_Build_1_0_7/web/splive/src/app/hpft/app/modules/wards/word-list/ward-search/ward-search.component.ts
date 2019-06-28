import { Component, OnInit } from '@angular/core';
import { ServicepointFilterModel } from '../../../../../../prod-shared/models/ui/servicepoint-models';
import { ProdServicepointService } from '../../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';
import { SPCategoriesShortDataResponse } from '../../../../../../prod-shared/models/api/servicepoint-models';

@Component({
  selector: 'app-ward-search',
  templateUrl: './ward-search.component.html',
  styleUrls: ['./ward-search.component.css']
})
export class WardSearchComponent implements OnInit {
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
