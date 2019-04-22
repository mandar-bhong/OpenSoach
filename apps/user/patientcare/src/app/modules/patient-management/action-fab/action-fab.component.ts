import { Component, OnInit } from '@angular/core';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';
import { DoctorOrdersComponent } from '../doctor-orders/doctor-orders.component';

@Component({
	moduleId: module.id,
	selector: 'action-fab',
	templateUrl: './action-fab.component.html',
	styleUrls: ['./action-fab.component.css']
})

export class ActionFabComponent implements OnInit {
	constructor(private params: ModalDialogParams) { }
	ngOnInit() { }
	doctorsOrders() {
		this.params.closeCallback('DoctorOrdersComponent');
	}
	closeDialog(){
		this.params.closeCallback();
	}
}