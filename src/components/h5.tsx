import { PropsWithChildren } from 'react'

export type H5Props = {}

export function H5({ children }: PropsWithChildren<H5Props>) {
  return (
    <h5 className="scroll-m-20 text-base font-semibold tracking-tight">
      {children}
    </h5>
  )
}
