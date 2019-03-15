import { Component, OnInit, Input } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { PassDataService } from '~/app/services/pass-data-service';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { JSONBaseDataModel } from '~/app/models/ui/json-base-data-model';
import { DataList, PersonalHistoryModel } from '~/app/models/ui/patient-details';


@Component({
	moduleId: module.id,
	selector: 'patient-details',
	templateUrl: './patient-details.component.html',
	styleUrls: ['./patient-details.component.css']
})

export class PatientDetailsComponent implements OnInit {
	patientname: string;
	jsonpresentComplaintsList = new JSONBaseDataModel<DataList[]>();
	reason_for_admission = new JSONBaseDataModel<DataList[]>();
	history_present_illness = new JSONBaseDataModel<DataList[]>();
	past_history = new JSONBaseDataModel<DataList[]>();
	treatment_before_admission = new JSONBaseDataModel<DataList[]>();
	investigation_before_admission = new JSONBaseDataModel<DataList[]>();
	family_history = new JSONBaseDataModel<DataList[]>();
	allergies = new JSONBaseDataModel<DataList[]>();
	personal_history = new PersonalHistoryModel();
	constructor(private routerExtensions: RouterExtensions,
		private passDataService: PassDataService,
		private patientListService: PatientListService) {
	}
	ngOnInit() {
		this.getMedicalDetailsByUUID();
		this.patientname = this.passDataService.getHeaderName();
	}
	goBackPage() {
		this.routerExtensions.back();
	}
	public getMedicalDetailsByUUID() {
		this.patientListService.getMedicalDetailsByUUID(this.passDataService.getAdmissionID()).then(
			(val) => {
				this.jsonpresentComplaintsList = new JSONBaseDataModel<DataList[]>();
				this.jsonpresentComplaintsList.data = [];
				Object.assign(this.jsonpresentComplaintsList, JSON.parse(val[0].present_complaints));

				this.reason_for_admission = new JSONBaseDataModel<DataList[]>();
				this.reason_for_admission.data = [];
				Object.assign(this.reason_for_admission, JSON.parse(val[0].reason_for_admission));


				this.history_present_illness = new JSONBaseDataModel<DataList[]>();
				this.history_present_illness.data = [];
				Object.assign(this.history_present_illness, JSON.parse(val[0].history_present_illness));

				this.past_history = new JSONBaseDataModel<DataList[]>();
				this.past_history.data = [];
				Object.assign(this.past_history, JSON.parse(val[0].past_history));

				this.treatment_before_admission = new JSONBaseDataModel<DataList[]>();
				this.treatment_before_admission.data = [];
				Object.assign(this.treatment_before_admission, JSON.parse(val[0].treatment_before_admission));

				this.investigation_before_admission = new JSONBaseDataModel<DataList[]>();
				this.investigation_before_admission.data = [];
				Object.assign(this.investigation_before_admission, JSON.parse(val[0].investigation_before_admission));

				this.family_history = new JSONBaseDataModel<DataList[]>();
				this.family_history.data = [];
				Object.assign(this.family_history, JSON.parse(val[0].family_history));

				this.allergies = new JSONBaseDataModel<DataList[]>();
				this.allergies.data = [];
				Object.assign(this.allergies, JSON.parse(val[0].allergies));

				this.personal_history = new PersonalHistoryModel();
				Object.assign(this.personal_history, JSON.parse(val[0].personal_history));
			},
			(error) => {
				console.log("Medial details error:", error);
			}
		);
	}
}