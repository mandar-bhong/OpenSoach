import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { StepperBlockComponent } from './stepper-block/stepper-block.component';
import { StepperListComponent } from './stepper-list/stepper-list.component';
import { StepperService } from './stepper.service';
import { MaterialModules } from '../material/material-modules';

@NgModule({
    imports: [
        CommonModule],
    declarations: [
        StepperBlockComponent,
        StepperListComponent,
        
    ],
    exports: [StepperListComponent],
    providers: [StepperService]
})
export class StepperModule { }
