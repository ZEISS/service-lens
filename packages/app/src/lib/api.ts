import type { NextRequest, NextResponse } from 'next/server'
import type { User } from './common'
import { COOKIE_NAME } from './common'
import { Query } from '@/components/data-table'

export const randomId = () => crypto.randomUUID()

export const getUsers = (req: NextRequest): User[] => {
  // we store all data in cookies for demo purposes
  const cookie = req.cookies.get(COOKIE_NAME)
  if (cookie) {
    return JSON.parse(cookie.value)
  }
  return []
}

export const saveUsers = (res: NextResponse, users: User[]) => {
  res.cookies.set(COOKIE_NAME, JSON.stringify(users), {
    path: '/'
  })
  return res
}

// export type Query<T> = (
//   pagination: PaginationState
// ) => Promise<{ rows: T[]; count: number }>

export const useQuery =
  <Q extends unknown>(query: Query<Q>) =>
  () =>
    query
