import { Component, OnInit } from '@angular/core';

import { DEFAULT_PAGE_MENU } from '../../../../../shared/app-common-constants';
import { FloatingMenu, FloatingMenuItem } from '../../../../../shared/models/ui/floating-menu';
import { FloatingButtonMenuService } from '../../../../../shared/services/floating-button-menu.service';

@Component({
  selector: 'app-customer-list',
  templateUrl: './customer-list.component.html',
  styleUrls: ['./customer-list.component.css']
})
export class CustomerListComponent implements OnInit {

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
    item.title = 'Customer';
    item.navigate = true;
    item.url = 'customers/add';
    item.data = { queryParams: { callbackurl: 'customers' }, skipLocationChange: true };
    floatingMenu.items.push(item);
    this.floatingButtonMenuService.setFloatingMenu(floatingMenu);
  }
}
