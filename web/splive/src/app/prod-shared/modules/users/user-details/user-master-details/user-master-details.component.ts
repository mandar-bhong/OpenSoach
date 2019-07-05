import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { APP_SHARED_DATA_STORE_KEYS } from '../../../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../../../shared/environment-provider';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppDataStoreService } from '../../../../../shared/services/app-data-store/app-data-store-service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { AppUserService } from '../../../../../shared/services/user/app-user.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { ProductcodeRequest, UserMasterUpdateRequest, UserRoleidListItemResponse } from '../../../../models/api/user-models';
import { UserDetailsModel, UserInfo } from '../../../../models/ui/user-models';
import { ProdUserService } from '../../../../services/user/prod-user.service';

@Component({
  selector: 'app-user-master-details',
  templateUrl: './user-master-details.component.html',
  styleUrls: ['./user-master-details.component.css']
})
export class UserMasterDetailsComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new UserDetailsModel();
  routeSubscription: Subscription;
  userStates: EnumDataSourceItem<number>[];
  uroleids: UserRoleidListItemResponse[];
  constructor(private appUserService: AppUserService,
    private prodUserService: ProdUserService,
    private route: ActivatedRoute,
    private appDataStoreService: AppDataStoreService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'User Account Information';
  }

  ngOnInit() {
    this.createControls();
    this.getRoleList();
    this.userStates = this.appUserService.getUserStates();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.usrid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getUserMasterDetails();
      } else {
        this.subTitle = 'Add Master Details of User';
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
        this.getUserProfileMasterDetails();
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      emailControl: new FormControl('', [Validators.email]),
      userroleControl: new FormControl('', [Validators.required]),
      userStateControl: new FormControl('', [Validators.required]),
    });
  }
  getRoleList() {
    const prodectcodeRequest = new ProductcodeRequest();
    prodectcodeRequest.prodcode = EnvironmentProvider.prodcode;
    this.prodUserService.getRoleMasterDataList(prodectcodeRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.uroleids = payloadResponse.data.filter(r => r.prodcode === null);
      }
    });
  }
  getUserMasterDetails() {
    this.prodUserService.getUserMasterDetails({ recid: this.dataModel.usrid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFromMaster(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.usrname;
          const userInfo = this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_INFO)
            .getObject<UserInfo>(APP_SHARED_DATA_STORE_KEYS.USER_INFO);
          if (this.dataModel.usrname === userInfo.usrname) {
            this.isEditable = false;
            this.showBackButton = false;
          } else {
            this.isEditable = true;
            this.showBackButton = false;
          }
        } else {
          this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
  }
  getUserProfileMasterDetails() {
    this.prodUserService.getUserProfileMasterDetails().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFromMaster(payloadResponse.data);
          this.prodUserService.userID = payloadResponse.data.userid;
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.usrname;
          const userInfo = this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_INFO)
            .getObject<UserInfo>(APP_SHARED_DATA_STORE_KEYS.USER_INFO);
          if (this.dataModel.usrname === userInfo.usrname) {
            this.isEditable = false;
            this.showBackButton = false;
          } else {
            this.isEditable = true;
            this.showBackButton = false;
          }
        } else {
          this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
  }
  savemaster() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    const request = new UserMasterUpdateRequest();
    this.dataModel.copyToUpdateRequest(request);
    this.prodUserService.updateUserEdit(request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.subTitle = this.dataModel.usrname;
      }
      this.inProgress = false;
    });
  }
  getuserrole(value: number) {
    if (this.uroleids && value) {
      return this.uroleids.find(a => a.uroleid === value).urolename;
    }
  }
  getuserstate(value: number) {
    if (this.userStates && value) {
      return this.userStates.find(a => a.value === value).text;
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
