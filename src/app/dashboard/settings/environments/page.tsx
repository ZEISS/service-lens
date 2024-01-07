import { Separator } from '@/components/ui/separator'

export default function Page() {
  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Environments</h3>
        <p className="text-sm text-muted-foreground">
          Customize the environments of the service lens.
        </p>
      </div>
      <Separator />
    </div>
  )
}
