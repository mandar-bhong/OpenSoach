import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EditRecordBase } from '../../../../shared/views/edit-record-base';
import {
  ServicepointAssociateRequest,
  ServicepointConfigureListResponse,
  ServicepointConfigureTemplateListRequest,
  ServicepointDataListResponse,
} from '../../../models/api/service-configuration-models';
import { ConfigureAssociateModel } from '../../../models/ui/service-configuration-models';
import { SpServiceConfService } from '../../../services/spservice/sp-service-conf.service';

@Component({
  selector: 'app-servicepoint-details',
  templateUrl: './servicepoint-details.component.html',
  styleUrls: ['./servicepoint-details.component.css']
})
export class ServicepointDetailsComponent extends EditRecordBase implements OnInit, OnDestroy {
  routeSubscription: Subscription;
  editableForm: FormGroup;
  selectedStatus: number;
  dataModel = new ConfigureAssociateModel();
  spconfigures: ServicepointConfigureListResponse[] = [];
  servicepointDataListResponse: ServicepointDataListResponse;
  servconfid: number;

  constructor(public spServiceConfService: SpServiceConfService,
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
      queryParams: { spid: this.dataModel.spid, callbackurl: 'servicepoints' }, skipLocationChange: true
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
        this.associate();
      }
    });
  }

  associate() {
    // associate serivcepoint configure
    const request = new ServicepointAssociateRequest();
    this.dataModel.copyToAssociateRequest(request);
    this.spServiceConfService.associateConfigure(request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
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
  // chartwindow() {
  //   this.router.navigate(['charts', 'configure'], {
  //     queryParams: {servconfid: this.dataModel.servconfid, callbackurl: 'servicepoints' }, skipLocationChange: true
  //   });
  // }
}
