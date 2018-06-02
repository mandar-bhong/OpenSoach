import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { ProdServicepointsModule } from '../../../../prod-shared/modules/servicepoints/prod-servicepoints.module';
import { ServicepointsRoutingModule } from './servicepoints-routing.module';

@NgModule({
  imports: [
    CommonModule,
    ServicepointsRoutingModule,
    ProdServicepointsModule
  ],
  declarations: []
})
export class ServicepointsModule { }
