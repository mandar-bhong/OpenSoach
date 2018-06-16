import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormControl, FormGroup, Validators } from '@angular/forms';

import { AppSpecificDataProvider } from '../../../app-specific-data-provider';
import { EnvironmentProvider } from '../../../environment-provider';
import { AuthRequest } from '../../../models/api/auth-models';
import { TranslatePipe } from '../../../pipes/translate/translate.pipe';
import { AppDataStoreService } from '../../../services/app-data-store/app-data-store-service';
import { AuthService } from '../../../services/auth.service';
import { LoginHandlerService } from '../../../services/login-handler.service';
import { AppNotificationService } from '../../../services/notification/app-notification.service';
import { EditRecordBase } from '../../../views/edit-record-base';

@Component({
  selector: 'hkt-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent extends EditRecordBase implements OnInit {
  hide = true;
  username: string;
  password: string;
  constructor(private appDataStoreService: AppDataStoreService,
    private loginHandlerService: LoginHandlerService,
    private router: Router,
    private authService: AuthService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe
  ) {
    super();
   }
  ngOnInit() {
    this.createControls();
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      emailControl: new FormControl('', [Validators.required]),
      passwordControl: new FormControl('', [Validators.required]),
    });
  }
  login() {
    const authRequest = new AuthRequest();
    authRequest.username = this.username;
    authRequest.password = this.password;
    authRequest.prodcode = EnvironmentProvider.prodcode;

    this.authService.login(authRequest).subscribe(response => {
      if (response && response.issuccess) {
        if (AppSpecificDataProvider.userCateory === response.data.usrcategory) {
          this.loginHandlerService.login(response.data);
          this.router.navigate([''], { skipLocationChange: true });
        } else {
          this.appNotificationService.error(this.translatePipe.transform('ERROR_LOGIN_INVALID_CATEGORY'));
        }
      }
    });
  }
  closeForm() {
  }
}
