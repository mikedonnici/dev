import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ObsTwoComponent } from './obs-two.component';

describe('ObsTwoComponent', () => {
  let component: ObsTwoComponent;
  let fixture: ComponentFixture<ObsTwoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ObsTwoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ObsTwoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
