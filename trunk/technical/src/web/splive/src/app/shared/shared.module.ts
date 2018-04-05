import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { Shared1Component } from './shared1/shared1.component';
import { AppContainerComponent } from './layouts/app-layout/app-container/app-container.component';
import {ContentComponent} from './layouts/app-layout/content/content.component';
import {FooterComponent} from './layouts/app-layout/footer/footer.component';
import {SideBarComponent} from './layouts/app-layout/side-bar/side-bar.component';
import {TopHeaderComponent} from './layouts/app-layout/top-header/top-header.component';
import { RouterModule, Routes } from '@angular/router';

@NgModule({
  imports: [
    CommonModule,
    RouterModule
  ],
  declarations: [
    Shared1Component,
    AppContainerComponent,
    ContentComponent,
    FooterComponent,
    SideBarComponent,
    TopHeaderComponent,
  ],
  exports: [
    AppContainerComponent
  ]
})
export class SharedModule { }
