import AsyncStorage from "@react-native-async-storage/async-storage";

export async function setItem(key: string, value: unknown): Promise<void> {
   const data = typeof value === 'string' ? value : JSON.stringify(value);
   await AsyncStorage.setItem(key, data);
 }
 
 export async function getItem<T>(key: string): Promise<T | null> {
   const item = await AsyncStorage.getItem(key);
   if (item?.startsWith('{') || item?.startsWith('[')) {
     return JSON.parse(item || 'null') as T;
   }
   return item as T;
 }
 
 export async function removeItem(key: string): Promise<void> {
   await AsyncStorage.removeItem(key);
 }
 
 export async function clear(): Promise<void> {
   await AsyncStorage.clear();
 }