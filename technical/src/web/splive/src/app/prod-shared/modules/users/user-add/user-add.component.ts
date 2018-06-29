import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { EnvironmentProvider } from '../../../../shared/environment-provider';
import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../shared/views/edit-record-base';
import { ProductcodeRequest, UserAddRequest, UserRoleidListItemResponse } from '../../../models/api/user-models';
import { UserAddModel } from '../../../models/ui/user-models';
import { ProdUserService } from '../../../services/user/prod-user.service';

@Component({
  selector: 'app-user-add',
  templateUrl: './user-add.component.html',
  styleUrls: ['./user-add.component.css']
})
export class UserAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new UserAddModel();
  routeSubscription: Subscription;
  prodcode: ProductcodeRequest;
  uroleids: UserRoleidListItemResponse[];
  constructor(private prodUserService: ProdUserService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'User Details';
  }
  ngOnInit() {
    this.createControls();
    this.getRoleList();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.userid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
      } else {
        this.subTitle = 'Add Details of User';
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  getRoleList() {
    const prodectcodeRequest = new ProductcodeRequest();
    prodectcodeRequest.prodcode = EnvironmentProvider.prodcode;
    this.prodUserService.getRoleDataList(prodectcodeRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.uroleids = payloadResponse.data.filter(r => r.prodcode === null);
      }
    });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      emailControl: new FormControl('', [Validators.required]),
      userroleControl: new FormControl('', [Validators.required]),
    });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const userAddRequest = new UserAddRequest();
      this.dataModel.copyTo(userAddRequest);
      this.prodUserService.addUser(userAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.userid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.usrname;
        }
        this.inProgress = false;
      });
    }
  }
  getuserrole(value: number) {
    if (this.uroleids && value) {
      return this.uroleids.find(a => a.uroleid === value).urolename;
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

