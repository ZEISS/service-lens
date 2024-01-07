'use client'

import * as React from 'react'
import { X } from 'lucide-react'

import { Badge } from '@/components/ui/badge'
import { Command, CommandGroup, CommandItem } from '@/components/ui/command'
import { Command as CommandPrimitive } from 'cmdk'

export type FancyMultiSelectValue<T> = {
  value: T
  label: string
}

export interface FancyMultiSelectProps<T> {
  defaultValue?: string
  placeholder?: string
  onValueChange?(value: any): void
  dataValues?: FancyMultiSelectValue<T>[]
}

export function FancyMultiSelect<T>({
  placeholder,
  onValueChange,
  dataValues = []
}: FancyMultiSelectProps<T>) {
  const inputRef = React.useRef<HTMLInputElement>(null)
  const [open, setOpen] = React.useState(false)
  const [selected, setSelected] = React.useState<FancyMultiSelectValue<T>[]>([
    dataValues[0]
  ])
  const [inputValue, setInputValue] = React.useState('')

  React.useEffect(() => {
    onValueChange?.(selected.map(s => s.value))
  }, [selected, onValueChange])

  const handleSelect = React.useCallback((value: FancyMultiSelectValue<T>) => {
    setSelected(prev => [...prev, value])
  }, [])

  const handleUnselect = React.useCallback(
    (select: FancyMultiSelectValue<T>) => {
      setSelected(prev => prev.filter(s => s.value !== select.value))
    },
    []
  )

  const handleKeyDown = React.useCallback(
    (e: React.KeyboardEvent<HTMLDivElement>) => {
      const input = inputRef.current
      if (input) {
        if (e.key === 'Delete' || e.key === 'Backspace') {
          if (input.value === '') {
            setSelected(prev => {
              const newSelected = [...prev]
              newSelected.pop()

              return newSelected
            })
          }
        }
        // This is not a default behaviour of the <input /> field
        if (e.key === 'Escape') {
          input.blur()
        }
      }
    },
    []
  )

  const selectables = dataValues.filter(select =>
    selected.every(s => s.value !== select.value)
  )

  return (
    <Command
      onKeyDown={handleKeyDown}
      className="overflow-visible bg-transparent"
    >
      <div className="group border border-input px-3 py-2 text-sm ring-offset-background rounded-md focus-within:ring-2 focus-within:ring-ring focus-within:ring-offset-2">
        <div className="flex gap-1 flex-wrap">
          {selected.map(select => {
            return (
              <Badge key={select.label} variant="secondary">
                {select.label}
                <button
                  className="ml-1 ring-offset-background rounded-full outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
                  onKeyDown={e => {
                    if (e.key === 'Enter') {
                      handleUnselect(select)
                    }
                  }}
                  onMouseDown={e => {
                    e.preventDefault()
                    e.stopPropagation()
                  }}
                  onClick={() => handleUnselect(select)}
                >
                  <X className="h-3 w-3 text-muted-foreground hover:text-foreground" />
                </button>
              </Badge>
            )
          })}
          {/* Avoid having the "Search" Icon */}
          <CommandPrimitive.Input
            ref={inputRef}
            value={inputValue}
            onValueChange={setInputValue}
            onBlur={() => setOpen(false)}
            onFocus={() => setOpen(true)}
            placeholder={placeholder}
            className="ml-2 bg-transparent outline-none placeholder:text-muted-foreground flex-1"
          />
        </div>
      </div>
      <div className="relative mt-2">
        {open && selectables.length > 0 ? (
          <div className="absolute w-full z-10 top-0 rounded-md border bg-popover text-popover-foreground shadow-md outline-none animate-in">
            <CommandGroup className="h-full overflow-auto">
              {selectables.map(select => {
                return (
                  <CommandItem
                    key={select.label}
                    onMouseDown={e => {
                      e.preventDefault()
                      e.stopPropagation()
                    }}
                    onSelect={value => {
                      setInputValue('')
                      handleSelect(select)
                    }}
                    className={'cursor-pointer'}
                  >
                    {select.label}
                  </CommandItem>
                )
              })}
            </CommandGroup>
          </div>
        ) : null}
      </div>
    </Command>
  )
}
