import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";
import { PreloadAllModules } from "@angular/router";

const routes: Routes = [
    // { path: "", redirectTo: "/home", pathMatch: "full" },
    // {
    //     path: "", redirectTo: "/login", pathMatch: "full"
    // },
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
