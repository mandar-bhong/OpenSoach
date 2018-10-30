import { TestBed, inject } from '@angular/core/testing';

import { DetailsComponent } from './details.component';

describe('a details component', () => {
	let component: DetailsComponent;

	// register all needed dependencies
	beforeEach(() => {
		TestBed.configureTestingModule({
			providers: [
				DetailsComponent
			]
		});
	});

	// instantiation through framework injection
	beforeEach(inject([DetailsComponent], (DetailsComponent) => {
		component = DetailsComponent;
	}));

	it('should have an instance', () => {
		expect(component).toBeDefined();
	});
});