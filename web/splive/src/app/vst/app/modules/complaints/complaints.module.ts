import { NgModule } from '@angular/core';

import { ComplaintsRoutingModule } from './complaints-routing.module';
import { ComplaintDetailsComponent } from '../complaints/complaint-detalis/complaint-details.component';
import { ComplaintListComponent } from '../complaints/complaint-list/complaint-list.component';
import { ComplaintViewComponent } from '../complaints/complaint-list/complaint-view/complaint-view.component';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule, MatInputModule, MatSelectModule, MatExpansionModule, MatTableModule } from '@angular/material';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { ComplaintSearchComponent } from './complaint-list/complaint-search/complaint-search.component';


@NgModule({
  imports: [
    ComplaintsRoutingModule,
    CommonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    FormsModule,
    ReactiveFormsModule,
    MatExpansionModule,
    MatTableModule,
    AppCommonModule,
    MaterialModules
  ],
  declarations: [
    ComplaintDetailsComponent,
    ComplaintListComponent,
    ComplaintSearchComponent,
    ComplaintViewComponent
  ]
})
export class ComplaintsModule { }
