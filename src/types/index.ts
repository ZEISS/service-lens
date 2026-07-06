import type { SQL } from "drizzle-orm"
import type z from "zod"

export type Prettify<T> = {
  [K in keyof T]: T[K]
} & {}

export type EmptyProps<T extends React.ElementType> = Omit<React.ComponentProps<T>, keyof React.ComponentProps<T>>

export interface SearchParams {
  [key: string]: string | string[] | undefined
}

export interface QueryBuilderOpts {
  where?: SQL
  orderBy?: SQL
  distinct?: boolean
  nullish?: boolean
}

export type ZodTreeifyError<T> = ReturnType<typeof z.treeifyError<T>>
export type ZodFormState<T> = {
  values?: z.infer<T>
  errors?: ZodTreeifyError<T>
  success: boolean
}
