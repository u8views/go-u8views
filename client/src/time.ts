export function formatDay(timestamp: number): string {
    return new Date(timestamp * 1000).toLocaleDateString('en-us', {day: "2-digit", month: "long"})
}
