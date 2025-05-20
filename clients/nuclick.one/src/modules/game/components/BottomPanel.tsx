// components/BottomPanel.tsx

import React, { useEffect, useState } from "react"

const BottomPanel: React.FC = () => {
    const [progress, setProgress] = useState(0)
    const [timeLeft, setTimeLeft] = useState("")

    // Helper to calculate milliseconds until next UTC 14:00
    const getMsUntilNextUTC14 = (): number => {
        const now = new Date()
        const nextUTC14 = new Date(Date.UTC(
            now.getUTCFullYear(),
            now.getUTCMonth(),
            now.getUTCDate(),
            14, 0, 0, 0
        ))

        if (now.getUTCHours() >= 14) {
            nextUTC14.setUTCDate(nextUTC14.getUTCDate() + 1)
        }

        return nextUTC14.getTime() - now.getTime()
    }

    useEffect(() => {
        const updateProgress = () => {
            const now = new Date()
            const total = 24 * 60 * 60 * 1000 // 24h in ms
            const msSinceLastUTC14 = total - getMsUntilNextUTC14()
            const percent = (msSinceLastUTC14 / total) * 100

            setProgress(Math.min(percent, 100))

            // Format time left
            const seconds = Math.floor(getMsUntilNextUTC14() / 1000)
            const hrs = Math.floor(seconds / 3600)
            const mins = Math.floor((seconds % 3600) / 60)
            const secs = seconds % 60

            setTimeLeft(`${hrs}h ${mins}m ${secs}s`)
        }

        updateProgress()
        const interval = setInterval(updateProgress, 1000)

        return () => clearInterval(interval)
    }, [])

    return (
        <div className="bg-white p-4 border-t h-[120px] shrink-0 flex flex-col justify-center">
            <p className="text-gray-700 mb-2">
                Next reset at UTC 14:00 — Time left: <strong>{timeLeft}</strong>
            </p>
            <div className="w-full bg-gray-200 h-4 rounded-full overflow-hidden">
                <div
                    className="bg-green-500 h-full transition-all"
                    style={{ width: `${progress}%` }}
                />
            </div>
        </div>
    )
}

export default BottomPanel
