import React from 'react';
import { View, StyleSheet, Text, TouchableOpacity, Image } from 'react-native';
import Header from '../components/Header';
import { useGetExplore } from '../libs';
import Swiper from 'react-native-swiper';
import { useState } from 'react';

const Explore = ({ navigation }: { navigation: any }) => {
  const { isError, isLoading, data, refetch, isRefetching } = useGetExplore();
  const [currentIndex, setCurrentIndex] = useState(0);

  const handleSwipe = (index: number) => {
    setCurrentIndex(index);
    if (index > 4) {
      refetch();
    }
  };

  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
        {isLoading && <Text>Loading...</Text>}
        {isError && <Text>Error fetching data</Text>}
        {isRefetching && <Text>Loading more...</Text>}

        <Swiper
          onIndexChanged={handleSwipe}
          loop={false}
          index={currentIndex}
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
