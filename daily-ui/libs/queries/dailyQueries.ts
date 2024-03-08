import { useMutation, useQuery } from "@tanstack/react-query"
import {queryClient} from "."
import {CreateDailyRequest, DailyResponse, EditDailyImageRequest, createDaily, deleteDaily, editDailyImage, favDaily, getDailies, getDaily, viewDaily} from ".."

export const dailyQueryKeys = {
   createDaily:'#daily/createDaily',
   getDaily:'#daily/getDaily',
   getDailies: '#daily/getDailies',
   favDaily:'#daily/favDaily',
   viewDaily:'#daily/viewDaily',
   editDailyImage:'#daily/editDailyImage',
   reportDaily:'#daily/reportDaily',
   deleteDaily:'#daily/deleteDaily'
}

export const useGetDailies = () =>{
   return useQuery({
      queryKey:[dailyQueryKeys.getDailies],
      queryFn: ({signal}) => getDailies(signal),
   })
}

export const useGetDaily = (
   id:string
) =>{
   return useQuery({
      queryKey:[dailyQueryKeys.getDaily],
      queryFn: ({signal}) => getDaily(id,signal),
   })
}

export const useCreateDaily = () =>{
   return useMutation({
      mutationFn:(daily:CreateDailyRequest)=>createDaily(daily),
      onSuccess: () => {
        queryClient.invalidateQueries({queryKey:[dailyQueryKeys.getDailies,dailyQueryKeys.getDaily]});
      },
    })
}

export const useFavDaily = () =>{
   return useMutation({
      mutationFn:(id:string)=>favDaily(id)
    })
}

export const useViewDaily = () =>{
   return useMutation({
      mutationFn:(id:string)=>viewDaily(id)
    })
}

export const useDeleteDaily = () =>{
   return useMutation({
      mutationFn:(id:string)=>deleteDaily(id),
      onSuccess: () => {
        queryClient.invalidateQueries({queryKey:[dailyQueryKeys.getDailies,dailyQueryKeys.getDaily]});
      },
    })
}

export const useEditDailyImage = () =>{
   return useMutation({
      mutationFn:(editDaily:EditDailyImageRequest)=>editDailyImage(editDaily),
      onSuccess: () => {
        queryClient.invalidateQueries({queryKey:[dailyQueryKeys.getDailies,dailyQueryKeys.getDaily]});
      },
    })
}