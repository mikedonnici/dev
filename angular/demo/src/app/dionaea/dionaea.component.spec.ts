import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DionaeaComponent } from './dionaea.component';

describe('DionaeaComponent', () => {
  let component: DionaeaComponent;
  let fixture: ComponentFixture<DionaeaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DionaeaComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DionaeaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
