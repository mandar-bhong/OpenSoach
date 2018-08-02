import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PatientsRoutingModule } from './patients-routing.module';
import { PatientListComponent } from './patient-list/patient-list.component';
import { PatientAddComponent } from './patient-add/patient-add.component';
import { PatientViewComponent } from './patient-list/patient-view/patient-view.component';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ProdCommonModule } from '../../../../prod-shared/prod-common.module';
import { AppCommonModule } from '../../../../shared/app-common.module';

@NgModule({
  imports: [
    CommonModule,
    PatientsRoutingModule,
    MaterialModules,
    ProdCommonModule,
    FormsModule,
    ReactiveFormsModule,
    AppCommonModule
  ],
  declarations: [PatientListComponent, PatientAddComponent, PatientViewComponent]
})
export class PatientsModule { }
