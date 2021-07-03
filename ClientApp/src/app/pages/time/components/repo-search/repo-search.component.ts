import { Component, EventEmitter, OnInit, Output, OnDestroy, ViewChild  } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatMenuTrigger } from '@angular/material/menu';

import { Subscription } from 'rxjs';
import { distinctUntilChanged, debounceTime } from 'rxjs/operators';

import { RepoSearchResult } from 'src/app/models/repos';
import { RepoService } from 'src/app/services/repo/repo.service';

@Component({
  selector: 'repo-search',
  templateUrl: './repo-search.component.html',
  styleUrls: ['./repo-search.component.scss']
})
export class RepoSearchComponent implements OnInit, OnDestroy {
  @Output() repoSelected: EventEmitter<any> = new EventEmitter<any>();
  @ViewChild(MatMenuTrigger, {static: true}) matMenuTrigger: MatMenuTrigger;
  private searchSub: Subscription = new Subscription();
  searchFc: FormControl = new FormControl();
  result: RepoSearchResult;

  constructor(private readonly repoService: RepoService) { }

  ngOnInit(): void {
    this.searchSub = this.searchFc.valueChanges.pipe(
      distinctUntilChanged(),
      debounceTime(500)
    ).subscribe({
      next: (value: string) => {
        if (value.length > 4) {
          this.search(value);
        }
      }
    });

    this.matMenuTrigger.menuOpened.subscribe({
      next: () => {
        const searchValue: string = this.searchFc.value || '';
        if (searchValue.length <= 3) {
          this.matMenuTrigger.closeMenu();
        }
      }
    });
  }

  ngOnDestroy(): void {
    this.searchSub.unsubscribe();
    this.matMenuTrigger.closeMenu();
  }

  search(query: string): void {
    this.repoService.searchGitHub(query)
      .subscribe((res: RepoSearchResult) => {
        this.result = res;
        this.matMenuTrigger.openMenu();
      });
  }

}