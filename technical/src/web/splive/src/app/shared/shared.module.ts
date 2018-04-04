import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { Shared1Component } from './shared1/shared1.component';
import { AppContainerComponent } from './applayout/app-container/app-container.component';
import {ContentComponent} from './applayout/content/content.component';
import {FooterComponent} from './applayout/footer/footer.component';
import {SideBarComponent} from './applayout/side-bar/side-bar.component';
import {TopHeaderComponent} from './applayout/top-header/top-header.component';
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
  ]
})
export class SharedModule { }
