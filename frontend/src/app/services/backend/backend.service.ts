import { Injectable } from '@angular/core';
import { AddEntryRequest, AddEntryResponse, GetEmbeddingRequest, GetEmbeddingResponse, Index, IndexDataResponse, IndexEntry } from 'src/app/models/interfaces';

// Save question and answers as vectors
// Compare question to already saved question and return an already generated response if similarity is > phi

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor() { }

  getIndices(): Index[] {
    return [
      { name: "Mock", url: "http://localhost:8080/", entries: 1, size: "100 Kb", replicas: 1 },
      { name: "Radstadt", url: "http://localhost:8081/", entries: 1, size: "100 Kb", replicas: 1 }
    ]
  }

  getIndex(name: string): Index {
    return this.getIndices().filter((index: Index) => { return index.name == name })[0]
  }

  async indexGetRequest(index: Index, endpoint: string, params: URLSearchParams | undefined = undefined) {
    let url = `/api/${index.name.toLocaleLowerCase()}/${endpoint}`
    if (params != undefined) {
      url += '?' + params
    }
    let response = await (await fetch(url)).json()
    return response
  }

  async indexPostRequest(index: Index, endpoint: string, body: any) {
    let url = `/api/${index.name.toLocaleLowerCase()}/${endpoint}`
    console.log(url);

    let response =  await (await fetch(url, { method: "POST", body: JSON.stringify(body) })).json()
    return response
  }

  async getHealth(index: Index): Promise<boolean> {
    try {
      let response = await this.indexGetRequest(index, "health")
      return response == "Im Ok! Thanks for asking!"
    } catch {
      return false
    }
  }

  async getEmbeddingAda2(index: Index, content: string): Promise<GetEmbeddingResponse> {
    let reqBody: GetEmbeddingRequest = {
      Content: content
    }
    let response = await this.indexPostRequest(index, "embedding/ada", reqBody) as GetEmbeddingResponse
    return response
  }

  async getIndexData(index: Index): Promise<IndexDataResponse> {
    let response = await this.indexGetRequest(index, "stats") as IndexDataResponse
    return response
  }

  async getIndexEntries(index: Index, from: number, to: number): Promise<IndexEntry[]> {
    let params = new URLSearchParams({
      "from": from.toString(),
      "to": to.toString(),
    })
    let response = await this.indexGetRequest(index, "entry/list", params) as IndexEntry[]
    return response
  }

  async addEntry(index: Index, entry: IndexEntry): Promise<AddEntryResponse> {
    // let index = this.getIndices().filter((index: Index) => { return index.name == name})[0]
    let reqBody: AddEntryRequest = { Entry: entry }
    let response = await this.indexPostRequest(index, "entry", reqBody) as AddEntryResponse
    return response
  }
}
