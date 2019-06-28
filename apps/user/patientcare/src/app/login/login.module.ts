import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import {NativeScriptFormsModule} from "nativescript-angular/forms"
import { LoginComponent } from "~/app/login/login.component";
import { LoginRoutingModule } from "~/app/login/login-routing.module";
import { ForgotPasswordComponent } from "./forgot-password/forgot-password.component";


@NgModule({
    imports: [
        NativeScriptCommonModule,
        LoginRoutingModule,
        NativeScriptFormsModule,
    ],
    declarations: [
        LoginComponent,
        ForgotPasswordComponent
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class LoginModule { }
