import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { APP_SHARED_DATA_STORE_KEYS } from '../../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../../shared/environment-provider';
import { EnumDataSourceItem } from '../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppDataStoreService } from '../../../../shared/services/app-data-store/app-data-store-service';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { AppUserService } from '../../../../shared/services/user/app-user.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../shared/views/edit-record-base';
import {
  ProductcodeRequest,
  UserAddDetailsRequest,
  UserMasterResponse,
  UserMasterUpdateRequest,
  UserRoleidListItemResponse,
} from '../../../models/api/user-models';
import { UserDetailsModel, UserInfo } from '../../../models/ui/user-models';
import { ProdUserService } from '../../../services/user/prod-user.service';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new UserDetailsModel();
  routeSubscription: Subscription;
  userStates: EnumDataSourceItem<number>[];
  currentRecord: UserMasterResponse;
  userGenders: EnumDataSourceItem<number>[];
  uroleids: UserRoleidListItemResponse[];
  showForm = false;
  showFormedit = false;
  constructor(private appUserService: AppUserService,
    private prodUserService: ProdUserService,
    private route: ActivatedRoute,
    private userSharedService: AppUserService,
    private appDataStoreService: AppDataStoreService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.getRoleList();
    this.userStates = this.appUserService.getUserStates();
    this.userGenders = this.appUserService.getUsersGender();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.usrid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getUserMasterDetails();
        this.getUserDetails();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      emailControl: new FormControl(''),
      userroleControl: new FormControl(''),
      userStateControl: new FormControl('', [Validators.required]),
      fnameControl: new FormControl('', [Validators.required]),
      lnameControl: new FormControl('', [Validators.required]),
      mobilenoControl: new FormControl(''),
      userGenderControl: new FormControl('', [Validators.required]),
      alternateContactControl: new FormControl('')
    });
  }
  getUserMasterDetails() {
    this.prodUserService.getUserMasterDetails({ recid: this.dataModel.usrid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFromMaster(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          const userInfo = this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_INFO)
            .getObject<UserInfo>(APP_SHARED_DATA_STORE_KEYS.USER_INFO);
          if (this.dataModel.usrname === userInfo.usrname) {
            this.showForm = false;
            this.showFormedit = true;
          } else {
            this.showFormedit = false;
            this.showForm = true;

          }
        } else {
          // this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
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

  getUserDetails() {
    this.prodUserService.getUserDetails({ recid: this.dataModel.usrid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        } else {
          // this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    const userAddDetailsRequest = new UserAddDetailsRequest();
    this.dataModel.copyTo(userAddDetailsRequest);
    this.prodUserService.updateUserDetails(userAddDetailsRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_USERS_DETAILS_SAVED'));
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
      }
    });
  }
  savemaster() {
    const request = new UserMasterUpdateRequest();
    this.dataModel.copyToUpdateRequest(request);
    this.prodUserService.updateUserEdit(request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform(''));
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
      }
    });
  }

  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  editInfo() {
    this.editForm();
    this.editableForm.controls['emailControl'].disable();
    this.editableForm.controls['userStateControl'].disable();
    this.editableForm.controls['userroleControl'].disable();

  }
  editMaster() {
    this.editForm();
    this.editableForm.controls['emailControl'].disable();
    this.editableForm.controls['fnameControl'].disable();
    this.editableForm.controls['lnameControl'].disable();
    this.editableForm.controls['mobilenoControl'].disable();
    this.editableForm.controls['userGenderControl'].disable();
    this.editableForm.controls['userGenderControl'].disable();
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
