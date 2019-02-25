import { Component, OnInit } from "@angular/core";
import { Router } from "@angular/router";
import { DatabaseService } from "../services/offline-store/database.service";
import { Page } from "ui/page";
import { RouterExtensions } from "nativescript-angular/router";

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
		private page: Page) {
		// Use the component constructor to inject providers.
		this.input = {
			"username": "admin@customer1.com",
			"password": "admin",
			"prodcode": "SPL_VST"
		}
		// hide action bar in page 
		page.actionBarHidden = true;
	}

	ngOnInit(): void {
		// Init your component properties here.
	}
	toggleForm() {
		this.isLoggingIn = !this.isLoggingIn;
	}
	login() {
		console.log(this.input.username);
		console.log(this.input.password);
		console.log(this.input.prodcode);

		// TODO: Dummy code for database testing
		this.databaseService.getdbConnection()
			.then(db => {
				db.all("SELECT id, item_name FROM items WHERE user_id = ?", ["Sanjay"]).then(rows => {
					for (var row in rows) {
						console.log("SELECT", { id: rows[row][0], name: rows[row][1] });
					}
				}, error => {
					console.log("SELECT ERROR", error);
				});
			});
		this.routerExtensions.navigate(['home'], { clearHistory: true });

		//TODO: user ServerAPIInterfaceService
		// this.router.navigate(['/list'], { skipLocationChange: true });
		// if (this.input.username && this.input.password) {
		//     this.httpClient.post("http://172.105.232.148/api/v1/login",
		//         {
		//             'username': this.input.username,
		//             'password': this.input.password,
		//             'prodcode': this.input.prodcode
		//         })
		//     .subscribe(
		//         data => {
		//             console.log("POST Request is successful ", data);
		//             this.router.navigate(['/home/list'], { skipLocationChange: true });
		//             // this.routerExtensions.navigate(["/home/list"], { clearHistory: true });
		//             // this.clearFields();
		//         }
		//     );
		// }
	}
	clearFields() {
		this.input.username = '';
		this.input.password = '';
	}
}