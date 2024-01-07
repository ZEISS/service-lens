export type SectionProps = {
  children?: React.ReactNode
  className?: string
}

export function Section({ children, ...props }: SectionProps) {
  return <section className="h-full px-4 py-6 lg:px-8">{children}</section>
}
