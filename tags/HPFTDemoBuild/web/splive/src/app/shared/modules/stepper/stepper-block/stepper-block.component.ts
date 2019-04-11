import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-stepper-block',
  templateUrl: './stepper-block.component.html',
  styleUrls: ['./stepper-block.component.css']
})
export class StepperBlockComponent implements OnInit {
  @Input()
  stepperBlock;
  constructor() { }

  ngOnInit() {
  }
}
