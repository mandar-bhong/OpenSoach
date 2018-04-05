import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { Shared1Component } from './shared1/shared1.component';
import { AppContainerComponent } from './layouts/applayout/app-container/app-container.component';
import {ContentComponent} from './layouts/applayout/content/content.component';
import {FooterComponent} from './layouts/applayout/footer/footer.component';
import {SideBarComponent} from './layouts/applayout/side-bar/side-bar.component';
import {TopHeaderComponent} from './layouts/applayout/top-header/top-header.component';
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
