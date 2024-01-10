import { PropsWithChildren } from 'react'

export type H4Props = {}

export function H4({ children }: PropsWithChildren<H4Props>) {
  return (
    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
      {children}
    </h4>
  )
}
