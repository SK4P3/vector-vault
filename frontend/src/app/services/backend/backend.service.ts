import { Injectable } from '@angular/core';
import { Index, IndexDataResponse } from 'src/app/models/interfaces';

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor() { }

  getIndices(): Index[] {
    return [
      { name: "Mock Data", url: "http://localhost:8080/", entries: 1, size: "100 Kb", replicas: 1 }
    ]
  }

  getIndex(name: string): Index {
    return this.getIndices().filter((index: Index) => { return index.name == name})[0]
  }

  async getIndexData(name: string) {
    let index = this.getIndices().filter((index: Index) => { return index.name == name})[0]
    let response = await (await fetch("/api/stats")).json() as IndexDataResponse
    return response
  }
}
