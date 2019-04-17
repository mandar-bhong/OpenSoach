import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { ActivationChangePassword, ChangeUserPasswordRequest } from '../../../models/api/user-models';
import { ActivationChangePasswordModel, ConfirmPasswordModel } from '../../../models/ui/user-models';
import { TranslatePipe } from '../../../pipes/translate/translate.pipe';
import { AuthService } from '../../../services/auth.service';
import { LoginHandlerService } from '../../../services/login-handler.service';
import { AppNotificationService } from '../../../services/notification/app-notification.service';
import { EditRecordBase, FORM_MODE } from '../../../views/edit-record-base';
import { ROUTE_LOGIN } from '../../../app-common-constants';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent extends EditRecordBase implements OnInit, OnDestroy {
  routeSubscription: Subscription;
  dataModel = new ConfirmPasswordModel();
  successHide = true;
  code: string;
  activationCodeSubscription: Subscription;
  activationChangePassword: ActivationChangePassword;
  activateSubscription: Subscription;
  queryParameter: any;
  receivedCode: string;
  userId: number;
  activationChangePasswordModel = new ActivationChangePasswordModel();
  newpassword: string;
  hideconfirm = true;
  hidenew = true;
  firstView = false;
  secondView = false;
  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private authService: AuthService,
    private loginHandlerService: LoginHandlerService,
  ) {
    super();
  }

  ngOnInit() {
    this.showBackButton = false;    
    this.createControls();
    this.setFormMode(FORM_MODE.EDITABLE);
    this.activateSubscription = this.route.params.subscribe(params => {
      this.queryParameter = JSON.parse(JSON.stringify(params));
      this.receivedCode = this.queryParameter.code;
    });

    this.getActivationParams();

  }
  createControls(): void {
    this.editableForm = new FormGroup({
      newpasswordControl: new FormControl('', [Validators.required]),
      confirmpasswordControl: new FormControl('', [Validators.required]),
    });
  }

  save() {
    const changeUserPasswordRequest = new ChangeUserPasswordRequest();
    this.dataModel.usrid = this.userId;
    this.dataModel.copyTo(changeUserPasswordRequest);
    if (this.dataModel.newpassword === this.dataModel.confirmpassword) {
      this.authService.changeUserPassword(changeUserPasswordRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.successHide = !this.successHide;
          setTimeout(() => {
            // this.loginHandlerService.logout();
            this.router.navigate([ROUTE_LOGIN], { skipLocationChange: true });
        }, 5000);
        }
      });
    }
    else {
      this.appNotificationService.error(this.translatePipe.transform('CHANGE_PASS'));
    }
  }


  getActivationParams() {
    // write code of get 
    // if response success then manage if conditions.
    const activationChangePassword = new ActivationChangePassword;
    activationChangePassword.code = this.receivedCode;
    this.authService.getActivationPerams(activationChangePassword).subscribe(PayloadResponse => {
      if (PayloadResponse && PayloadResponse.issuccess) {
        this.userId = PayloadResponse.data.recid;
        this.firstView = true;
        this.secondView = false;
      } else{
        this.firstView = false;
        this.secondView = true;
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
    if (this.activateSubscription) {
      this.activateSubscription.unsubscribe();
    }

  }
}
