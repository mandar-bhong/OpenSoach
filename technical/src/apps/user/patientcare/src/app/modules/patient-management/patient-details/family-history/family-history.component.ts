import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '../patient-details.component';

@Component({
	moduleId: module.id,
	selector: 'family-history',
	templateUrl: './family-history.component.html',
	styleUrls: ['./family-history.component.css']
})

export class FamilyHistoryComponent implements OnInit {

	getData = false;
	noData = false;
	constructor() { }
	@Input() familylistitem: DataList[];

	ngOnInit() {
		setTimeout(() => {
			if (this.familylistitem.length > 0) {
				this.getData = true;
				this.noData = false;
			} else {
				this.noData = true;
				this.getData = false;
			}
		});
	}
}