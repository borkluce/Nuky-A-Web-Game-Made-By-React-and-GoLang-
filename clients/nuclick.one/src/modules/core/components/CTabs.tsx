import React, { useState } from "react"

export interface Tab {
    title: string // Title of the tab
    content: React.ReactNode // Content of the tab
}

interface CTabsProps {
    className?: string
    tabs: Tab[] // Array of Tab objects
}

const CTabs: React.FC<CTabsProps> = ({ tabs, className = "" }) => {
    const [activeIndex, setActiveIndex] = useState(0) // Track the active tab

    return (
        <div className={`flex flex-col ${className}`}>
            {/* Tab Buttons */}
            <div className="flex border-b border-black border-opacity-10">
                {tabs.map((tab, index) => (
                    <button
                        key={tab.title}
                        className={`px-10 py-2 transition-all duration-300 
                            ${
                                index === activeIndex
                                    ? "bg-white border-b-[4px] border-5 font-bold"
                                    : "hover:bg-gray-200"
                            }`}
                        onClick={() => setActiveIndex(index)}
                    >
                        {tab.title} {/* Display tab title */}
                    </button>
                ))}
            </div>

            {/* Tab Content */}
            <div className="mt-4">
                {tabs[activeIndex] && tabs[activeIndex].content}
            </div>
        </div>
    )
}

export default CTabs
