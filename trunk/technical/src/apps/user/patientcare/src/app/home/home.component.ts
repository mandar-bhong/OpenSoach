import { Component, OnInit } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Router } from "@angular/router";

@Component({
    selector: "Home",
    moduleId: module.id,
    templateUrl: "./home.component.html"
})
export class HomeComponent implements OnInit {
    input: any;
    isLoggingIn = true;
    constructor(private httpClient: HttpClient,
        private router: Router,) {
        // Use the component constructor to inject providers.
        this.input = {
            "username": "admin@customer1.com",
            "password": "admin",
            "prodcode": "SPL_VST"
        }
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
        this.router.navigate(['/list'], { skipLocationChange: true });
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
