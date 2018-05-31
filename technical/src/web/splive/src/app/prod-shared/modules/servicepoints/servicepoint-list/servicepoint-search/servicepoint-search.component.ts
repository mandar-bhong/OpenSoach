import { Component, OnInit } from '@angular/core';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { ServicepointFilterModel } from '../../../../models/ui/service-configuration-models';
import { SpServiceConfService } from '../../../../services/spservice/sp-service-conf.service';

@Component({
  selector: 'app-servicepoint-search',
  templateUrl: './servicepoint-search.component.html',
  styleUrls: ['./servicepoint-search.component.css']
})
export class ServicepointSearchComponent implements OnInit {
  dataModel = new ServicepointFilterModel();
  isExpanded = false;
  spStates: EnumDataSourceItem<number>[];
  constructor(public spServiceConfService: SpServiceConfService) { }

  ngOnInit() {
    this.spStates = this.spServiceConfService.getServicepointStates();
  }
  search() {
    this.isExpanded = false;
    const servicepointFilterModel = new ServicepointFilterModel();
    this.dataModel.copyTo(servicepointFilterModel);
    this.spServiceConfService.dataListSubjectTrigger(servicepointFilterModel);
  }
  panelOpened() {
    this.isExpanded = true;
  }
}
