import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginStatusService } from '../../../services/login-status.service';

@Component({
  selector: 'app-top-header',
  templateUrl: './top-header.component.html',
  styleUrls: ['./top-header.component.css']
})
export class TopHeaderComponent implements OnInit {

  constructor(private loginStatusService: LoginStatusService) { }

  ngOnInit() {
  }

  logout() {
    this.loginStatusService.logout();
  }

}
