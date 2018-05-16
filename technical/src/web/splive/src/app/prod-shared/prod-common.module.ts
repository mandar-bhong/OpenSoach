import { CommonModule } from '@angular/common';
import { ModuleWithProviders, NgModule } from '@angular/core';

import { ProdDeviceService } from './services/device/prod-device.service';

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
                ProdDeviceService
            ]
        };
    }
}
