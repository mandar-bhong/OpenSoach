import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ServicepointsRoutingModule } from './servicepoints-routing.module';
import { ProdServicepointsModule } from '../../../../prod-shared/modules/servicepoints/prod-servicepoints.module';

@NgModule({
  imports: [
    CommonModule,
    ServicepointsRoutingModule,
    ProdServicepointsModule
  ],
  declarations: []
})
export class ServicepointsModule { }
