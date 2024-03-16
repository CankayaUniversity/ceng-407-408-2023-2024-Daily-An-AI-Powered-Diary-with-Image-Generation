import React, { useState, useEffect, useRef } from 'react';
import {Switch, View, StyleSheet, Text, TextInput, Image, Pressable, Alert, Platform, Keyboard, KeyboardAvoidingView, TouchableWithoutFeedback} from 'react-native';
import Header from '../components/Header';
import { useCreateDaily } from '../libs';
import { CreateDailyRequest } from "../libs"

const WriteADaily = ({navigation}:{navigation:any}) => {
   const [daily, setDaily] = useState('');
   const [shared, setShared] = useState(false);
   const mutation = useCreateDaily()
   
   useEffect(() => {
      if (mutation.isError) {
         const errorMessage = mutation.error instanceof Error ? mutation.error.toString() : String(mutation.error);
         Alert.alert("Please try again. Error: ", errorMessage);
         console.log(errorMessage);
      }
      if (mutation.isSuccess) {
         navigation.navigate("YourDaily");
         console.log(mutation.isSuccess);
         setDaily('');
      }
   }, [mutation])

   const handleDailyChange = (text:string) => {
      setDaily(text);
   }
   
   const handleShared = () => {
      console.log(!shared);
      setShared(!shared);
   }
 
   const handleDailySubmit = () => {
     if (daily.trim() === '') {
       Alert.alert('Error', 'Please enter some text before submitting.');
       return;
     }
 
     // Here you can perform any action with the tweet, such as sending it to a server or saving it locally.
     // For demonstration, we'll just log the tweet to the console.
     mutation.mutate({text:daily, isShared:shared})
    };

   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <KeyboardAvoidingView style={styles.container} behavior={Platform.OS === 'ios' ? 'padding' : 'height'} keyboardVerticalOffset={80}>
            {
               mutation.isPending &&
               <Text style={{ alignItems: 'center', justifyContent: 'center', fontSize: 40, color: 'white' }}>Loading</Text>
            }

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
                  <Switch 
                     onValueChange={handleShared}
                     trackColor={{ false: "#767577", true: "#81b0ff" }}
                     thumbColor={shared ? "#f5dd4b" : "#f4f3f4"}
                     value={shared}
                  />
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
   },

   tickIcon: {
      width: 36,
      height: 36
   },

   save: {
      width: "90%",
      flexDirection: "row",
      justifyContent: "space-between"
   }
 });

export default WriteADaily;
