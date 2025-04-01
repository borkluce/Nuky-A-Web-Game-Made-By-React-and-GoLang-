import axios from "axios"

export const CAxios = axios.create({
    baseURL: process.env.REACT_APP_API_URL,
})
