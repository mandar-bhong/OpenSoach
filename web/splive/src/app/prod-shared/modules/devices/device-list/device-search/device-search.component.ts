import { Component, OnInit } from '@angular/core';

import { DeviceFilterRequest } from '../../../../models/api/device-models';
import { DeviceFilterModel } from '../../../../models/ui/device-models';
import { ProdDeviceService } from '../../../../services/device/prod-device.service';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';

@Component({
  selector: 'app-device-search',
  templateUrl: './device-search.component.html',
  styleUrls: ['./device-search.component.css']
})
export class DeviceSearchComponent implements OnInit {
  dataModel = new DeviceFilterModel();
  isExpanded = false;
  connectionStates: EnumDataSourceItem<number>[];
  constructor(private prodDeviceService: ProdDeviceService) { }

  ngOnInit() {
    this.connectionStates = this.prodDeviceService.getConnectionStates();
  }
  search() {
    this.isExpanded = false;
    const deviceFilterRequest = new DeviceFilterRequest();
    this.dataModel.copyTo(deviceFilterRequest);
    this.prodDeviceService.dataListSubjectTrigger(deviceFilterRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }

}
