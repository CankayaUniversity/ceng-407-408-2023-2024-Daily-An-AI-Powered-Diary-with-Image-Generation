import React from 'react';
import { View, StyleSheet, Text, TouchableOpacity, Image } from 'react-native';
import Header from '../components/Header';
import { useGetExplore, useGetExploreInfinite } from '../libs';
import Swiper from 'react-native-swiper';
import { useState, useEffect, useRef } from 'react';
import { AxiosError } from 'axios';

const Explore = ({ navigation }: { navigation: any }) => {
  // const { error, isError, isLoading, data, refetch, isRefetching } = useGetExplore();
  const { error, isError, isLoading, data, fetchNextPage, isFetchingNextPage } =
    useGetExploreInfinite();
  const currentIndex = useRef(0);
  const [currentPage, setCurrentPage] = useState(0);

  useEffect(() => {
    if (error) {
      const axiosError = error as AxiosError;
      if (axiosError.response?.status === 401) {
        console.log("Unauthorized, redirecting to login");
        navigation.navigate('Login');
      }
    }
  }, [error, navigation]);


  const pages = data?.pages.flat() ?? [];
  const handleSwipe = (index: number) => {
    console.log(index);
    currentIndex.current += 1;
    console.log(currentIndex.current + " " + currentPage);
    console.log(pages.length);
    if (index > 1) {
      if (data && index >= pages.length - 1) {
        console.log("fetching next page");
        setCurrentPage(data?.pages.length + 1);
        fetchNextPage();
      }
    }
  };


  // {isRefetching && <Text style={{ color: "white", fontSize: 16 }}>Loading more...</Text>}
  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
        {isLoading && <Text style={{ color: "white", fontSize: 16 }}>Loading...</Text>}
        {isError && <Text style={{ color: "white", fontSize: 16 }}>Error fetching data</Text>}
        {isFetchingNextPage && <Text style={{ color: "white", fontSize: 16 }}>Loading more...</Text>}

        <Swiper
          onIndexChanged={handleSwipe}
          loop={false}
          index={currentIndex.current}
          horizontal={false}
        >
          {pages !== undefined &&
            pages.map((daily: any) => (
              <View key={daily.id} style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
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
