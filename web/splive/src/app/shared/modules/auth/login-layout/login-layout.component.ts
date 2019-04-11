import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'hkt-login-layout',
  templateUrl: './login-layout.component.html',
  styleUrls: ['./login-layout.component.css']
})
export class LoginLayoutComponent implements OnInit {
  flipped = false;
  constructor() { }

  ngOnInit() {
  }
  flipIt() {
    this.flipped = !this.flipped;
  }
}
