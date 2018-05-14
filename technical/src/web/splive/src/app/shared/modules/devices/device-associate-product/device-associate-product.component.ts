import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatTableDataSource } from '@angular/material';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../shared/views/edit-record-base';
import {
  DeviceAssociateProductListItemResponse, DeviceAssociateProductRequest
} from '../../../../spl/app/models/api/device-models';
import { DeviceAssociateProductModel } from '../../../../spl/app/models/ui/device-models';
import { ProductListItemResponse } from '../../../../spl/app/models/api/product-models';
import { DeviceService } from '../../../../spl/app/services/device.service';
import { ProductService } from '../../../../spl/app/services/product.service';
import { CustomerService } from '../../../../spl/app/services/customer.service';
import { CustomerListItemResponse } from '../../../../spl/app/models/api/customer-models';
import { CustomerAssociateProductListItemResponse } from '../../../../spl/app/models/api/customer-models';
import { TranslatePipe } from '../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../shared/services/notification/app-notification.service';
@Component({
  selector: 'app-device-associate-product',
  templateUrl: './device-associate-product.component.html',
  styleUrls: ['./device-associate-product.component.css']
})
export class DeviceAssociateProductComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new DeviceAssociateProductModel();
  products: ProductListItemResponse[];
  cutomers: CustomerListItemResponse[];
  showForm = false;
  routeSubscription: Subscription;
  dcpmlist: DeviceAssociateProductListItemResponse[];
  cpmlist: CustomerAssociateProductListItemResponse[];

  dataSource: MatTableDataSource<DeviceAssociateProductListItemResponse>;
  displayedColumns = ['custname', 'prodcode'];
  addVisiblity = false;
  constructor(private productService: ProductService,
    private deviceService: DeviceService,
    private customerService: CustomerService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.devid = Number(params['id']);
      this.callbackUrl = params['callbackurl'];
    });
    this.productService.getDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.products = payloadResponse.data;
        this.getDcpmList();
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
      productControl: new FormControl('', [Validators.required]),
      custnameControl: new FormControl('', [Validators.required])
    });
  }

  getDcpmList() {
    this.deviceService.getDeviceProductAssociation({ recid: this.dataModel.devid })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dcpmlist = payloadResponse.data;
          this.dataSource = new MatTableDataSource<DeviceAssociateProductListItemResponse>(this.dcpmlist);
          if (this.dcpmlist.length < this.products.length) {
            this.addVisiblity = true;
          } else {
            this.addVisiblity = false;
          }
        }
      });
  }

  getCpmList() {
    this.customerService.getCustomerProductAssociation({ recid: this.dataModel.custid })
      .subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.cpmlist = payloadResponse.data;
          // TODD: display informative message saying customer is not been associated with any product yet.
        }
      });
  }

  save() {
    if (this.editableForm.invalid) { return; }
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const request = new DeviceAssociateProductRequest();
      this.dataModel.copyToAddRequest(request);
      this.deviceService.associateDeviceToProduct(request).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success(this.translatePipe.transform('SUCCESS_ADD_DEVICE_ASSOCIATE_SAVED'));
          this.getDcpmList();
          this.closeForm();
        }
      });
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
