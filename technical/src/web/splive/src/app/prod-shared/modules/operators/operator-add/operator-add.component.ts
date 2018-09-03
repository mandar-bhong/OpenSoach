import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { EnumDataSourceItem } from '../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../shared/views/edit-record-base';
import { OperatorAddRequest, OperatorUpdateRequest } from '../../../models/api/operator-models';
import { OperatorAddModel } from '../../../models/ui/operator-models';
import { ProdOperatorService } from '../../../services/operator/prod-operator.service';

@Component({
  selector: 'app-operator-add',
  templateUrl: './operator-add.component.html',
  styleUrls: ['./operator-add.component.css']
})
export class OperatorAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new OperatorAddModel();
  routeSubscription: Subscription;
  operatorStates: EnumDataSourceItem<number>[];
  operatorAreas: EnumDataSourceItem<number>[];

  constructor(private prodOperatorService: ProdOperatorService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe
  ) {
    super();
    this.iconCss = 'fa fa-meh-o';
    this.pageTitle = this.translatePipe.transform('OPERATOR_ADD_TITLE');
  }

  ngOnInit() {
    this.createControls();
    this.operatorStates = this.prodOperatorService.getOperatorStates();
    this.operatorAreas = this.prodOperatorService.getOperatorAreas();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.fopid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getOperatorDetails();
      } else {
        this.subTitle = this.translatePipe.transform('OPERATOR_ADD_MODE_TITLE');
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      fopnameControl: new FormControl('', [Validators.required]),
      emailidControl: new FormControl('', [Validators.email]),
      mobilenoControl: new FormControl('', [Validators.required, Validators.pattern(/^\d+$/)]),
      fopcodeControl: new FormControl('', [ Validators.pattern(/^\d+$/)]),
      shortdescControl: new FormControl(''),
      fopstateControl: new FormControl('', [Validators.required]),
      fopareaControl: new FormControl('', [Validators.required]),
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const operatorAddRequest = new OperatorAddRequest();
      this.dataModel.copyTo(operatorAddRequest);
      this.prodOperatorService.addOperator(operatorAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.fopid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.fopname;
        }
        this.inProgress = false;
      });
    } else {
      const operatorUpdateRequest = new OperatorUpdateRequest();
      this.dataModel.copyToUpdateRequest(operatorUpdateRequest);
      this.prodOperatorService.updateOperatorDetails(operatorUpdateRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.fopname;
        }
        this.inProgress = false;
      });
    }
  }
  getOperatorDetails() {
    this.prodOperatorService.getOperatorDetails({ recid: this.dataModel.fopid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.subTitle = this.dataModel.fopname;
        }
      }
    });
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  getoperatorarea(value: number) {
    if (this.operatorAreas && value) {
      return this.operatorAreas.find(a => a.value === value).text;
    }
  }
  getoperatorstate(value: number) {
    if (this.operatorStates && value) {
      return this.operatorStates.find(a => a.value === value).text;
    }
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
