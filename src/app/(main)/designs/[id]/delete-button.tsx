"use client"

import { Button } from "@/components/ui/button"
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import type { TDesign } from "@/db/schemas/design"
import { TrashIcon } from "lucide-react"
import { useRouter } from "next/navigation"
import { useState } from "react"
import { toast } from "sonner"
import { deleteDesignAction } from "../_components/data-rows-actions.action"

interface DeleteButtonProps {
  design: TDesign
}

export function DeleteButton({ design }: DeleteButtonProps) {
  const router = useRouter()
  const [isDeleting, setIsDeleting] = useState(false)

  const handleDelete = async () => {
    setIsDeleting(true)
    try {
      // Create FormData to match the expected action signature
      const formData = new FormData()
      formData.append("id", design.id)

      const result = await deleteDesignAction(null, formData)

      if (result?.success) {
        toast.success("Design deleted successfully")
        router.push("/designs")
      } else {
        toast.error("Failed to delete design")
      }
    } catch (error) {
      console.error("Error deleting design:", error)
      toast.error("Failed to delete design")
    } finally {
      setIsDeleting(false)
    }
  }

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="outline" disabled={isDeleting}>
          <TrashIcon className="h-4 w-4" />
          Delete
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Delete Design</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to delete &quot;{design.title}&quot;? This action cannot be undone.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel disabled={isDeleting}>Cancel</AlertDialogCancel>
          <AlertDialogAction
            onClick={handleDelete}
            disabled={isDeleting}
            className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
          >
            {isDeleting ? "Deleting..." : "Delete"}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
