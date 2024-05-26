import React from 'react';
import { ScrollView, Pressable, Text, Image,View } from 'react-native';
import Header from '../components/Header';
import { useGetDailies } from '../libs';

const YourDaily = ({ navigation }: { navigation: any }) => {
  const { isError, isLoading, data } = useGetDailies(10);
  const dateDict = (dateStr:string) => {
    const date = new Date(dateStr);
    const monthArr = ["JAN",	"FEB",	"MAR",	"APR",	"MAY",	"JUN",	"JUL",	"AUG",	"SEPT",	"OCT",	"NOV",	"DEC"]

    return monthArr[date.getMonth()+1]+"\n"+date.getDay();
  };

  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={{ flexDirection: "row", flexWrap: 'wrap', justifyContent: 'flex-start', gap: 10, alignItems: 'center', marginTop: 10, paddingStart: 15, paddingBottom: 90 }}>
        {
          isLoading &&
          <Text style={{ alignItems: 'center', justifyContent: 'center', fontSize: 40, color: 'white' }}>Loading</Text>
        }
        {data?.map((el, index) => {
          return (
            <Pressable
            key={index}
            onPress={() => navigation.navigate("ReadDaily", { data: el })}
            style={{
              aspectRatio: 1,
              width: '30%',
              borderWidth: 0.5,
              borderColor: 'gray',
              alignItems: "center",
              justifyContent: "center",
              opacity: 0.85,
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
