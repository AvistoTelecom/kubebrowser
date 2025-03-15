export function copyToClipboard(text: string): void {
  navigator.clipboard.writeText(text).catch(err => {
    console.error('Failed to copy text: ', err)
  })
}
