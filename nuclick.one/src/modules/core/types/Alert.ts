export default interface Alert {
    id: string // Unique identifier for the alert
    message: string // Message displayed in the alert
    type: "success" | "error" | "warning" | "info" // Type of the alert
    countdown?: number // Optional countdown timer in seconds
    className?: string // Optional CSS class for styling
    onClose: Function
}
