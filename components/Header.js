import React from 'react';
import {View, StyleSheet,TouchableOpacity,Text,ImageBackground} from 'react-native';

const Header = ({navigation, children, previous, homepage}) => {
   return (
      <View style={styles.container}>
         <View style={styles.background}>
            <View style={styles.defaultpage}>
               <View style={{flexDirection:'row',borderWidth:2,justifyContent:'space-between'}}>
                  {
                     homepage &&
                     <TouchableOpacity>
                        <Text style={{fontSize:60,fontWeight:'400',color:'white'}}>≡</Text>
                     </TouchableOpacity>
                  }
                  {
                     !homepage &&
                     <TouchableOpacity onPress={()=>navigation.navigate(previous)}>
                        <Text style={{fontSize:60,fontWeight:'400',color:'white'}}>←</Text>
                     </TouchableOpacity>
                  }

                  <TouchableOpacity onPress={()=>navigation.navigate("Home")}>
                     <Text style={{fontSize:60,fontWeight:'400',color:'white'}}>d</Text>
                  </TouchableOpacity>
               </View>
               <View style={{height:'100%', justifyContent: 'center'}}>
                  <ImageBackground source={require('../assets/background-main.png')} resizeMode="cover" imageStyle={{borderTopLeftRadius:16,borderTopRightRadius:16}} style={{height: '100%', width: '100%',position: 'absolute'}}>
                  {children}
                  </ImageBackground>
               </View>
            </View>
         </View>
      </View>
   );
}

const styles = StyleSheet.create({
   container: {
     flex: 1,
     width: '100%',
     height: '100%',
   },
   defaultpage: {
     maxWidth: 500,
     height: '100%',
     width: '100%',
     // Additional homepage styles...
   },
   background: {
     alignContent: 'center',
     alignItems: 'center',
     width: '100%',
     position: 'fixed',
     top: 0,
     left: 0,
     right: 0,
     bottom: 0,
     backgroundColor: 'black', // Example background color
   },
 });

export default Header;
