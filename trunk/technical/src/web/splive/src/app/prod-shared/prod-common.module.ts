import { CommonModule } from '@angular/common';
import { ModuleWithProviders, NgModule } from '@angular/core';

import { ProdDeviceService } from './services/device/prod-device.service';
import { SpServiceConfService } from './services/spservice/sp-service-conf.service';

@NgModule({
    imports: [
        CommonModule,
    ],
    declarations: [
    ],
    exports: [
    ]
})
export class ProdCommonModule {
    static forRoot(): ModuleWithProviders {
        return {
            ngModule: ProdCommonModule,
            providers: [
                ProdDeviceService,
                SpServiceConfService
            ]
        };
    }
}
