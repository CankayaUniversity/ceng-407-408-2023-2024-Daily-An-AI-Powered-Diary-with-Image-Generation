import { AxiosResponse } from 'axios';
import { CommonResponse, CreateDailyRequest, DailyResponse, EditDailyImageRequest, ReportDailyRequest, UserInfo, UserToken, serviceConsumer, storageKeys, storageUtils } from '..';

export const createDaily = async (daily: CreateDailyRequest) => {
   const url = "/daily";
   const response = await serviceConsumer.post<
   DailyResponse,
   AxiosResponse<DailyResponse>,
   CreateDailyRequest>(url,daily);
   return response;
}

export const getDailies = async () => {
   const url = "/daily";
   const response = await serviceConsumer.get<
   DailyResponse,
   AxiosResponse<DailyResponse>>(url);
   return response;
}

export const favDaily = async (id: string) => {
   const url = `/daily/fav/${id}`;
   const response = await serviceConsumer.put<
   CommonResponse,
   AxiosResponse<CommonResponse>>(url);
   return response;
}

export const viewDaily = async (id: string) => {
   const url = `/daily/view/${id}`;
   const response = await serviceConsumer.put<
   CommonResponse,
   AxiosResponse<CommonResponse>>(url);
   return response;
}

export const editDailyImage = async (daily: EditDailyImageRequest) => {
   const url = "/daily/image/";
   const response = await serviceConsumer.put<
   CommonResponse,
   AxiosResponse<CommonResponse>,
   EditDailyImageRequest>(url,daily);
   return response;
}

export const reportDaily = async (dailyReport: ReportDailyRequest) => {
      const url = "/daily/report/";
      const response = await serviceConsumer.post<
      CommonResponse,
      AxiosResponse<CommonResponse>,
      ReportDailyRequest>(url,dailyReport);
      return response;
}

export const getDaily = async (id: string) => {
   const url = `/daily/${id}`;
   const response = await serviceConsumer.get<
   DailyResponse,
   AxiosResponse<DailyResponse>>(url);
   return response;
}

export const deleteDaily = async (id: string) => {
   const url = `/daily/${id}`;
   const response = await serviceConsumer.delete<
   CommonResponse,
   AxiosResponse<CommonResponse>>(url);
   return response;
}
