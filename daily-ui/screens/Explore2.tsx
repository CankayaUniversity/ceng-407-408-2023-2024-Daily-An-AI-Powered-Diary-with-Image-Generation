import React from 'react';
import { View, StyleSheet, Text, TouchableOpacity, Image } from 'react-native';
import Header from '../components/Header';
import { useGetExplore } from '../libs';
import Swiper from 'react-native-swiper';
import { useState, useEffect, useRef } from 'react';
import { AxiosError } from 'axios';

const Explore2 = ({ navigation }: { navigation: any }) => {
  const { error, isError, isLoading, data, refetch, isRefetching } = useGetExplore();

  useEffect(() => {
    if (error) {
      const axiosError = error as AxiosError;
      if (axiosError.response?.status === 401) {
        console.log("Unauthorized, redirecting to login");
        navigation.navigate('Login');
      }
    }
  }, [error, navigation]);

  const [pageIndex, setPageIndex] = useState(0);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [swiperData, setData] = useState(data);

  // Effect to trigger fetching when pageIndex changes, beyond the initial load
  useEffect(() => {
    if (pageIndex > currentIndex) {
      refetch();
    }
  }, [pageIndex, refetch]);

  // Example swiper onIndexChanged logic to trigger loading more data
  const handleIndexChange = (swiperIndex: number) => {
    if (swiperIndex === 5 * currentIndex) { // Adjust as needed for your use case
      setPageIndex(pageIndex + 1);
    }
  };
  (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error fetching data</div>;
  }

  // Render your swiper with the data
  return (
    <Swiper
      onIndexChanged={handleIndexChange}
    // Include other swiper props as needed
    >
      {swiperData.map((item, index) => (
        <div key={index} style={{/* Style for each swiper slide /}}>
{/ Render your item content here */}
</div>
  ))
}
</Swiper >
);
};
}

const styles = StyleSheet.create({})

export Explore2;
