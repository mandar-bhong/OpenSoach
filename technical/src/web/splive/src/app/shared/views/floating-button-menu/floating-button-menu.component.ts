import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { FloatingMenu, FloatingMenuItem } from '../../models/ui/floating-menu';
import { FloatingButtonMenuService } from '../../services/floating-button-menu.service';

@Component({
  selector: 'app-floating-button-menu',
  templateUrl: './floating-button-menu.component.html',
  styleUrls: ['./floating-button-menu.component.css']
})
export class FloatingButtonMenuComponent implements OnInit, OnDestroy {

  @Input()
  menuInstanceKey: string;
  items = new Array<FloatingMenuItem>();
  floatingMenuSubscription: Subscription;
  routerEventSubscription: Subscription;
  constructor(private floatingButtonMenuService: FloatingButtonMenuService,
    private router: Router) {
    this.routerEventSubscription = this.router.events
      .subscribe((event) => {
        if (event instanceof NavigationEnd) {
          this.items.length = 0;
        }
      });

    this.floatingMenuSubscription = this.floatingButtonMenuService.floatingMenuSubject.subscribe(floatingMenu => {
      this.setFloatingMenu(floatingMenu);
    });
  }

  setFloatingMenu(floatingMenu: FloatingMenu) {
    if (this.menuInstanceKey === floatingMenu.menuInstanceKey) {
      floatingMenu.items.forEach(item => this.items.push(item));
    }
  }

  ngOnInit() {
  }

  ngOnDestroy() {
    this.routerEventSubscription.unsubscribe();
    this.floatingMenuSubscription.unsubscribe();
  }

  menuItemClick(item: FloatingMenuItem) {
    if (item.navigate) {
      this.router.navigate([item.data], { skipLocationChange: true });
    }
  }
}
