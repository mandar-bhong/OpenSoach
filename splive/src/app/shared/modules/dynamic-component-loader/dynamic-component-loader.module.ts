import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { DynamicContentComponent } from './dynamic-content.component';
import { DynamicContextService } from './dynamic-context.service';

@NgModule({
    imports: [
        CommonModule],
    declarations: [
        DynamicContentComponent
    ],
    exports: [DynamicContentComponent],
    providers: [DynamicContextService]

})
export class DynamicComponentLoaderModule { }
