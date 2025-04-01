import React, { useState, useEffect } from "react"

// Types
import Alert from "../types/alert"

// Components
import CAlert from "../components/CAlert"

interface AlertViewProps {}

const AlertView: React.FC<AlertViewProps> = () => {
    const [alerts, setAlerts] = useState<Alert[]>([])

    // Adds an alert to the stack with a countdown

    // Removes an alert by ID
    const removeAlert = (id: string) => {
        setAlerts((prevAlerts) => prevAlerts.filter((alert) => alert.id !== id))
    }

    // Countdown effect for each alert
    useEffect(() => {
        const interval = setInterval(() => {
            setAlerts((prevAlerts) =>
                prevAlerts.map((alert) =>
                    alert.countdown! > 0
                        ? { ...alert, countdown: alert.countdown! - 1 }
                        : alert
                )
            )
        }, 1000)

        return () => clearInterval(interval)
    }, [])

    // Automatically remove alerts with countdown reaching 0
    useEffect(() => {
        alerts
            .filter((alert) => alert.countdown === 0)
            .forEach((expiredAlert) => removeAlert(expiredAlert.id))
    }, [alerts])

    return (
        <div className="flex flex-col fixed top-10 left-0 right-0 z-50 gap-8">
            {alerts.map((alert) => (
                <CAlert
                    key={alert.id}
                    message={`${alert.message} (${alert.countdown}s remaining)`}
                    onClose={() => removeAlert(alert.id)}
                    type={alert.type}
                    id={alert.id}
                />
            ))}
        </div>
    )
}

export default AlertView
