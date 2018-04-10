import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs/Subscription';

import { SidebarToggleService } from '../../../services/sidebar-toggle.service';

@Component({
  selector: 'app-side-bar',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css'],

})
export class SideBarComponent implements OnInit, OnDestroy {
  menuToggleSubscription: Subscription;
  setclass: any;
  toggleCssClass: boolean;
  constructor(private sidebarToggleService: SidebarToggleService) { }

  ngOnInit() {
    this.toggleCssClass = true;
    this.menuToggleSubscription = this.sidebarToggleService.menuToggleSubject.subscribe(value => {
      this.toggleCssClass = !this.toggleCssClass;
      this.toggleMenu();
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

}
