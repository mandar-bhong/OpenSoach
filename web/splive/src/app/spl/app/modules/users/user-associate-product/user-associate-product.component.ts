import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatTableDataSource } from '@angular/material';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import {
  UserAssociateProductListItemResponse,
  UserAssociateProductRequest,
  UserAssociateProductUpdateRequest,
} from '../../../../../shared/models/api/user-models';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { UserAssociateProductModel } from '../../../../../shared/models/ui/user-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { AppUserService } from '../../../../../shared/services/user/app-user.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import {
  CustomerAssociateProductListItemResponse,
  CustomerListItemResponse,
  CustomerRoleidListItemResponse,
} from '../../../models/api/customer-models';
import { ProductListItemResponse } from '../../../models/api/product-models';
import { CustomerService } from '../../../services/customer.service';
import { ProductService } from '../../../services/product.service';
import { UserService } from '../../../services/user.service';


@Component({
  selector: 'app-user-associate-product',
  templateUrl: './user-associate-product.component.html',
  styleUrls: ['./user-associate-product.component.css']
})
export class UserAssociateProductComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new UserAssociateProductModel();
  products: ProductListItemResponse[];
  cutomers: CustomerListItemResponse[];
  ucpmstates: EnumDataSourceItem<number>[];
  showForm = false;
  routeSubscription: Subscription;
  ucpmlist: UserAssociateProductListItemResponse[];
  uroleids: CustomerRoleidListItemResponse[];
  cpmlist: CustomerAssociateProductListItemResponse[];
  dataSource: MatTableDataSource<UserAssociateProductListItemResponse>;
  displayedColumns = ['custname', 'prodcode', 'urolecode', 'ucpmstate', 'action'];
  addVisiblity = false;
  currentRecord: UserAssociateProductListItemResponse;
  constructor(
    private productService: ProductService,
    private userService: UserService,
    private customerService: CustomerService,
    public userSharedService: AppUserService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private route: ActivatedRoute,
    private router: Router
  ) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.showBackButton = false;
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.usrid = Number(params['id']);
      this.callbackUrl = params['callbackurl'];
    });
    this.ucpmstates = this.userSharedService.getUcpmStates();
    this.productService.getDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.products = payloadResponse.data;
        this.getUcpmList();
      }
    });
    this.customerService.getCustomerNameList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.cutomers = payloadResponse.data;
        if (this.cutomers === null || this.cutomers.length === 0) {
          this.appNotificationService.info(this.translatePipe.transform('CUSTOMER_IS_NOT_BEEN_ASSOCIATED_WITH_ANY_PRODUCT_USER'));
        }
      }
    });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      custnameControl: new FormControl('', [Validators.required]),
      productControl: new FormControl('', [Validators.required]),
      userroleControl: new FormControl('', [Validators.required]),
      ucpmStateControl: new FormControl('', [Validators.required])
    });
  }
  getUcpmList() {
    this.userService.getUserProductAssociation({ recid: this.dataModel.usrid })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.ucpmlist = payloadResponse.data;
          this.dataSource = new MatTableDataSource<UserAssociateProductListItemResponse>(this.ucpmlist);
          if (this.ucpmlist.length < this.products.length) {
            this.addVisiblity = true;
          } else {
            this.addVisiblity = false;
          }
        }
      });
  }

  getCustomerData() {
    if (this.dataModel.custid && this.dataModel.custid > 0) {
      this.getCpmList();
    }
  }
  getCpmList() {
    this.customerService.getCustomerProductAssociation({ recid: this.dataModel.custid })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.cpmlist = payloadResponse.data;
          if (this.cpmlist === null || this.cpmlist.length === 0) {
            this.appNotificationService.info(this.translatePipe.transform('PRODUCT_IS_NOT_ASSOCIATED'));
          }
        }
      });
  }
  getCustRoleList() {
    if (this.dataModel.cpm && this.dataModel.cpm !== null) {
      if (this.dataModel.cpm.prodcode && this.dataModel.cpm.prodcode !== null) {
        this.customerService.getCustRoleDataList({ prodcode: this.dataModel.cpm.prodcode })
          .subscribe(payloadResponse => {
            if (payloadResponse && payloadResponse.issuccess) {
              this.uroleids = payloadResponse.data;
            }
          });
      }
    }
  }

  add() {
    this.editableForm.reset();
    this.recordState = EDITABLE_RECORD_STATE.ADD;
    this.setFormMode(FORM_MODE.EDITABLE);
    this.showForm = true;
  }

  closeForm() {
    this.showForm = false;
    this.editableForm.reset();
    // this.currentRecord = null;
  }

  editRecord(ucpm: UserAssociateProductListItemResponse) {
    this.currentRecord = ucpm;
    this.editableForm.reset();
    this.recordState = EDITABLE_RECORD_STATE.UPDATE;
    this.setFormMode(FORM_MODE.EDITABLE);
    this.editableForm.controls['custnameControl'].disable();
    this.editableForm.controls['productControl'].disable();
    this.editableForm.controls['userroleControl'].disable();
    this.dataModel.copyFrom(ucpm);
    this.showForm = true;
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const request = new UserAssociateProductRequest();
      this.dataModel.copyToAddRequest(request);
      this.userService.associateUserToProduct(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.ucpmid = payloadResponse.data.recid;
          this.getUcpmList();
          this.closeForm();
        }
      });
      this.inProgress = false;
    } else {
      const request = new UserAssociateProductUpdateRequest();
      this.dataModel.copyToUpdateRequest(request);
      this.userService.updateAssociateUserToProduct(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.currentRecord.ucpmstate = this.dataModel.ucpmstate;
          this.closeForm();
        }
      });
      this.inProgress = false;
    }
  }

  closeForms() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  getuserrole(value: number) {
    if (this.uroleids && value) {
      return this.uroleids.find(a => a.uroleid === value).urolename;
    }
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
