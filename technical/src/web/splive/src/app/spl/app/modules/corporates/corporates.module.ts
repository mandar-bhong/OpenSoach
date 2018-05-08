import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CorporatesRoutingModule } from './corporates-routing.module';
import { CorporateListComponent } from './corporate-list/corporate-list.component';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { CorporateAddComponent } from './corporate-add/corporate-add.component';
import { AppCommonModule } from '../../../../shared/app-common.module';
import { CorporateSearchComponent } from './corporate-list/corporate-search/corporate-search.component';
import { CorporateListViewComponent } from './corporate-list/corporate-list-view/corporate-list-view.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
@NgModule({
  imports: [
    CommonModule,
    CorporatesRoutingModule,
    MaterialModules,
    AppCommonModule,
    FormsModule,
    ReactiveFormsModule

  ],
  declarations: [
    CorporateListComponent,
    CorporateAddComponent,
    CorporateSearchComponent,
    CorporateListViewComponent
  ]
})
export class CorporatesModule { }
