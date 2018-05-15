import { Component, OnInit, Input } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';

@Component({
  selector: 'app-chart-configure-basic',
  templateUrl: './chart-configure-basic.component.html',
  styleUrls: ['./chart-configure-basic.component.css']
})
export class ChartConfigureBasicComponent implements OnInit {


  editableForm: FormGroup;
  constructor(private dynamicContextService: DynamicContextService) {
  }

  ngOnInit() {
    this.createControls();
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      configNameControl: new FormControl('', [Validators.required]),
      categoryControl: new FormControl('', [Validators.required])
    });
  }

  nextClick() {
    if (this.editableForm.invalid) { return; }
    this.dynamicContextService.changeComponent('app-chart-configure-time');
  }

}
