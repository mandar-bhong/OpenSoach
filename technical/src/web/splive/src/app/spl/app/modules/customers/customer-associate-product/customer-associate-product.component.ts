import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatTableDataSource } from '@angular/material';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { CustomerSharedService } from '../../../../../shared/services/customer/customer-shared.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import {
  CustomerAssociateProductListItemResponse,
  CustomerAssociateProductRequest,
  CustomerAssociateProductUpdateRequest,
} from '../../../models/api/customer-models';
import { DbInstanceListItemResponse } from '../../../models/api/db-instance-models';
import { ProductListItemResponse } from '../../../models/api/product-models';
import { CustomerAssociateProductModel } from '../../../models/ui/customer-models';
import { CustomerService } from '../../../services/customer.service';
import { DBInstanceService } from '../../../services/db-instance.service';
import { ProductService } from '../../../services/product.service';

@Component({
  selector: 'app-customer-associate-product',
  templateUrl: './customer-associate-product.component.html',
  styleUrls: ['./customer-associate-product.component.css']
})
export class CustomerAssociateProductComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new CustomerAssociateProductModel();
  products: ProductListItemResponse[];
  dbinstances: DbInstanceListItemResponse[];
  cpmstates: EnumDataSourceItem<number>[];
  showForm = false;
  routeSubscription: Subscription;
  cpmlist: CustomerAssociateProductListItemResponse[];
  dataSource: MatTableDataSource<CustomerAssociateProductListItemResponse>;
  displayedColumns = ['prodcode', 'dbiname', 'cpmstate', 'action'];
  addVisiblity = false;
  currentRecord: CustomerAssociateProductListItemResponse;
  constructor(private productService: ProductService,
    private customerService: CustomerService,
    public customerSharedService: CustomerSharedService,
    private dbInstanceService: DBInstanceService,
    private route: ActivatedRoute,
    private router: Router) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.cpmstates = this.customerSharedService.getCpmStates();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.custid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
    this.productService.getDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.products = payloadResponse.data;
        this.getCpmList();
      }
    });

    this.dbInstanceService.getDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dbinstances = payloadResponse.data;
      }
    });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      productControl: new FormControl('', [Validators.required]),
      dbInstanceControl: new FormControl('', [Validators.required]),
      cpmStateControl: new FormControl('', [Validators.required])
    });
  }

  getCpmList() {
    this.customerService.getCustomerProductAssociation({ recid: this.dataModel.custid })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.cpmlist = payloadResponse.data;
          this.dataSource = new MatTableDataSource<CustomerAssociateProductListItemResponse>(this.cpmlist);
          if (this.cpmlist.length < this.products.length) {
            this.addVisiblity = true;
          } else {
            this.addVisiblity = false;
          }
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
    this.currentRecord = null;
    this.editableForm.reset();
  }

  editRecord(cpm: CustomerAssociateProductListItemResponse) {
    this.currentRecord = cpm;
    // this.editableForm.reset();
    this.recordState = EDITABLE_RECORD_STATE.UPDATE;
    this.setFormMode(FORM_MODE.EDITABLE);
    this.editableForm.controls['productControl'].disable();
    this.editableForm.controls['dbInstanceControl'].disable();
    this.dataModel.copyFrom(cpm);
    this.showForm = true;
  }

  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const request = new CustomerAssociateProductRequest();
      this.dataModel.copyToAddRequest(request);
      this.customerService.associateCustomerToProduct(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.cpmid = payloadResponse.data.recid;
          this.getCpmList();
          this.closeForm();
        }
      });
      this.inProgress = false;
    } else {
      const request = new CustomerAssociateProductUpdateRequest();
      this.dataModel.copyToUpdateRequest(request);
      this.customerService.updateAssociateCustomerToProduct(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.currentRecord.cpmstate = this.dataModel.cpmstate;
          this.closeForm();
        }
      });
      this.inProgress = false;
    }
  }
  getproductcode(value: number) {
    if (this.products && value) {
      return this.products.find(a => a.prodid === value).prodcode;
    }
  }
  getdbinstances(value: number) {
    if (this.dbinstances && value) {
      return this.dbinstances.find(a => a.dbiid === value).dbiname;
    }
  }
  getcpmstates(value: number) {
    if (this.cpmstates && value) {
      return this.cpmstates.find(a => a.value === value).text;
    }
  }
  closeForms() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
