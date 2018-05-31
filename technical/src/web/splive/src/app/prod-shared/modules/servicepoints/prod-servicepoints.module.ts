import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ServicepointDetailsComponent } from './servicepoint-details/servicepoint-details.component';
import { ServicepointListComponent } from './servicepoint-list/servicepoint-list.component';
import { ServicepointSearchComponent } from './servicepoint-list/servicepoint-search/servicepoint-search.component';
import { ServicepointListViewComponent } from './servicepoint-list/servicepoint-list-view/servicepoint-list-view.component';
import { MaterialModules } from '../../../shared/modules/material/material-modules';
import { AppCommonModule } from '../../../shared/app-common.module';
import { ProdCommonModule } from '../../prod-common.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

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
