import { Injectable } from '@angular/core';
import { NavigationStart, Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { Subscription } from 'rxjs/Subscription';

import { TranslatePipe } from '../../pipes/translate/translate.pipe';

@Injectable()
export class AppNotificationService {
    routerEventSubscription: Subscription;
    constructor(private translatePipe: TranslatePipe,
        private toastr: ToastrService,
        private router: Router) {
            this.clearMessagesOnNavigation();
         }

    clearMessagesOnNavigation() {
        this.routerEventSubscription = this.router.events
            .subscribe((event) => {
                if (event instanceof NavigationStart) {
                    this.toastr.clear();
                }
            });
    }
    success(content?: any, override?: any): any {
        return this.toastr.success(content, this.translatePipe.transform('AppNotificationSuccess'),
            { timeOut: 2000, extendedTimeOut: 1000 });
    }

    error(content?: any, override?: any): any {
        console.log('Notification Error', content);
        return this.toastr.error(content, this.translatePipe.transform('AppNotificationError'), override);
    }

    info(content?: any, override?: any): any {
        return this.toastr.info(content, this.translatePipe.transform('AppNotificationInformation'),
            { timeOut: 3000 });
    }

    warn(content?: any, override?: any): any {
        return this.toastr.warning(content, this.translatePipe.transform('AppNotificationWarning'), override);
    }
}

