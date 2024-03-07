import React, { useState, useEffect, useRef } from 'react';
import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View, TextInput, TouchableOpacity, Pressable } from 'react-native';
import Header from '../components/Header';
import { UserInfo, register } from '../libs';

const Register = ({ navigation }:{navigation:any}) => {
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');

  const handleRegister = async () => {
    try {
      const userInfo:UserInfo = {
        email: email,
        password: password
      }
      const response = await register(userInfo);
      navigation.navigate('Login');
    } catch (error:any) {
      navigation.navigate('Register');
      console.log(error.message)
    } 
  };

  return (
    <Header navigation={navigation} previous="Register" homepage={false}>
    <View style={styles.container}>
      <Text style={styles.logo}>daily</Text>
      <View style={styles.inputView}>
        <TextInput
          style={styles.inputText}
          placeholder="email@example.com"
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
      <Pressable style={styles.registerBtn} onPress={handleRegister}>
        <Text style={styles.buttonText}>Register</Text>
      </Pressable>
      <Pressable style={styles.alreadyHaveAccount} onPress={() => navigation.navigate("Login")}>
        <Text style={styles.altText}>Already have an account? Login</Text>
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
  registerBtn: {
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
  buttonText: {
    color: 'white',
    fontSize: 20,
  },
  alreadyHaveAccount: {
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 15,
    color: 'white',
  },

  altText: {
    color: 'white',
    fontSize: 15,
  },
});

export default Register;
