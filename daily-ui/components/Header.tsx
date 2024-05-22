import React from 'react';
import { View, StyleSheet, TouchableOpacity, Image, Text, ImageBackground } from 'react-native';

const Header = ({ navigation, children, previous, homepage }: { navigation: any, children: any, previous: any, homepage: any }) => {
   return (
      <View style={styles.container}>
         <View style={styles.background}>
            <View style={styles.defaultpage}>
               <View style={{ flexDirection: 'row', borderWidth: 2, justifyContent: 'space-between', alignItems: 'center', paddingBottom: 8}}>
                  {
                     homepage &&
                     <TouchableOpacity style={{ paddingTop: 30, paddingLeft: 15 }}>
                        <Image source={require("../assets/menu.png")}></Image>
                     </TouchableOpacity>
                  }
                  {
                     !homepage &&
                     <TouchableOpacity style={{ paddingTop: 30, paddingLeft: 15 }} onPress={() => navigation.navigate(previous)}>
                        <Image source={require("../assets/back.png")}></Image>
                     </TouchableOpacity>
                  }

                  <TouchableOpacity style={{ paddingTop: 30, paddingLeft: 15 }} onPress={() => navigation.navigate("Login")}>
                     <Text style={{ fontSize: 20, fontWeight: '400', color: 'white' }}>logout</Text>
                  </TouchableOpacity>
                  <TouchableOpacity style={{ paddingTop: 30, paddingRight: 15 }} onPress={() => navigation.navigate("Home")}>
                     <Image source={require("../assets/main-logo-small.png")}></Image>
                  </TouchableOpacity>
               </View>
               <View style={{ height: '100%', justifyContent: 'center' }}>
                  <ImageBackground source={require('../assets/background-main.png')} resizeMode="cover" imageStyle={{ borderTopLeftRadius: 16, borderTopRightRadius: 16 }} style={{ height: '100%', width: '100%', position: 'absolute' }}>
                     {children}
                  </ImageBackground>
               </View>
            </View>
         </View >
      </View >
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
      position: 'absolute',
      top: 0,
      left: 0,
      right: 0,
      bottom: 0,
      backgroundColor: 'black', // Example background color
   },
});

export default Header;
