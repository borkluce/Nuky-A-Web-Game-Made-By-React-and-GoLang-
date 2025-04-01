import React from "react"

interface CLogoSectionProps {
    lightMode?: boolean
    className?: string
    isHorizontal?: boolean
    logoSize?: number
}

const CLogoSection: React.FC<CLogoSectionProps> = ({
    lightMode = false,
    className = "",
    isHorizontal = false,
}: CLogoSectionProps) => {
    return (
        <div
            className={`flex items-center
        ${isHorizontal ? "flex-row" : "flex-col"}
         gap-2 py-1`}
        >
            <img
                src="/brand/logo/raw.png"
                alt="Logo"
                className={`w-44 -my-12  ${lightMode ? "" : ""} ${className}`}
            />
        </div>
    )
}

export default CLogoSection
