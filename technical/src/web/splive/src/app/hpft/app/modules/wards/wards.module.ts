import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { ProdCommonModule } from '../../../../prod-shared/prod-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { WardListViewComponent } from './word-list/ward-list-view/ward-list-view.component';
import { WardSearchComponent } from './word-list/ward-search/ward-search.component';
import { WordListComponent } from './word-list/word-list.component';
import { WardDeviceAssociateComponent } from './ward-device-associate/ward-device-associate.component';

@NgModule({
  imports: [
    CommonModule,
    MaterialModules,
    ProdCommonModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  declarations: [WordListComponent, WardListViewComponent, WardSearchComponent, WardDeviceAssociateComponent],
  entryComponents: [
    WardDeviceAssociateComponent
  ]
})
export class WardsModule { }
