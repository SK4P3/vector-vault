import { Component, Input } from '@angular/core';
import { IndexStatus } from '../index-status/index-status.component';


@Component({
  selector: 'app-index-entry',
  templateUrl: './index-entry.component.html',
  styleUrls: ['./index-entry.component.scss']
})
export class IndexEntryComponent {

  @Input() indexName!: string;
  @Input() url!: string;

  @Input() entries!: string;
  @Input() size!: string;
  @Input() replicas!: string;

  public get indexStatus(): typeof IndexStatus {
    return IndexStatus;
  }
}
