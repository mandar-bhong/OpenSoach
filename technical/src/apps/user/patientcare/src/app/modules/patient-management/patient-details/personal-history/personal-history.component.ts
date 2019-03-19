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


		setTimeout(() => {
			// this.personallistitem = new PersonalHistoryModel;
			// this.personallistitem.weight = new WeightData;
			// this.personallistitem.alcohol = new AlcoholData();
			// this.personallistitem.smoking = new SmokData();
			console.log('get data personallistitem', this.personallistitem);
			// console.log('get data weight', this.personallistitem.weight);
			if (this.personallistitem.alcohol.alcoholquantity != null) {
				this.alcoholtest = true;
				this.alcoholtest1 = false;
			} else {
				console.log('test data')
				this.alcoholtest = false;
				this.alcoholtest1 = true
			}
			if (this.personallistitem.smoking.smokingquantity != null) {
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
		});



	}
}