import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";

import { ListComponent } from "./list.component";
import { DetailsComponent } from "./details/details.component";
import { CamerasComponent } from "~/app/modules/list/cameras/cameras.component";
import { PatientDetailsComponent } from "~/app/modules/list/patient-details/patient-details.component";
import { MonitorComponent } from "~/app/modules/list/monitor/monitor.component";
import { ActionComponent } from "~/app/modules/list/action/action.component";
import { ChartsComponent } from "~/app/modules/list/charts/charts.component";


const routes: Routes = [
    { 
        path: "", component: ListComponent 
    },
    { 
        path: "details", component: DetailsComponent 
    },
    { 
        path: "cameras", component: CamerasComponent
    },
    { 
        path: "patient", component: PatientDetailsComponent
    },
    { 
        path: "monitor", component: MonitorComponent
    },
    { 
        path: "action", component: ActionComponent
    },
    { 
        path: "charts", component: ChartsComponent
    },
];

@NgModule({
    imports: [NativeScriptRouterModule.forChild(routes)],
    exports: [NativeScriptRouterModule]
})
export class ListRoutingModule { }
