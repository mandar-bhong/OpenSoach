import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { ServicepointConfigureListResponse } from '../../../../../prod-shared/models/api/service-configuration-models';
import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { MedicalDetailAddRequest, PatientDataAddRequest, PatientDetailAddRequest } from '../../../models/api/patient-models';
import { PatientDataModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { PatientAdmissionComponent } from '../patient-admission/patient-admission.component';
import { FormGroup, FormControl } from '@angular/forms';
import { JSONBaseModel } from 'app/models/ui/json-base-model';
import { modelGroupProvider } from '@angular/forms/src/directives/ng_model_group';
@Component({
  selector: 'app-patient-medical',
  templateUrl: './patient-medical.component.html',
  styleUrls: ['./patient-medical.component.css']
})
export class PatientMedicalComponent extends EditRecordBase implements OnInit,OnDestroy {

  dataModel = new PatientDataModel();
  routeSubscription: Subscription;
  spconfigures: ServicepointConfigureListResponse[] = [];
  savebutton = false;

  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Medical Detail';
   }

  ngOnInit() {
    // this.createControls();
    // this.setFormMode(FORM_MODE.VIEW);
    this.showBackButton = false;

    // example code
    // const model = JSON.parse('{version:1,da');
    // const persons=new JSONBaseModel<PersonAccompanying[]>()
    // Object.assign(persons,model);
    
  }

  // textValue = 'initial value';
  // log = '';

  // logText(value: string): void {
  //   this.log += `Text changed to '${value}'\n`;
  // }

  // createControls(): void {
  //   this.editableForm = new FormGroup({
  //     presentcomplains: new FormControl(''),
  //     pasthistory: new FormControl(''),
  //     persondate: new FormControl(''),
  //     doctorincharge: new FormControl('')
  //   });
  // }
  
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
