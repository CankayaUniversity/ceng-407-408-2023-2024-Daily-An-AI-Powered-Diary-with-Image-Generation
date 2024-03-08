import { AxiosResponse } from 'axios';
import { CommonResponse, CreateDailyRequest, DailyResponse, EditDailyImageRequest, ReportDailyRequest,serviceConsumer} from '..';

export const createDaily = async (
   daily: CreateDailyRequest,
) => {
   const url = "/daily";
   const response = await serviceConsumer.post<
   DailyResponse,
   AxiosResponse<DailyResponse>,
   CreateDailyRequest>(url,daily);
   return response.data;
}

export const getDailies = async (
   signal?:AbortSignal
) => {
   const url = "/daily/list";
   const response = await serviceConsumer.get<
   DailyResponse[],
   AxiosResponse<DailyResponse[]>>(url,{signal});
   return response.data;
}

export const favDaily = async (
   id: string,
) => {
   const url = `/daily/fav/${id}`;
   const response = await serviceConsumer.put<
   CommonResponse,
   AxiosResponse<CommonResponse>>(url);
   return response.data;
}

export const viewDaily = async (
   id: string,
) => {
   const url = `/daily/view/${id}`;
   const response = await serviceConsumer.put<
   CommonResponse,
   AxiosResponse<CommonResponse>>(url);
   return response.data;
}

export const editDailyImage = async (
   daily: EditDailyImageRequest,
) => {
   const url = "/daily/image";
   const response = await serviceConsumer.put<
   CommonResponse,
   AxiosResponse<CommonResponse>,
   EditDailyImageRequest>(url,daily);
   return response.data;
}

export const reportDaily = async (
   dailyReport: ReportDailyRequest,
) => {
      const url = "/daily/report";
      const response = await serviceConsumer.post<
      CommonResponse,
      AxiosResponse<CommonResponse>,
      ReportDailyRequest>(url,dailyReport);
      return response.data;
}

export const getDaily = async (
   id: string,
   signal?:AbortSignal
) => {
   const url = `/daily/${id}`;
   const response = await serviceConsumer.get<
   DailyResponse,
   AxiosResponse<DailyResponse>>(url,{signal});
   return response.data;
}

export const deleteDaily = async (
   id: string,
) => {
   const url = `/daily/${id}`;
   const response = await serviceConsumer.delete<
   CommonResponse,
   AxiosResponse<CommonResponse>>(url);
   return response.data;
}
