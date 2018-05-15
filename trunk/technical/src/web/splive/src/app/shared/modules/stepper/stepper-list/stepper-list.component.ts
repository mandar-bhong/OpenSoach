import { Component, OnInit, Input } from '@angular/core';
import { StepperService } from '../stepper.service';

@Component({
  selector: 'app-stepper-list',
  templateUrl: './stepper-list.component.html',
  styleUrls: ['./stepper-list.component.css']
})
export class StepperListComponent implements OnInit {

  count = 0;
  isDone = 0;
  @Input() stepperList;

  constructor(private stepperService: StepperService) {
    this.stepperService.stepperCountSubject.subscribe((value) => {
      this.count = value;
      this.doneStepper();
    });
  }
 ngOnInit() {
    this.doneStepper();
    console.log('stepperList', this.stepperList);
  }
  doneStepper() {
 if (this.count > this.isDone) {
      this.isDone++;
    }
    this.changeState();
  }
  changeState() {
    // loop into stepper object
    if (this.stepperList) {
      for (let i = 0; i < this.stepperList.length; i++) {
        const stepperBlock = this.stepperList[i];
        stepperBlock.color = this.getCssClass(i);
      }
    }
  }
  getCssClass(index) {
    if (this.count > index) {
      return 'done';
    } else if (this.count === index) {
      return 'processing';
    } else if (this.count < index) {
      if (this.isDone > index) {
        if (this.isDone === index) {
          return '';
        }
        return 'done';
      }
    } else if (index - 1 > this.isDone) {
      return '';
    }
  }
}
