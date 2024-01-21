import axios, { AxiosInstance, AxiosResponse } from 'axios'
import { toast } from 'react-hot-toast'
import { GetAPIConfig } from '../constants'

const apiInstance: AxiosInstance = axios.create({
    baseURL: GetAPIConfig().API_URL,
    headers: { 'content-type': 'application/json' },
    withCredentials: true,
})

apiInstance.interceptors.response.use(
    async function (response: AxiosResponse<any>) {
        return Promise.resolve(response)
    },
    async function (error: any) {
        if (error.code === "ERR_NETWORK") {
            toast.error("There was some problem connecting slashbase server")
            return Promise.reject(error)
        }
        const status = error.status || error.response.status;
        if (status === 401) {
            return Promise.resolve(error.response)
        }
        if (status === 500) {
            // TODO: move error toasts here
            return Promise.resolve(error.response)
        }
        return Promise.reject(error)
    }
)
const Request = { apiInstance }

export default Request