import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { APP_SHARED_DATA_STORE_KEYS } from '../../../../../shared/app-common-constants';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppDataStoreService } from '../../../../../shared/services/app-data-store/app-data-store-service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { AppUserService } from '../../../../../shared/services/user/app-user.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { UserAddDetailsRequest } from '../../../../models/api/user-models';
import { UserDetailsModel, UserInfo } from '../../../../models/ui/user-models';
import { ProdUserService } from '../../../../services/user/prod-user.service';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.css']
})
export class UserInfoComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new UserDetailsModel();
  routeSubscription: Subscription;
  userGenders: EnumDataSourceItem<number>[];
  constructor(private appUserService: AppUserService,
    private prodUserService: ProdUserService,
    private route: ActivatedRoute,
    private appDataStoreService: AppDataStoreService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'User Profile';
  }

  ngOnInit() {
    this.showBackButton = false;
    this.createControls();
    this.userGenders = this.appUserService.getUsersGender();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.usrid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getUserDetails();
        this.getUserMasterDetails();
      } else {
        this.subTitle = 'Add Details of User';
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
        this.getUserProfileDetails();
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      fnameControl: new FormControl(''),
      lnameControl: new FormControl(''),
      mobilenoControl: new FormControl('', Validators.pattern(/^\d+$/)),
      userGenderControl: new FormControl('', [Validators.required]),
      alternateContactControl: new FormControl('', Validators.pattern(/^\d+$/))
    });
  }
  getUserDetails() {
    this.prodUserService.getUserDetails({ recid: this.dataModel.usrid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = (this.dataModel.fname + ' ' + this.dataModel.lname);
        } else {
          this.appNotificationService.info(this.translatePipe.transform('USER_INFO_DETAILS_NOT_AVAILABLE'));
        }
      }
    });
  }
  getUserProfileDetails() {
    this.prodUserService.getUserProfileDetails().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = (this.dataModel.fname + ' ' + this.dataModel.lname);
          this.showBackButton = false;
        }
         else {
          this.appNotificationService.info(this.translatePipe.transform('USER_INFO_DETAILS_NOT_AVAILABLE'));
        }
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
          const userInfo = this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_INFO)
            .getObject<UserInfo>(APP_SHARED_DATA_STORE_KEYS.USER_INFO);
          // console.log('userInfo', userInfo);
          if (this.dataModel.usrname === userInfo.usrname) {
            this.isEditable = true;
            this.showBackButton = false;
          } else {
            this.isEditable = false;
            this.showBackButton = false;
          }
        } else {
          this.appNotificationService.info(this.translatePipe.transform('USER_INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    const userAddDetailsRequest = new UserAddDetailsRequest();
    this.dataModel.usrid = this.prodUserService.userID;
    this.dataModel.copyTo(userAddDetailsRequest);
    this.prodUserService.updateUserDetails(userAddDetailsRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.subTitle = (this.dataModel.fname + ' ' + this.dataModel.lname);
      }
      this.inProgress = false;
    });
  }
  getgender(value: number) {
    if (this.userGenders && value) {
      return this.userGenders.find(a => a.value === value).text;
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
