import axios from "axios";
import { API_BASE_URL } from "@env";
import { storageUtils,storageKeys, login } from "..";
import { UserInfo, UserToken } from "..";

const defaultRequestTimeout = 30000;

export const serviceConsumer = axios.create({
   baseURL: API_BASE_URL,
   timeout: defaultRequestTimeout,
   headers:{
      "Content-Type":"application/json",
   },
   responseType: 'json'
})

serviceConsumer.interceptors.request.use(
   async (config) => {
      if(!config.url?.includes("login") && !config.url?.includes("register")){
         const userToken = await storageUtils.getItem<UserToken>(storageKeys.bearerToken);
         if (userToken){
            config.headers.Authorization = "Bearer " + userToken.token;
         }
      }
      return config;
   },
   (error) => {
      return Promise.reject(error);
   }
);

serviceConsumer.interceptors.response.use(
   (response) =>{
      return response;
   },
   async (error) =>{
      const userInfo = await storageUtils.getItem<UserInfo>(storageKeys.userInfo);
      const status = error.response?.status;
      const request = error.config;
      if (request._retry) {
         return Promise.reject(error);
      }
      if(status == 401){
         if(userInfo){
            request._retry = true;
            await login(userInfo);
            return serviceConsumer(request);
         }
      }
      return Promise.reject(error);
   }
)
