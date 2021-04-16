import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';

@Component({
  selector: 'app-obs-two',
  templateUrl: './obs-two.component.html',
  styleUrls: ['./obs-two.component.css']
})
export class ObsTwoComponent implements OnInit, OnDestroy {

  private counterSubs = new Subscription();

  constructor() {
  }

  ngOnInit(): void {

    const customIntervalObservable = new Observable(observer => {
      let count = 0;
      setInterval(() => {
        count++;
        if (Math.floor(Math.random() * 5) === 3) {
          observer.error(new Error('1 in 5 chance of an error'));
        }
        if (count === 10) {
          observer.complete();
        }
        observer.next(count);
      }, 1000);
    });

    this.counterSubs = customIntervalObservable.subscribe(
      n => {
        console.log(`Obs2 count = ${n}`);
      },
      e => {
        alert(`Error: ${e.message}`);
      },
      () => {
        console.log('Complete');
      }
    );
  }

  ngOnDestroy(): void {
    console.log(`unsubscribe()`);
    this.counterSubs.unsubscribe();
  }
}
