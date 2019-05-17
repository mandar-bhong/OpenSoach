import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NetworkStatusComponent } from '~/app/network-status.component';
import { PassDataService } from '~/app/services/pass-data-service';

@NgModule({
    declarations: [NetworkStatusComponent],
    imports: [],
    exports: [NetworkStatusComponent],
    providers: [],
})
export class sharedModule { }