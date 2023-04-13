import { SizeUnit } from "./enums"

export interface Index {
  name: string
  url: string
  entries: number
  size: string
  replicas: number
}

export interface IndexDataResponse {
  Name: string
  Entries: number
  Size: string
  VectorType: string
  VectorLength: number
}

export interface Size {
  value: number
  unit: SizeUnit
}
