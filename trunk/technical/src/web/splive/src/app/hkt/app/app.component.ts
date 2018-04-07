import { Component, OnInit } from '@angular/core';
import {LoginStatusService} from '../../shared/services/login-status.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'HKT';

  constructor(private loginStatusService: LoginStatusService) {}

  ngOnInit() {

console.log('app init');
this.loginStatusService.init();

  }
}
