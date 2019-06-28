import { Component, OnInit } from '@angular/core';

import { SpServiceConfService } from '../../../../../../prod-shared/services/spservice/sp-service-conf.service';
import { ServiceConfigureFilterRequest } from '../../../../../../prod-shared/models/api/service-configuration-models';
import { ServiceConfigurationModel } from '../../../../../../prod-shared/models/ui/service-configuration-models';


@Component({
  selector: 'app-chart-search',
  templateUrl: './chart-search.component.html',
  styleUrls: ['./chart-search.component.css']
})
export class ChartSearchComponent implements OnInit {
  dataModel = new ServiceConfigurationModel();
  isExpanded = false;
  constructor(private spServiceConfService: SpServiceConfService) { }

  ngOnInit() {
  }
  search() {
    this.isExpanded = false;
    const serviceConfigureFilterRequest = new ServiceConfigureFilterRequest();
    this.dataModel.copyTo(serviceConfigureFilterRequest);
    this.spServiceConfService.dataListSubjectTrigger(serviceConfigureFilterRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }
}
