import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { IndexComponent } from './index/index.component';
import { OneComponent } from './one/one.component';
import { TwoComponent } from './two/two.component';
import { AComponent } from './one/a/a.component';
import { DogComponent } from './one/a/dog/dog.component';
import { BComponent } from './one/b/b.component';

@NgModule({
  declarations: [
    AppComponent,
    IndexComponent,
    OneComponent,
    TwoComponent,
    AComponent,
    DogComponent,
    BComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
