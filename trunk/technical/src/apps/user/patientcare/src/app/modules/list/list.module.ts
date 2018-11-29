import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { HttpClientModule } from '@angular/common/http';

import { ListRoutingModule } from "./list-routing.module";
import { ListComponent } from "./list.component";
import { DetailsComponent } from "./details/details.component";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular";

@NgModule({
    imports: [
        NativeScriptCommonModule,
        ListRoutingModule,
        HttpClientModule,
        NativeScriptUIListViewModule
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
