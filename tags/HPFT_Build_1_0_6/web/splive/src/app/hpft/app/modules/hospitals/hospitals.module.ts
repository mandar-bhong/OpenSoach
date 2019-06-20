import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { HospitalsRoutingModule } from './hospitals-routing.module';
import { HospitalListComponent } from './hospital-list/hospital-list.component';
import { HospitalSearchComponent } from './hospital-list/hospital-search/hospital-search.component';
import { HospitalViewComponent } from './hospital-list/hospital-view/hospital-view.component';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { ProdCommonModule } from '../../../../prod-shared/prod-common.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppCommonModule } from '../../../../shared/app-common.module';

@NgModule({
  imports: [
    CommonModule,
    MaterialModules,
    ProdCommonModule,
    FormsModule,
    ReactiveFormsModule,
    AppCommonModule,
    HospitalsRoutingModule
  ],
  declarations: [
    HospitalListComponent,
    HospitalSearchComponent,
    HospitalViewComponent
  ]
})
export class HospitalsModule { }
