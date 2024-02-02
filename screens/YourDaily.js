import React from 'react';
import { View,Text } from 'react-native';
import Header from '../components/Header';

const YourDaily = ({navigation}) => {
   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <View style={{flex:1,alignItems:'center',justifyContent:'center'}}>
            <Text style={{fontSize:40,fontWeight:'200',color:'white'}}>Your Daily</Text>
         </View>
      </Header>
   );
}

export default YourDaily;
