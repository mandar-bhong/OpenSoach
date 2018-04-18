import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {
  myForm: FormGroup;
  userfisname: string;
  userlasname: string;
  usermobilenumber: number;
  usernumbersecond: number;
  usergend: string;
  constructor() {
    this.createControls();
  }

  ngOnInit() {
  }
  createControls(): void {
    this.myForm = new FormGroup({
      firstName: new FormControl('', [Validators.required]),
      lastName: new FormControl('', [Validators.required]),
      mobileNumber: new FormControl('', [Validators.required]),
      userGender: new FormControl('', [Validators.required])
    });
  }
  save() {

  }
}
