import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToastrModule } from 'ngx-toastr';

import { AppCommonModule } from '../../shared/app-common.module';
import { LayoutModule } from '../../shared/layouts/layout.module';
import { MaterialModules } from '../../shared/modules/material/material-modules';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CorporateService } from './services/corporate.service';
import { CustomerService } from './services/customer.service';
import { UserService } from './services/user.service';
import { DeviceService } from './services/device.service';
import { DBInstanceService } from './services/db-instance.service';
import { ProductService } from './services/product.service';


@NgModule({
  declarations: [
    AppComponent
  ],

  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSlideToggleModule,
    LayoutModule,
    AppCommonModule.forRoot(),
    FormsModule,
    ReactiveFormsModule,
    ToastrModule.forRoot({
      closeButton: true,
      timeOut: 15000,
      extendedTimeOut: 5000,
      progressBar: true,
      positionClass: 'toast-bottom-right',
      preventDuplicates: true,
      tapToDismiss: false
    }),
    MaterialModules.forRoot()
  ],
  providers: [
    CustomerService,
    CorporateService,
    UserService,
    DeviceService,
    ProductService,
    DBInstanceService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
