import {Component, OnDestroy, OnInit} from '@angular/core';
import {SubjectService} from '../subject.service';

@Component({
  selector: 'app-subject-example',
  templateUrl: './subject-example.component.html',
  styleUrls: ['./subject-example.component.css']
})
export class SubjectExampleComponent implements OnInit, OnDestroy {

  subjectHasBeenActivated = false;

  constructor(private subjectService: SubjectService) { }

  ngOnInit(): void {
    // This could be in another component, of course!
    this.subjectService.subjectActivated.subscribe(activated => {
      this.subjectHasBeenActivated = activated;
    });
  }

  ngOnDestroy(): void {
    this.subjectService.subjectActivated.unsubscribe();
  }

  toggleSubject(): void {
    this.subjectService.subjectActivated.next(!this.subjectHasBeenActivated);
  }
}
