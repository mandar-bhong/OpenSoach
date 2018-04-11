import { CommonModule } from '@angular/common';
import { ModuleWithProviders, NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

import { AuthGuard } from './auth-guard';
import { AppDataStoreService } from './services/app-data-store/app-data-store-service';
import { LoginStatusService } from './services/login-status.service';
import { AuthService } from './services/auth.service';
import { ServerApiInterfaceService } from './services/api/server-api-interface.service';
import { TranslatePipe } from './pipes/translate/translate.pipe';
import { SidebarToggleService } from './services/sidebar-toggle.service';
import { TranslateService } from './pipes/translate/translate.service';
import { ApiErrorService } from './services/api/api-error.service';
import { AppNotificationService } from './services/notification/app-notification.service';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
  ],
  declarations: [
    TranslatePipe
  ],
  exports: [
    TranslatePipe
  ]
})
export class AppCommonModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: AppCommonModule,
      providers: [
        AppDataStoreService,
        LoginStatusService,
        ServerApiInterfaceService,
        AuthService,
        AuthGuard,
        SidebarToggleService,
        TranslatePipe,
        TranslateService,
        ApiErrorService,
        AppNotificationService
      ]
    };
  }
}
