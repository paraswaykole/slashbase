import axios, { AxiosInstance, AxiosResponse } from 'axios'
import { toast } from 'react-hot-toast'
import { GetAPIConfig } from '../constants'
import { SecurityKey } from '../../wailsjs/go/app/App'

declare var window: any;

const getApiInstance = async (): Promise<AxiosInstance> => {

    if (window.apiInstance) {
        return window.apiInstance
    }

    const securityKey: string = await SecurityKey()

    const apiInstance: AxiosInstance = axios.create({
        baseURL: GetAPIConfig().API_URL,
        headers: { 'content-type': 'text/json', 'x-security-key': securityKey },
    })

    apiInstance.interceptors.response.use(
        async function (response: AxiosResponse<any>) {
            return Promise.resolve(response)
        },
        async function (error: any) {
            if (error.code === "ERR_NETWORK") {
                toast.error("There was some problem starting Slashbase. Please report this on out GitHub Repository.")
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

    window.apiInstance = apiInstance

    return apiInstance
}

const Request = { getApiInstance }

export default Request