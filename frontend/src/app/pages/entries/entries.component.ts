import { animate, style, transition, trigger } from '@angular/animations';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { IndexDataResponse, IndexEntry } from 'src/app/models/interfaces';
import { BackendService } from 'src/app/services/backend/backend.service';

interface IndexEntryListItem {
  entry: IndexEntry
  expanded: boolean
}

@Component({
  selector: 'app-entries',
  templateUrl: './entries.component.html',
  styleUrls: ['./entries.component.scss'],
  animations: [
    trigger(
      'inOutAnimation',
      [
        transition(
          ':enter',
          [
            style({ opacity: 0 }),
            animate('200ms ease-out',
              style({ opacity: 1 }))
          ]
        ),
        transition(
          ':leave',
          [
            style({ opacity: 1 }),
            animate('200ms ease-in',
              style({ opacity: 0 }))
          ]
        )
      ]
    )
  ]
})

export class EntriesComponent implements OnInit, OnDestroy {
  name!: string;
  private sub: any;

  data: IndexDataResponse = { Name: "", Entries: 0, Size: "", VectorDimension: 0, VectorType: "" }
  entries: IndexEntryListItem[] = []
  expandedEntries: number[] = []

  entryModalVisible: boolean = false

  constructor(private route: ActivatedRoute, private backend: BackendService) {

  }

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.name = params['name'];
      this.initData();
    });
  }

  async initData() {
    this.data = await this.backend.getIndexData(this.name);
    this.getEntries(995, 1000)
  }

  async getEntries(from: number, to: number) {
    this.entries.push(...(await this.backend.getIndexEntries(this.name, from, to)).map(item => { return { entry: item, expanded: false } }))
  }

  expandItem(event: any, item: IndexEntryListItem) {
    item.expanded = !item.expanded
  }

  showAddEntryModal() {
    this.entryModalVisible = !this.entryModalVisible
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
