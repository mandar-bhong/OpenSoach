import { Component, OnInit } from '@angular/core';

@Component({
	moduleId: module.id,
	selector: 'monitor',
	templateUrl: './monitor.component.html',
	styleUrls: ['./monitor.component.css']
})

export class MonitorComponent implements OnInit {
	categoricalSource: { Country: string, Amount: any }[] = [
		{ Country: "10:00 AM", Amount: 35 },
		{ Country: "11:00 AM", Amount: 20 },
		{ Country: "12:00 AM", Amount: 28 },
		{ Country: "1:00 PM", Amount: 10 },
		{ Country: "2:00 PM", Amount: 40 }
	];

	constructor() { }

	ngOnInit() {
		// alert('monitor');
	}

}