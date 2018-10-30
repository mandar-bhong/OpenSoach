import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";

import { ListComponent } from "./list.component";
import { DetailsComponent } from "./details/details.component";


const routes: Routes = [
    { path: "", component: ListComponent },
    { path: "details", component: DetailsComponent }
];

@NgModule({
    imports: [NativeScriptRouterModule.forChild(routes)],
    exports: [NativeScriptRouterModule]
})
export class ListRoutingModule { }
