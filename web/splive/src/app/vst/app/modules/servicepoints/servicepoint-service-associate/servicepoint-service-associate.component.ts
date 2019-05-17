import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { SpServiceConfService } from '../../../../../prod-shared/services/spservice/sp-service-conf.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase } from '../../../../../shared/views/edit-record-base';
import {
  ServicepointConfigureListResponse,
  ServicepointConfigureTemplateListRequest,
} from '../../../models/api/service-configuration-models';
import { ServicepointAssociateRequest, ServicepointDataListResponse } from '../../../models/api/servicepoint-models';
import { ServicePointServiceConfigureAssociateModel } from '../../../models/ui/servicepoint-models';
import { ProdServicepointService } from '../../../services/servicepoint/prod-servicepoint.service';

@Component({
  selector: 'app-servicepoint-service-associate',
  templateUrl: './servicepoint-service-associate.component.html',
  styleUrls: ['./servicepoint-service-associate.component.css']
})
export class ServicepointServiceAssociateComponent extends EditRecordBase implements OnInit, OnDestroy {
  routeSubscription: Subscription;
  editableForm: FormGroup;
  selectedStatus: number;
  dataModel = new ServicePointServiceConfigureAssociateModel();
  spconfigures: ServicepointConfigureListResponse[] = [];
  servicepointDataListResponse: ServicepointDataListResponse;
  servconfid: number;

  constructor(public spServiceConfService: SpServiceConfService,
    public prodServicepointService: ProdServicepointService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private route: ActivatedRoute,
    private router: Router,
  ) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.getServicepointConfigureList();
    this.selectedStatus = 1;
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.spid = Number(params['id']);
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      configNameControl: new FormControl('')
    });
  }
  save() {
    switch (this.selectedStatus) {
      case 1:
        this.add();
        break;
      case 2:
        this.associate();
        break;
      case 3:
        this.copyTemplateList();
        break;
    }
  }
  add() {
    // add new chart nevigate chart configure window with spid
    this.router.navigate(['charts', 'configure'], {
      queryParams: { spid: this.dataModel.spid, mode: 2, callbackurl: 'servicepoints' }, skipLocationChange: true
    });

  }

  getServicepointConfigureList() {
    // get Existing servicepoint configure short list
    this.spServiceConfService.getServicepointConfigureList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.spconfigures = payloadResponse.data;
      }
    });
  }

  copyTemplateList() {
    // copy to template list
    const servicepointConfigureTemplateListRequest = new ServicepointConfigureTemplateListRequest();
    this.dataModel.copyTo(servicepointConfigureTemplateListRequest);
    this.spServiceConfService.copyTemplateList(servicepointConfigureTemplateListRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.servconfid = payloadResponse.data.recid;
        this.associate(true);
      }
    });
  }

  associate(navigate?: boolean) {
    // associate serivcepoint configure
    const request = new ServicepointAssociateRequest();
    this.dataModel.copyToAssociateRequest(request);
    this.prodServicepointService.associateConfigure(request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (navigate) {
          this.router.navigate(['charts', 'configure'], {
            queryParams: { id: this.dataModel.servconfid, mode: 1, callbackurl: 'servicepoints' }, skipLocationChange: true
          });
        } else {
          this.appNotificationService.success();
          this.closeForm();
        }
      }
    });
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
