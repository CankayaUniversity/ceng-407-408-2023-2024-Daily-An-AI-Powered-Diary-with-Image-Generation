import React, { useState } from 'react';
import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View, TextInput, TouchableOpacity, Pressable } from 'react-native';
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
      await login(userInfo);
      navigation.navigate('Home');
    } catch (error: any) {
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
    backgroundColor: '#0D1326',
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
