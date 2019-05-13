import { Component, OnInit } from '@angular/core';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { AdmissionDetailsModel } from '~/app/models/ui/patient-details';



@Component({
	moduleId: module.id,
	selector: 'admission-details',
	templateUrl: './admission-details.component.html',
	styleUrls: ['./admission-details.component.css']
})

export class AdmissionDetailsComponent implements OnInit {
	admissionDetailsModel = new AdmissionDetailsModel();
	constructor(private patientListService: PatientListService,
		private passdataservice: PassDataService) { }

	ngOnInit() {
		this.getAdmissionDetailsByUUID();
	}
	public getAdmissionDetailsByUUID() {
		this.patientListService.getAdmissionDetailsByUUID(this.passdataservice.getPatientID()).then(
			(val) => {
				val.forEach(item => {
					this.admissionDetailsModel = item;
				});
			},
			(error) => {
				console.log("admistion details error:", error);
			}
		);
	}

}