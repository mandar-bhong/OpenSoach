import { Component, OnInit } from '@angular/core';

import { ServicepointListResponse } from '../../../../../../prod-shared/models/api/servicepoint-models';
import { ProdServicepointService } from '../../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';
import { EnumDataSourceItem } from '../../../../../../shared/models/ui/enum-datasource-item';
import { ComplaintFiltrRequest } from '../../../../models/api/complaint-models';
import { ComplaintFilterModel } from '../../../../models/ui/complaint-models';
import { ProdComplaintService } from '../../../../services/complaint/prod-complaint.service';

@Component({
  selector: 'app-complaint-search',
  templateUrl: './complaint-search.component.html',
  styleUrls: ['./complaint-search.component.css']
})
export class ComplaintSearchComponent implements OnInit {
  dataModel = new ComplaintFilterModel();
  isExpanded = false;
  splist: ServicepointListResponse[] = [];
  compStates: EnumDataSourceItem<number>[];
  constructor(private prodComplaintService: ProdComplaintService,
    private prodServicePointService: ProdServicepointService) { }

  ngOnInit() {
    this.compStates = this.prodComplaintService.getComplaintStates();
    this.getServicepointList();
  }
  getServicepointList() {
    this.prodServicePointService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.splist = payloadResponse.data;
      }
    });
  }
  search() {
    this.isExpanded = false;
    const complaintFiltrRequest = new ComplaintFiltrRequest();
    this.dataModel.copyTo(complaintFiltrRequest);
    this.prodComplaintService.dataListSubjectTrigger(complaintFiltrRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }

}
