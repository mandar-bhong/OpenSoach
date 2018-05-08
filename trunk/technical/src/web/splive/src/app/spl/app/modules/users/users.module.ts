import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UsersRoutingModule } from './users-routing.module';
import { UserListComponent } from './user-list/user-list.component';
import { UserAddComponent } from './user-add/user-add.component';
import { UserDetailsComponent } from './user-details/user-details.component';

import { MatTableModule } from '@angular/material/table';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { UserSearchComponent } from './user-list/user-search/user-search.component';
import { UserViewComponent } from './user-list/user-view/user-view.component';
import { AppCommonModule } from '../../../../shared/app-common.module';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';

@NgModule({
  imports: [
    CommonModule,
    UsersRoutingModule,
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
    UserDetailsComponent,
    UserSearchComponent,
    UserViewComponent]
})
export class UsersModule { }
