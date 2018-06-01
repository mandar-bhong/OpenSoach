import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../shared/views/edit-record-base';
import { OperatorServicePointListModel } from '../../../models/api/operator-models';
import { OperatorServicePointsDataModel } from '../../../models/ui/operator-models';
import { ProdOperatorService } from '../../../services/operator/prod-operator.service';

@Component({
  selector: 'app-operator-associate',
  templateUrl: './operator-associate.component.html',
  styleUrls: ['./operator-associate.component.css']
})
export class OperatorAssociateComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new OperatorServicePointsDataModel();
  dataSource;
  routeSubscription: Subscription;
  constructor(public prodOperatorService: ProdOperatorService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe
  ) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.fopid = Number(params['id']);
        this.dataModel.list = [];
        this.dataModel.previouslist = [];
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getServicepointList();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      // fopnameControl: new FormControl(''),
    });
  }
  getOperatorServicpoint() {
    this.prodOperatorService.getOperatorServicpoint({ recid: this.dataModel.fopid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          payloadResponse.data.forEach(item => {
            const spSelected = this.dataModel.list.find(sp => sp.spid === item.spid);
            if (spSelected) {
              spSelected.ischecked = true;
            }
          });
        }
        this.dataModel.previouslist = payloadResponse.data;

      }
    });
  }
  getServicepointList() {
    this.prodOperatorService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const operatorServicePointListModel = new OperatorServicePointListModel();
          operatorServicePointListModel.spid = item.spid;
          operatorServicePointListModel.spname = item.spname;
          this.dataModel.list.push(operatorServicePointListModel);
        });
        this.getOperatorServicpoint();
      }
    });
  }
  save() {

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
