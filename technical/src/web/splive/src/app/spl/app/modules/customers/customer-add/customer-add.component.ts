import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';

import { CustomerAddRequest } from '../../../models/api/customer-models';
import { CorporateShortDataResponse } from '../../../models/api/corporate-models';
import { CustomerAddModel } from '../../../models/ui/customer-models';
import { CustomerService } from '../../../services/customer.service';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { CorporateService } from '../../../services/corporate.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';

@Component({
  selector: 'app-customer-add',
  templateUrl: './customer-add.component.html',
  styleUrls: ['./customer-add.component.css']
})
export class CustomerAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new CustomerAddModel();
  customerStates: EnumDataSourceItem<number>[];
  corporates: CorporateShortDataResponse[] = [];
  routeSubscription: Subscription;
  constructor(private customerService: CustomerService,
    private corporateService: CorporateService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.customerStates = this.customerService.getCustomerStates();
    this.recordState = EDITABLE_RECORD_STATE.ADD;
    this.setFormMode(FORM_MODE.EDITABLE);
    this.getCorporateList();

  }
  createControls(): void {
    this.editableForm = new FormGroup({
      emailControl: new FormControl('', [Validators.required]),
      userState: new FormControl('', [Validators.required]),
      corprateId: new FormControl('', [Validators.required])
    });
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.recordState = EDITABLE_RECORD_STATE.ADD;
      this.setFormMode(FORM_MODE.EDITABLE);
      this.callbackUrl = params['callbackurl'];
    });
  }
  getCorporateList() {
    this.corporateService.getCorporateShortDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.corporates = payloadResponse.data;
      }
    });
  }
  save() {
    const customerAddRequest = new CustomerAddRequest();
    this.dataModel.copyTo(customerAddRequest);
    this.customerService.addCustomer(customerAddRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_ADD_CUTOMERS_SAVED'));
      }
    });
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  editRecord(id: number) {
    this.router.navigate(['customers', 'add'], { queryParams: { id: id, callbackurl: 'customers' }, skipLocationChange: true });
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
