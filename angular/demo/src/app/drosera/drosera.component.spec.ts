import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DroseraComponent } from './drosera.component';

describe('DroseraComponent', () => {
  let component: DroseraComponent;
  let fixture: ComponentFixture<DroseraComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DroseraComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DroseraComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
