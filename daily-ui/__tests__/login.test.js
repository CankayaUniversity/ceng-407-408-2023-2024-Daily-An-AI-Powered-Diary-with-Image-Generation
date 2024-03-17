import React from 'react';
import { render, fireEvent, waitFor } from '@testing-library/react-native';
import Login from '../screens/Login';
import { login as mockLogin } from '../libs'; // Import login function from a mock or stub

jest.mock('../libs', () => ({
  login: jest.fn(),
}));

const mockNavigate = jest.fn();
const navigation = {
  navigate: jest.fn(),
};

describe('Login component', () => {
  it('renders correctly', () => {
    const { getByText, getByPlaceholderText } = render(<Login navigation={navigation} />);
    expect(getByText('daily')).toBeDefined();
    expect(getByPlaceholderText('username@mailprovider.com')).toBeDefined();
    expect(getByPlaceholderText('Password')).toBeDefined();
    expect(getByText('LOGIN')).toBeDefined();
  });

  it('calls login function with user info when login button is pressed', async () => {
    const { getByPlaceholderText, getByText } = render(<Login navigation={navigation} />);
    const emailInput = getByPlaceholderText('username@mailprovider.com');
    const passwordInput = getByPlaceholderText('Password');
    const loginButton = getByText('LOGIN');

    fireEvent.changeText(emailInput, 'test@example.com');
    fireEvent.changeText(passwordInput, 'password');
    fireEvent.press(loginButton);

    // Wait for the asynchronous operation to complete
    await waitFor(() => expect(mockLogin).toHaveBeenCalled());
    console.log('Mock login function called:', mockLogin.mock.calls);
    
    expect(getByText('LOGIN')).toBeDefined(); // Check for login button
    const emailInputValue = emailInput.props.value;
    console.log('Email input value after login:', emailInputValue);
  });

  it("doesn't call the login function if the email doesn't match a valid pattern", async () => {
    const { getByPlaceholderText, getByText } = render(<Login navigation={navigation} />);
    const emailInput = getByPlaceholderText('username@mailprovider.com');
    const passwordInput = getByPlaceholderText('Password');
    const loginButton = getByText('LOGIN');

    fireEvent.changeText(emailInput, 'fail');
    fireEvent.changeText(passwordInput, 'password');
    fireEvent.press(loginButton);

    // Wait for the asynchronous operation to complete
    await waitFor(() => expect(mockLogin).not.toHaveBeenCalled());
    console.log('Mock login function called:', mockLogin.mock.calls);
  });
});
