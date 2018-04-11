// import { NotificationsService, Notification } from 'angular2-notifications';
import { Injectable } from '@angular/core';
import { TranslatePipe } from '../../pipes/translate/translate.pipe';

@Injectable()
export class AppNotificationService {
    constructor(private translatePipe: TranslatePipe) { }

    success(content?: any, override?: any): any {
        console.log(this.translatePipe.transform('AppNotificationSuccess'), content);
        // return this.notificationService.success(this.translatePipe.transform('AppNotificationSuccess'), content, override);
    }

    error(content?: any, override?: any): any {
        console.log(this.translatePipe.transform('AppNotificationError'), content);
        // return this.notificationService.error(this.translatePipe.transform('AppNotificationError'), content, override);
    }

    alert(content?: any, override?: any): any {
        console.log(this.translatePipe.transform('AppNotificationAlert'), content);
        // return this.notificationService.alert(this.translatePipe.transform('AppNotificationAlert'), content, override);
    }

    info(content?: any, override?: any): any {
        console.log(this.translatePipe.transform('AppNotificationInformation'), content);
        // return this.notificationService.info(this.translatePipe.transform('AppNotificationInformation'), content, override);
    }

    warn(content?: any, override?: any): any {
        console.log(this.translatePipe.transform('AppNotificationWarning'), content);
        // return this.notificationService.warn(this.translatePipe.transform('AppNotificationWarning'), content, override);
    }
}
