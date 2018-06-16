import { Component, OnInit } from '@angular/core';

import { DEFAULT_PAGE_MENU } from '../../../../shared/app-common-constants';
import { FloatingMenu, FloatingMenuItem } from '../../../../shared/models/ui/floating-menu';
import { FloatingButtonMenuService } from '../../../../shared/services/floating-button-menu.service';

@Component({
  selector: 'app-operator-list',
  templateUrl: './operator-list.component.html',
  styleUrls: ['./operator-list.component.css']
})
export class OperatorListComponent implements OnInit {

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
    item.title = 'Operator';
    item.navigate = true;
    item.url = '/foperators/add';
    item.data = { queryParams: { callbackurl: 'foperators' }, skipLocationChange: true };
    floatingMenu.items.push(item);
    this.floatingButtonMenuService.setFloatingMenu(floatingMenu);
  }
}
