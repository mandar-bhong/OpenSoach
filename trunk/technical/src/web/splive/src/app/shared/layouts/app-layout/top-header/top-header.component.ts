import { Component, OnInit } from '@angular/core';

import { LoginStatusService } from '../../../services/login-status.service';
import { SidebarToggleService } from '../../../services/sidebar-toggle.service';

@Component({
  selector: 'app-top-header',
  templateUrl: './top-header.component.html',
  styleUrls: ['./top-header.component.css']
})
export class TopHeaderComponent implements OnInit {
  menuFull = true;
  constructor(private loginStatusService: LoginStatusService, private sidebarToggleService: SidebarToggleService) { }

  ngOnInit() {
  }

  logout() {
    this.loginStatusService.logout();
  }
  toggleChange() {
    this.menuFull = !this.menuFull;
    this.sidebarToggleService.toggleMenu(this.menuFull);
  }
}
