import { Component, Input } from '@angular/core';

export enum IndexStatus {
  "Online",
  "Offline"
}


@Component({
  selector: 'app-index-status',
  templateUrl: './index-status.component.html',
  styleUrls: ['./index-status.component.scss']
})
export class IndexStatusComponent {
  @Input() status!: IndexStatus

  getStatusStr(status: IndexStatus): string  {
    return IndexStatus[status]
  }
}
