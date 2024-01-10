'use client'

import * as React from 'react'
import { SessionProvider as NextAuthSessionProvider } from 'next-auth/react'
import { type SessionProviderProps } from 'next-auth/react'

export function SessionProvider({ children, ...props }: SessionProviderProps) {
  return (
    <NextAuthSessionProvider {...props}>{children}</NextAuthSessionProvider>
  )
}

export default SessionProvider
