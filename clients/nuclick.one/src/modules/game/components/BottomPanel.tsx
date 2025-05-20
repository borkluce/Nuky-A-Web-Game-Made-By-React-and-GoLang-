import React, { useEffect, useState } from "react"
import { GiMineExplosion } from "react-icons/gi"
import { useUser } from "../../auth/hooks/useUser"

const BottomPanel: React.FC = () => {
    const { getRemainingCooldownSeconds } = useUser()
    const [cooldownSeconds, setCooldownSeconds] = useState<number>(0)
    const [progressPercent, setProgressPercent] = useState(0)

    // Update cooldown every second
    useEffect(() => {
        const interval = setInterval(() => {
            setCooldownSeconds(getRemainingCooldownSeconds())
        }, 1000)

        return () => clearInterval(interval)
    }, [getRemainingCooldownSeconds])

    // Update daily progress bar to next UTC 14:00
    useEffect(() => {
        const updateProgress = () => {
            const now = new Date()
            const utcNow = new Date(
                Date.UTC(
                    now.getUTCFullYear(),
                    now.getUTCMonth(),
                    now.getUTCDate(),
                    now.getUTCHours(),
                    now.getUTCMinutes(),
                    now.getUTCSeconds()
                )
            )

            let nextUTC14 = new Date(utcNow)
            nextUTC14.setUTCHours(14, 0, 0, 0)

            // If already past UTC 14 today, set to next day's 14:00
            if (utcNow >= nextUTC14) {
                nextUTC14.setUTCDate(nextUTC14.getUTCDate() + 1)
            }

            const totalDay = 24 * 60 * 60 * 1000
            const elapsed = utcNow.getTime() - (nextUTC14.getTime() - totalDay)
            const progress = Math.min((elapsed / totalDay) * 100, 100)
            setProgressPercent(progress)
        }

        updateProgress()
        const interval = setInterval(updateProgress, 60000) // update every minute
        return () => clearInterval(interval)
    }, [])

    return (
        <div className="flex w-full h-20 border-t bg-white">
            {/* Progress to UTC 14:00 */}
            <div className="w-1/2 p-2">
                <div className="text-sm text-gray-600 mb-1 flex items-center justify-between">
                    <span>Until next UTC 14:00</span>
                    <GiMineExplosion className="text-red-500 ml-2" size={18} />
                </div>
                <div className="w-full h-4 bg-gray-200 rounded">
                    <div
                        className="h-full bg-green-500 rounded"
                        style={{ width: `${progressPercent}%`, transition: "width 0.5s" }}
                    />
                </div>
            </div>

            {/* Cooldown counter */}
            <div className="w-1/2 p-2 flex flex-col justify-center items-end">
                <div className="text-sm text-gray-600 mb-1">Cooldown</div>
                <div className="text-lg font-semibold">
                    {cooldownSeconds > 0 ? `${cooldownSeconds}s` : "Ready"}
                </div>
            </div>
        </div>
    )
}

export default BottomPanel
