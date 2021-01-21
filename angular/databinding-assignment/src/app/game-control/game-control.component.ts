import {Component, EventEmitter, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-game-control',
  templateUrl: './game-control.component.html',
  styleUrls: ['./game-control.component.css']
})
export class GameControlComponent implements OnInit {

  timer = 0;
  count = 0;
  @Output() counter = new EventEmitter<number>();

  constructor() { }

  ngOnInit(): void {
  }

  start(): void {
    this.timer = setInterval(() => {
      this.counter.emit(this.count++);
    }, 1000);
  }

  stop(): void {
    clearInterval(this.timer);
    this.count = 0;
  }

}
