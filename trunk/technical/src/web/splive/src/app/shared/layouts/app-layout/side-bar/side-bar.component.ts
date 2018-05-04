import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { AppSpecificDataProvider } from '../../../app-specific-data-provider';
import { SideMenuModel } from '../../../models/ui/routing-model';
import { SidebarToggleService } from '../../../services/sidebar-toggle.service';

@Component({
  selector: 'app-side-bar',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css'],

})
export class SideBarComponent implements OnInit, OnDestroy {
  sideMenuLinks: SideMenuModel[];
  menuToggleSubscription: Subscription;
  setclass: any;
  toggleCssClass: boolean;
  constructor(private sidebarToggleService: SidebarToggleService,
    private router: Router) { }

  ngOnInit() {
    this.toggleCssClass = true;
    this.menuToggleSubscription = this.sidebarToggleService.menuToggleSubject.subscribe(value => {
      this.toggleCssClass = !this.toggleCssClass;
      this.toggleMenu();
    });

    this.sideMenuLinks = AppSpecificDataProvider.sideMenuRoutes;
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
      console.log('this.router.url', this.router.url);
      return 'selectedMenuItem';
    }
  }
}
