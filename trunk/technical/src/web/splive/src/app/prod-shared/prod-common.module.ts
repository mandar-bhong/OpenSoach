import { CommonModule } from '@angular/common';
import { ModuleWithProviders, NgModule } from '@angular/core';

import { ProdDeviceService } from './services/device/prod-device.service';
import { ProdUserService } from './services/user/prod-user.service';
import { SplConfService } from './services/spl-conf.service';
import { SpServiceConfService } from './services/spservice/sp-service-conf.service';
import { ProdOperatorService } from './services/operator/prod-operator.service';
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
                ProdUserService,
                SpServiceConfService,
                ProdOperatorService,
                SplConfService

            ]
        };
    }
}
