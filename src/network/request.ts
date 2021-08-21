import axios, { AxiosInstance } from 'axios'
import Constants from '../constants'
import storage from '../data/storage';

const API_URL = Constants.API_URL

const apiInstance: AxiosInstance = axios.create({
    baseURL: API_URL,
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

export default { apiInstance}