import { Component, OnInit } from '@angular/core';

import { DEFAULT_PAGE_MENU } from '../../../../../shared/app-common-constants';
import { FloatingMenu, FloatingMenuItem } from '../../../../../shared/models/ui/floating-menu';
import { FloatingButtonMenuService } from '../../../../../shared/services/floating-button-menu.service';

@Component({
  selector: 'app-patient-list',
  templateUrl: './patient-list.component.html',
  styleUrls: ['./patient-list.component.css']
})
export class PatientListComponent implements OnInit {

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
    item.title = 'Patient';
    item.navigate = true;
    item.url = 'patients/patient_search';
    item.data = { queryParams: { callbackurl: 'patients' }, skipLocationChange: true };
    floatingMenu.items.push(item);
    this.floatingButtonMenuService.setFloatingMenu(floatingMenu);
  }

}
