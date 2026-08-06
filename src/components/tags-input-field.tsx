"use client"

import { Badge } from "@/components/ui/badge"
import { FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { cn } from "@/lib/utils"
import { AnimatePresence, motion } from "framer-motion"
import { Plus, Tag, X } from "lucide-react"
import { type KeyboardEvent, memo, type ReactNode, useEffect, useRef, useState } from "react"
import { type FieldValues, type Path, useFormContext } from "react-hook-form"

interface TagsInputFieldProps<TFieldValues extends FieldValues> {
  name: Path<TFieldValues>
  beautifyName?: string
  description?: string
  label: string
  placeholder?: string
  disabled?: boolean
  className?: string
  maxTags?: number
  maxLength?: number
  autoFocus?: boolean
  startIcon?: ReactNode
  endIcon?: ReactNode
  allowDuplicates?: boolean
  suggestions?: string[]
  variant?: "default" | "enterprise" | "minimal"
  tagVariant?: "default" | "secondary" | "outline" | "destructive"
}

const TagsInputFieldBase = <TFieldValues extends FieldValues>({
  name,
  beautifyName,
  description,
  label,
  placeholder,
  disabled = false,
  className,
  maxTags,
  maxLength = 100,
  autoFocus = false,
  startIcon,
  endIcon,
  allowDuplicates = false,
  suggestions = [],
  variant = "enterprise",
  tagVariant = "default",
}: TagsInputFieldProps<TFieldValues>) => {
  const { control } = useFormContext()
  const [inputValue, setInputValue] = useState("")
  const [showSuggestions, setShowSuggestions] = useState(false)
  const inputRef = useRef<HTMLInputElement>(null)
  const containerRef = useRef<HTMLDivElement>(null)

  const filteredSuggestions = suggestions.filter(
    (suggestion) => suggestion.toLowerCase().includes(inputValue.toLowerCase()) && inputValue.length > 0,
  )

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (containerRef.current && !containerRef.current.contains(event.target as Node)) {
        setShowSuggestions(false)
      }
    }

    document.addEventListener("mousedown", handleClickOutside)
    return () => document.removeEventListener("mousedown", handleClickOutside)
  }, [])

  const addTag = (tag: string, currentTags: string[], onChange: (tags: string[]) => void) => {
    const trimmedTag = tag.trim()
    if (!trimmedTag) return

    if (!allowDuplicates && currentTags.includes(trimmedTag)) return
    if (maxTags && currentTags.length >= maxTags) return
    if (trimmedTag.length > maxLength) return

    onChange([...currentTags, trimmedTag])
    setInputValue("")
    setShowSuggestions(false)
  }

  const removeTag = (index: number, currentTags: string[], onChange: (tags: string[]) => void) => {
    const newTags = currentTags.filter((_, i) => i !== index)
    onChange(newTags)
  }

  const handleKeyDown = (
    e: KeyboardEvent<HTMLInputElement>,
    currentTags: string[],
    onChange: (tags: string[]) => void,
  ) => {
    if (e.key === "Enter" || e.key === ",") {
      e.preventDefault()
      addTag(inputValue, currentTags, onChange)
    } else if (e.key === "Backspace" && !inputValue && currentTags.length > 0) {
      removeTag(currentTags.length - 1, currentTags, onChange)
    } else if (e.key === "Escape") {
      setShowSuggestions(false)
    }
  }

  const getVariantStyles = () => {
    switch (variant) {
      case "enterprise":
        return {
          container:
            "dark:border-gray-600 border-2 border-muted hover:border-primary/30 transition-all duration-300 backdrop-blur-sm",
          input: "bg-transparent border-0 focus:ring-0 placeholder:text-muted-foreground/60",
          suggestions: "bg-background/95 backdrop-blur-md border border-border/50 shadow-xl",
        }
      case "minimal":
        return {
          container: "bg-background border border-border hover:border-primary/50 transition-colors",
          input: "bg-transparent border-0 focus:ring-0",
          suggestions: "bg-background border border-border shadow-lg",
        }
      default:
        return {
          container: "bg-background border border-input hover:border-primary/50 transition-colors",
          input: "bg-transparent border-0 focus:ring-0",
          suggestions: "bg-background border border-border shadow-lg",
        }
    }
  }

  const styles = getVariantStyles()

  return (
    <FormField
      control={control}
      name={name}
      render={({ field }) => {
        const tags = field.value || []

        return (
          <FormItem className={cn("space-y-2", className)}>
            <FormLabel className="flex items-center gap-2">
              {label}
              {maxTags && (
                <Badge variant="outline" className="text-xs dark:border-gray-500">
                  {tags.length}/{maxTags}
                </Badge>
              )}
            </FormLabel>

            <FormControl>
              <div ref={containerRef} className="relative">
                <div
                  className={cn(
                    "flex min-h-[2.5rem] flex-wrap items-center gap-2 rounded-md p-2",
                    styles.container,
                    disabled && "cursor-not-allowed opacity-50",
                  )}
                >
                  {startIcon && <span className="text-muted-foreground">{startIcon}</span>}

                  <AnimatePresence>
                    {tags.map((tag: string, index: number) => (
                      <motion.div
                        key={`${tag}-${index}`}
                        initial={{ opacity: 0, scale: 0.8 }}
                        animate={{ opacity: 1, scale: 1 }}
                        exit={{ opacity: 0, scale: 0.8 }}
                        transition={{ duration: 0.2 }}
                      >
                        <Badge
                          variant={tagVariant}
                          className={cn(
                            "group flex items-center gap-1 pr-1 transition-colors",
                            variant === "enterprise" && "border-primary/20 bg-primary",
                          )}
                        >
                          <span className="max-w-[150px] truncate">{tag}</span>
                          {!disabled && (
                            <button
                              type="button"
                              onClick={() => removeTag(index, tags, field.onChange)}
                              className="ml-1 rounded-full p-0.5 transition-colors hover:bg-destructive/20"
                            >
                              <X className="h-3 w-3" />
                            </button>
                          )}
                        </Badge>
                      </motion.div>
                    ))}
                  </AnimatePresence>

                  <div className="min-w-[120px] flex-1">
                    <Input
                      ref={inputRef}
                      value={inputValue}
                      onChange={(e) => {
                        setInputValue(e.target.value)
                        setShowSuggestions(e.target.value.length > 0 && suggestions.length > 0)
                      }}
                      onKeyDown={(e) => handleKeyDown(e, tags, field.onChange)}
                      onFocus={() => setShowSuggestions(inputValue.length > 0 && suggestions.length > 0)}
                      placeholder={
                        tags.length === 0
                          ? (placeholder ?? `Enter ${label.toLowerCase()} and press Enter`)
                          : maxTags && tags.length >= maxTags
                            ? `Maximum ${maxTags} tags reached`
                            : "Add another..."
                      }
                      className={styles.input}
                      disabled={disabled || (maxTags ? tags.length >= maxTags : false)}
                      autoFocus={autoFocus}
                      maxLength={maxLength}
                    />
                  </div>

                  {endIcon && <span className="text-muted-foreground">{endIcon}</span>}

                  {!disabled && (
                    <button
                      type="button"
                      onClick={() => {
                        if (inputValue.trim()) {
                          addTag(inputValue, tags, field.onChange)
                        }
                        inputRef.current?.focus()
                      }}
                      className="rounded p-1 transition-colors hover:bg-muted"
                      disabled={!inputValue.trim() || (maxTags ? tags.length >= maxTags : false)}
                    >
                      <Plus className="h-4 w-4 text-muted-foreground" />
                    </button>
                  )}
                </div>

                {/* Suggestions Dropdown */}
                <AnimatePresence>
                  {showSuggestions && filteredSuggestions.length > 0 && (
                    <motion.div
                      initial={{ opacity: 0, y: -10 }}
                      animate={{ opacity: 1, y: 0 }}
                      exit={{ opacity: 0, y: -10 }}
                      transition={{ duration: 0.2 }}
                      className={cn("absolute z-50 mt-1 w-full overflow-auto rounded-md py-1", styles.suggestions)}
                    >
                      {filteredSuggestions.map((suggestion) => (
                        <button
                          key={suggestion}
                          type="button"
                          onClick={() => addTag(suggestion, tags, field.onChange)}
                          className="w-full px-3 py-2 text-left text-sm transition-colors hover:bg-muted"
                        >
                          <div className="flex items-center gap-2">
                            {startIcon ? startIcon : <Tag className="h-3 w-3 text-muted-foreground" />}
                            {suggestion}
                          </div>
                        </button>
                      ))}
                    </motion.div>
                  )}
                </AnimatePresence>
              </div>
            </FormControl>

            {description && <FormDescription>{description}</FormDescription>}

            <div className="flex items-center justify-between">
              <FormMessage />
              <div className="flex items-center gap-2 text-muted-foreground text-xs">
                {maxLength && inputValue && (
                  <span>
                    {inputValue.length}/{maxLength}
                  </span>
                )}
                {tags.length > 0 && (
                  <Badge variant="outline" className="text-xs">
                    {tags.length} {beautifyName ?? "tags"}
                  </Badge>
                )}
              </div>
            </div>
          </FormItem>
        )
      }}
    />
  )
}

export const TagsInputField = memo(TagsInputFieldBase) as typeof TagsInputFieldBase
