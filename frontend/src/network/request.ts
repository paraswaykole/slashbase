import axios, { AxiosInstance } from 'axios'
import { GetAPIConfig } from '../constants'

const getApiInstance = () => {
    const apiInstance: AxiosInstance = axios.create({
        baseURL: GetAPIConfig().API_URL,
        headers: {'content-type': 'text/json'},
        withCredentials: true,
    })
    return apiInstance
}

export default { getApiInstance }