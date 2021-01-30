import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DrosophyllumComponent } from './drosophyllum.component';

describe('DrosophyllumComponent', () => {
  let component: DrosophyllumComponent;
  let fixture: ComponentFixture<DrosophyllumComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DrosophyllumComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DrosophyllumComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
