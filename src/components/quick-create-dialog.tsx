"use client"

import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { PlusCircleIcon } from "lucide-react"
import { useRouter } from "next/navigation"
import { useActionState, useEffect } from "react"
import { createDesignAction } from "./quick-create-dialog.action"
import { Label } from "./ui/label"

export function QuickCreateDialog() {
  const router = useRouter()

  const [state, formAction, pending] = useActionState(createDesignAction, {
    errors: [],
    success: false,
  })

  useEffect(() => {
    if (state.success) {
      router.push(`/designs/${state.designId}`)
    }
  }, [router, state.success, state.designId])

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="w-full justify-start">
          <PlusCircleIcon />
          <span>Quick Create</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Create Design</DialogTitle>
          <DialogDescription>Fill in the information below to create a new design.</DialogDescription>
        </DialogHeader>
        <form id="quick-create-form" action={formAction} className="w-full">
          <div className="grid gap-4">
            <Label htmlFor="title">Title</Label>
            <Input
              id="title"
              name="title"
              type="text"
              placeholder="Indiana Jones and the Fate of Atlantis"
              autoComplete="off"
              disabled={pending}
              required
            />
            {/* {formErrors.map((error, index) => (
                            <FormMessage key={index}>{error.path}</FormMessage>
                        ))} */}
          </div>
        </form>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button type="submit" form="quick-create-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
