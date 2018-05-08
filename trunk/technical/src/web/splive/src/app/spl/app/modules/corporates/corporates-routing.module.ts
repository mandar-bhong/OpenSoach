import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CorporateListComponent } from './corporate-list/corporate-list.component';
import { CorporateAddComponent } from './corporate-add/corporate-add.component';

const routes: Routes = [
  {
    path: '',
    component: CorporateListComponent
  },
  {
    path: 'add',
    component: CorporateAddComponent
  }
];
@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CorporatesRoutingModule { }
