import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { HttpClientModule } from '@angular/common/http';

import { ListRoutingModule } from "./list-routing.module";
import { ListComponent } from "./list.component";
import { DetailsComponent } from "./details/details.component";
// import { NativeScriptUISideDrawerModule } from "nativescript-ui-sidedrawer/angular";


@NgModule({
    imports: [
        NativeScriptCommonModule,
        ListRoutingModule,
        HttpClientModule,
        // NativeScriptUISideDrawerModule
    ],
    declarations: [
        ListComponent,
        DetailsComponent
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class ListModule { }
