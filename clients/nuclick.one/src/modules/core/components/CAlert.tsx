import React from "react"

// Types
import Alert from "../types/alert"

const CAlert: React.FC<Alert> = ({ id, message, onClose, className, type }) => {
    return (
        <div
            className={`
                bg-black bg-opacity-80 
                text-white font-bold
                border-b-8
                ${
                    type === "success"
                        ? "border-green-500"
                        : type === "error"
                        ? "border-red-500"
                        : "border-blue-500"
                }
                p-5
                cursor-pointer z-50
                mx-auto w-2/6 h-fit
                ${className}
                `}
            onClick={() => onClose(id)}
        >
            {message}
        </div>
    )
}

export default CAlert
