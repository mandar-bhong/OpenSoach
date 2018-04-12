import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs/Subscription';

import { SidebarToggleService } from '../../../services/sidebar-toggle.service';
import { RoutingModel } from '../../../models/ui/routing-model';

@Component({
  selector: 'app-side-bar',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css'],

})
export class SideBarComponent implements OnInit, OnDestroy {
  sideMenuLinks: RoutingModel[];
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

    this.sideMenuLinks = [];
    let abc = new RoutingModel();
    abc.url = '/devices';
    abc.displayinsidemenu = true;
    abc.linkiconcss = 'fa fa-dashboard';
    abc.title = 'Devices';

    this.sideMenuLinks.push(abc);

    abc = new RoutingModel();
    abc.url = '/charts';
    abc.displayinsidemenu = true;
    abc.linkiconcss = 'fa fa-film';
    abc.title = 'Charts';

    this.sideMenuLinks.push(abc);

    abc = new RoutingModel();
    abc.url = '/item 3';
    abc.displayinsidemenu = true;
    abc.linkiconcss = 'fa fa-book';
    abc.title = 'item3';

    this.sideMenuLinks.push(abc);

    abc = new RoutingModel();
    abc.url = '/item 4';
    abc.displayinsidemenu = true;
    abc.linkiconcss = 'fa fa-heart';
    abc.title = 'item 4';

    this.sideMenuLinks.push(abc);
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
