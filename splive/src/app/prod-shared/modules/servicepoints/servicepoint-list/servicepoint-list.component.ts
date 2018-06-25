import { Component, OnInit } from '@angular/core';
import { FloatingMenu, FloatingMenuItem } from '../../../../shared/models/ui/floating-menu';
import { DEFAULT_PAGE_MENU } from '../../../../shared/app-common-constants';
import { FloatingButtonMenuService } from '../../../../shared/services/floating-button-menu.service';

@Component({
  selector: 'app-servicepoint-list',
  templateUrl: './servicepoint-list.component.html',
  styleUrls: ['./servicepoint-list.component.css']
})
export class ServicepointListComponent implements OnInit {

  constructor(private floatingButtonMenuService: FloatingButtonMenuService) { }

  ngOnInit() {
    this.setFloatingMenu();
  }

  setFloatingMenu() {
    const floatingMenu = new FloatingMenu();
    floatingMenu.menuInstanceKey = DEFAULT_PAGE_MENU;
    floatingMenu.items = new Array<FloatingMenuItem>();
    const item = new FloatingMenuItem();
    item.icon = 'view_list';
    item.title = 'Charts';
    item.navigate = true;
    item.url = 'servicepoints/charts';
    item.data = { queryParams: { callbackurl: 'servicepoints' }, skipLocationChange: true };
    floatingMenu.items.push(item);
    this.floatingButtonMenuService.setFloatingMenu(floatingMenu);
  }

}
