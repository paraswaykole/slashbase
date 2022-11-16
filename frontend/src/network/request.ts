import axios, { AxiosInstance, AxiosResponse } from 'axios'
import { GetAPIConfig } from '../constants'
import { clearLogin } from '../redux/currentUserSlice'
import reduxStore from '../redux/store'

const getApiInstance = () => {
    const apiInstance: AxiosInstance = axios.create({
        baseURL: GetAPIConfig().API_URL,
        headers: { 'content-type': 'text/json' },
        withCredentials: true,
    })

    apiInstance.interceptors.response.use(
        async function (response: AxiosResponse<any>) {
            return Promise.resolve(response)
        },
        async function (error: any) {
            const status = error.status || error.response.status;
            if (status === 401) {
                const { dispatch } = reduxStore
                await dispatch(clearLogin())
                return Promise.resolve(error.response)
            }
            if (status === 500) {
                // TODO: move error toasts here
                return Promise.resolve(error.response)
            }
            return Promise.reject(error)
        }
    )
    return apiInstance
}

export default { getApiInstance }