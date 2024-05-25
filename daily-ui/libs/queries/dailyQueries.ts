import { useMutation, useQuery, useInfiniteQuery } from "@tanstack/react-query"
import { queryClient } from "."
import { CreateDailyRequest, getStatistics, ReportDailyRequest, EditDailyImageRequest, createDaily, deleteDaily, editDailyImage, favDaily, getDailies, getDaily, getExplore, reportDaily, viewDaily } from ".."
import { Alert } from "react-native"

export const dailyQueryKeys = {
   createDaily: '#daily/createDaily',
   getDaily: '#daily/getDaily',
   getDailies: '#daily/getDailies',
   favDaily: '#daily/favDaily',
   viewDaily: '#daily/viewDaily',
   editDailyImage: '#daily/editDailyImage',
   reportDaily: '#daily/reportDaily',
   deleteDaily: '#daily/deleteDaily',
   getExplore: '#daily/getExplore',
   getExploreInfinite: '#daily/getExploreInfinite',
   getStatistics: '#daily/getStatistics'
}

export const useGetExploreInfinite = () => {
   return useInfiniteQuery({
      queryKey: [dailyQueryKeys.getExploreInfinite],
      queryFn: ({ signal }) => getExplore(signal),
      initialPageParam: 1,
      getNextPageParam: () => null,
   })
}

export const useGetExplore = () => {
   return useQuery({
      queryKey: [dailyQueryKeys.getExplore],
      queryFn: ({ signal }) => getExplore(signal),
   })
}

export const useGetStatistics = () => {
   return useQuery({
      queryKey: [dailyQueryKeys.getStatistics],
      queryFn: ({ signal }) => getStatistics(signal),
   })
}

export const useGetDailies = (limit?: number) => {
   return useQuery({
      queryKey: [dailyQueryKeys.getDailies],
      queryFn: ({ signal }) => getDailies(signal, limit),
   })
}

export const useGetDaily = (
   id: string
) => {
   return useQuery({
      queryKey: [dailyQueryKeys.getDaily],
      queryFn: ({ signal }) => getDaily(id, signal),
   })
}

export const useCreateDaily = (navigation: any) => {
   return useMutation({
      mutationFn: (daily: CreateDailyRequest) => createDaily(daily),
      onError: (error) => {
         Alert.alert('Error', error.message);
      },
      onSuccess: () => {
         queryClient.invalidateQueries({ queryKey: [dailyQueryKeys.getDailies] });
         navigation.navigate("YourDaily");
      },
   })
}

export const useFavDaily = () => {
   return useMutation({
      mutationFn: (id: string) => favDaily(id)
   })
}

export const useViewDaily = () => {
   return useMutation({
      mutationFn: (id: string) => viewDaily(id)
   })
}

export const useDeleteDaily = () => {
   return useMutation({
      mutationFn: (id: string) => deleteDaily(id),
      onSuccess: () => {
         queryClient.invalidateQueries({ queryKey: [dailyQueryKeys.getDailies, dailyQueryKeys.getDaily] });
      },
   })
}

export const useEditDailyImage = () => {
   return useMutation({
      mutationFn: (editDaily: EditDailyImageRequest) => editDailyImage(editDaily),
      onSuccess: () => {
         queryClient.invalidateQueries({ queryKey: [dailyQueryKeys.getDailies, dailyQueryKeys.getDaily] });
      },
   })
}

export const useReportDaily = () => {
   return useMutation({
      mutationFn: (report: ReportDailyRequest) => reportDaily(report),
      onSuccess: () => {
         queryClient.invalidateQueries({ queryKey: [dailyQueryKeys.reportDaily] });
      },
   })
}
