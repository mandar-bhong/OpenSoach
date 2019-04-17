import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { AppSpecificDataProvider } from '../../../app-specific-data-provider';
import { SideMenuModel } from '../../../models/ui/routing-model';
import { SidebarToggleService } from '../../../services/sidebar-toggle.service';
import { AppUserService } from '../../../services/user/app-user.service';
import { PayloadResponse } from '../../../models/api/payload-models';
import { USER_ADMIN } from '../../../app-common-constants';

@Component({
  selector: 'app-side-bar',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css'],

})
export class SideBarComponent implements OnInit, OnDestroy {
  sideMenuLinks: SideMenuModel[] = [];
  menuToggleSubscription: Subscription;
  setclass: any;
  toggleCssClass: boolean;
  constructor(private sidebarToggleService: SidebarToggleService,
    private router: Router,
    private appUserService: AppUserService) { }

  ngOnInit() {
    this.toggleCssClass = true;
    this.menuToggleSubscription = this.sidebarToggleService.menuToggleSubject.subscribe(value => {
      this.toggleCssClass = !this.toggleCssClass;
      this.toggleMenu();
    });

    this.sideMenuLinks = AppSpecificDataProvider.sideMenuRoutes;
    this.adminForPatientCare();
  }

  adminForPatientCare() {
    this.appUserService.getLoginInfo().subscribe(PayloadResponse => {
      if (PayloadResponse.data.cpmrole === USER_ADMIN) {
        this.sideMenuLinks = AppSpecificDataProvider.sideMenuRoutes;
      } else {
        let specificLink = AppSpecificDataProvider.sideMenuRoutes.filter(a => a.url === '/dashboard' || a.url === '/patients');
        this.sideMenuLinks = specificLink;
      }
    });
  }

  ngOnDestroy(): void {
    if (this.menuToggleSubscription) {
      this.menuToggleSubscription.unsubscribe();
    }
  }
  toggleMenu(): void {
    if (!this.toggleCssClass) {
      this.setclass = 'toggle';
    } else {
      this.setclass = '';
    }
  }

  getSelectedMenuCss(url: string) {
    if (this.router.url.startsWith(url)) {
      return 'selectedMenuItem';
    }
  }
}
