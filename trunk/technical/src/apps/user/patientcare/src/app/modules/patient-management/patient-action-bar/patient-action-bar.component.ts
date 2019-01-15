import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';

@Component({
	moduleId: module.id,
	selector: 'patient-action-bar',
	templateUrl: './patient-action-bar.component.html',
	styleUrls: ['./patient-action-bar.component.css']
})

export class PatientActionBarComponent implements OnInit {
	//  process variable 
	patientname: string;

	constructor(private routerExtensions: RouterExtensions) {

	}
	@Input() patientName: string;
	ngOnInit() {
		this.patientname = this.patientName;
	}

	goBackPage() {
		this.routerExtensions.back();
	}
	
	patientdetail(){
		this.routerExtensions.navigate(['patientmgnt', 'patient'], { clearHistory: false });
	}
}