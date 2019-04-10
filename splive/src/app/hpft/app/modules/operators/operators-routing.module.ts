import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { OperatorListComponent } from '../../../../prod-shared/modules/operators/operator-list/operator-list.component';
import { OperatorAddComponent } from '../../../../prod-shared/modules/operators/operator-add/operator-add.component';
import { OperatorAssociateComponent } from '../../../../prod-shared/modules/operators/operator-associate/operator-associate.component';

const routes: Routes = [
  {
    path: '',
    component: OperatorListComponent
  },
  {
    path: 'add',
    component: OperatorAddComponent
  },
  {
    path: 'detail',
    component: OperatorAddComponent
  },
  {
    path: 'associate',
    component: OperatorAssociateComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class OperatorRoutingModule { }
