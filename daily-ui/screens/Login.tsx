import React, { useState } from 'react';
import { StatusBar } from 'expo-status-bar';
import { Image, StyleSheet, Text, View, TextInput, TouchableOpacity, Pressable } from 'react-native';
import Header from '../components/Header';
import { UserInfo, login } from '../libs';

const Login = ({ navigation }: { navigation: any }) => {
  const [email, setEmail] = useState(''); const [password, setPassword] = useState('');

  const handleLogin = async () => {
    try {
      const userInfo: UserInfo = {
        email: email,
        password: password
      }
      console.log(password);
      await login(userInfo);
      navigation.navigate('Home');
    } catch (error: any) {
      console.log(error.message)
    }
  };

  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <View style={styles.container}>
        <Image style={styles.logo} source={require("../assets/main-logo-big.png")}>
        </Image>
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
        <Pressable
          style={({ pressed }) => [
            styles.loginBtn,
            { opacity: pressed ? 0.5 : 1 }
          ]}
          onPress={handleLogin}>
          <Text style={styles.loginText}>Log in</Text>
        </Pressable>
        <StatusBar style="auto" />
      </View>
    </Header >

  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  logo: {
    marginTop: 160,
    fontSize: 50,
    marginBottom: 40,
  },
  inputView: {
    width: '80%',
    backgroundColor: '#ffffff',
    borderRadius: 15,
    height: 50,
    marginBottom: 15,
    justifyContent: 'center',
    padding: 20,
  },
  inputText: {
    height: 50,
    color: 'black',
  },
  loginBtn: {
    borderWidth: 1,
    backgroundColor: '#6A51BE',
    paddingHorizontal: 10,
    paddingVertical: 10,
    width: '80%',
    opacity: 0.8,
    borderRadius: 15,
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 10,
    borderColor: '#6A51BE',
    border: 10
  },
  loginText: {
    color: 'white',
    fontWeight: 'bold',
    fontSize: 20,
  },
});

export default Login;
