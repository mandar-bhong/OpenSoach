import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";
import { LoginModule } from "~/app/login/login.module";
import { HomeModule } from "~/app/home/home.module";
import { PatientMgntModule } from "~/app/modules/patient-management/patient-mgnt.module";

const routes: Routes = [
    // { path: "", redirectTo: "/home", pathMatch: "full" },
    {
        path: "", redirectTo: "/login", pathMatch: "full"
    },
    // {
    //     path: "login", loadChildren: "~/app/login/login.module#LoginModule"
    // },
    // {
    //     path: "home", loadChildren: "~/app/home/home.module#HomeModule"
    // },
    // {
    //     path: "patientmgnt", loadChildren: "~/app/modules/patient-management/patient-mgnt.module#PatientMgntModule"
    // },


    {path:'login' ,loadChildren: ()=>LoginModule},
    {path:'home' ,loadChildren: ()=>HomeModule},
    {path:'patientmgnt' ,loadChildren: ()=>PatientMgntModule},
];

@NgModule({
    imports: [NativeScriptRouterModule.forRoot(routes)],
    exports: [NativeScriptRouterModule]
})
export class AppRoutingModule { }
