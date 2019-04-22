import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";
import { PreloadAllModules } from "@angular/router";
import { StartupComponent } from "./startup-page/startup.component";

const routes: Routes = [
    { path: "", redirectTo: "/startup", pathMatch: "full" },
    {
        path: "startup", component: StartupComponent
    },
    {
        path: "login", loadChildren: "~/app/login/login.module#LoginModule"
    },
    {
        path: "home", loadChildren: "~/app/home/home.module#HomeModule"
    },
    {
        path: "patientmgnt", loadChildren: "~/app/modules/patient-management/patient-mgnt.module#PatientMgntModule"
    },
];

@NgModule({
    imports: [NativeScriptRouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules })],
    exports: [NativeScriptRouterModule]
})
export class AppRoutingModule { }
