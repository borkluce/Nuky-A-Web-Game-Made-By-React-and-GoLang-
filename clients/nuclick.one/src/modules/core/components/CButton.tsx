import React from "react"

interface CButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    children: React.ReactNode
    secondary?: boolean
}

const CButton: React.FC<CButtonProps> = ({
    children,
    className,
    secondary,
    ...props
}) => {
    return (
        <button
            {...props}
            className={`
                    py-3 px-5
                    ${
                        secondary
                            ? "border-2 border-opacity-60 border-black text-black"
                            : "bg-black bg-opacity-100 text-white"
                    }
                    text-sm font-bold hover:bg-primary
                    ${className || ""}
                `}
        >
            {children}
        </button>
    )
}

export default CButton
