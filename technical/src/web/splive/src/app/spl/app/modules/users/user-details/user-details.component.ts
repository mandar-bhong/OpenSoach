import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { UserAddDetailsRequest } from '../../../../../shared/models/api/user-models';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { UserDetailsModel } from '../../../../../shared/models/ui/user-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { UserService } from '../../../services/user.service';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new UserDetailsModel();
  routeSubscription: Subscription;
  userGenders: EnumDataSourceItem<number>[];
  constructor(private userService: UserService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe
  ) {
    super();
  }

  ngOnInit() {
    this.userGenders = this.userService.getUsersGender();
    this.createControls();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.usrid = Number(params['id']);
      this.callbackUrl = params['callbackurl'];
      this.getUserDetails();
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      fnameControl: new FormControl('', [Validators.required]),
      lnameControl: new FormControl('', [Validators.required]),
      mobilenoControl: new FormControl(''),
      userGenderControl: new FormControl(''),
      alternateContactControl: new FormControl('')
    });
  }
  save() {
    if (this.editableForm.invalid) { return; }
    const userAddDetailsRequest = new UserAddDetailsRequest();
    this.dataModel.copyTo(userAddDetailsRequest);
    this.userService.updateUserDetails(userAddDetailsRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_USERS_DETAILS_SAVED'));
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
      } else { }
    });
  }
  getUserDetails() {
    this.userService.getUserDetails({ recid: this.dataModel.usrid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        } else {
          this.appNotificationService.info(this.translatePipe.transform('INFO_DETAILS_NOT_AVAILABLE'));
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }
    });
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }

  editRecord(id: number) {
    this.router.navigate(['users', 'user-detail'], { queryParams: { id: id, callbackurl: 'users' }, skipLocationChange: true });
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
