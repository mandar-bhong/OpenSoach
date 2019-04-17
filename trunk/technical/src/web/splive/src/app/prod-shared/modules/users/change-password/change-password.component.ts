import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EditRecordBase, FORM_MODE, EDITABLE_RECORD_STATE } from '../../../../shared/views/edit-record-base';
import { ChangeUserPasswordRequest } from '../../../models/api/user-models';
import { ConfirmPasswordModel } from '../../../models/ui/user-models';
import { ProdUserService } from '../../../services/user/prod-user.service';
import { LoginHandlerService } from '../../../../shared/services/login-handler.service';
import { ROUTE_LOGIN } from '../../../../shared/app-common-constants';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent extends EditRecordBase implements OnInit, OnDestroy {
  routeSubscription: Subscription;
  dataModel = new ConfirmPasswordModel();
  hide = true;
  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private prodUserService: ProdUserService,
    private loginHandlerService: LoginHandlerService,
  ) {
    super();
    this.iconCss = 'fa fa-key';
    this.pageTitle = 'Change Password';
  }

  ngOnInit() {
    this.showBackButton = false;
    this.createControls();
    this.setFormMode(FORM_MODE.EDITABLE);
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      curpasswordControl: new FormControl('', [Validators.required]),
      newpasswordControl: new FormControl('', [Validators.required]),
      confirmpasswordControl: new FormControl('', [Validators.required]),
    });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    const changeUserPasswordRequest = new ChangeUserPasswordRequest();
    this.dataModel.copyTo(changeUserPasswordRequest);
    if (this.dataModel.newpassword === this.dataModel.confirmpassword && this.dataModel.oldpassword != this.dataModel.newpassword) {
      this.prodUserService.changeUserPassword(changeUserPasswordRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.router.navigate([ROUTE_LOGIN], { skipLocationChange: true });
        }
      });
    }
    else {
      this.appNotificationService.error(this.translatePipe.transform('CHANGE_PASS'));
    }
    this.inProgress = false;
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
