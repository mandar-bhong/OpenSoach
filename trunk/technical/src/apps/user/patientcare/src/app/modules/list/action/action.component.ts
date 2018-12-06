import { Component, OnInit } from '@angular/core';

@Component({
	moduleId: module.id,
	selector: 'action',
	templateUrl: './action.component.html',
	styleUrls: ['./action.component.css']
})

export class ActionComponent implements OnInit {
	data = [];
	constructor() { }

	ngOnInit() { 
		for (let i = 1; i < 50; i++) {
			let newName = { ward: "3A/312", name: "Sumeet karande", mobile: "9878978980" };
			this.data.push(newName);
		}
		// alert('action');
	}
}