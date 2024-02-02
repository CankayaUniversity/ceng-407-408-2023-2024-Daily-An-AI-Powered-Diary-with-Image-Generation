import React from 'react';
import {View, StyleSheet,Text} from 'react-native';
import Header from '../components/Header';

const WriteADaily = () => {
   return (
      <Header previous="Home" homepage={false}>
         <View style={{flex:1,alignItems:'center',justifyContent:'center'}}>
            <Text style={{fontSize:40,fontWeight:'200',color:'white'}}>Write a Daily</Text>
         </View>
      </Header>
   );
}

const styles = StyleSheet.create({})

export default WriteADaily;
