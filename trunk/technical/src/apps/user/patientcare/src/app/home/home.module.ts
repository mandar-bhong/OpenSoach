import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { HttpClientModule } from '@angular/common/http';

import { HomeRoutingModule } from "./home-routing.module";
import { HomeComponent } from "./home.component";
import {NativeScriptFormsModule} from "nativescript-angular/forms"
import { NetworkStatusComponent } from "~/app/network-status.component";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular/listview-directives";
import { AppModule } from "~/app/app.module";
import { sharedModule } from "~/app/modules/shared-mudule";

@NgModule({
    imports: [
        NativeScriptCommonModule,
        HomeRoutingModule,
        NativeScriptFormsModule,
        HttpClientModule,
        NativeScriptUIListViewModule,     
        sharedModule 
    ],
    declarations: [
        HomeComponent
      
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class HomeModule { }
