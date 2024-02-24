import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import Home from './screens/Home';
import YourDaily from './screens/YourDaily';
import Explore from './screens/Explore';
import Profile from './screens/Profile';
import Statistics from './screens/Statistics';
import WriteADaily from './screens/WriteADaily';
import Login from './screens/Login';
import React, { useState, useEffect, useRef } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';

const Stack = createNativeStackNavigator();

const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  // Check authentication status on app load
  useEffect(() => {
    // Here you can check if the user has a bearer key stored in AsyncStorage or any other storage mechanism
    const checkAuthentication = async () => {
      // Example: Check if the user has a bearer key stored
      const bearerKey = await AsyncStorage.getItem('bearerKey');
      setIsAuthenticated(!!bearerKey); // Update isAuthenticated based on the presence of bearer key
    };

    checkAuthentication();
  }, []);


  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName='Login' screenOptions={{headerShown:false}}>
        <Stack.Screen name="Home" component={Home}/>
        <Stack.Screen name="YourDaily" component={YourDaily}/>
        <Stack.Screen name="Explore" component={Explore}/>
        <Stack.Screen name="Profile" component={Profile}/>
        <Stack.Screen name="Statistics" component={Statistics}/>
        <Stack.Screen name="WriteADaily" component={WriteADaily}/>
        <Stack.Screen name="Login" component={Login}/>
      </Stack.Navigator>
    </NavigationContainer>
  );
}

export default App;
