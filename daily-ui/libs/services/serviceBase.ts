import axios from "axios";
import { API_BASE_URL } from "@env";
import { storageUtils, storageKeys, login } from "..";
import { UserInfo, UserToken } from "..";

const defaultRequestTimeout = 1200000;

export const serviceConsumer = axios.create({
   baseURL: API_BASE_URL,
   timeout: defaultRequestTimeout,
   headers: {
      "Content-Type": "application/json",
   },
   responseType: 'json'
})

serviceConsumer.interceptors.request.use(
   async (config) => {
      if (!config.url?.includes("login") && !config.url?.includes("register")) {
         const userToken = await storageUtils.getItem<UserToken>(storageKeys.bearerToken);
         if (userToken) {
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
   (response) => {
      return response;
   },
   async (error) => {
      const status = error.response?.status;
      const request = error.config;
      // Check for 401 Unauthorized and if we haven't retried yet
      if (status === 401 && !request._retry) {
         const userInfo = await storageUtils.getItem<UserInfo>(storageKeys.userInfo);
         if (userInfo) {
            // Mark that we're attempting a retry
            request._retry = true;
            // Attempt to re-login and retry the request
            try {
               await login(userInfo);
               return serviceConsumer(request);
            } catch (loginError) {
               // Login attempt failed, set a flag in storage indicating navigation to login is needed
               await storageUtils.clear();
               request._retry = false;
               return Promise.reject(error);
            }
         }
      } else if (request._retry) {
         // If we've already retried, indicate a need for navigation to login
         request._retry = false;
         await storageUtils.clear();
      }
      return Promise.reject(error);
   }
);

