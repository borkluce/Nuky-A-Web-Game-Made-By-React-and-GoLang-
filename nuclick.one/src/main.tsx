import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import "./index.css"
import { AuthView } from "./modules/auth"
import { BrowserRouter, Route, Routes } from "react-router"

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<LayoutView />}>
                    <Route index element={<AuthView />} />
                    <Route path="home" element={<MapView />} />
                    <Route path="*" element={<NotFoundView />} />
                </Route>
            </Routes>
        </BrowserRouter>
    </StrictMode>
)
