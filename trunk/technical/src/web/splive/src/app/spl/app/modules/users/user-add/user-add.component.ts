import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';

import { UserAddRequest, UserRoleidListItemResponse } from '../../../../../shared/models/api/user-models';

import { UserAddModel } from '../../../../../shared/models/ui/user-models';
import { UserService } from '../../../services/user.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { RecordIDResponse, RecordIDRequest } from '../../../../../shared/models/api/common-models';

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

  uroleids: UserRoleidListItemResponse[];

  constructor(private userService: UserService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }
  ngOnInit() {
    this.userStates = this.userService.getUserStates();
    this.userCategories = this.userService.getUsersCategories();
    this.createControls();
    this.getRoleList();
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      userCategory: new FormControl('', [Validators.required]),
      emailControl: new FormControl('', [Validators.required]),
      userroleControl: new FormControl(''),
      userStateControl: new FormControl('', [Validators.required])
    });
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.recordState = EDITABLE_RECORD_STATE.ADD;
      this.setFormMode(FORM_MODE.EDITABLE);
      this.callbackUrl = params['callbackurl'];
    });
  }
  getRoleList() {
    this.userService.getRoleDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.uroleids = payloadResponse.data;
      }
    });
  }
  save() {
    const userAddRequest = new UserAddRequest();
    this.dataModel.copyTo(userAddRequest);
    this.userService.addUser(userAddRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_ADD_CUTOMERS_SAVED'));
      } else {
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.EDITABLE);
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
