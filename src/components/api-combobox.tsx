"use client"

import * as React from "react"
import { Check, ChevronsUpDown } from "lucide-react"
import { useDebouncedCallback } from "use-debounce"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@/components/ui/command"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"

import { ALIGN_OPTIONS } from "@radix-ui/react-popper"

export type ComboBoxItem = {
  value: string
  label: string
}

export type ApiComboBoxFetchFunc<R> = (e: string, setItems: React.Dispatch<React.SetStateAction<ComboBoxItem[]>>) => R

type ComboboxProps<R = any> = {
  selectedItem: ComboBoxItem
  onSelect: (item: ComboBoxItem) => void
  searchPlaceholder?: string
  className?: string
  disabled?: boolean
  align?: (typeof ALIGN_OPTIONS)[number]
  onFetch: ApiComboBoxFetchFunc<R>
}

export function ApiComboBox({
  selectedItem,
  onSelect,
  searchPlaceholder = "Search...",
  className,
  disabled = false,
  align,
  onFetch,
}: ComboboxProps) {
  const [open, setOpenState] = React.useState(false)
  const [items, setItems] = React.useState<ComboBoxItem[]>([selectedItem])

  const debouncedFetchItems = useDebouncedCallback(onFetch, 300)
  const handleOnSearchChange = (e: string) => (e === "" && onFetch(e, setItems)) || debouncedFetchItems(e, setItems)

  function setOpen(isOpen: boolean) {
    if (isOpen) {
      setItems([])
      handleOnSearchChange("")
    }
    setOpenState(isOpen)
  }

  return (
    <Popover open={open} onOpenChange={setOpen} modal={true}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className={cn("justify-between", className)}
          disabled={disabled}
        >
          <span className="truncate flex items-center">{selectedItem.label || "Select an item"}</span>
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent style={{ width: "var(--radix-popover-trigger-width)" }} className={cn("p-0")} align={align}>
        <Command shouldFilter={false}>
          <CommandInput placeholder={searchPlaceholder} onValueChange={handleOnSearchChange} />
          <CommandList>
            <CommandEmpty>No results found.</CommandEmpty>
            <CommandGroup>
              {items.map((item) => {
                if (!item.value) {
                  return null
                }
                const isSelected = selectedItem.value === item.value
                return (
                  <CommandItem
                    key={item.value}
                    value={item.value}
                    keywords={[item.label]}
                    onSelect={() => {
                      onSelect(item)
                      setOpen(false)
                    }}
                  >
                    {item.label}
                    <Check className={cn("ml-auto h-4 w-4", isSelected ? "opacity-100" : "opacity-0")} />
                  </CommandItem>
                )
              })}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}
