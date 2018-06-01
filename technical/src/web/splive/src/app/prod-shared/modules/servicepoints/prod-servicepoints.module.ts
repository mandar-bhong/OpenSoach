import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppCommonModule } from '../../../shared/app-common.module';
import { MaterialModules } from '../../../shared/modules/material/material-modules';
import { ProdCommonModule } from '../../prod-common.module';
import { ServicepointDetailsComponent } from './servicepoint-details/servicepoint-details.component';
import { ServicepointListViewComponent } from './servicepoint-list/servicepoint-list-view/servicepoint-list-view.component';
import { ServicepointListComponent } from './servicepoint-list/servicepoint-list.component';
import { ServicepointSearchComponent } from './servicepoint-list/servicepoint-search/servicepoint-search.component';

@NgModule({
  imports: [
    CommonModule,
    AppCommonModule,
    MaterialModules,
    ProdCommonModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  declarations: [
    ServicepointDetailsComponent,
    ServicepointListComponent,
    ServicepointListViewComponent,
    ServicepointSearchComponent,
  ]
})
export class ProdServicepointsModule { }
