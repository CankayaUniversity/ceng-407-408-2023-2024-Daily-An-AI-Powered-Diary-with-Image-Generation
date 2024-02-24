import React, { useState, useEffect, useRef } from 'react';
import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View, TextInput, TouchableOpacity, Pressable } from 'react-native';
import Header from '../components/Header';
import loginRequest from '../services/loginService';
import AsyncStorage from '@react-native-async-storage/async-storage';

const Login = ({ navigation }) => {
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [error, setError] = React.useState('');

  const handleLogin = async () => {
    // login logic here, now it just navigates to Home
    try {
        const bearerToken = await loginRequest(email, password); // Call loginRequest function from the service
        console.log(bearerToken)

        await AsyncStorage.setItem('bearerToken', bearerToken)

        navigation.navigate('Home');
    } catch (error) {
      navigation.navigate('Login');
      setError(error.message);
      console.log(error.message)
    } 
  };

  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
    <View style={styles.container}>
      <Text style={styles.logo}>daily</Text>
      <View style={styles.inputView}>
        <TextInput
          style={styles.inputText}
          placeholder="username@mailprovider.com"
          placeholderTextColor="#003f5c"
          onChangeText={(text) => setEmail(text)}
          value={email}
        />
      </View>
      <View style={styles.inputView}>
        <TextInput
          style={styles.inputText}
          placeholder="Password"
          placeholderTextColor="#003f5c"
          onChangeText={(text) => setPassword(text)}
          value={password}
          secureTextEntry={true}
        />
      </View>
      <Pressable style={styles.loginBtn} onPress={handleLogin}>
        <Text style={styles.loginText}>LOGIN</Text>
      </Pressable>
      <StatusBar style="auto" />
    </View>
    </Header>

  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  logo: {
    color: "white",
    fontWeight: 'bold',
    fontSize: 50,
    marginBottom: 40,
  },
  inputView: {
    width: '80%',
    backgroundColor: '#ffffff',
    borderRadius: 25,
    height: 50,
    marginBottom: 20,
    justifyContent: 'center',
    padding: 20,
  },
  inputText: {
    height: 50,
    color: 'black',
  },
  loginBtn: {
    borderWidth: 1,
    backgroundColor:'#0D1326',
    paddingHorizontal: 10,
    paddingVertical: 10,  
    opacity: 0.8,
    borderRadius: 30,
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 10,
    borderColor: "white",
    border: 10
  },
  loginText: {
    color: 'white',
    fontSize: 20,
  },
});

export default Login;
