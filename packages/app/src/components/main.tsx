import { PropsWithChildren } from 'react'

export interface MainProps extends React.HTMLAttributes<HTMLElement> {}

export function Main({ children, ...props }: PropsWithChildren<MainProps>) {
  return <main {...props}>{children}</main>
}
