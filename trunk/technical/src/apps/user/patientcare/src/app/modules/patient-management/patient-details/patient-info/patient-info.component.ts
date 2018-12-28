import { Component, OnInit } from '@angular/core';

@Component({
	moduleId: module.id,
	selector: 'patient-info',
	templateUrl: './patient-info.component.html',
	styleUrls: ['./patient-info.component.css']
})

export class PatientInfoComponent implements OnInit {
	_dataItems =[];
	constructor() { }

	ngOnInit() { 
		this._dataItems.push({ name: "Amol Patil", rno: "RHC-2018-3456", hospitalised: "23/11/2018", bedno: "#A/312", weight: "64Kgs" });
	}
}