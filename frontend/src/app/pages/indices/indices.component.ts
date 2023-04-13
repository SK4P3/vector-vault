import { Component } from '@angular/core';
import { Index } from 'src/app/models/interfaces';
import { BackendService } from 'src/app/services/backend/backend.service';

@Component({
  selector: 'app-indices',
  templateUrl: './indices.component.html',
  styleUrls: ['./indices.component.scss']
})
export class IndicesComponent {

  indecies: Index[]

  constructor(private backend: BackendService) {
    this.indecies = backend.getIndices()
  }


}
