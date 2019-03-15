import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '~/app/models/ui/patient-details';

@Component({
	moduleId: module.id,
	selector: 'medical-details',
	templateUrl: './medical-details.component.html',
	styleUrls: ['./medical-details.component.css']
})

export class MedicalDetailsComponent implements OnInit {

	getData = false;
	noData =  false;
	constructor() { }

	@Input() listItem: DataList[];
	ngOnInit() {
		setTimeout(()=>{
			if(this.listItem.length > 0){
				console.log(' if');
				this.getData = true;
				this.noData = false;
			}else{
				console.log(' else');
				this.noData = true;
				this.getData = false;
			}
	
		});

	}
	
}