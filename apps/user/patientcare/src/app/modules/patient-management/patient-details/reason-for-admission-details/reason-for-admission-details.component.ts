import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '~/app/models/ui/patient-details';

@Component({
	moduleId: module.id,
	selector: 'reason-for-admission-details',
	templateUrl: './reason-for-admission-details.component.html',
	styleUrls: ['./reason-for-admission-details.component.css']
})

export class ReasonForAdmissionDetailsComponent implements OnInit {
	getData = false;
	noData = false;
	constructor() { }
	@Input() reasonlistItem: DataList[];
	ngOnInit() {

		if (this.reasonlistItem.length > 0) {
			this.getData = true;
			this.noData = false;
		} else {
			this.noData = true;
			this.getData = false;
		}

	}
}