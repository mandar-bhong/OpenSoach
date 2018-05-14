
import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatTableDataSource } from '@angular/material';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { UserSharedService } from '../../../../../shared/services/user/user-shared.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import {
  UserAssociateProductListItemResponse,
  UserAssociateProductRequest,
  UserAssociateProductUpdateRequest,
  UserRoleidListItemResponse
} from '../../../../../shared/models/api/user-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { CustomerListItemResponse } from '../../../models/api/customer-models';
import { ProductListItemResponse } from '../../../models/api/product-models';
import { UserAssociateProductModel } from '../../../../../shared/models/ui/user-models';
import { UserService } from '../../../services/user.service';
import { ProductService } from '../../../services/product.service';
import { CustomerService } from '../../../services/customer.service';
import { CustomerAssociateProductListItemResponse, CustomerRoleidListItemResponse } from '../../../models/api/customer-models';
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
  showEditForm = false;
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
    public userSharedService: UserSharedService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private route: ActivatedRoute,
    private router: Router
  ) {
    super();
  }

  ngOnInit() {
    this.createControls();
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
    this.getCpmList();
  }

  getCpmList() {
    this.customerService.getCustomerProductAssociation({ recid: this.dataModel.custid })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.cpmlist = payloadResponse.data;
        } else {
          this.appNotificationService.info(this.translatePipe.transform('CUSTOMER_IS_NOT_BEEN_ASSOCIATED_WITH_ANY_PRODUCT'));
          // TODD: display informative message saying customer is not been associated with any product yet.
        }
      });
  }
  getCustRoleList() {
    this.customerService.getCustRoleDataList({ prodcode: this.dataModel.cpm.prodcode })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.uroleids = payloadResponse.data;
        }
      });
  }

  add() {
    this.editableForm.reset();
    this.recordState = EDITABLE_RECORD_STATE.ADD;
    this.setFormMode(FORM_MODE.EDITABLE);
    this.showForm = true;
  }

  closeForm() {
    this.showForm = false;
    this.showEditForm = false;
    this.currentRecord = null;
    this.editableForm.reset();
  }

  editRecord(ucpm: UserAssociateProductListItemResponse) {
    this.currentRecord = ucpm;
    this.editableForm.reset();
    console.log(this.editableForm.controls['productControl']);
    this.recordState = EDITABLE_RECORD_STATE.UPDATE;
    this.setFormMode(FORM_MODE.EDITABLE);
    // this.editableForm.controls['productControl'].disable();
    // this.editableForm.controls['custnameControl'].disable();
    // this.editableForm.controls['urolecodeControl'].disable();
    this.dataModel.copyFrom(ucpm);
    this.showForm = false;
    this.showEditForm = true;

  }


  save() {
    if (this.editableForm.invalid) { return; }
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const request = new UserAssociateProductRequest();
      this.dataModel.copyToAddRequest(request);
      this.userService.associateUserToProduct(request).subscribe(payloadResponse => {
        console.log('payloadResponse', payloadResponse);
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.ucpmid = payloadResponse.data.recid;
          this.getUcpmList();
          this.closeForm();
        }
      });
    } else {
      const request = new UserAssociateProductUpdateRequest();
      this.dataModel.copyToUpdateRequest(request);
      this.userService.updateAssociateUserToProduct(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          console.log('this.dataModel.ucpmstate', this.dataModel.ucpmstate);
          this.currentRecord.ucpmstate = this.dataModel.ucpmstate;
          this.closeForm();
        }
      });
    }
  }

  closeWindow() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
