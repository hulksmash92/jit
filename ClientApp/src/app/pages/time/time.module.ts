import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatChipsModule } from '@angular/material/chips';

import { TimeComponent } from './time.component';
import { FormComponent } from './form/form.component';
import { TableComponent } from './table/table.component';
import { TimeRoutingModule } from './time-routing.module';
import { MaterialModule } from 'src/app/modules/material/material.module';
import { RepoSearchComponent } from './components/repo-search/repo-search.component';
import { RepoItemSearchComponent } from './components/repo-item-search/repo-item-search.component';
import { TimeFormTagsComponent } from './components/time-form-tags/time-form-tags.component';

@NgModule({
  declarations: [
    TimeComponent,
    FormComponent,
    TableComponent,
    RepoSearchComponent,
    RepoItemSearchComponent,
    TimeFormTagsComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatChipsModule,
    MaterialModule,
    TimeRoutingModule
  ]
})
export class TimeModule { }
