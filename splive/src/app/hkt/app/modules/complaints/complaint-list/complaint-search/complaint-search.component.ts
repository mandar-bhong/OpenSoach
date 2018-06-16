import { Component, OnInit } from '@angular/core';

import { ComplaintFiltrRequest } from '../../../../models/api/complaint-models';
import { ComplaintFilterModel } from '../../../../models/ui/complaint-models';
import { ProdComplaintService } from '../../../../services/complaint/prod-complaint.service';
import { EnumDataSourceItem } from '../../../../../../shared/models/ui/enum-datasource-item';

@Component({
  selector: 'app-complaint-search',
  templateUrl: './complaint-search.component.html',
  styleUrls: ['./complaint-search.component.css']
})
export class ComplaintSearchComponent implements OnInit {
  dataModel = new ComplaintFilterModel();
  isExpanded = false;
  compStates: EnumDataSourceItem<number>[];
  constructor(private prodComplaintService: ProdComplaintService) { }

  ngOnInit() {
    this.compStates = this.prodComplaintService.getComplaintStates();
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
