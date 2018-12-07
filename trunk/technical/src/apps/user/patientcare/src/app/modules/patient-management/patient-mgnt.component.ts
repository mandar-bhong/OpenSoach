import { Component, OnInit } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { Page } from 'ui/page';



@Component({
	moduleId: module.id,
	selector: 'patient-mgnt',
	templateUrl: './patient-mgnt.component.html',
	styleUrls: ['./patient-mgnt.component.css']
})


export class PatientMgntComponent implements OnInit {

	constructor(private routerExtensions: RouterExtensions,
		private page: Page) { }


	ngOnInit() {

	}
	
}