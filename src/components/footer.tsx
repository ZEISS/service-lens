import { cn } from '@/lib/utils'

export type FooterProps = {
  className?: string
}

export default function Footer({ className, ...props }: FooterProps) {
  return (
    <footer
      className={cn('flex items-center space-x-4 lg:space-x-6', className)}
      {...props}
    >
      <div className="border-b"></div>
    </footer>
  )
}
