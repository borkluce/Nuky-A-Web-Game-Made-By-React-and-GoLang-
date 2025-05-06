import axios from "axios"

export const CAxios = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
})
