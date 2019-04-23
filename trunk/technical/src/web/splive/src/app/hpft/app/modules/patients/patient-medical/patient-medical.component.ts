import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { Allergies, FamilyHistory, HistoryPresentIllness, InvestigationBeforeAdmission, JSONBaseDataModel, JSONInnerData, PastHistory, PersonalHistory, PersonalHistoryInfo, PresentComplaint, ReasonForAdmission, TreatmentBeforeAdmission } from '../../../../app/models/api/patient-data-models';
import { MedicalDetailsModel } from '../../../../app/models/ui/patient-models';
import { PatientService } from '../../../../app/services/patient.service';


@Component({
  selector: 'app-patient-medical',
  templateUrl: './patient-medical.component.html',
  styleUrls: ['./patient-medical.component.css']
})
export class PatientMedicalComponent implements OnInit, OnDestroy {
  dataModel = new MedicalDetailsModel();
  routeSubscription: Subscription;
  isResponsereceived = false;
  medicaldetialsid: number;
  constructor(private patientService: PatientService,
    private appNotificationService: AppNotificationService) {
  }

  ngOnInit() {

    this.dataModel.presentComplaintsData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.presentComplaintsData.data = [];

    this.dataModel.reasonForAdmissionData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.reasonForAdmissionData.data = [];

    this.dataModel.historyPresentIllnessData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.historyPresentIllnessData.data = [];

    this.dataModel.pastHistoryData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.pastHistoryData.data = [];

    this.dataModel.treatmentBeforeAdmissionData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.treatmentBeforeAdmissionData.data = [];

    this.dataModel.investigationBeforeAdmissionData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.investigationBeforeAdmissionData.data = [];

    this.dataModel.familyHistoryData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.familyHistoryData.data = [];


    this.dataModel.allergiesData = new JSONBaseDataModel<JSONInnerData[]>();
    this.dataModel.allergiesData.data = [];

    this.dataModel.personalHistoryData = new JSONBaseDataModel<PersonalHistoryInfo>();
    this.dataModel.personalHistoryData.version = 1;
    this.dataModel.personalHistoryData.data =new PersonalHistoryInfo();

    if (this.patientService.admissionid) {
      this.getPatientMedicalId();
    }

  }

  //Getting data from database
  getPatientMedicalId() {
    this.patientService.getPatientMedicalID({ recid: this.patientService.admissionid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.medicaldetialsid = payloadResponse.data.medicaldetails.medicaldetialsid;
          if (this.medicaldetialsid) {
            this.getPatientMedical();
          }
        }
      }
    });
  }

  //Getting data from database
  getPatientMedical() {
    this.patientService.getPatientMedicalDetail({ recid: this.medicaldetialsid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {    
        if (payloadResponse.data) {
          this.medicaldetialsid = payloadResponse.data.medicaldetialsid;
          this.dataModel.copyFrom(payloadResponse.data);
          this.isResponsereceived = true;
        }
      }
    });
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
  // reveived events which is generated from child component
  // Save function for Present Complaints.

  onAddedComplaint(outputValue: string) {
    // to do write code for saving received value.
    this.dataModel.medicaldetialsid = this.patientService.medicaldetialsid;
    this.dataModel.presentComplaintsData.data.push({ text: outputValue, date: new Date() });
    const presentComplaint = new PresentComplaint();
    presentComplaint.medicaldetialsid = this.medicaldetialsid;
    presentComplaint.presentcomplaints = JSON.stringify(this.dataModel.presentComplaintsData);
    this.patientService.medicalAddPatientComplaint(presentComplaint).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });

  }
  // Save function for Reason For Admission.
  onAddedAdmission(outputValue: string) {
    this.dataModel.reasonForAdmissionData.data.push({ text: outputValue, date: new Date() });
    const reasonForAdmission = new ReasonForAdmission();
    reasonForAdmission.medicaldetialsid = this.medicaldetialsid;
    reasonForAdmission.reasonforadmission = JSON.stringify(this.dataModel.reasonForAdmissionData);
    this.patientService.medicalAddPatientReasonForAdmission(reasonForAdmission).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });
  }

  // Save function for History Present Illness.
  onAddedHistoryPresentIllness(outputValue: string) {
    this.dataModel.historyPresentIllnessData.data.push({ text: outputValue, date: new Date() });
    const historyPresentIllness = new HistoryPresentIllness();
    historyPresentIllness.medicaldetialsid = this.medicaldetialsid;
    historyPresentIllness.historypresentillness = JSON.stringify(this.dataModel.historyPresentIllnessData);
    this.patientService.medicalAddPatientHistoryPresentIllness(historyPresentIllness).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });
  }

  // Save function for History Present Illness.
  onAddedPastHistory(outputValue: string) {
    this.dataModel.pastHistoryData.data.push({ text: outputValue, date: new Date() });
    const pastHistory = new PastHistory();
    pastHistory.medicaldetialsid = this.medicaldetialsid;
    pastHistory.pasthistory = JSON.stringify(this.dataModel.pastHistoryData);
    this.patientService.medicalAddPatientPastHistory(pastHistory).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });
  }

  // Save function for History Present Illness.
  onAddedTreatmentBeforeAdmission(outputValue: string) {
    this.dataModel.treatmentBeforeAdmissionData.data.push({ text: outputValue, date: new Date() });
    const treatmentBeforeAdmission = new TreatmentBeforeAdmission();
    treatmentBeforeAdmission.medicaldetialsid = this.medicaldetialsid;
    treatmentBeforeAdmission.treatmentbeforeadmission = JSON.stringify(this.dataModel.treatmentBeforeAdmissionData);
    this.patientService.medicalAddPatientTreatmentBeforeAdmission(treatmentBeforeAdmission).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });
  }

  // Save function for History Present Illness.
  onAddedInvestigationBeforeAdmission(outputValue: string) {
    this.dataModel.investigationBeforeAdmissionData.data.push({ text: outputValue, date: new Date() });
    const investigationBeforeAdmission = new InvestigationBeforeAdmission();
    investigationBeforeAdmission.medicaldetialsid = this.medicaldetialsid;
    investigationBeforeAdmission.investigationbeforeadmission = JSON.stringify(this.dataModel.investigationBeforeAdmissionData);
    this.patientService.medicalAddPatientInvestigationBeforeAdmission(investigationBeforeAdmission).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });
  }

  // Save function for Family History.
  onAddedFamilyHistory(outputValue: string) {
    this.dataModel.familyHistoryData.data.push({ text: outputValue, date: new Date() });
    const familyHistory = new FamilyHistory();
    familyHistory.medicaldetialsid = this.medicaldetialsid;
    familyHistory.familyhistory = JSON.stringify(this.dataModel.familyHistoryData);
    this.patientService.medicalAddPatientFamilyHistory(familyHistory).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });
  }

  // Save function for Allergies.
  onAddedAllergies(outputValue: string) {
    this.dataModel.allergiesData.data.push({ text: outputValue, date: new Date() });
    const allergiess = new Allergies();
    allergiess.medicaldetialsid = this.medicaldetialsid;
    allergiess.allergies = JSON.stringify(this.dataModel.allergiesData);
    this.patientService.medicalAddPatientAllergies(allergiess).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
      }
    });
  }

  // Save function for Person History.
  onAddedPersonHistory(outputValue: string) {
    const personalHistory = new PersonalHistory();
    personalHistory.medicaldetialsid = this.medicaldetialsid;
    personalHistory.personalhistory = JSON.stringify(outputValue);
    this.patientService.medicalAddPatientPersonalHistory(personalHistory).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        this.getPatientMedicalId();
      }
    });
  }
}
