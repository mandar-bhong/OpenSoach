import { NgModule, ModuleWithProviders } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AppDataStoreService } from './services/app-data-store/app-data-store-service';

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
            AppDataStoreService
          ]
      };
  }
}
