import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NetworkStatusComponent } from '~/app/network-status.component';
import { PassDataService } from '~/app/services/pass-data-service';
import { OsSelectionListComponent } from '../os-selection-list.component';

import { NativeScriptUIListViewModule } from 'nativescript-ui-listview/angular/listview-directives';
import { NativeScriptFormsModule } from 'nativescript-angular/forms';
import { NativeScriptCommonModule } from 'nativescript-angular/common';

@NgModule({
    declarations: [
        NetworkStatusComponent,
        OsSelectionListComponent
    ],
    imports: [
        NativeScriptUIListViewModule,
        NativeScriptFormsModule,
        NativeScriptCommonModule
    ],
    exports: [
        NetworkStatusComponent,
        OsSelectionListComponent
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ],
    entryComponents: [
        OsSelectionListComponent
    ], 
    providers: [],
})
export class sharedModule { }