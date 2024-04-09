import { AxiosResponse } from 'axios';
import { CommonResponse, UserInfo, UserToken, serviceConsumer, storageKeys, storageUtils } from '..';

export const login = async (
   userInfo: UserInfo,
   signal?: AbortSignal
) => {
   const url = "/login";
   const response = await serviceConsumer.post<
      UserToken,
      AxiosResponse<UserToken>,
      UserInfo>(url, userInfo, { signal });
   if (response.data.token) {
      storageUtils.clear(); // clear the cache before setting a new token
      await storageUtils.setItem(storageKeys.bearerToken, response.data);
   }
   return response;
}

export const register = async (
   userInfo: UserInfo,
   signal?: AbortSignal
) => {
   const url = "/register";
   const response = await serviceConsumer.post<
      CommonResponse,
      AxiosResponse<CommonResponse>,
      UserInfo>(url, userInfo, { signal });
   return response;
}
