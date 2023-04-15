import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { SearchComponent } from './pages/search/search.component';
import { IndicesComponent } from './pages/indices/indices.component';
import { EntriesComponent } from './pages/entries/entries.component';
import { IndexEntryComponent } from './components/index-entry/index-entry.component';
import { IndexStatusComponent } from './components/index-status/index-status.component';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    NavbarComponent,
    SearchComponent,
    IndicesComponent,
    EntriesComponent,
    IndexEntryComponent,
    IndexStatusComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
