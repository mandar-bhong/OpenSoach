import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export class StepperState {
  completed: true;
  stepindex: number;
}

@Injectable()
export class StepperService {
  stepperCountSubject: Subject<StepperState> = new Subject<StepperState>();

  constructor() { }
  stepperCount(stepperState: StepperState) {
    this.stepperCountSubject.next(stepperState);
  }
}
