export type CreateDailyRequest = {
   image?: string,
   text: string,
   isShared: boolean
};

export type DailyResponse = {
   id: string,
   text: string,
   author: string,
   keywords: string[],
   emotions: {
     anger: number,
     happiness: number,
     sadness: number,
     shock: number
   },
   image: string,
   favourites: number,
   createdAt: string,
   viewers: [],
   isShared: boolean
};

export type EditDailyImageRequest = {
   id: string,
   image: string
}; 

export type ReportDailyRequest = {
   content: string,
   dailyId: string,
   id: string,
   reportedAt: number,
   reports: number,
   title: string
};