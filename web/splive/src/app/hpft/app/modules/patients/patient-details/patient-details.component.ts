import { Component, EventEmitter, Input, OnDestroy, OnInit, Output } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { EditRecordBase } from '../../../../../shared/views/edit-record-base';
import { PatientDataModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';

@Component({
  selector: 'app-patient-details',
  templateUrl: './patient-details.component.html',
  styleUrls: ['./patient-details.component.css']
})
export class PatientDetailsComponent extends EditRecordBase implements OnInit, OnDestroy {

  @Input()
  editRecordBase: EditRecordBase;
  @Output()
  editClick = new EventEmitter<null>();
  routeSubscription: Subscription;
  dataModal = new PatientDataModel();
  selectedIndex: 0;
  disableTab: boolean;
  admissionid: number;
  admissionIdSubscription: Subscription;
  patientName: string;
  patientNameSubscription: Subscription;

  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private translatePipe: TranslatePipe,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Personal Detail';
  }

  ngOnInit() {
    this.patientNameSubscription = this.patientService.patientName.subscribe((value) => {
        this.patientName = value;
    });
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.subTitle = this.translatePipe.transform('OPERATOR_ADD_MODE_TITLE');
      this.callbackUrl = params['callbackurl'];
    });
    this.admissionIdSubscription = this.patientService.admissionIdReceived.subscribe((value) => {
      this.admissionid = value;
      this.HideTab();
    });
    if (this.patientService.selcetdIndex && this.patientService.selcetdIndex != null) {
      this.changeTab(this.patientService.selcetdIndex);
    }
    this.route.queryParams.subscribe(params => {
      const id = params['id'];
      this.admissionid = params['admissionid'];
      this.HideTab();
    });
  }

  HideTab() {
    if (this.admissionid != null) {
      this.disableTab = true;
    } else {
      this.disableTab = false;
    }
  }

  changeTab(value) {
    this.selectedIndex = value;
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
    if (this.admissionIdSubscription) {
      this.admissionIdSubscription.unsubscribe();
    }
    if (this.patientNameSubscription) {
      this.patientNameSubscription.unsubscribe();
    }
  }

}
