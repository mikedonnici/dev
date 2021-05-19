import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {map, filter} from 'rxjs/operators';

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
        if (Math.floor(Math.random() * 20) === 3) {
          observer.error(new Error('1 in 20 chance of an error'));
        }
        if (count === 20) {
          observer.complete();
        }
        observer.next(count);
      }, 500);
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

    // Add an operator to transform the observable data on the fly
    const newObservable = customIntervalObservable.pipe(
      filter( (data: any) => {
          return data % 3 !== 0;
      }),
      map((data: any) => {
        return `Transformed ${data}`;
      })
    );
    newObservable.subscribe(n => {
      console.log(n);
    });
  }

  ngOnDestroy(): void {
    console.log(`unsubscribe()`);
    this.counterSubs.unsubscribe();
  }
}
