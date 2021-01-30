import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NepenthesComponent } from './nepenthes.component';

describe('NepenthesComponent', () => {
  let component: NepenthesComponent;
  let fixture: ComponentFixture<NepenthesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NepenthesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NepenthesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
