import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { HPFTRouteHelper } from "../../../../hpft/app/helpers/route-helper";
import { PROD_HPFT } from '../../../app-common-constants';
import { AppRepoShared } from '../../../app-repo/app-repo';
import { AppSpecificDataProvider } from '../../../app-specific-data-provider';
import { EnvironmentProvider } from '../../../environment-provider';
import { AuthRequest } from '../../../models/api/auth-models';
import { TranslatePipe } from '../../../pipes/translate/translate.pipe';
import { AuthService } from '../../../services/auth.service';
import { LoginHandlerService } from '../../../services/login-handler.service';
import { AppNotificationService } from '../../../services/notification/app-notification.service';



@Component({
  selector: 'hkt-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  hide = true;
  username: string;
  password: string;
  loginform: FormGroup;
  flipped = false;
  userHomeRoute: any;


  constructor(private loginHandlerService: LoginHandlerService,
    private router: Router,
    private authService: AuthService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe
  ) {
  }


  ngOnInit() {

    switch (AppRepoShared.appProductCode) {
      case PROD_HPFT:
        this.userHomeRoute = HPFTRouteHelper.getUserHomeRoute;
        break;
      default:
        this.userHomeRoute = this.userHomeRoutHandler;
        break;
    }

    this.createControls();
  }


  createControls(): void {
    this.loginform = new FormGroup({
      emailControl: new FormControl('', [Validators.required]),
      passwordControl: new FormControl('', [Validators.required]),
    });
  }

  flipIt() {
    this.flipped = !this.flipped;
  }

  login() {
    if (this.loginform.invalid) {
      return;
    }
    const authRequest = new AuthRequest();
    authRequest.username = this.username;
    authRequest.password = this.password;
    authRequest.prodcode = EnvironmentProvider.prodcode;

    this.authService.login(authRequest).subscribe(response => {
      if (response && response.issuccess) {
        if (AppSpecificDataProvider.userCateory === response.data.usrcategory) {
          this.loginHandlerService.login(response.data);
          this.router.navigate([this.userHomeRoute(response.data.urolecode)], { skipLocationChange: true });
        } else {
          this.appNotificationService.error(this.translatePipe.transform('ERROR_LOGIN_INVALID_CATEGORY'));
        }
      }
    });
  }


  userHomeRoutHandler(userrole : string) {
    this.router.navigate([''], { skipLocationChange: true });
  }

}
