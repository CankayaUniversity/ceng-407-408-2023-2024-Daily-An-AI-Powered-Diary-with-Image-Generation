import axios from 'axios';

const loginRequest = async (email, password) => {
  try {
    const response = await axios.post('http://localhost:9090/api/login', {
      email: email,
      password: password,
    });

    // Assuming that your backend returns a status code of 200 for successful login
    if (response.status === 200) {
      // Save the bearer token to AsyncStorage
      return response.data.token
    } else {
      throw new Error('Failed to login. Please check your credentials.');
    }
  } catch (error) {
    throw new Error('Failed to login. Please check your credentials.');
  }
};

export default loginRequest;
