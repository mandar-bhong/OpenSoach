import { Component, OnInit } from "@angular/core";
import { Router } from "@angular/router";
import { DatabaseService } from "../services/offline-store/database.service";
import { Page } from "tns-core-modules/ui/page";
import { RouterExtensions } from "nativescript-angular/router";
import { HttpClient } from "@angular/common/http";
import { API_SPL_BASE_URL, APP_MODE } from "../app-constants";
import { AppStartupService } from "../services/app-startup.service";
import { DatabaseHelper } from "../helpers/database-helper";
import { ServerApiInterfaceService } from "../services/server-api-interface.service";
import * as appSettings from "tns-core-modules/application-settings";
import { AppGlobalContext } from "../app-global-context";

@Component({
	moduleId: module.id,
	selector: 'login',
	templateUrl: './login.component.html',
	styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit {
	input: any;
	isLoggingIn = true;
	constructor(private routerExtensions: RouterExtensions,
		private databaseService: DatabaseService,
		private router: Router,
		private serverApiInterfaceService: ServerApiInterfaceService,
		private appStartUpService: AppStartupService,
		private httpClient: HttpClient,
		private page: Page) {
		// Use the component constructor to inject providers.
		this.input = {
			// "username": "admin@customer1.com",
			// "password": "admin",
			"prodcode": "SPL_HPFT"
		}
		// hide action bar in page 

		page.actionBarHidden = true;
	}

	ngOnInit(): void {
		this.bindUserData();
		// Init your component properties here.
	}
	toggleForm() {
		this.isLoggingIn = !this.isLoggingIn;
	}
	async login() {
		console.log(this.input.username);
		console.log(this.input.password);
		console.log(this.input.prodcode);
		if (this.input.username && this.input.password) {
			const authRequest = new AuthRequest();
			authRequest.username = this.input.username;
			authRequest.password = this.input.password;
			authRequest.prodcode = this.input.prodcode;
			console.log('auth request', authRequest);
			this.serverApiInterfaceService.post<any>(API_SPL_BASE_URL + "/v1/endpoint/user/login", authRequest).then(
				async (success) => {
					console.log("POST Request is successful ", success);
					appSettings.setNumber("USER_ID", success.userid);
					const usres = await this.getUsers();
					console.log('usres len', usres);

					switch (usres.length) {
						case 0://no record is available in db
						this.addUser();
							break;
						case 1: // only 1 record is available in db
						if (!(usres[0].user_name === this.input.username && usres[0].password === this.input.password)) {
							const deleteUser = await this.deleteUsers();
							if (deleteUser) {
								await this.addUser();
							}
						}
							break;
						default: // more than 1 record is available in db
						const deluser = await this.deleteUsers();
							if (deluser) {
								await this.addUser();
							}
							break;
					}
					console.log('final statement executed');
					this.appStartUpService.handleDevAuthResponse(success, APP_MODE.USER_DEVICE);
					//   this.router.navigate(['/home/list'], { skipLocationChange: true });
					//   this.routerExtensions.navigate(["/home/list"], { clearHistory: true });
					//   this.clearFields();
				}, (error) => {
					console.log('POST Request is Failed', error);
				});
		}
	}// end of class

	getUsers(): Promise<any[]> {
		return new Promise((resolve, reject) => {
			DatabaseHelper.selectAll('getuser').then((success) => {
				console.log('user list', success);
				resolve(success);
			}, (error) => {
				console.log('getuser response Failed', error);
				reject(error);
			});
		});
	}// end of fucntions.
	deleteUsers(): Promise<any> {
		return new Promise((resolve, reject) => {
			DatabaseHelper.selectAll('deleteuser').then((success) => {
				console.log('success', success);
				resolve(success);
			}, (error) => {
				console.log('getuser response Failed', error);
				reject(error);
			});
		});
	}// end of fucntions.

	addUser(): Promise<any> {
		console.log('in add user function ')
		return new Promise((resolve, reject) => {
			let param = [];
			param.push(this.input.username);
			param.push(this.input.password);
			param.push(this.input.prodcode);
			DatabaseHelper.update('user_login_tbl_insert', param).then((success) => {
				console.log('user added into database ', success);
				resolve();
			}, (error) => {
				console.log('getuser response Failed', error);
				reject();
			});
		});
	}

	async bindUserData() {
		//TODO: Get Remember from db and same db row
		//
		const rememberMe = 1;
		const item = await this.getUsers();
		if (item.length > 0) {
			this.input.username = item[0].user_name
			this.input.password = item[0].password;
			this.input.prodcode = item[0].auth_code;
			this.login();
		}


		//if rember me = 0

	}

	clearFields() {
		this.input.username = '';
		this.input.password = '';
	}
}
export class AuthReponse {
	issuccess: boolean;
	error: any;
	data: any;
}
export class AuthRequest {
	username: string;
	password: string;
	prodcode: string;
}
