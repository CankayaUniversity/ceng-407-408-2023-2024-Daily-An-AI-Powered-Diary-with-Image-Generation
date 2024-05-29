import React, { useState, useEffect, useRef } from 'react';
import { Switch, Button, Modal, View, StyleSheet, Text, TextInput, Image, Pressable, Alert, Platform, Keyboard, KeyboardAvoidingView, TouchableWithoutFeedback } from 'react-native';
import Header from '../components/Header';
import { useCreateDaily } from '../libs';

const WriteADaily = ({ navigation }: { navigation: any }) => {
   const [daily, setDaily] = useState('');
   const [shared, setShared] = useState(false);
   const [isModalVisible, setModalVisible] = useState(false);
   const [promptText, setPromptText] = useState("");
   const { mutate, isPending } = useCreateDaily(navigation);

   const handleDailyChange = (text: string) => {
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

      mutate({ text: daily, isShared: shared, prompt: promptText })
      console.log("Daily: " + daily + " Shared: " + shared);
      setModalVisible(!isModalVisible);
   };

   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <KeyboardAvoidingView style={styles.container} behavior={Platform.OS === 'ios' ? 'padding' : 'height'} keyboardVerticalOffset={80}>
            {
               isPending == true &&
               <Text style={{ alignItems: 'center', justifyContent: 'center', fontSize: 40, color: 'white' }}>Loading</Text>
            }

            <Pressable onPress={Keyboard.dismiss} style={{ flex: 1, alignItems: "center", height: "100%", width: "100%" }}>
               <Text style={{ fontSize: 40, fontWeight: '200', color: 'white', paddingBottom: 12 }}>Write a Daily</Text>
               <TextInput
                  multiline
                  style={styles.input}
                  onChangeText={handleDailyChange}
                  value={daily}
                  placeholder="Tell us about your day"
                  inputMode="text"
                  placeholderTextColor="white"
               />
               <View style={styles.save}>
                  {isPending == false &&
                     <Pressable style={{ backgroundColor: '#6A51BE', width: "100%", borderRadius: 12, paddingVertical: 10 }} onPress={() => { setModalVisible(!isModalVisible); console.log(isPending); }}>
                        <Text style={{ color: 'white', fontSize: 24, fontWeight: "200", textAlign: "center" }}>Submit</Text>
                     </Pressable>
                  }
               </View>
            </Pressable>
         </KeyboardAvoidingView>
         <Modal
            visible={isModalVisible}
            transparent={true}
            animationType="slide"
            onRequestClose={() => setModalVisible(false)}
         >
            <View style={styles.modalOverlay}>
               <TouchableWithoutFeedback onPress={Keyboard.dismiss} style={{ height: "100%", width: "100%" }}>
                  <View style={styles.modalContent}>
                     <TextInput
                        style={styles.textInput}
                        placeholder="Do you want to add something?"
                        placeholderTextColor="#999"
                        value={promptText}
                        onChangeText={setPromptText}
                        enablesReturnKeyAutomatically
                        multiline={true}
                     />
                     <View style={styles.save}>
                        <Text style={{ fontSize: 14 }}>Do you want to share?</Text>
                        <Switch

                           onValueChange={handleShared}
                           trackColor={{ false: "#767577", true: "#6A51BE" }}
                           thumbColor={shared ? "#white" : "#f4f3f4"}
                           value={shared}
                        />
                     </View>
                     <View style={styles.modalButtons}>
                        <Button title="Cancel" onPress={() => setModalVisible(!isModalVisible)} />
                        <Button title="Submit" onPress={() => { handleDailySubmit(); }} />
                     </View>
                  </View>
               </TouchableWithoutFeedback>
            </View>
         </Modal>
      </Header>
   );
}

const styles = StyleSheet.create({
   container: {
      flex: 1,
      height: "100%",
      marginTop: 20,
      marginBottom: 100,
      marginHorizontal: 20,
      justifyContent: "flex-start",
      alignItems: 'center',
      color: "white"
   },

   input: {
      padding: 10,
      marginBottom: 20,
      color: "white",
      height: "80%",
      width: "100%",
      borderRadius: 10,
      backgroundColor: "#0D1326",
      opacity: 0.75,
      fontSize: 20,
      fontWeight: "200"
   },

   tickIcon: {
      width: 36,
      height: 36
   },

   save: {
      width: "100%",
      flexDirection: "row",
      justifyContent: "space-between",
      alignItems: "center",
      marginBottom: 10
   },
   modalOverlay: {
      flex: 1,
      backgroundColor: 'rgba(0, 0, 0, 0.5)',
      justifyContent: 'center',
      alignItems: 'center'
   },
   modalContent: {
      width: '80%',
      backgroundColor: 'white',
      borderRadius: 10,
      padding: 20,
      alignItems: 'center'
   },
   modalTitle: {
      fontSize: 20,
      marginBottom: 20
   },
   textInput: {
      width: '100%',
      height: 100,
      borderColor: '#ccc',
      borderWidth: 1,
      borderRadius: 5,
      padding: 10,
      marginBottom: 20,
      textAlignVertical: 'top'
   },
   modalButtons: {
      flexDirection: 'row',
      justifyContent: 'space-between',
      width: '100%'
   }
});

export default WriteADaily;
