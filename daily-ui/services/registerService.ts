import axios from 'axios';
import { API_BASE_URL } from "@env";

const registerRequest = async (email:string, password:string) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/register`, {
      email: email,
      password: password,
    });
    console.log(API_BASE_URL)
    if (response.status === 200) {
      return response.data.token
    } else {
      throw new Error('Failed to register.');
    }
  } catch (error) {
    throw new Error('Failed to register.');
  }
};

export default registerRequest;
