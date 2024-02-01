import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import Home from './screens/Home';
import YourDaily from './screens/YourDaily';
import Explore from './screens/Explore';
import Profile from './screens/Profile';
import Statistics from './screens/Statistics';
import WriteADaily from './screens/WriteADaily';

const Stack = createNativeStackNavigator();

const App= () => {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName='Home' screenOptions={{headerShown:false}}>
        <Stack.Screen name="Home" component={Home}/>
        <Stack.Screen name="YourDaily" component={YourDaily}/>
        <Stack.Screen name="Explore" component={Explore}/>
        <Stack.Screen name="Profile" component={Profile}/>
        <Stack.Screen name="Statistics" component={Statistics}/>
        <Stack.Screen name="WriteADaily" component={WriteADaily}/>
      </Stack.Navigator>
    </NavigationContainer>
  );
}

export default App;
