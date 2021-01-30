import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  message = '';
  colours = ['red', 'green', 'blue', 'orange', 'grey'];
  colourCSS: {[key: string]: string} = {
    red: 'alert-danger',
    green: 'alert-success',
    blue: 'alert-info',
    orange: 'alert-warning',
    grey: 'alert-dark'
  };

  constructor() {

  }

  ngOnInit(): void {
  }

  clearMessage(): void {
    this.message = '';
  }

  contains(substr: string): boolean {
    if (this.message.includes(substr)) {
      return true;
    }
    return false;
  }

}
