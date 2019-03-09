import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '../patient-details.component';

@Component({
	moduleId: module.id,
	selector: 'past-history-about-health',
	templateUrl: './past-history-about-health.component.html',
	styleUrls: ['./past-history-about-health.component.css']
})

export class PastHistoryAboutHealthComponent implements OnInit {

	getData = false;
	noData = false;
	constructor() { }
	@Input() pasthistorylistitem: DataList[];
	ngOnInit() { 
		setTimeout(() => {

			if (this.pasthistorylistitem.length > 0) {
				this.getData = true;
				this.noData = false;
			} else {
				this.noData = true;
				this.getData = false;
			}

		});
	}
}