import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { ModuleWithProviders, NgModule } from '@angular/core';

import { AuthGuard } from './auth-guard';
import { TranslatePipe } from './pipes/translate/translate.pipe';
import { TranslateService } from './pipes/translate/translate.service';
import { ApiErrorService } from './services/api/api-error.service';
import { ServerApiInterfaceService } from './services/api/server-api-interface.service';
import { AppDataStoreService } from './services/app-data-store/app-data-store-service';
import { AuthService } from './services/auth.service';
import { CustomerSharedService } from './services/customer/customer-shared.service';
import { LoginHandlerService } from './services/login-handler.service';
import { LoginStatusProviderService } from './services/login-status-provider.service';
import { AppNotificationService } from './services/notification/app-notification.service';
import { SidebarToggleService } from './services/sidebar-toggle.service';
import { UserSharedService } from './services/user/user-shared.service';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule
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
        LoginHandlerService,
        LoginStatusProviderService,
        ServerApiInterfaceService,
        AuthService,
        AuthGuard,
        SidebarToggleService,
        TranslatePipe,
        TranslateService,
        ApiErrorService,
        AppNotificationService,
        UserSharedService,
        CustomerSharedService
      ]
    };
  }
}
