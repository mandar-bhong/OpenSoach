import { Component, OnInit } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
@Component({
	moduleId: module.id,
	selector: 'patient-details',
	templateUrl: './patient-details.component.html',
	styleUrls: ['./patient-details.component.css']
})

export class PatientDetailsComponent implements OnInit {

	constructor(private routerExtensions: RouterExtensions) { }

	ngOnInit() { }
	goBackPage(){
		this.routerExtensions.navigate(['list', 'details'], { clearHistory: true });
	}
}