import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubjectExampleComponent } from './subject-example.component';

describe('SubjectExampleComponent', () => {
  let component: SubjectExampleComponent;
  let fixture: ComponentFixture<SubjectExampleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SubjectExampleComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SubjectExampleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
