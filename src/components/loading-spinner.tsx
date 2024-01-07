import { CgSpinner } from 'react-icons/cg'

export function LoadingSpinner() {
  return (
    <div className="flex items-center justify-center text-sm text-muted-foreground px-4">
      <CgSpinner className="mr-2 h-4 w-4 animate-spin" />
      Loading...
    </div>
  )
}
