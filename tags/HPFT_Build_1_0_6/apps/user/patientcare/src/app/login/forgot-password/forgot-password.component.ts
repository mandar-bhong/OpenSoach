import { Component, OnInit } from "@angular/core";
import { Router } from "@angular/router";
import { Page } from "tns-core-modules/ui/page";
import { RouterExtensions } from "nativescript-angular/router";
import { HttpClient } from "@angular/common/http";
import { APP_MODE } from "../../app-constants";
import { AppStartupService } from "../../services/app-startup.service";
import { DatabaseHelper } from "../../helpers/database-helper";
import { ServerApiInterfaceService } from "../../services/server-api-interface.service";
import * as appSettings from "tns-core-modules/application-settings";
import { AppGlobalContext } from "../../app-global-context";
import { AppRepoService } from "../../services/app-repo.service";

@Component({
	moduleId: module.id,
	selector: 'forgot-password',
	templateUrl: './forgot-password.component.html',
	styleUrls: ['./forgot-password.component.css']
})

export class ForgotPasswordComponent implements OnInit {

	email: string;
	isLoggingIn = true;
	constructor(private routerExtensions: RouterExtensions,
		private router: Router,
		private serverApiInterfaceService: ServerApiInterfaceService,
		private appStartUpService: AppStartupService,
		private httpClient: HttpClient,
		private page: Page) {

		page.actionBarHidden = true;
	}

	ngOnInit(): void {

		// Init your component properties here.
	}

	forgotPassword() {

		if (this.email) {
			const authRequest = new ChangePasswordReq();
			authRequest.usrname = this.email;
			console.log('auth request', authRequest);
			console.log('(AppRepoService.Instance.API_SPL_BASE_URL', AppRepoService.Instance.API_SPL_BASE_URL);
			var dialogs = require("tns-core-modules/ui/dialogs");
			this.serverApiInterfaceService.post<any>(AppRepoService.Instance.API_SPL_BASE_URL + "/v1/user/forgot/password", authRequest).then(
				(success) => {
					dialogs.alert({
						title: "Forgot Password",
						message: "We've sent an OTP to the email ",
						okButtonText: "Ok"
					}).then(this.routerExtensions.navigate(['login'], { clearHistory: true }));
				}, (error) => {
					if(error.error.code==10006){
						dialogs.alert({
							title: "Forgot Password",
							message: "invalid Email",
							okButtonText: "Ok"
						});
					}
				});
		}
	}
	redirect() {
		this.routerExtensions.navigate(['login'], { clearHistory: true });
	}
	toggleForm() {
		this.isLoggingIn = !this.isLoggingIn;
	}
}
export class ChangePasswordReq {
	usrname: string;
}