import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs/Subscription';

import { StepperService, StepperState } from '../stepper.service';

@Component({
  selector: 'app-stepper-list',
  templateUrl: './stepper-list.component.html',
  styleUrls: ['./stepper-list.component.css']
})
export class StepperListComponent implements OnInit, OnDestroy {

  count = 0;
  done = -1;
  stepperSubscription: Subscription;
  @Input() stepperList;

  constructor(private stepperService: StepperService) {
    this.stepperSubscription = this.stepperService.stepperCountSubject.subscribe((value) => {
      this.doneStepper(value);
    });
  }
  ngOnInit() {
    this.changeState();
  }
  doneStepper(currentState: StepperState) {
    this.count = currentState.stepindex;
    if (this.count > this.done || (this.count < this.done && currentState.completed)) {
      this.done++;
    }
    this.changeState();
  }
  changeState() {
    if (this.stepperList) {
      for (let i = 0; i < this.stepperList.length; i++) {
        const stepperBlock = this.stepperList[i];
        stepperBlock.color = this.getCssClass(i);
      }
    }
  }
  getCssClass(index) {
    if (index === this.count && index !== this.done) {
      return 'processing';
    } else if (index <= this.done) {
      return 'done';
    }

    return '';
  }

  ngOnDestroy() {
    this.stepperSubscription.unsubscribe();
  }
}
