import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'hkt-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css']
})
export class ForgotPasswordComponent implements OnInit {
  flipped = false;
  username;
  forgotpasswordform: FormGroup;
  otpfield = true;
  newpassword;
  confirmpassword;
  // passwordfield = true;
  userid = false;
  constructor(private router: Router) { }

  ngOnInit() {
    this.createControls();
  }
  createControls(): void {
    this.forgotpasswordform = new FormGroup({
      emailControl: new FormControl('', [Validators.required])
    });
  }
  flipIt() {
    // this.flipped = !this.flipped;
    // window.location.reload();
  }
  otpsend() {
    this.otpfield = false;
  }
  changepassword() {
    // this.otpfield = true;
    // this.userid = true;
    // this.passwordfield = false;

  }

}
