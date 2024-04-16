import React from 'react';
import { View, StyleSheet, Text, TouchableOpacity, Image } from 'react-native';
import Header from '../components/Header';
import { getExplore } from '../libs';
import Swiper from 'react-native-swiper';
import { useState, useEffect, useRef } from 'react';
import { AxiosError } from 'axios';
import uuidv4 from 'uuid/v4';
import FlipCard from './FlipCard';

const Explore2 = ({ navigation }: { navigation: any }) => {
  const [error, setError] = useState<AxiosError | null>(null);
  const [data, setData] = useState<any[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const currentIndex = useRef(0);

  const handleSwipe = (index: number) => {
    console.log("swiped");
    currentIndex.current = index;

    console.log(index);
    if (index >= (currentPage) * 5 - 1) {
      setCurrentPage((currentPage) => currentPage + 1);
    }
  };

  useEffect(() => {
    const abortController = new AbortController();
    const fetchData = async () => {
      try {
        const newData = await getExplore();
        setData(data => [...data, ...newData]);
        console.log(data);
        setError(null);
      } catch (error: any) {
        setError(error);
        console.error('Failed to fetch', error);
      }
    };

    fetchData();
    return () => abortController.abort();
  }, [currentPage]);

  useEffect(() => {
    if (error) {
      const axiosError = error as AxiosError;
      if (axiosError.response?.status === 401) {
        console.log("Unauthorized, redirecting to login");
        navigation.navigate('Login');
      }
    }
  }, [error, navigation]);

  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
        <Swiper
          key={uuidv4()}
          loop={false}
          onIndexChanged={handleSwipe}
          horizontal={false}
          index={currentIndex.current}
        >
          {
            (data.length > currentPage * 5 - 1) && data.map((daily: any, index: number) => (
              <FlipCard key={uuidv4()} dailyUrl={daily.image} dailyContent={daily.text} />
            ))
          }

        </Swiper>
      </View>
    </Header>
  );

  /*          {
              (data.length > currentPage * 5 - 1) && data.map((daily: any, index: number) => (
                <View key={uuidv4()} style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
                  <Image source={{ uri: daily.image }} style={{ width: '100%', height: '80%' }} />
                </View>
              ))
            }
            */

}
const styles = StyleSheet.create({})

export default Explore2;
