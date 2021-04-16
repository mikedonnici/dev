import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ObsOneComponent } from './obs-one.component';

describe('ObsOneComponent', () => {
  let component: ObsOneComponent;
  let fixture: ComponentFixture<ObsOneComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ObsOneComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ObsOneComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
