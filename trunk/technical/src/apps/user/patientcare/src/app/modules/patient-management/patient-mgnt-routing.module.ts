import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";
import { PatientMgntComponent } from "~/app/modules/patient-management/patient-mgnt.component";
import { DetailsComponent } from "~/app/modules/patient-management/details/details.component";
import { CamerasComponent } from "~/app/modules/patient-management/cameras/cameras.component";
import { PatientDetailsComponent } from "~/app/modules/patient-management/patient-details/patient-details.component";
import { MonitorComponent } from "~/app/modules/patient-management/monitor/monitor.component";
import { ActionComponent } from "~/app/modules/patient-management/action/action.component";
import { ChartsComponent } from "~/app/modules/patient-management/charts/charts.component";
import { MedicineChartComponent } from "~/app/modules/patient-management/charts/medicine-chart/medicine-chart.component";
import { IntakeChartComponent } from "~/app/modules/patient-management/charts/intake-chart/intake-chart.component";
import { MonitorChartComponent } from "~/app/modules/patient-management/charts/monitor-chart/monitor-chart.component";
import { ReportsComponent } from "~/app/modules/patient-management/reports/reports.component";


const routes: Routes = [
    { 
        path: "", component: PatientMgntComponent 
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
    { 
        path: "medicine-chart", component: MedicineChartComponent
    },
    { 
        path: "intake-chart", component: IntakeChartComponent
    },
    { 
        path: "monitor-chart", component: MonitorChartComponent
    },
    { 
        path: "reports", component: ReportsComponent
    },
];

@NgModule({
    imports: [NativeScriptRouterModule.forChild(routes)],
    exports: [NativeScriptRouterModule]
})
export class PatientMgntRoutingModule { }
