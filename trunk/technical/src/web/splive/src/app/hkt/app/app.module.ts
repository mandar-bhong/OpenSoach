import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from '../../hkt/app/app-routing.module';

import { AppComponent } from './app.component';
import { Hkttest1Component } from './hkttest1/hkttest1.component';

import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import { SharedModule } from '../../shared/shared.module';

//content import API
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
@NgModule({
  declarations: [
    AppComponent,
    Hkttest1Component,
   
  ],

  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSlideToggleModule,
    SharedModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
