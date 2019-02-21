import { Component, OnInit, EventEmitter, Input, Output } from '@angular/core';
import { PatientService } from '../../../../app/services/patient.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { EditRecordBase } from '../../../../../shared/views/edit-record-base';
import { Subscription } from 'rxjs';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { ServicepointConfigureListResponse } from '../../../../../prod-shared/models/api/service-configuration-models';
import { Router } from '@angular/router';
import { PatientCheckModal } from '../../../models/ui/patient-models';
import { PatientAddService } from '../../../services/patient-add.service';
import { PatientSearchRequestFilter } from '../../../models/api/patient-models';

@Component({
  selector: 'app-patient-check-search',
  templateUrl: './patient-check-search.component.html',
  styleUrls: ['./patient-check-search.component.css']
})
export class PatientCheckSearchComponent extends EditRecordBase implements OnInit {

  @Input()
  editRecordBase: EditRecordBase;
  @Output()
  editClick = new EventEmitter<null>();
  dataModel = new PatientCheckModal();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  splist: ServicepointListResponse[] = [];
  spconfigures: ServicepointConfigureListResponse[] = [];
  showlist = false;
  isExpanded = false;

  constructor(public patientService: PatientService,
    public patientAddService: PatientAddService,
    private router: Router,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Search Patient';

  }

  ngOnInit() {
    this.createControls();
    this.dataModel = new PatientCheckModal();
  }

  search() {
    this.isExpanded = false;
    const patientSearchRequestFilter = new PatientSearchRequestFilter();
    this.dataModel.copyTo(patientSearchRequestFilter);
    this.patientService.fname = this.dataModel.fname;
    this.patientService.lname = this.dataModel.lname;
    if (this.dataModel.fname != null) {
      this.showlist = true;
      // to do disscuss with sanjay.
      setTimeout(() => {
        this.patientAddService.dataListSubjectTrigger(patientSearchRequestFilter);
      }, 10);

    }

  }

  newPatient() {
    this.router.navigate(['patients', 'add'], { queryParams: { callbackurl: 'patients' }, skipLocationChange: true });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      fnameControl: new FormControl('', [Validators.required]),
      lnameControl: new FormControl('')
    });
  }

  close() {
    this.router.navigate(['patients'], { queryParams: { callbackurl: 'patients' }, skipLocationChange: true });
  }

  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
