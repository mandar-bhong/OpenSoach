import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
import { EditRecordBase } from '../../../../shared/views/edit-record-base';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent extends EditRecordBase implements OnInit, OnDestroy {
  routeSubscription: Subscription;
  constructor(
    private route: ActivatedRoute,
    private router: Router,
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
      curpasswordControl: new FormControl('', [Validators.required]),
      newpasswordControl: new FormControl('', [Validators.required]),
      confirmpasswordControl: new FormControl('', [Validators.required]),
    });
  }

  save() {

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
