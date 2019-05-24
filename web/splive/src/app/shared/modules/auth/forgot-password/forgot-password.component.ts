import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Subscription } from 'rxjs';
import { ForgotPasswordRequest, ResetPasswordRequest } from '../../../models/api/user-models';
import { ForgotPasswordModel } from '../../../models/ui/user-models';
import { TranslatePipe } from '../../../pipes/translate/translate.pipe';
import { AuthService } from '../../../services/auth.service';
import { AppNotificationService } from '../../../services/notification/app-notification.service';

@Component({
  selector: 'hkt-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css']
})
export class ForgotPasswordComponent implements OnInit {
  flipped = false;
  forgotpasswordform: FormGroup;
  resetPasswordForm: FormGroup;
  otpfield = true;
  dataModel = new ForgotPasswordModel();
  passwordfield = true;
  userid = false;
  hideconfirm = true;
  hidenewPass = true;
  setHide: boolean;
  removeImageSubscription: Subscription;
  imageHide = false;

  constructor(private authService: AuthService,
    private translatePipe: TranslatePipe,
    private appNotificationService: AppNotificationService) { }

  ngOnInit() {
    this.createControls();
    this.createControlsForPassword();
  }
  // FOR EMAIL VALIDATION
  createControls(): void {
    this.forgotpasswordform = new FormGroup({
      emailControl: new FormControl('', [Validators.required]),
      otpControl: new FormControl('', [Validators.required])
    });
  }

  // FOR PASSWORD VALIDATION
  createControlsForPassword(): void {
    this.resetPasswordForm = new FormGroup({
      newPasswordControl: new FormControl('', [Validators.required]),
      confirmPasswordControl: new FormControl('', [Validators.required])
    });
  }

  // USED FOR SEND OTP TO EMAIL
  sendOtp() {
    const forgotPasswordRequest = new ForgotPasswordRequest();
    this.dataModel.copyToForgotPass(forgotPasswordRequest);
    this.authService.forgotUserPassword(forgotPasswordRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.otpfield = false;
        this.appNotificationService.success(this.translatePipe.transform('EMAIL_VERIFIED_SUCCESS'));
      }
    });
  }

  //USED TO CHANGE PASSWORD
  resetPassword() {

   if(this.resetPasswordForm.valid){
    const resetPasswordRequest = new ResetPasswordRequest();
    this.dataModel.copyToResetPass(resetPasswordRequest);
    if (this.dataModel.newpassword != null) {
      if (this.dataModel.newpassword === this.dataModel.confirmpassword) {

        this.authService.resetUserPassword(resetPasswordRequest).subscribe(payloadResponse => {
          if (payloadResponse && payloadResponse.issuccess) {
            this.appNotificationService.success(this.translatePipe.transform('RESET_PASSWORD_SUCCESS'));
            //for reading success message on UI set timeout
            setTimeout(() => {
              this.setHide = true;
              this.otpfield = true;
              this.userid = true;

              this.authService.setImageVisibility(false);
            }, 2500);
          }
        });
      }
      else {
        this.appNotificationService.error(this.translatePipe.transform('CHANGE_PASS'));
      }
    }
   }
   
  }// end of fucntion

}



