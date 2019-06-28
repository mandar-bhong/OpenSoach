import { Component, OnInit } from '@angular/core';
import { MatTooltipModule } from '@angular/material/tooltip';
@Component({
  selector: 'app-medical-container',
  templateUrl: './medical-container.component.html',
  styleUrls: ['./medical-container.component.css']
})
export class MedicalContainerComponent implements OnInit {
  selectedView: string;
  constructor() {
    this.selectedView = 'details';
   }
  ngOnInit() {
   
  }
  selectedViewClick(value: string) {
    this.selectedView = value;
  }
}
