import React from 'react';
import {View, StyleSheet,TouchableOpacity,Text,ImageBackground} from 'react-native';

const Header = ({navigation, children, previous, homepage}) => {
   return (
      <View style={styles.container}>
         <View style={styles.defaultpage}>
            <View style={styles.background}>
               <View style={{flexDirection:'row',borderWidth:2,justifyContent:'space-between'}}>
                  {
                     homepage &&
                     <TouchableOpacity>
                        <Text style={{marginStart:10,fontSize:60,fontWeight:'400',color:'white'}}>≡</Text>
                     </TouchableOpacity>
                  }
                  {
                     !homepage &&
                     <TouchableOpacity onPress={()=>navigation.navigate(previous)}>
                        <Text style={{marginStart:10,fontSize:60,fontWeight:'400',color:'white'}}>←</Text>
                     </TouchableOpacity>
                  }

                  <TouchableOpacity onPress={()=>Alert.alert("selam")}>
                     <Text style={{marginStart:10,fontSize:60,fontWeight:'400',color:'white'}}>d</Text>
                  </TouchableOpacity>
               </View>
               <ImageBackground source={require('../assets/background-main.png')} resizeMode="cover" imageStyle={{borderTopLeftRadius:16,borderTopRightRadius:16}} style={{marginTop:60,height: '100%',width: '100%',position: 'absolute',paddingBottom:20}}>
               {children}
               </ImageBackground>
            </View>
      </View>
      </View>
   );
}

const styles = StyleSheet.create({
   container: {
     flex: 1,
     alignItems: 'center',
   },
   defaultpage: {
     maxWidth: 500,
     height: '100%',
     width: '100%',
     alignItems: 'center',
     padding: 20,
     // Additional homepage styles...
   },
   background: {
     position: 'absolute',
     top: 0,
     left: 0,
     right: 0,
     bottom: 0,
     backgroundColor: 'black', // Example background color
   },
 });

export default Header;
