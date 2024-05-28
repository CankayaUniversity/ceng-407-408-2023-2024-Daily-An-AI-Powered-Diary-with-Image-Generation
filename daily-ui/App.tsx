import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import Home from './screens/Home';
import YourDaily from './screens/YourDaily';
import Explore from './screens/Explore';
import Explore2 from './screens/Explore2';
import Profile from './screens/Profile';
import Statistics from './screens/Statistics';
import WriteADaily from './screens/WriteADaily';
import Login from './screens/Login';
import Register from './screens/Register';
import React from 'react'; import AuthLoadingScreen from './screens/AuthLoadingScreen';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { queryClient } from './libs';
import ReadDaily from './screens/ReadDaily';

const Stack = createNativeStackNavigator();


const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <NavigationContainer>
        <Stack.Navigator initialRouteName="AuthLoadingScreen" screenOptions={{ headerShown: false }}>
          <Stack.Screen name="AuthLoadingScreen" component={AuthLoadingScreen} />
          <Stack.Screen name="Home" component={Home} />
          <Stack.Screen name="YourDaily" component={YourDaily} />
          <Stack.Screen name="ReadDaily" component={ReadDaily} />
          <Stack.Screen name="Explore" component={Explore2} />
          <Stack.Screen name="Profile" component={Profile} />
          <Stack.Screen name="Statistics" component={Statistics} />
          <Stack.Screen name="WriteADaily" component={WriteADaily} />
          <Stack.Screen name="Login" component={Login} options={{gestureEnabled: false}}/>
          <Stack.Screen name="Register" component={Register}/>
        </Stack.Navigator>
      </NavigationContainer>
    </QueryClientProvider>
  );
}

export default App;
