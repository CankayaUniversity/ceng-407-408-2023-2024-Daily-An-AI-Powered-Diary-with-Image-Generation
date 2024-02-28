import React from 'react';
import { ScrollView,TouchableOpacity,Text } from 'react-native';
import Header from '../components/Header';

const arr = ["gürkan","semih","fırat","bilal","gürkan","semih","fırat","bilal","gürkan","semih","fırat","bilal","gürkan","semih","fırat","bilal"]

const YourDaily = ({navigation}:{navigation:any}) => {
   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
        <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={{ flexDirection: "row", flexWrap: 'wrap', justifyContent:'flex-start',gap:10, alignItems: 'center',marginTop:10,paddingStart:15,paddingBottom:90}}>
          {arr.map((el, index) => {
            return (
              <TouchableOpacity key={index} style={{ aspectRatio: 1 / 1, width: '30%', borderWidth: 0.5, borderColor: 'gray', opacity: 0.85, borderRadius: 10, backgroundColor: '#0D1326' }}>
                <Text>{el}</Text>
              </TouchableOpacity>
            );
          })}
        </ScrollView>
      </Header>
   );
}

export default YourDaily;