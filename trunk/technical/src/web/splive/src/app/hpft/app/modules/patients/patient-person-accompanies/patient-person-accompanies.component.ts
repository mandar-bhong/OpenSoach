import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { PatientDataModel, PatientPersonAccompanying } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { FormGroup, FormControl } from '@angular/forms';
import { JSONBaseDataModel, PatientPersonDetail } from 'app/models/api/patient-models';
// import { PersonAccompanying, JSONBaseDataModel } from 'app/models/ui/person-accompanying-model';

@Component({
  selector: 'app-patient-person-accompanies',
  templateUrl: './patient-person-accompanies.component.html',
  styleUrls: ['./patient-person-accompanies.component.css']
})
export class PatientPersonAccompaniesComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new PatientPersonAccompanying();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  personGender: EnumDataSourceItem<number>[];
  contact: any;
  name: any;
  gender: any;
  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Person Accompanying';
  }

  ngOnInit() {
    this.createControls();
    this.showBackButton = false;
    this.personGender = this.patientService.getPersonGender();
    console.log('type of datamodel', typeof (this.dataModel.testdata));
    // this.dataModel.testdata = new  PatientPersonDetail();
    this.dataModel.testdata = new JSONBaseDataModel<PatientPersonDetail[]>();
    // this.setFormMode(FORM_MODE.VIEW);
    // this.showBackButton = false;
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['addid']) {
        this.dataModel.admissionid = Number(params['addid']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getPatientPersonDetail();
      } else {
        // this.subTitle = 'Add Details of Corporate';
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });


    // // example code
    // const model = JSON.parse('{version:1,da');
    // const persons=new JSONBaseModel<PersonAccompanying[]>()
    // Object.assign(persons,model);
  }


  createControls(): void {
    this.editableForm = new FormGroup({
      personname: new FormControl(''),
      personage: new FormControl(''),
      genderControl: new FormControl(''),
      personaladdress: new FormControl(''),
      reletionpatients: new FormControl(''),
      mobilenoControl: new FormControl(''),
      alternateContactControl: new FormControl('')
    });
  }


  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  //gender value
  getgender(value: number) {
    if (this.personGender && value) {
      return this.personGender.find(a => a.value === value).text;
    }
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }


  getPatientPersonDetail() {
    this.patientService.getPatientPersonDetail({ recid: this.dataModel.admissionid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {

          console.log('payloadResponse.data', payloadResponse.data);
          this.dataModel.copyFrom(payloadResponse.data);
          console.log('this.dataModel', this.dataModel);
          if(this.dataModel.testdata.data.length>0){
          console.log('testdata check', this.dataModel.testdata.data[0].name);
          this.name = this.dataModel.testdata.data[0].name;
          this.gender = this.dataModel.testdata.data[0].gender;
          this.contact = this.dataModel.testdata.data[0].contact;
          }
          // const test = this.dataModel;
          // const model = JSON.parse('test');
          // console.log('model',model);
          // const persons=new JSONBaseModel<PersonAccompanying[]>()
          // Object.assign(persons,model);
          // console.log('object',Object.assign(persons,model));

        }
      }
    });
  }
}
