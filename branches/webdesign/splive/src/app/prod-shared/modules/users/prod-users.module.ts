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
import { UserAddComponent } from './user-add/user-add.component';
import { UserDetailsComponent } from './user-details/user-details.component';
import { UserListComponent } from './user-list/user-list.component';
import { UserSearchComponent } from './user-list/user-search/user-search.component';
import { UserViewComponent } from './user-list/user-view/user-view.component';


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
    UserListComponent,
    UserAddComponent,
    UserSearchComponent,
    UserViewComponent,
    UserDetailsComponent
  ]
})
export class ProdUsersModule { }
