import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";

import { HomeComponent } from "./home.component";
import { PatientMonitoreUnmonitoreListComponent } from "./patient-monitore-unmonitore-list/patient-monitore-unmonitore-list.component";


const routes: Routes = [
    { path: "", component: HomeComponent },
    { path: "monitore", component: PatientMonitoreUnmonitoreListComponent },
];

@NgModule({
    imports: [NativeScriptRouterModule.forChild(routes)],
    exports: [NativeScriptRouterModule]
})
export class HomeRoutingModule { }
