import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
// import { BrowserRouter, Route, Routes } from "react-router"

import "./index.css"

// Views
// import { AuthView } from "./modules/auth"
import { GameView } from "./modules/game"
import { AuthView } from "./modules/auth"

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <div className="w-screen min-h-screen flex items-center">
            {/* <GameView /> */}
            <AuthView />
        </div>
    </StrictMode>
)

/*
    For now the routing is disabled because everything like authentication page as well are handled in the
    single viewport without any routing, instead using pop-ups, modals etc.
*/
