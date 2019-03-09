import { Component, OnInit, Input } from '@angular/core';
import { DataList } from '../patient-details.component';

@Component({
	moduleId: module.id,
	selector: 'investigation-before-admission',
	templateUrl: './investigation-before-admission.component.html',
	styleUrls: ['./investigation-before-admission.component.css']
})

export class InvestigationBeforeAdmissionComponent implements OnInit {

	getData = false;
	noData = false;
	constructor() { }
	@Input() investigationlistitem: DataList[];

	ngOnInit() {
		setTimeout(() => {
			if (this.investigationlistitem.length > 0) {
				this.getData = true;
				this.noData = false;
			} else {
				this.noData = true;
				this.getData = false;
			}
		});
	}
}