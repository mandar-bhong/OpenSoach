import { NgModule } from '@angular/core';

import { OperatorRoutingModule } from './operators-routing.module';
import { ProdOperatorsModule } from '../../../../prod-shared/modules/operators/prod-operators.module';

@NgModule({
  imports: [
    OperatorRoutingModule,
    ProdOperatorsModule
  ],
  declarations: []
})
export class OperatorsModule { }
