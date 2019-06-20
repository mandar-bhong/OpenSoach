import { Component, OnInit, Input } from '@angular/core';
import { PersonalHistoryModel, WeightData, AlcoholData, SmokData } from '~/app/models/ui/patient-details';
import { JSONBaseDataModel } from '~/app/models/ui/json-base-data-model';

@Component({
	moduleId: module.id,
	selector: 'personal-history',
	templateUrl: './personal-history.component.html',
	styleUrls: ['./personal-history.component.css']
})

export class PersonalHistoryComponent implements OnInit {
	alcoholtest = false;
	alcoholtest1 = false;
	quantityYes = false;
	quantityNo = false;
	other = false;
	// quantity = false;
	constructor() {
		// 
	}
	@Input() personallistitem: PersonalHistoryModel;

	// this.personallistitem.data : 
	ngOnInit() {


		if (this.personallistitem.alcohol.quantity != null) {
			this.alcoholtest = true;
			this.alcoholtest1 = false;
		} else {
			console.log('test data')
			this.alcoholtest = false;
			this.alcoholtest1 = true
		}
		if (this.personallistitem.smoking.quantity != null) {
			this.quantityYes = true;
			this.quantityNo = false;
		} else {
			this.quantityYes = false;
			this.quantityNo = true;
		}
		if (this.personallistitem.other != null) {
			this.other = true;
		} else {
			this.other = false;
		}
	}
}