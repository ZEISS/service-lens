export type Column = {
  name: string
  uid: string
  sortable?: boolean
}

export type Columns<TDEF> = TDEF[]

const columns: Columns<Column> = [
  { name: 'ID', uid: 'id', sortable: true },
  { name: 'NAME', uid: 'name', sortable: true },
  { name: 'DESCRIPTION', uid: 'description' },
  { name: 'ENVIRONMENT', uid: 'environment', sortable: true },
  { name: 'PROFILE', uid: 'profile', sortable: true },
  { name: 'ACTIONS', uid: 'actions' }
]

const statusOptions = [
  { name: 'Active', uid: 'active' },
  { name: 'Paused', uid: 'paused' },
  { name: 'Vacation', uid: 'vacation' }
]

export type TableUser = {
  id: number
  name: string
  role: string
  team: string
  status: string
  age: string
  avatar: string
  email: string
}

export { columns, statusOptions }
