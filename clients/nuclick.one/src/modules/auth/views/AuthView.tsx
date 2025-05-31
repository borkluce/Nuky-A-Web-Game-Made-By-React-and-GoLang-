import { useEffect, useState } from "react"
import { useUser } from "../hooks/useUser"
import { LoginForm } from "@/components/login-form"
import { Button } from "@/components/ui/button"
import GameView from "@/modules/game/views/GameView"

interface AuthViewProps {
    children?: React.ReactNode
}

export function AuthView({ children }: AuthViewProps) {
    const { user, isAuthenticated, logout } = useUser()
    const [isLoading, setIsLoading] = useState(true)

    useEffect(() => {
        // Check if user is authenticated on component mount
        setIsLoading(false)
    }, [])

    if (isLoading) {
        return (
            <div className="flex items-center justify-center min-h-screen">
                <div className="text-center">
                    <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900 mx-auto"></div>
                    <p className="mt-4 text-gray-600">Loading...</p>
                </div>
            </div>
        )
    }

    if (!isAuthenticated() || !user) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="w-full max-w-md">
                    <LoginForm />
                </div>
            </div>
        )
    }

    // If logging in is successful
    return <GameView />
}

export default AuthView
