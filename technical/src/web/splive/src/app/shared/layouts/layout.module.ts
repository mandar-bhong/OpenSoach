import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AppContainerComponent } from './app-layout/app-container/app-container.component';
import {ContentComponent} from './app-layout/content/content.component';
import {FooterComponent} from './app-layout/footer/footer.component';
import {SideBarComponent} from './app-layout/side-bar/side-bar.component';
import {TopHeaderComponent} from './app-layout/top-header/top-header.component';
import { RouterModule, Routes } from '@angular/router';
import { AuthLayoutComponent } from './auth-layout/auth-layout.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule
  ],
  declarations: [
    AppContainerComponent,
    ContentComponent,
    FooterComponent,
    SideBarComponent,
    TopHeaderComponent,
    AuthLayoutComponent
  ],
})
export class LayoutModule { }
