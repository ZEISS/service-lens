import { Separator } from '@/components/ui/separator'

export default async function Page() {
  return (
    <>
      <div className="space-y-6">
        <div>
          <h3 className="text-lg font-medium">General</h3>
          <p className="text-sm text-muted-foreground"></p>
        </div>
        <Separator />
      </div>
    </>
  )
}
