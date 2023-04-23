import { BackendService } from './../../services/backend/backend.service';
import { Component, Input, OnInit } from '@angular/core';
import { IndexStatus } from '../index-status/index-status.component';
import { Index } from 'src/app/models/interfaces';


@Component({
  selector: 'app-index-entry',
  templateUrl: './index-entry.component.html',
  styleUrls: ['./index-entry.component.scss']
})
export class IndexEntryComponent implements OnInit {

  constructor(private backend: BackendService) {

  }
  @Input() index!: Index;
  indexName!: string;
  url!: string;
  entries!: string;
  size!: string;
  replicas!: string;

  status!: IndexStatus

  ngOnInit(): void {
    this.indexName = this.index.name
    this.url = this.index.url
    this.entries = this.index.entries.toString()
    this.size = this.index.size
    this.replicas = this.index.replicas.toString()
    this.status = IndexStatus.Offline
    this.getIndexStatus();
  }

  async getIndexStatus() {
    let health = await this.backend.getHealth(this.index);

    if(health) {
      this.status = IndexStatus.Online
    }
  }

  public get indexStatus(): typeof IndexStatus {
    return IndexStatus;
  }
}
