import {Component, OnDestroy, OnInit} from '@angular/core';
import {interval, Subscription} from 'rxjs';

@Component({
  selector: 'app-obs-one',
  templateUrl: './obs-one.component.html',
  styleUrls: ['./obs-one.component.css']
})
export class ObsOneComponent implements OnInit, OnDestroy {

  private counterSubs = new Subscription();

  constructor() { }

  ngOnInit(): void {
    this.counterSubs = interval(1000).subscribe(n => {
      console.log(`Obs1 count = ${n}`);
    });
  }

  ngOnDestroy(): void {
    console.log('Unsubscribe()');
    this.counterSubs.unsubscribe();
  }
}
