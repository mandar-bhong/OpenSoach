import { Component, OnInit } from '@angular/core';

import { DEFAULT_PAGE_MENU } from '../../../../shared/app-common-constants';
import { FloatingMenu, FloatingMenuItem } from '../../../../shared/models/ui/floating-menu';
import { FloatingButtonMenuService } from '../../../../shared/services/floating-button-menu.service';
@Component({
  selector: 'app-device-list',
  templateUrl: './device-list.component.html',
  styleUrls: ['./device-list.component.css']
})
export class DeviceListComponent implements OnInit {

  constructor(private floatingButtonMenuService: FloatingButtonMenuService) { }

  ngOnInit() {
    this.setFloatingMenu();
  }
  setFloatingMenu() {
    const floatingMenu = new FloatingMenu();
    floatingMenu.menuInstanceKey = DEFAULT_PAGE_MENU;
    floatingMenu.items = new Array<FloatingMenuItem>();
    const item = new FloatingMenuItem();
    item.icon = 'add_circle';
    item.title = 'Device';
    item.navigate = true;
    item.url = 'devices/add';
    item.data = { queryParams: { callbackurl: 'devices' }, skipLocationChange: true };
    floatingMenu.items.push(item);
    this.floatingButtonMenuService.setFloatingMenu(floatingMenu);
  }
}
