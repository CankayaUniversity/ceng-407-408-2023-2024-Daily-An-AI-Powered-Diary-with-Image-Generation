import { AxiosResponse } from 'axios';
import { CommonResponse, UserInfo, UserToken, serviceConsumer, storageKeys, storageUtils } from '..';

export const login = async (userInfo: UserInfo) => {
   const url = "/login";
   const response = await serviceConsumer.post<
   UserToken,
   AxiosResponse<UserToken>,
   UserInfo>(url,userInfo);
   if(response.data.token){
      await storageUtils.setItem(storageKeys.bearerToken,response.data);
   }
   return response;
}

export const register = async (userInfo: UserInfo) => {
   const url = "/register";
   const response = await serviceConsumer.post<
   CommonResponse,
   AxiosResponse<CommonResponse>,
   UserInfo>(url,userInfo);
   return response;
}