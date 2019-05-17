import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatTableDataSource } from '@angular/material';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { CustomerServiceAssociateListResponse, CustomerServiceAssociateUpdateRequest } from '../../../models/api/customer-models';
import { ProductListItemResponse } from '../../../models/api/product-models';
import { CustomerSeviceAssociateProductModel } from '../../../models/ui/customer-models';
import { CustomerService } from '../../../services/customer.service';
import { ProductService } from '../../../services/product.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';

@Component({
  selector: 'app-customer-servicepoint-associate',
  templateUrl: './customer-servicepoint-associate.component.html',
  styleUrls: ['./customer-servicepoint-associate.component.css']
})
export class CustomerServicepointAssociateComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new CustomerSeviceAssociateProductModel();
  routeSubscription: Subscription;
  displayedColumns = ['prodcode', 'spcount', 'action'];
  showForm = false;
  cpmlist: CustomerServiceAssociateListResponse[];
  dataSource: MatTableDataSource<CustomerServiceAssociateListResponse>;
  currentRecord: CustomerServiceAssociateListResponse;
  constructor(private customerService: CustomerService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.custid = Number(params['id']);
        this.getCustomerService();
      this.callbackUrl = params['callbackurl'];
    });
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      productControl: new FormControl(''),
      updatecountControl: new FormControl('', [Validators.required])
    });
  }
  closeForm() {
    this.showForm = false;
    this.currentRecord = null;
    this.editableForm.reset();
  }

  getCustomerService() {
    this.customerService.getCustomerServiceProductAssociation({ recid: this.dataModel.custid })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.cpmlist = payloadResponse.data;
          this.dataSource = new MatTableDataSource<CustomerServiceAssociateListResponse>(this.cpmlist);
          if (this.cpmlist === null || this.cpmlist.length === 0) {
            this.appNotificationService.info(this.translatePipe.transform('CUSTOMER_IS_NOT_ASSOCIATED'));
          }
        }
      });
  }
  closeForms() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    const request = new CustomerServiceAssociateUpdateRequest();
    this.dataModel.copyToUpdateRequest(request);
    this.customerService.updateCustomerServicePointProduct(request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.closeForm();
        this.getCustomerService();
      }
    });
    this.inProgress = false;
  }
  editRecord(cpm: CustomerServiceAssociateListResponse) {
    this.currentRecord = cpm;
    this.editableForm.reset();
    this.recordState = EDITABLE_RECORD_STATE.UPDATE;
    this.setFormMode(FORM_MODE.EDITABLE);
    this.createControls();
    this.editableForm.controls['productControl'].disable();
    this.dataModel.copyFrom(cpm);
    this.showForm = true;
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
