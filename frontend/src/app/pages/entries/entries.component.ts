import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { IndexDataResponse } from 'src/app/models/interfaces';
import { BackendService } from 'src/app/services/backend/backend.service';

@Component({
  selector: 'app-entries',
  templateUrl: './entries.component.html',
  styleUrls: ['./entries.component.scss']
})
export class EntriesComponent implements OnInit, OnDestroy {
  name!: string;
  private sub: any;

  data: IndexDataResponse = { Name: "", Entries: 0, Size: "", VectorLength: 0, VectorType: "" }

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
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
