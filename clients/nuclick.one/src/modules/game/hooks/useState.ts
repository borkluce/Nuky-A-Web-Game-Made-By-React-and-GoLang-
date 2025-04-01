import { create } from "zustand"
import { CAxios } from "../../core/configs/cAxios"

// Types
import { State } from "../types/state"

interface useStateState {
    stateList: State[] | null
    getStates: () => Promise<void>
}

export const useState = create<useStateState>((set, get) => ({
    stateList: null,
    getStates: async () => {},
}))
