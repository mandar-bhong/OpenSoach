import { Component, OnInit } from '@angular/core';

import { DeviceFilterModel } from '../../../../../spl/app/models/ui/device-models';
import { DeviceFilterRequest } from '../../../../../spl/app/models/api/device-models';
import { DeviceService } from '../../../../../spl/app/services/device.service';

@Component({
  selector: 'app-device-search',
  templateUrl: './device-search.component.html',
  styleUrls: ['./device-search.component.css']
})
export class DeviceSearchComponent implements OnInit {
  dataModel = new DeviceFilterModel();
  isExpanded = false;
  constructor(private deviceService: DeviceService) { }

  ngOnInit() {
  }
  search() {
    this.isExpanded = false;
    const deviceFilterRequest = new DeviceFilterRequest();
    this.dataModel.copyTo(deviceFilterRequest);
    this.deviceService.dataListSubjectTrigger(deviceFilterRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }

}
