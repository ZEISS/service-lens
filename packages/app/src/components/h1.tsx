import { PropsWithChildren } from 'react'

export type H1Props = {}

export function H1({ children }: PropsWithChildren<H1Props>) {
  return (
    <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
      {children}
    </h1>
  )
}
