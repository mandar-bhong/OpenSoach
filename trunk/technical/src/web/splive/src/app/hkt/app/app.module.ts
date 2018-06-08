import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DragulaModule } from 'ng2-dragula';
import { ToastrModule } from 'ngx-toastr';

import { ProdCommonModule } from '../../prod-shared/prod-common.module';
import { AppCommonModule } from '../../shared/app-common.module';
import { LayoutModule } from '../../shared/layouts/layout.module';
import { MaterialModules } from '../../shared/modules/material/material-modules';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ChartConfigureService } from './services/chart-configure.service';
import { ProdComplaintService } from './services/complaint/prod-complaint.service';

@NgModule({
  declarations: [
    AppComponent,
  ],

  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSlideToggleModule,
    LayoutModule,
    AppCommonModule.forRoot(),
    ProdCommonModule.forRoot(),
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
    MaterialModules.forRoot(),
    DragulaModule
  ],
  providers: [ChartConfigureService, ProdComplaintService],
  bootstrap: [AppComponent]
})
export class AppModule { }
