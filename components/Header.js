import React from 'react';
import {View, StyleSheet,TouchableOpacity,Text,ImageBackground} from 'react-native';

const Header = ({children, previous, homepage}) => {
   return (
      <View style={{flex:1}}>
         <View style={{height:'100%',backgroundColor:'black'}}>
            <View style={{flexDirection:'row',borderWidth:2,justifyContent:'space-between'}}>
               {
                  homepage &&
                  <TouchableOpacity>
                     <Text style={{marginStart:10,fontSize:60,fontWeight:'400',color:'white'}}>←</Text>
                  </TouchableOpacity>
               }
               {
                  !homepage &&
                  <TouchableOpacity onPress={()=>navigation.navigate(previous)}>
                     <Text style={{marginStart:10,fontSize:60,fontWeight:'400',color:'white'}}>≡</Text>
                  </TouchableOpacity>
               }

               <TouchableOpacity onPress={()=>Alert.alert("selam")}>
                  <Text style={{marginStart:10,fontSize:60,fontWeight:'400',color:'white'}}>d</Text>
               </TouchableOpacity>
            </View>
            <ImageBackground source={require('../assets/test.jpg')} resizeMode="cover" imageStyle={{borderTopLeftRadius:16,borderTopRightRadius:16}} style={{marginTop:60,height: '100%',width: '100%',position: 'absolute',paddingBottom:20}}>
            {children}
            </ImageBackground>
         </View>
      </View>
   );
}

const styles = StyleSheet.create({})

export default Header;
