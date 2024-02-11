import { PropsWithChildren } from 'react'

export type LeadProps = {}

export function Lead({ children }: PropsWithChildren<LeadProps>) {
  return <p className="text-xl text-muted-foreground">{children}</p>
}
