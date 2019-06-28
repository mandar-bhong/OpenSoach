import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { PassDataService } from '~/app/services/pass-data-service';

@Component({
	moduleId: module.id,
	selector: 'patient-action-bar',
	templateUrl: './patient-action-bar.component.html',
	styleUrls: ['./patient-action-bar.component.css']
})

export class PatientActionBarComponent implements OnInit {
	//  process variable 
	patientname: string;
	constructor(private routerExtensions: RouterExtensions,
		private passDataService: PassDataService) {
	}
	@Input() patientName: string;
	ngOnInit() {
		this.patientname = this.patientName;
	}

	goBackPage() {
		// const test =  true;
		// console.log('this.passDataService.backalert', this.passDataService.backalert);
		// if (this.passDataService.backalert === test) {
		// 	console.log(this.passDataService.showNotification());

		// } else {
		// 	this.routerExtensions.back();
		// }
		// this.passDataService.backalert = false;
		this.routerExtensions.back();

	}

	patientdetail() {
		this.routerExtensions.navigate(['patientinfo'], { clearHistory: false });
		// this.routerExtensions.navigate(['patientmgnt', 'patient'], { clearHistory: false });
		// const test =  true;
		// console.log('this.passDataService.backalert', this.passDataService.backalert);
		// if (this.passDataService.backalert === test) {
		// 	this.passDataService.showNotification();
		// 	this.routerExtensions.navigate(['patientmgnt', 'patient'], { clearHistory: false });
		// } else {
		// 	this.routerExtensions.back();
		// }
		// this.passDataService.backalert = false;

	}

}