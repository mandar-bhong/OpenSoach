import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormBuilder, FormGroup, ReactiveFormsModule, NgForm } from '@angular/forms';
@Component({
  selector: 'app-user-add',
  templateUrl: './user-add.component.html',
  styleUrls: ['./user-add.component.css']
})
export class UserAddComponent implements OnInit {
  myForm: FormGroup;
  username: string;
  usercategery: string;
  userrolee: string;
  userstate: string;

  constructor() { }

  ngOnInit() {
    this.createControls();
  }
  createControls(): void {
    this.myForm = new FormGroup({
      userCategory: new FormControl('', [Validators.required]),
      emailControl: new FormControl('', [Validators.required]),
      userRole: new FormControl('', [Validators.required]),
      userState: new FormControl('', [Validators.required])
    });
  }
  save() {
     }
}
