import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import { AppRoutingModule } from './app-routing.module';

import {AppComponent} from './app.component';
import { HomeComponent } from './home/home.component';
import { ObsOneComponent } from './obs-one/obs-one.component';
import { ObsTwoComponent } from './obs-two/obs-two.component';
import { SubjectExampleComponent } from './subject-example/subject-example.component';


const appRoutes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'obs1', component: ObsOneComponent},
  {path: 'obs2', component: ObsTwoComponent},
  {path: 'subject', component: SubjectExampleComponent},
];

@NgModule({
  declarations: [
    AppComponent,
    ObsOneComponent,
    ObsTwoComponent,
    HomeComponent,
    SubjectExampleComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes),
    AppRoutingModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
