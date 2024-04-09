import React from 'react';
import { View, StyleSheet, Text, TouchableOpacity, Image } from 'react-native';
import Header from '../components/Header';
import { useGetExplore } from '../libs';
import Swiper from 'react-native-swiper';
import { useState } from 'react';
import { AxiosError } from 'axios';

const Explore = ({ navigation }: { navigation: any }) => {
  const { error, isError, isLoading, data, refetch, isRefetching } = useGetExplore();
  const [currentIndex, setCurrentIndex] = useState(0);

  if (error) {
    const axiosError = error as AxiosError;
    if (axiosError.response?.status === 401) {
      navigation.navigate('Login');
      console.log("Unauthorized, redirecting to login");
    }
  }

  const handleSwipe = (index: number) => {
    setCurrentIndex(index);
    if (index > 3) {
      refetch();
    }
  };

  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
        {isLoading && <Text style={{ color: "white", fontSize: 16 }}>Loading...</Text>}
        {isError && <Text style={{ color: "white", fontSize: 16 }}>Error fetching data</Text>}
        {isRefetching && <Text style={{ color: "white", fontSize: 16 }}>Loading more...</Text>}

        <Swiper
          onIndexChanged={handleSwipe}
          loop={false}
          index={currentIndex}
          horizontal={false}
        >
          {data != undefined &&
            data.map((daily: any, index: number) => (
              <View key={index} style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
                <Image source={{ uri: daily.image }} style={{ width: '100%', height: '80%' }} />
              </View>
            ))
          }
        </Swiper>
      </View>
    </Header>
  );

  /**
  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
        <Text style={{ fontSize: 40, fontWeight: '200', color: 'white' }}>Explore</Text>
        <TouchableOpacity onPress={() => navigation.navigate("ReadDaily", { data })} style={{ aspectRatio: 1 / 1, width: '30%', borderWidth: 0.5, borderColor: 'gray', opacity: 0.85, borderRadius: 10, backgroundColor: '#0D1326' }}>
          <Image source={{ uri: data.image }} resizeMode='contain' style={{ width: '100%', height: '100%', borderRadius: 10 }} />
        </TouchableOpacity>
      </View>
    </Header>
  );
  */
}

const styles = StyleSheet.create({})

export default Explore;
