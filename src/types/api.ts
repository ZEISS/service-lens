// Utility type to extract keys from an object type T whose values match type V
type KeysMatching<T extends object, V> = {
  [K in keyof T]-?: T[K] extends V ? K : never
}[keyof T]

// Generic interface for API request parameters
export type ApiRequest<T> = {
  offset?: number
  limit?: number
  filter?: Partial<T>
}
declare const ApiRequest: new <T>() => ApiRequest<T>

// Generic interface for API response structure
export type ApiResponse<T> = {
  items: T[]
  totalCount: number
  offset: number
  limit: number
}
declare const ApiResponse: new <T>() => ApiResponse<T>

// // Type for creating a new instance of T, excluding auto-generated fields like 'id' and timestamps
// export type CreateParams<T> = Omit<T, KeysMatching<T, number | Date>> & Partial<Pick<T, KeysMatching<T, number | Date>>>
// // Type for updating an instance of T, making all fields optional except for 'id'
// export type UpdateParams<T> = Partial<Omit<T, 'id'>> & { id: T['id'] }
// // Type for identifying an instance of T by its 'id'
// export type IdentifierParams<T> = Pick<T, 'id'>
// // Type for filtering instances of T based on its fields
// export type FilterParams<T> = Partial<T>
// // Type for sorting instances of T based on its fields
// export type SortParams<T> = Partial<Record<keyof T, 'asc' | 'desc'>>

// export type DesignCreateParams = CreateParams<import('@/db/models/design').Design>
// export type DesignUpdateParams = UpdateParams<import('@/db/models/design').Design>
// export type DesignIdentifierParams = IdentifierParams<import('@/db/models/design').Design>
// export type DesignFilterParams = FilterParams<import('@/db/models/design').Design>
// export type DesignSortParams = SortParams<import('@/db/models/design').Design>
// export interface DesignApiResponse extends ApiResponse<import('@/db/models/design').Design> {}
