import axios from 'axios';
import AsyncStorage from '@react-native-async-storage/async-storage';

/* 
  USE THIS API INTERCEPTOR TO BUILD
  REQUESTS IN THE FUTURE, IT ADDS
  THE CONFIG TO THE REQUEST
*/

const api = axios.create({
  baseURL: 'http://your-api-url.com', // Replace with your API base URL
});

// Add a request interceptor to include the bearer token in all requests
api.interceptors.request.use(
  async function (config) {
    const bearerToken = await AsyncStorage.getItem('bearerToken');
    if (bearerToken) {
      config.headers.Authorization = `Bearer ${bearerToken}`;
    }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

export default api;