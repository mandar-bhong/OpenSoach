import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EditRecordBase } from '../../../../shared/views/edit-record-base';
import { OperatorServicePointListModel } from '../../../models/api/operator-models';
import { OperatorServicePointsDataModel } from '../../../models/ui/operator-models';
import { ProdOperatorService } from '../../../services/operator/prod-operator.service';
import { ProdServicepointService } from '../../../services/servicepoint/prod-servicepoint.service';

@Component({
  selector: 'app-operator-associate',
  templateUrl: './operator-associate.component.html',
  styleUrls: ['./operator-associate.component.css']
})
export class OperatorAssociateComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new OperatorServicePointsDataModel();
  dataSource;
  routeSubscription: Subscription;
  selectedsp: OperatorServicePointListModel;
  callbackUrl: string;
  constructor(public prodOperatorService: ProdOperatorService,
    public prodServicepointService: ProdServicepointService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe
  ) {
    super();
    this.dataModel.splist = [];
    this.dataModel.associatedsplist = [];
    this.dataModel.availablesplist = [];
    this.iconCss = 'material-icons';
    this.iconName = 'link';
    this.pageTitle = this.translatePipe.transform('OPERATOR_ASSOCIATE');
  }

  ngOnInit() {
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.fopid = Number(params['id']);
      this.dataModel.fopname = String(params['fopname']);
      this.subTitle = 'Associate ' + this.dataModel.fopname + ' with service point(s)';
      this.getServicepointList();

      this.callbackUrl = params['callbackurl'];
    });
  }

  getOperatorServicpoint() {
    this.prodOperatorService.getOperatorServicepoint({ recid: this.dataModel.fopid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.splist.forEach(item => {
          const association = payloadResponse.data.find(sp => sp.spid === item.spid);
          if (association) {
            this.dataModel.associatedsplist.push(item);
          } else {
            this.dataModel.availablesplist.push(item);
          }
        });
      }
    });
  }
  getServicepointList() {
    this.prodServicepointService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const operatorServicePointListModel = new OperatorServicePointListModel();
          operatorServicePointListModel.spid = item.spid;
          operatorServicePointListModel.spname = item.spname;

          this.dataModel.splist.push(operatorServicePointListModel);
        });
        this.getOperatorServicpoint();
      }
    });
  }

  add() {
    this.prodOperatorService.associateOperatorToServicepoint(
      { fopid: this.dataModel.fopid, spid: this.selectedsp.spid }).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.associatedsplist.push(this.selectedsp);
          this.dataModel.availablesplist.splice(this.dataModel.availablesplist.indexOf(this.selectedsp), 1);
        }
      });
  }

  remove(item: OperatorServicePointListModel) {
    if (confirm('Are you sure you want to remove the operator association from the service point?')) {
      this.prodOperatorService.removeOperatorServicepointAssociation(
        { fopid: this.dataModel.fopid, spid: item.spid }).subscribe(payloadResponse => {
          if (payloadResponse && payloadResponse.issuccess) {
            this.dataModel.availablesplist.push(item);
            this.dataModel.associatedsplist.splice(this.dataModel.associatedsplist.indexOf(item), 1);
          }
        });
    }
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
