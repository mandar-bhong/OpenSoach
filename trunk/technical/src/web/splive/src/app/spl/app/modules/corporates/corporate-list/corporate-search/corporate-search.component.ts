import { Component, OnInit } from '@angular/core';
import { FormsModule, ReactiveFormsModule, FormControl, FormGroup, Validators } from '@angular/forms';

import { CorporateShortDataResponse, CorporateFilterRequest } from '../../../../models/api/corporate-models';
import { CorporateFilterModel } from '../../../../models/ui/corporate-models';
import { CorporateService } from '../../../../services/corporate.service';

@Component({
  selector: 'app-corporate-search',
  templateUrl: './corporate-search.component.html',
  styleUrls: ['./corporate-search.component.css']
})
export class CorporateSearchComponent implements OnInit {
  dataModel = new CorporateFilterModel();
  isExpanded = false;
  corporates: CorporateShortDataResponse[] = [];
  constructor(private corporateService: CorporateService) { }

  ngOnInit() {
    this.getCorporateList();
  }
  getCorporateList() {
    this.corporateService.getCorporateShortDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.corporates = payloadResponse.data;
      }
    });
  }
  search() {
    this.isExpanded = false;
    const corporateFilterRequest = new CorporateFilterRequest();
    this.dataModel.copyTo(corporateFilterRequest);
    this.corporateService.dataListSubjectTrigger(corporateFilterRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }

}
