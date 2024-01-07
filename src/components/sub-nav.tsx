import React from 'react'

export type SubNavProps = {
  children?: React.ReactNode
  className?: string
}

export type SubNavTitleProps = {
  children?: React.ReactNode
}

export type SubNavActionsProps = {
  children?: React.ReactNode
}

export type SubNavSubtitleProps = {
  children?: string
}

export function SubNav({ children, ...props }: SubNavProps) {
  return (
    <aside
      className="flex border-b items-center justify-between space-y-2 p-8 pt-6"
      {...props}
    >
      {children}
    </aside>
  )
}

export function SubNavTitle({ children, ...props }: SubNavTitleProps) {
  return (
    <h3
      className="scroll-m-20 text-2xl font-semibold tracking-tight"
      {...props}
    >
      {children}
    </h3>
  )
}

export function SubNavSubtitle({ children, ...props }: SubNavSubtitleProps) {
  return (
    <p className="text-sm text-muted-foreground" {...props}>
      {children}
    </p>
  )
}

export function SubNavActions({ children, ...props }: SubNavProps) {
  return (
    <div className="flex items-center space-x-2" {...props}>
      {children}
    </div>
  )
}
