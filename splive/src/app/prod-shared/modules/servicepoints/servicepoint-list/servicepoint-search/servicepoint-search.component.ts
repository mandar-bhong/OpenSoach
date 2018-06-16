import { Component, OnInit } from '@angular/core';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
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
  spStates: EnumDataSourceItem<number>[];
  constructor(public prodServicepointService: ProdServicepointService) { }

  ngOnInit() {
    this.spStates = this.prodServicepointService.getServicepointStates();
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
