import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { HttpClientModule } from '@angular/common/http';

import { HomeRoutingModule } from "./home-routing.module";
import { HomeComponent } from "./home.component";
import {NativeScriptFormsModule} from "nativescript-angular/forms"
import { NetworkStatusComponent } from "~/app/network-status.component";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular/listview-directives";

@NgModule({
    imports: [
        NativeScriptCommonModule,
        HomeRoutingModule,
        NativeScriptFormsModule,
        HttpClientModule,
        NativeScriptUIListViewModule
    ],
    declarations: [
        HomeComponent,
        NetworkStatusComponent
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class HomeModule { }
