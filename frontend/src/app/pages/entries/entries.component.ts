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
  styleUrls: ['./entries.component.scss']
})
export class EntriesComponent implements OnInit, OnDestroy {
  name!: string;
  private sub: any;

  data: IndexDataResponse = { Name: "", Entries: 0, Size: "", VectorLength: 0, VectorType: "" }
  entries: IndexEntryListItem[] = []
  expandedEntries: number[] = []

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
    this.getEntries(0, 10)
  }

  async getEntries(from: number, to: number) {
    this.entries.push(...(await this.backend.getIndexEntries(this.name, from, to)).map(item => { return { entry: item, expanded: false } }))
  }

  expandItem(event: any, item: IndexEntryListItem) {
    item.expanded = !item.expanded
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
