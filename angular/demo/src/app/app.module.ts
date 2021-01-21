import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DionaeaComponent } from './dionaea/dionaea.component';
import { SarraceniaComponent } from './sarracenia/sarracenia.component';
import { NepenthesComponent } from './nepenthes/nepenthes.component';
import { DroseraComponent } from './drosera/drosera.component';
import { DrosophyllumComponent } from './drosophyllum/drosophyllum.component';
import { HomeComponent } from './home/home.component';
import {FormsModule} from "@angular/forms";

@NgModule({
  declarations: [
    AppComponent,
    DionaeaComponent,
    SarraceniaComponent,
    NepenthesComponent,
    DroseraComponent,
    DrosophyllumComponent,
    HomeComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
