import React, { useState } from 'react';
import { ScrollView, Pressable, Text, Image,View, Alert} from 'react-native';
import Header from '../components/Header';
import { EditDailyImageRequest, useEditDailyImage, useGetDailies } from '../libs';
import * as ImagePicker from 'expo-image-picker';

const YourDaily = ({ navigation }: { navigation: any }) => {
  const {isLoading, data,isRefetching } = useGetDailies();
  const {isPending,mutate} = useEditDailyImage();
  const dateDict = (dateStr:string) => {
    const date = new Date(dateStr);
    const monthArr = ["JAN",	"FEB",	"MAR",	"APR",	"MAY",	"JUN",	"JUL",	"AUG",	"SEPT",	"OCT",	"NOV",	"DEC"]

    return monthArr[date.getMonth()]+"\n"+date.getDate();
  };

  const pickImage = async (id:string) => {
    // No permissions request is necessary for launching the image library
    let result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: true,
      aspect: [1, 1],
      quality: 1,
      base64:true
    });

    console.log(result);

    if (!result.canceled) {
      const editReq: EditDailyImageRequest ={
        id:id,
        image:'data:image/jpeg;base64,'+result.assets[0].base64 as string
      }
      mutate(editReq)
    }
  };


  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={{ flexDirection: "row", flexWrap: 'wrap', justifyContent: 'flex-start', gap: 10, alignItems: 'center', marginTop: 10, paddingStart: 15, paddingBottom: 90 }}>
        {
          (isLoading || isPending) &&
          <Text style={{ alignItems: 'center', justifyContent: 'center', fontSize: 40, color: 'white' }}>Loading</Text>
        }
        {(!isLoading && !isPending) && data?.map((el, index) => {
          return (
            <Pressable
            key={index}
            onPress={() => navigation.navigate("ReadDaily", { data: el })}
            onLongPress={()=>Alert.alert(
              '',
              'Do you want to change the image?',  
              [
                 {text: 'Yes', onPress: () => pickImage(el.id), style: 'cancel'},
                 {text: 'No'},
              ],
              { cancelable: false }
         )}
            style={{
              aspectRatio: 1,
              width: '30%',
              borderWidth: 0.5,
              borderColor: 'gray',
              alignItems: "center",
              justifyContent: "center",
              borderRadius: 10,
              backgroundColor: '#0D1326',
            }}
          >
            <View style={{ width: '100%', height: '100%', alignItems: 'center', justifyContent: 'center' }}>
              <Text style={{ color: "white", fontSize: 42, textAlign:"center", fontWeight: '200'}}>{dateDict(el.createdAt)}</Text>
            </View>
            <Image
              source={{ uri: el.image }}
              resizeMode='contain'
              style={{
                width: '100%',
                height: '100%',
                borderRadius: 10,
                position: "absolute",
                opacity: 0.6,
              }}
            />
          </Pressable>
          );
        })
      }
      </ScrollView>
    </Header>
  );
}

export default YourDaily;
