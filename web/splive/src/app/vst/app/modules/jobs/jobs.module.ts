import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { JobsRoutingModule } from './jobs-routing.module';
import { JobListComponent } from './job-list/job-list.component';
import { JobSearchComponent } from './job-list/job-search/job-search.component';
import { ListViewComponent } from './job-list/list-view/list-view.component';
import { MatFormFieldModule, MatInputModule, MatSelectModule, MatExpansionModule, MatTableModule } from '@angular/material';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { JobDetailsComponent } from './job-details/job-details.component';

@NgModule({
  imports: [
    CommonModule,
    JobsRoutingModule,
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
  declarations: [JobListComponent, JobSearchComponent, ListViewComponent, JobDetailsComponent]
})
export class JobsModule { }
