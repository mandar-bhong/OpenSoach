import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { USER_CATEGORY } from '../../../../../shared/app-common-constants';
import {
  UserAddRequest,
  UserMasterResponse,
  UserMasterUpdateRequest,
  UserRoleidListItemResponse,
} from '../../../../../shared/models/api/user-models';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { UserAddModel } from '../../../../../shared/models/ui/user-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { AppUserService } from '../../../../../shared/services/user/app-user.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { UserService } from '../../../services/user.service';

@Component({
  selector: 'app-user-add',
  templateUrl: './user-add.component.html',
  styleUrls: ['./user-add.component.css']
})
export class UserAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new UserAddModel();
  userStates: EnumDataSourceItem<number>[];
  userCategories: EnumDataSourceItem<number>[];
  routeSubscription: Subscription;
  currentRecord: UserMasterResponse;
  uroleids: UserRoleidListItemResponse[];
  showCat = false;
  constructor(private userService: UserService,
    private userSharedService: AppUserService,
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
    this.userStates = this.userSharedService.getUserStates();
    this.userCategories = this.userService.getUsersCategories();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.userid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getUserEdit();
        this.removeControl();
      } else {
        this.subTitle = 'Add Details of User';
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      userCategory: new FormControl('', [Validators.required]),
      emailControl: new FormControl('', [Validators.required]),
      userroleControl: new FormControl(''),
      userStateControl: new FormControl('', [Validators.required])
    });
  }
  userCategoryChange() {

    if (this.dataModel.usrcategory === USER_CATEGORY.OSU) {
      // this.showCat = true;
      this.showCat = false;
      this.getRoleList();
    } else {
      // this.showCat = false;
      this.showCat = true;
      this.removeControladd();
    }
  }

  getRoleList() {
    this.userService.getRoleDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.uroleids = payloadResponse.data.filter(r => r.prodcode === null);
      }
    });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const userAddRequest = new UserAddRequest();
      this.dataModel.copyTo(userAddRequest);
      this.userService.addUser(userAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.userid = payloadResponse.data.recid;
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.usrname;
          this.removeControl();
        }
      });
      this.inProgress = false;
    } else {
      const userMasterUpdateRequest = new UserMasterUpdateRequest();
      this.dataModel.copyToUpdateRequest(userMasterUpdateRequest);
      this.userService.updateUserMaster(userMasterUpdateRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.setFormMode(FORM_MODE.VIEW);
          this.subTitle = this.dataModel.usrname;
        }
      });
      this.inProgress = false;
    }
  }

  getUserEdit() {
    this.userService.getUserEdit({ recid: this.dataModel.userid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.userCategoryChange();
          this.subTitle = this.dataModel.usrname;
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
  getroleidlist(value: number) {
    if (this.uroleids && value) {
      return this.uroleids.find(a => a.uroleid === value).urolename;
    }
  }
  getuserCategorie(value: number) {
    if (this.userCategories && value) {
      return this.userCategories.find(a => a.value === value).text;
    }
  }
  getuserStates(value: number) {
    if (this.userStates && value) {
      return this.userStates.find(a => a.value === value).text;
    }
  }
  removeControl() {
    if (this.dataModel.usrcategory === USER_CATEGORY.CU) {
      this.editableForm.removeControl('userroleControl');
    }
    this.editableForm.removeControl('emailControl');
    this.editableForm.removeControl('userCategory');
  }
  removeControladd() {
    if (this.dataModel.usrcategory === USER_CATEGORY.CU) {
      this.editableForm.removeControl('userroleControl');
    }
  }
}
