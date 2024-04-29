import { StyleSheet, View, Text, Pressable, Image, ScrollView } from 'react-native'
import React, { useState } from 'react'
import Header from '../components/Header'
import { DailyResponse } from '../libs'

export default function ReadDaily({ route, navigation }: { route: any, navigation: any }) {
   const data: DailyResponse = route.params.data
   const [isVisible, setVisible] = useState(true)
   return (
      <Header navigation={navigation} previous="YourDaily" homepage={false}>
         <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
            <Pressable onPress={() => setVisible(!isVisible)} style={{ height: '100%', width: '100%', opacity: 1.0, backgroundColor: '#0D1326' }}>
               <ScrollView contentContainerStyle={{ flexGrow: 1, paddingBottom: 100 }}>
                  {
                     isVisible &&
                     <Image source={{ uri: data.image }} style={styles.image}></Image>
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
   image: {
      resizeMode: 'contain',
      width: '100%',
      height: '100%',
   },
});
