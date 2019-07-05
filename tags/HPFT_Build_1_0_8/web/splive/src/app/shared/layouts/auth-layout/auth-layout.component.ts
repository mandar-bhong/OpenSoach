import { Component, OnInit } from '@angular/core';

import { Input, OnDestroy, Inject, ViewEncapsulation } from '@angular/core';
import { AppSpecificDataProvider } from '../../app-specific-data-provider';

@Component({
  selector: 'hkt-auth-layout',
  templateUrl: './auth-layout.component.html',
  styleUrls: ['./auth-layout.component.css']
})
export class AuthLayoutComponent implements OnInit {

  logoprefix: string;
  constructor() {
    this.logoprefix = AppSpecificDataProvider.logoprefix;
   }

  ngOnInit() { }
}
