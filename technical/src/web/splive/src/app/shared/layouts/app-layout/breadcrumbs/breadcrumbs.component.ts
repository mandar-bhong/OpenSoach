import { Component, OnDestroy } from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { DEFAULT_PAGE_MENU } from '../../../app-common-constants';
import { AppSpecificDataProvider } from '../../../app-specific-data-provider';

@Component({
  selector: 'app-breadcrumbs',
  templateUrl: './breadcrumbs.component.html',
  styleUrls: ['./breadcrumbs.component.css']
})
export class BreadcrumbsComponent implements OnDestroy {
  tempState = [];
  breadcrumbs: Array<Object>;
  routerEventSubscription: Subscription;
  PAGE_MENU = DEFAULT_PAGE_MENU;
  constructor(private router: Router, private route: ActivatedRoute) {
    this.buildBreadCrumb();
  }

  buildBreadCrumb() {
    this.routerEventSubscription = this.router.events
      .subscribe((event) => {
        if (event instanceof NavigationEnd) {
          this.breadcrumbs = [];
          this.tempState = [];
          let currentRoute = this.route.root,
            url = '';
          do {
            const childrenRoutes = currentRoute.children;
            currentRoute = null;
            childrenRoutes.forEach(routes => {
              const routeSnapshot = routes.snapshot;
              const currentUrl = routeSnapshot.url.map(segment => segment.path).join('/');
              if (currentUrl !== '') {
                url += '/' + currentUrl;
                if (!this.tempState.includes(url)) {
                  this.tempState.push(url);
                  this.setBreadCrumbItem(url);
                }
              }

              currentRoute = routes;
            });
          } while (currentRoute);
          if (this.breadcrumbs.length === 0) {
            this.setBreadCrumbItem('/dashboard');
          }
        }
      });
  }

  setBreadCrumbItem(url: string) {
    // get route details from routeMap
    const routingModel = AppSpecificDataProvider.appRoutes.get(url);
    if (routingModel) {
      this.breadcrumbs.push({
        breadcrumbtitle: routingModel.linktitle,
        pagetitle: routingModel.pagetitle,
        url: url,
        clickable: true
      });
    }
  }

  ngOnDestroy() {
    if (this.routerEventSubscription) {
      this.routerEventSubscription.unsubscribe();
    }
  }
}
