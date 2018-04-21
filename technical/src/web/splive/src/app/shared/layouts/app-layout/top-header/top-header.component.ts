import { Component, OnInit } from '@angular/core';

import { LoginHandlerService } from '../../../services/login-handler.service';
import { SidebarToggleService } from '../../../services/sidebar-toggle.service';
import { UserInfo } from '../../../models/ui/user-models';
import { CustomerInfo } from '../../../models/ui/customer-models';

@Component({
  selector: 'app-top-header',
  templateUrl: './top-header.component.html',
  styleUrls: ['./top-header.component.css']
})
export class TopHeaderComponent implements OnInit {
  menuFull = true;
  username = 'User';
  customername: string;
  constructor(private loginHandlerService: LoginHandlerService, private sidebarToggleService: SidebarToggleService) {
    this.loginHandlerService.userInfoSubject.subscribe(userInfo => {
      this.setUserName(userInfo);
    });
    this.loginHandlerService.customerInfoSubject.subscribe(customerInfo => {
      this.setCustomerName(customerInfo);
    });
  }
  ngOnInit() {
  }
  logout() {
    this.loginHandlerService.logout();
  }
  toggleChange() {
    this.menuFull = !this.menuFull;
    this.sidebarToggleService.toggleMenu(this.menuFull);
  }

  setUserName(userInfo: UserInfo) {
    if (userInfo.fname && userInfo.lname) {
      this.username = userInfo.fname + ' ' + userInfo.lname;
    } else {
      this.username = userInfo.usrname;
    }
  }

  setCustomerName(customerInfo: CustomerInfo) {
    this.customername = customerInfo.custname + ' - ' + customerInfo.corpname;    
  }
}
