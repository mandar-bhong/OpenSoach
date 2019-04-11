import { Component, OnInit } from '@angular/core';

import { DEFAULT_PAGE_MENU } from '../../../../../shared/app-common-constants';
import { FloatingMenu, FloatingMenuItem } from '../../../../../shared/models/ui/floating-menu';
import { FloatingButtonMenuService } from '../../../../../shared/services/floating-button-menu.service';

@Component({
  selector: 'app-corporate-list',
  templateUrl: './corporate-list.component.html',
  styleUrls: ['./corporate-list.component.css']
})
export class CorporateListComponent implements OnInit {

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
    item.title = 'Corporate';
    item.navigate = true;
    item.url = 'corporates/add';
    item.data = { queryParams: { callbackurl: 'corporates' }, skipLocationChange: true };
    floatingMenu.items.push(item);
    this.floatingButtonMenuService.setFloatingMenu(floatingMenu);
  }
}
