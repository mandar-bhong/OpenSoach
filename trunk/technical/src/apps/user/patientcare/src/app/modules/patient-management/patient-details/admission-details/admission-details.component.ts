import { Component, OnInit } from '@angular/core';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { PassDataService } from '~/app/services/pass-data-service';


export class AdmissionDetailsModel {
	patient_uuid: string;
	patient_reg_no: string;
	bed_no: string;
	sp_uuid: number;
	dr_incharge: number;
	admitted_on: Date;

}
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
		// console.log('this.passdataservice.getAdmissionID()', this.passdataservice.getAdmissionID())
	}
	public getAdmissionDetailsByUUID() {
			this.patientListService.getAdmissionDetailsByUUID(this.passdataservice.getPatientID()).then(
			(val) => {
				val.forEach(item => {
					// console.log("admistion details", val);
					
					this.admissionDetailsModel.patient_reg_no = item.patient_reg_no;
					this.admissionDetailsModel.sp_uuid = item.sp_uuid;
					this.admissionDetailsModel.bed_no = item.bed_no;
					this.admissionDetailsModel.dr_incharge = item.dr_incharge;
					this.admissionDetailsModel.admitted_on = item.admitted_on;

					// console.log('admissionDetailsModel', this.admissionDetailsModel);
				});
			},
			(error) => {
				console.log("admistion details error:", error);
			}
		);
	}
	
}