import React from "react"
import CLogoSection from "../../core/components/CLogoSection"
import { CButton } from "../../core"

const NotFoundView: React.FC = () => {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen pb-80">
            <CLogoSection logoSize={32} />
            <h1 className="text-6xl font-bold text-gray-800 mt-4">404</h1>
            <p className="text-xl text-gray-600 mt-4">Page Not Found</p>
            <CButton
                onClick={() => window.location.replace("/")}
                className="mt-6"
            >
                Go back home
            </CButton>
        </div>
    )
}

export default NotFoundView
