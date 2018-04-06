import { Component, OnInit } from '@angular/core';
import { AppDataStoreService } from '../../../services/app-data-store/app-data-store-service';

@Component({
  selector: 'hkt-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  hide = true;
  username: string;
  password: string;
  constructor(private appDataStoreService: AppDataStoreService) { }
  // constructor() { }
  ngOnInit() {
    this.username = 'admin@servicepoint.live';
    this.password = 'admin';
  }

  login() {
    this.appDataStoreService.getDataStore('AUTH_TOKEN').setObject<string>('AUTH_TOKEN', 'Some token');
    console.log(this.appDataStoreService.getDataStore('AUTH_TOKEN').getObject<any>('AUTH_TOKEN'));
  }

}
