// types/common.d.ts

declare global {
  interface BigInt {
    toJSON(): string
  }
}
