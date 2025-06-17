import React from "react"

interface WinnerPageProps {
    winnerProvinceName: string
}

const WinnerPage: React.FC<WinnerPageProps> = ({ winnerProvinceName }) => {
    return (
        <div className="flex items-center justify-center w-screen h-screen bg-gradient-to-br from-yellow-400 via-yellow-500 to-orange-500">
            <div className="text-center">
                <h1 className="text-6xl font-bold text-white mb-8 drop-shadow-lg">
                    🎉 GAME OVER 🎉
                </h1>
                <p className="text-4xl font-semibold text-white drop-shadow-md">
                    The province <span className="font-bold underline">{winnerProvinceName}</span> is the winner!
                </p>
            </div>
        </div>
    )
}

export default WinnerPage