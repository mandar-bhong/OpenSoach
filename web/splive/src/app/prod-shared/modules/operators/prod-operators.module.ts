import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatTableModule } from '@angular/material/table';
import { AppCommonModule } from '../../../shared/app-common.module';
import { MaterialModules } from '../../../shared/modules/material/material-modules';
import { OperatorAddComponent } from './operator-add/operator-add.component';
import { OperatorListComponent } from './operator-list/operator-list.component';
import { OperatorSearchComponent } from './operator-list/operator-search/operator-search.component';
import { OperatorViewComponent } from './operator-list/operator-view/operator-view.component';
import { OperatorAssociateComponent } from './operator-associate/operator-associate.component';

@NgModule({
  imports: [
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
    OperatorAddComponent,
    OperatorListComponent,
    OperatorSearchComponent,
    OperatorViewComponent,
    OperatorAssociateComponent
  ]
})
export class ProdOperatorsModule { }
