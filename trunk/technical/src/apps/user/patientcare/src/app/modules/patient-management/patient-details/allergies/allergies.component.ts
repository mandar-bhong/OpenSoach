import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '../patient-details.component';

@Component({
	moduleId: module.id,
	selector: 'allergies',
	templateUrl: './allergies.component.html',
	styleUrls: ['./allergies.component.css']
})

export class AllergiesComponent implements OnInit {

	getData = false;
	noData =  false;
	constructor() { }
	@Input() allergieslistitem: DataList[];
	
	ngOnInit() {
		setTimeout(()=>{
			if(this.allergieslistitem.length > 0){
				this.getData = true;
				this.noData = false;
			}else{
				this.noData = true;
				this.getData = false;
			}
		});
	 }
}