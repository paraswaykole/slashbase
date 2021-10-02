import axios, { AxiosInstance } from 'axios'
import Constants, { GetAPIConfig } from '../constants'
import storage from '../data/storage';


const getApiInstance = () => {
    const apiInstance: AxiosInstance = axios.create({
        baseURL: GetAPIConfig().API_URL,
        headers: {'content-type': 'text/json'},
        withCredentials: true,
    })
    apiInstance.interceptors.request.use(async function (config) {
        const token = await storage.getCurrentUserToken()
        if(token){
            config.headers['Authorization'] = 'Bearer '+token
        }
        return config
      }, function (error) {
        return Promise.reject(error)
    })
    return apiInstance
}

export default { getApiInstance}