import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import { BrowserRouter, Route, Routes } from "react-router"

import "./index.css"

// Views
import { AuthView } from "./modules/auth"
import { GameView } from "./modules/game"

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <GameView />
    </StrictMode>
)
