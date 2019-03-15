import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '~/app/models/ui/patient-details';

@Component({
	moduleId: module.id,
	selector: 'history-of-present-Illness',
	templateUrl: './history-of-present-Illness.component.html',
	styleUrls: ['./history-of-present-Illness.component.css']
})

export class HistoryOfPresentIllnessComponent implements OnInit {
	getData = false;
	noData =  false;
	constructor() { }
	@Input() historylistitem: DataList[];
	ngOnInit() {
		setTimeout(()=>{
			if(this.historylistitem.length > 0){
				this.getData = true;
				this.noData = false;
			}else{
				this.noData = true;
				this.getData = false;
			}

		});
	}
}