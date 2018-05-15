import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';
import { FormGroup } from '@angular/forms/src/model';

/**
 * DynamicContextService, holds the transferring context to the components rendered at runtime.
 */
@Injectable()
export class DynamicContextService {
  triggerSaveDynamicComponent: Subject<any> = new Subject<any>();
  triggerChangeDynamicComponent: Subject<string> = new Subject<string>();
  triggerToggleSaveButton: Subject<boolean> = new Subject<boolean>();
  triggerSaveResult: Subject<boolean> = new Subject<boolean>();
  triggerPassForm: Subject<FormGroup> = new Subject<FormGroup>();

  private dynamicComponentMappings = {};

  constructor() {
  }

  passForm(formGroup: FormGroup) {
    this.triggerPassForm.next(formGroup);
  }

  /**
   * Save event trigger
   */
  save(): void {
    this.triggerSaveDynamicComponent.next(null);
  }

  /**
   * Executes the triggerChangeDynamicComponent, to navigate the changes to dynamic-component
   */
  changeComponent(type: string): void {
    this.triggerChangeDynamicComponent.next(type);
  }

  toggleSaveButton(state: boolean) {
    this.triggerToggleSaveButton.next(state);
  }

  saveResult(result: boolean): void {
    this.triggerSaveResult.next(result);
  }

  addDynamicComponentMaping(typeName: string, component: any) {
    this.dynamicComponentMappings[typeName] = component;
  }

  getDynamicComponentMaping(typeName: string) {
    return this.dynamicComponentMappings[typeName];
  }
}
