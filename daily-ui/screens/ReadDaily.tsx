import { StyleSheet, View, Text, Pressable, Image, ScrollView } from 'react-native'
import React, { useState } from 'react'
import Header from '../components/Header'
import { DailyResponse } from '../libs'

export default function ReadDaily({ route, navigation }: { route: any, navigation: any }) {
   const data: DailyResponse = route.params.data
   const [isVisible, setVisible] = useState(true)
   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
            <Pressable onPress={() => setVisible(!isVisible)} style={{ height: '100%', width: '100%', opacity: 1.0, backgroundColor: '#0D1326' }}>
               <ScrollView contentContainerStyle={{ flexGrow: 1 }}>
                  {
                     isVisible &&
                     <Image source={{ uri: data.image }} resizeMode='contain' style={{ width: '100%', height: '100%' }}></Image>
                  }
                  {
                     !isVisible &&
                     <Text style={styles.text}>{data.text}</Text>
                  }
               </ScrollView>
            </Pressable>
         </View>
      </Header>
   )
}
const styles = StyleSheet.create({
   container: {
      flex: 1,
      flexDirection: 'column',
   },
   text: {
      textAlign: 'left',
      marginEnd: 10,
      marginTop: 10,
      fontSize: 40,
      fontWeight: '200',
      color: 'white'
   }
});
