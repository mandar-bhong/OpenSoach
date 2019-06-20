import { Component, OnInit } from '@angular/core';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';
import { MedicineChartComponent } from '../charts/medicine-chart/medicine-chart.component';
import { IntakeChartComponent } from '../charts/intake-chart/intake-chart.component';
import { MonitorChartComponent } from '../charts/monitor-chart/monitor-chart.component';
import { OutputChartComponent } from '../charts/output-chart/output-chart.component';

@Component({
	moduleId: module.id,
	selector: 'schedular-fab',
	templateUrl: './schedular-fab.component.html',
	styleUrls: ['./schedular-fab.component.css']
})

export class SchedularFabComponent implements OnInit {
	constructor(private params: ModalDialogParams) { }
	ngOnInit() { }
	
	medicineForm() {
		this.params.closeCallback(MedicineChartComponent);
	}
	intakeForm() {
		this.params.closeCallback(IntakeChartComponent);
	}
	monitorForm() {
		this.params.closeCallback(MonitorChartComponent);
	}
	outputForm(){
		this.params.closeCallback(OutputChartComponent);
	}
	closeDialog() {
		console.log('in close Dialog');
		this.params.closeCallback();
	}

}