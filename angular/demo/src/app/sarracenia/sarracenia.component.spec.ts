import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SarraceniaComponent } from './sarracenia.component';

describe('SarraceniaComponent', () => {
  let component: SarraceniaComponent;
  let fixture: ComponentFixture<SarraceniaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SarraceniaComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SarraceniaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
