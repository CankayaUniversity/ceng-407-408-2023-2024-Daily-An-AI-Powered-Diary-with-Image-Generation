import { StyleSheet, View, Text, Pressable, Image, ScrollView,TouchableOpacity } from 'react-native'
import React, { useState } from 'react'
import Header from '../components/Header'
import { DailyResponse } from '../libs'

export default function ReadDaily({ route, navigation }: { route: any, navigation: any }) {
   const data: DailyResponse = route.params.data
   const [isVisible, setVisible] = useState(true)
   let maxEmotionValue = -Infinity;
   let maxEmotion = "";

   for (const [emotion, value] of Object.entries(data.emotions)) {
      if (value > maxEmotionValue) {
         maxEmotion = emotion;
      }
   }
   return (
      <Header navigation={navigation} previous="YourDaily" homepage={false}>
         <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
            <Pressable onPress={() => {setVisible(!isVisible);console.log(data)}} style={{ height: '100%', width: '100%', opacity: 1.0, backgroundColor: '#0D1326' }}>
               <ScrollView contentContainerStyle={{ flexGrow: 1, paddingBottom: 100 }}>
                  {
                     isVisible &&
                     <View>
                        <Image source={{ uri: data.image }} style={styles.image}></Image>
                        <TouchableOpacity style={{position:"absolute",bottom:0,left:5,borderWidth:1,alignItems:"center",justifyContent:"center", aspectRatio: 1 / 1, width: '35%', opacity: 0.95, marginTop: 10, borderRadius: 10, backgroundColor: '#2D1C40' }}>
                           <Text style={styles.cardText}>{maxEmotion.toUpperCase()}</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={{position:"absolute",bottom:0,right:5,alignItems:"center",justifyContent:"center",borderWidth:1, aspectRatio: 1 / 1, width: '35%', opacity: 0.95, marginTop: 10, borderRadius: 10, backgroundColor: '#2D1C40' }}>
                           <Text style={styles.cardText}>{"Topic".toUpperCase()}</Text>
                        </TouchableOpacity>
                     </View>
                  }
                  {
                     !isVisible &&
                     <Text style={styles.text}>{data.text}</Text>
                  }
               </ScrollView>
            </Pressable>
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
      paddingRight: 10,
      paddingBottom: 10,
      marginEnd: 10,
      marginTop: 10,
      fontSize: 25,
      fontWeight: '200',
      color: 'white'
   },
   cardText: {
      textAlign: 'center',
      fontSize: 25,
      fontWeight: '200',
      color: 'white'
   },
   image: {
      resizeMode: 'contain',
      width: '100%',
      height: '100%',
   },
});
