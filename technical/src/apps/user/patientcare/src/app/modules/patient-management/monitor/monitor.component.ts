import { Component, OnInit } from '@angular/core';

@Component({
	moduleId: module.id,
	selector: 'monitor',
	templateUrl: './monitor.component.html',
	styleUrls: ['./monitor.component.css']
})

export class MonitorComponent implements OnInit {
    categoricalSource: { Country: string, Amount: number, SecondVal: number, ThirdVal: number }[] = [
        { Country: "Germany", Amount: 15, SecondVal: 14, ThirdVal: 24 },
        { Country: "France", Amount: 13, SecondVal: 23, ThirdVal: 25 },
        { Country: "Bulgaria", Amount: 24, SecondVal: 17, ThirdVal: 23 },
        { Country: "Spain", Amount: 11, SecondVal: 19, ThirdVal: 24 },
        { Country: "USA", Amount: 18, SecondVal: 8, ThirdVal: 21 }
    ];
	constructor() { }

	ngOnInit() {
		// alert('monitor');
	 }
}