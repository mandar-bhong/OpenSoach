import { CommonModule } from '@angular/common';
import { ModuleWithProviders, NgModule } from '@angular/core';

import { AuthGuard } from './auth-guard';
import { AppDataStoreService } from './services/app-data-store/app-data-store-service';
import { LoginStatusService } from './services/login-status.service';
import { SidebarToggleService } from './services/sidebar-toggle.service';

@NgModule({
  imports: [
    CommonModule,
  ],
  declarations: [
  ],
})
export class AppCommonModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: AppCommonModule,
      providers: [
        AppDataStoreService,
        LoginStatusService,
        AuthGuard,
        SidebarToggleService
      ]
    };
  }
}
