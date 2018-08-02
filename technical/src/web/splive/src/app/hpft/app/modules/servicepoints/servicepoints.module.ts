import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { ProdServicepointsModule } from '../../../../prod-shared/modules/servicepoints/prod-servicepoints.module';
import { ServicepointsRoutingModule } from './servicepoints-routing.module';
import { WardsModule } from '../wards/wards.module';
import { WordListComponent } from '../wards/word-list/word-list.component';

@NgModule({
  imports: [
    CommonModule,
    ServicepointsRoutingModule,
    WardsModule
  ],
  declarations: []
})
export class ServicepointsModule { }
