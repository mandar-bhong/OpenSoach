import { Component, OnInit, Input } from '@angular/core';
import { DataList, ListItem } from '~/app/models/ui/patient-details';

@Component({
	moduleId: module.id,
	selector: 'medical-details',
	templateUrl: './medical-details.component.html',
	styleUrls: ['./medical-details.component.css']
})

export class MedicalDetailsComponent implements OnInit {

	getData = false;
	noData = false;
	constructor() { }

	@Input() listItem: ListItem[];
	ngOnInit() {
       console.log('listItems',this.listItem);
		if (this.listItem.length > 0) {
			this.getData = true;
			this.noData = false;
			this.listItem.sort((a, b) => {
				return (new Date(b.date).getTime() - new Date(a.date).getTime())
			  });
		} else {
			this.noData = true;
			this.getData = false;
		}
	}

}