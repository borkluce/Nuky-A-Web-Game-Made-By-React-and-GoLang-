import React from "react"
import CLogoSection from "../../core/components/CLogoSection"

const Footer: React.FC = () => {
    return (
        <footer className="bg-white  text-black py-16 z-20">
            <div className="container mx-auto px-4">
                <div className="flex flex-col items-center">
                    <CLogoSection isHorizontal />
                    <p className="mt-4">
                        &copy; {new Date().getFullYear()} Vafaill LTD. All
                        rights reserved.
                    </p>
                </div>
            </div>
        </footer>
    )
}

export default Footer
