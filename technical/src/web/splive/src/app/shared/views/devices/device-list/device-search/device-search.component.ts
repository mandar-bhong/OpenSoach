import { Component, OnInit } from '@angular/core';

import { DeviceFilterRequest } from '../../../../models/api/device-models';
import { DeviceFilterModel } from '../../../../models/ui/device-models';
import { DeviceSharedService } from '../../../../services/device-shared.service';

@Component({
  selector: 'app-device-search',
  templateUrl: './device-search.component.html',
  styleUrls: ['./device-search.component.css']
})
export class DeviceSearchComponent implements OnInit {
  dataModel = new DeviceFilterModel();
  isExpanded = false;
  constructor(private deviceSharedService: DeviceSharedService) { }

  ngOnInit() {
  }
  search() {
    this.isExpanded = false;
    const deviceFilterRequest = new DeviceFilterRequest();
    this.dataModel.copyTo(deviceFilterRequest);
    this.deviceSharedService.dataListSubjectTrigger(deviceFilterRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }

}
