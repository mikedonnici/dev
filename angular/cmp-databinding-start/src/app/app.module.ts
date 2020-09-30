import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';


import { AppComponent } from './app.component';
import { ServerManagerComponent } from './server-manager/server-manager.component';
import { ServerItemComponent } from './server-manager/server-item/server-item.component';
import { ParentComponent } from './parent/parent.component';
import { ChildComponent } from './parent/child/child.component';

@NgModule({
  declarations: [
    AppComponent,
    ServerManagerComponent,
    ServerItemComponent,
    ParentComponent,
    ChildComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
