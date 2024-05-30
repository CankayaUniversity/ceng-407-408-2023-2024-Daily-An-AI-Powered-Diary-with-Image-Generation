import { StyleSheet, View, Text, Pressable, Image, ScrollView, TouchableOpacity, ImageBackground } from 'react-native'
import React, { useState } from 'react'
import Header from '../components/Header'
import { DailyResponse } from '../libs'
import { Colors } from '../libs/colors.tsx';

export default function ReadDaily({ route, navigation }: { route: any, navigation: any }) {
   const data: DailyResponse = route.params.data
   const [isVisible, setVisible] = useState(true)
   let maxEmotionValue = -Infinity;
   let maxEmotion = "";

   for (const [emotion, value] of Object.entries(data.emotions)) {
      if (value > maxEmotionValue) {
         maxEmotionValue = value;
         maxEmotion = emotion;
      }
   }
   return (
      <Header navigation={navigation} previous="YourDaily" homepage={false}>
         <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
            <ImageBackground source={require('../assets/background-main.png')} resizeMode="cover" imageStyle={{ borderTopLeftRadius: 16, borderTopRightRadius: 16 }} blurRadius={20} style={{ height: '100%', width: '100%' }}>
               <View style={{ flexGrow: 1, paddingBottom: 100 }}>
                  {
                     isVisible &&
                     <TouchableOpacity onPress={() => { setVisible(!isVisible); console.log(data) }}>
                        <Image source={{ uri: data.image }} style={styles.image}></Image>
                     </TouchableOpacity>
                  }
                  {
                     !isVisible &&
                     <View style={{ height: '100%', width: '100%', justifyContent: 'center', alignItems: 'center', opacity: 0.7 }}>
                        <View style={{ height: '75%', width: '85%', borderRadius: 10, borderWidth: 1, backgroundColor: "black", marginBottom: 10 }}>
                           <ScrollView showsVerticalScrollIndicator={false} scrollEnabled>
                              <Text style={styles.text}>{data.text}</Text>
                           </ScrollView>
                        </View>
                        <Pressable onPress={() => setVisible(!isVisible)} style={{ height: '20%', width: '85%', flexDirection: 'row', justifyContent: 'space-between' }}>
                           <View style={{ justifyContent: "space-between", padding: 10, borderWidth: 1, height: '100%', width: '49%', opacity: 1, marginTop: 5, borderRadius: 10, backgroundColor: "black" }}>
                              <Text style={styles.cardText}>{"MOOD"}</Text>
                              <Text style={styles.cardText}>{maxEmotion.toUpperCase()}</Text>
                           </View>
                           <View style={{ justifyContent: "space-between", padding: 10, borderWidth: 1, height: '100%', width: '49%', opacity: 1, marginTop: 5, borderRadius: 10, backgroundColor: "black" }}>
                              <Text style={styles.cardText}>{"TOPIC".toUpperCase()}</Text>
                              <Text style={styles.cardText}>{data.Topics[0].toUpperCase()}</Text>
                           </View>
                        </Pressable>
                     </View>
                  }
               </View>
            </ImageBackground>
         </View>
      </Header >
   )
}
const styles = StyleSheet.create({
   container: {
      flex: 1,
      flexDirection: 'column',
   },
   text: {
      textAlign: 'left',
      paddingLeft: 10,

      paddingBottom: 10,
      marginEnd: 10,
      marginTop: 10,
      fontSize: 25,
      fontWeight: '200',
      color: 'white'
   },
   cardText: {
      fontSize: 25,
      fontWeight: '200',
      color: 'white',
   },
   image: {
      resizeMode: 'contain',
      width: '100%',
      height: '100%',
   },
});
