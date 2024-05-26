export type CreateDailyRequest = {
   image?: string,
   text: string,
   isShared: boolean
};

export type StatisticsResponse = {
   date: string[],
   likes: number,
   views: number,
   dailiesWritten: number,
   mood: string,
   streak: number,
   topic: string
};

export type ExploreResponse = {
   id: string,
   text: string,
   image: string,
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
   isshared: boolean
};

export type EditDailyImageRequest = {
   id: string,
   image: string
};

export type ReportDailyRequest = {
   content: string,
   dailyId: string,
   title: string
};
