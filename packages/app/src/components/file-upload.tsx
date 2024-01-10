'use client'

import { useForm, UseFormRegisterReturn } from 'react-hook-form'
import { ReactNode, useRef } from 'react'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

type FileUploadProps = {
  register: UseFormRegisterReturn
  accept?: string
  multiple?: boolean
  children?: ReactNode
}

export const FileUpload = (props: FileUploadProps) => {
  const { register, accept, multiple, children } = props
  const inputRef = useRef<HTMLInputElement | null>(null)
  const { ref, ...rest } = register as {
    ref: (instance: HTMLInputElement | null) => void
  }

  const handleClick = () => inputRef.current?.click()

  return (
    <div className="grid w-full max-w-sm items-center gap-1.5">
      <Label htmlFor="file">{children}</Label>
      <Input
        id="picture"
        type="file"
        {...rest}
        ref={e => {
          ref(e)
          inputRef.current = e
        }}
      />
    </div>
  )
}
