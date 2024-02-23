import React, { useState, useEffect, useRef } from 'react';
import {View, StyleSheet, Text, TextInput, Image, Pressable, Alert, Platform, Keyboard, KeyboardAvoidingView, TouchableWithoutFeedback} from 'react-native';
import Header from '../components/Header';

const WriteADaily = ({navigation}) => {
   const [text, onChangeText] = useState('Useless Text');
   const [daily, setDaily] = useState('');

   const handleDailyChange = (text) => {
     setDaily(text);
   };
 
   const handleDailySubmit = () => {
     if (daily.trim() === '') {
       Alert.alert('Error', 'Please enter some text before submitting.');
       return;
     }
 
     // Here you can perform any action with the tweet, such as sending it to a server or saving it locally.
     // For demonstration, we'll just log the tweet to the console.
     Alert.alert('Daily submitted:', daily);
    };

   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <KeyboardAvoidingView style={styles.container} behavior={Platform.OS === 'ios' ? 'padding' : 'height'} keyboardVerticalOffset={80}>
            <Pressable onPress={Keyboard.dismiss} style={{flex: 1, alignItems: "center", height:"100%", width:"100%"}}>
               <Text style={{fontSize:40,fontWeight:'200',color:'white',paddingBottom: 12}}>Write a Daily</Text>
               <TextInput
                  multiline
                  style={styles.input}
                  onChangeText={handleDailyChange}
                  value={daily}
                  placeholder="Tell us about your day"
                  inputMode="text"
               />
               <View style={styles.save}>
                  <Pressable onPress={handleDailySubmit}>
                     <Image source={require("../assets/tickIcon.png")} style={styles.tickIcon}></Image>
                  </Pressable>
               </View>
            </Pressable>
         </KeyboardAvoidingView>
      </Header>
   );
}

const styles = StyleSheet.create({
   container: {
      flex:1,
      height: "100%",
      marginTop: 20,
      marginBottom: 100,
      marginHorizontal: 20,
      justifyContent: "flex-start",
      alignItems: 'center',
      color: "white" 
   },

   input: {
     borderColor: '#ccc',
     borderWidth: 1,
     padding: 10,
     marginBottom: 20,
     color: "white",
     height: "80%",
     width: "100%",
     borderRadius: 10,
     borderWidth: 0.7,
     borderColor:'white'
   },

   tickIcon: {
      width: 36,
      height: 36
   },

   save: {
      width: "90%",
      flex: 1,
      alignItems: "flex-end"
   }
 });

export default WriteADaily;
