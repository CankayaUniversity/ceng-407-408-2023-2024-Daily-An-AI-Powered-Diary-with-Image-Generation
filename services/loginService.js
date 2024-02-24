import axios from 'axios';
import { API_BASE_URL } from "@env";

const loginRequest = async (email, password) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/login`, {
      email: email,
      password: password,
    });

    if (response.status === 200) {
      return response.data.token
    } else {
      throw new Error('Failed to login. Please check your credentials.');
    }
  } catch (error) {
    throw new Error('Failed to login. Please check your credentials.');
  }
};

export default loginRequest;
