import { Component, OnInit } from '@angular/core';

@Component({
	moduleId: module.id,
	selector: 'medical-details',
	templateUrl: './medical-details.component.html',
	styleUrls: ['./medical-details.component.css']
})

export class MedicalDetailsComponent implements OnInit {
	_dataItems = [];
	constructor() { }

	ngOnInit() {
		console.log("medical component load");
		this._dataItems.push({ reason: "Test body", treatment: "Under Diagnosis", medhistory: "Undergone treatment for bone fracture in right leg", allergies: "Allergy with peanuts" });
	}
}