import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '../patient-details.component';

@Component({
	moduleId: module.id,
	selector: 'treatment-before-admission',
	templateUrl: './treatment-before-admission.component.html',
	styleUrls: ['./treatment-before-admission.component.css']
})

export class TreatmentBeforeAdmissionComponent implements OnInit {

	getData = false;
	noData = false;
	constructor() { }
	@Input() treatmentlistitem: DataList[];
	
	ngOnInit() { 
		setTimeout(() => {
			if (this.treatmentlistitem.length > 0) {
				this.getData = true;
				this.noData = false;
			} else {
				this.noData = true;
				this.getData = false;
			}
		});
	}
}